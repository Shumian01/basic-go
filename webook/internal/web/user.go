package web

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
	"webook/internal/domain"
	"webook/internal/service"

	ijwt "webook/internal/web/jwt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	biz                  = "login"
)

var _ headler = &UserHandler{}

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            service.UserService
	codeSvc        service.CodeService
	cmd            redis.Cmdable
	ijwt.Handler
}

func NewUserHandler(svc service.UserService, codeSvc service.CodeService, cmd redis.Cmdable, jwtHdl ijwt.Handler) *UserHandler {
	return &UserHandler{
		svc:     svc,
		codeSvc: codeSvc,
		cmd:     cmd,
		Handler: jwtHdl,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	//ug.GET("/profile", u.Profile)
	ug.GET("/profile", u.ProfileJWT)
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.LoginJWT)
	ug.POST("/logout", u.LogoutJWT)
	//ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	//put "/login_sms/code/send" 发验证码
	//post "/login_sms/code/send" 验证验证码
	ug.POST("/login_sms/code/send", u.SendLoginSMSCode)
	ug.POST("/login_sms", u.LoginSMS)
	ug.POST("/refresh_token", u.RefreshToken)
}

func (u *UserHandler) LogoutJWT(ctx *gin.Context) {

	err := u.ClearToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Msg:  "退出登录失败",
			Code: 5,
		})
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Msg: "退出登录成功",
	})
}

// RefreshToken 可以同时刷新长短token 用redis来记录是否有效, 即refresh_token是一次性的
// 参考登录校验部分 比较userAgent来提高安全性
func (u *UserHandler) RefreshToken(ctx *gin.Context) {
	//只有这个接口 拿出来才是refreshToken 其他地方都是accesstoken
	refreshToken := u.ExtractToken(ctx)
	var rc ijwt.RefreshClaims
	token, err := jwt.ParseWithClaims(refreshToken, &rc, func(*jwt.Token) (interface{}, error) {
		return ijwt.AtKey, nil
	})
	if err != nil || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err = u.CheckSession(ctx, rc.Ssid)
	if err != nil {
		// 要么redis出问题 要么已经退出登录
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//搞个新的accesstoken
	err = u.SetJWTToken(ctx, rc.Uid, rc.Ssid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Msg: "刷新成功",
	})
}
func (u *UserHandler) LoginSMS(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	//这边可以加上各种校验
	ok, err := u.codeSvc.Verify(ctx, biz, req.Phone, req.Code)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "验证码有误",
		})
		return
	}
	user, err := u.svc.FindOrCreate(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}
	//jwt校验
	if err = u.SetLoginToken(ctx, user.Id); err != nil {
		//记录日志
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, Result{
		Code: 4,
		Msg:  "验证码校验通过",
	})

}
func (u *UserHandler) SendLoginSMSCode(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	if req.Phone == "" {
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "输入有误",
		})
		return
	}
	const biz = "login"
	err := u.codeSvc.Send(ctx, biz, req.Phone)
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Result{
			Msg: "发送成功",
		})
	case service.ErrCodeSendTooMany:
		ctx.JSON(http.StatusOK, Result{
			Msg: "发送太频繁,请稍后再试",
		})
	default:
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
	}
}

func (h *UserHandler) Signup(ctx *gin.Context) {
	type SignupReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignupReq
	//Bind方法 会根据Content-Type来解析
	if err := ctx.Bind(&req); err != nil {
		ctx.String(http.StatusBadRequest, "请求格式错误")
		return
	}
	// 邮箱正则
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		ctx.String(http.StatusBadRequest, "邮箱格式不对")
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusBadRequest, "两次密码不一致")
		return
	}

	// 密码校验
	if len(req.Password) < 9 {
		ctx.String(http.StatusBadRequest, "密码格式不对")
		return
	}
	if !regexp.MustCompile(`[a-zA-Z]`).MatchString(req.Password) {
		ctx.String(http.StatusBadRequest, "密码格式不对")
		return
	}
	if !regexp.MustCompile(`\d`).MatchString(req.Password) {
		ctx.String(http.StatusBadRequest, "密码格式不对")
		return
	}
	//调用svc方法进行注册
	err := h.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.ErrUserDuplicated {
		ctx.String(http.StatusBadRequest, "邮箱冲突")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}
	fmt.Printf("注册信息：%+v\n", req)
	ctx.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		ctx.String(http.StatusBadRequest, "请求格式错误")
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "账号或者密码不对")
		return
	}

	if err != nil {
		ctx.String(http.StatusBadRequest, "系统错误")
		return
	}
	//步骤2
	//使用JWT设置登录态
	//生成一个JWT token

	//JWT Token 里面携带我的个人信息
	//比如 带userID
	if err := u.SetLoginToken(ctx, user.Id); err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	fmt.Println(user)
	fmt.Printf("%v+", req)
	ctx.String(http.StatusOK, "登录成功")
	return
}

func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "账号或者密码不对")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	//在这里登录成功了
	//设置session
	sees := sessions.Default(ctx)
	//我可以随便设置值了
	//你要放在sess里面的值
	sees.Set("userId", user.Id)
	sees.Options(sessions.Options{
		//HttpOnly: true,
		//Secure:   true, 只能在https上使用
		MaxAge: 60,
	})
	sees.Save()
	fmt.Printf("%v+", req)
	ctx.String(http.StatusOK, "登录成功")
	return
}

func (h *UserHandler) Edit(ctx *gin.Context) {
	type Req struct {
		Nickname string `json:"nickname"`
		AboutMe  string `json:"aboutMe"`
		// 年月日: yyyy-MM-dd
		Birthday string `json:"birthday"`
	}
	var req Req

	// 1) 绑定错误 -> 400 + JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "请求格式错误",
		})
		return
	}
	fmt.Println(req)

	// 2) 未登录 -> 401 + JSON
	u, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 1,
			"msg":  "未登录",
		})
		return
	}
	uc, ok := u.(*ijwt.UserClaims) // 确保类型与中间件一致（指针/包名）
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 1,
			"msg":  "未登录",
		})
		return
	}

	// 3) 生日校验 -> 400+JSON（如果允许为空，可先判断空串再 Parse）
	birthday, err := time.Parse(time.DateOnly, req.Birthday)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "生日格式不对，需 yyyy-MM-dd",
		})
		return
	}

	// 4) 更新出错 -> 500+JSON
	err = h.svc.UpdateNonSensitiveInfo(ctx, domain.User{
		Id:       uc.Uid,
		Nickname: req.Nickname,
		AboutMe:  req.AboutMe,
		Birthday: birthday,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 5,
			"msg":  "系统异常",
		})
		return
	}

	// 5) 成功 -> 200+JSON
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "更新成功",
	})
}

func (u *UserHandler) LogOut(ctx *gin.Context) {
	sees := sessions.Default(ctx)
	//我可以随便设置值了
	//你要放在sess里面的值
	sees.Options(sessions.Options{
		//HttpOnly: true,
		//Secure:   true, 只能在https上使用
		MaxAge: -1,
	})
	sees.Save()
	ctx.String(http.StatusOK, "退出登录成功")
	return
}
func (u *UserHandler) Profile(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Get("userId")

	ctx.String(http.StatusOK, "czh 的个人主页")

}
func (h *UserHandler) ProfileJWT(ctx *gin.Context) {
	type Profile struct {
		Nickname string `json:"Nickname"`
		Email    string `json:"Email"`
		AboutMe  string `json:"AboutMe"`
		Birthday string `json:"Birthday"`
	}
	ctx.Header("Cache-Control", "no-store")

	uc, ok := ctx.MustGet("user").(*ijwt.UserClaims)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	u, err := h.svc.Profile(ctx, uc.Uid)
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}
	fmt.Println(u)
	ctx.JSON(http.StatusOK, Profile{
		Nickname: u.Nickname,
		Email:    u.Email,
		AboutMe:  u.AboutMe,
		Birthday: u.Birthday.Format(time.DateOnly),
	})
}
