package web

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"regexp"
	"time"
	"webook/internal/domain"
	"webook/internal/service"
)

const biz = "login"

var _ headler = &UserHandler{}

type UserHandler struct {
	svc     service.UserService
	codeSvc service.CodeService
}

func NewUserHandler(svc service.UserService, codeSvc service.CodeService) *UserHandler {
	return &UserHandler{
		svc:     svc,
		codeSvc: codeSvc,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	//ug.GET("/profile", u.Profile)
	ug.GET("/profile", u.ProfileJWT)
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.LoginJWT)
	//ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	//put "/login_sms/code/send" 发验证码
	//post "/login_sms/code/send" 验证验证码
	ug.POST("/login_sms/code/send", u.SendLoginSMSCode)
	ug.POST("/login_sms", u.LoginSMS)
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
	if err = u.SetJWTToken(ctx, user.Id); err != nil {
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

func (u *UserHandler) Signup(ctx *gin.Context) {
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
		ctx.String(http.StatusBadRequest, "密码必须至少9位")
		return
	}
	if !regexp.MustCompile(`[a-zA-Z]`).MatchString(req.Password) {
		ctx.String(http.StatusBadRequest, "密码必须包含字母")
		return
	}
	if !regexp.MustCompile(`\d`).MatchString(req.Password) {
		ctx.String(http.StatusBadRequest, "密码必须包含数字")
		return
	}
	//调用svc方法进行注册
	err := u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.ErrUserDuplicated {
		ctx.String(http.StatusBadRequest, "邮箱冲突")
		return
	}
	if err != nil {
		ctx.String(http.StatusBadRequest, "系统错误")
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
	if err := u.SetJWTToken(ctx, user.Id); err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	fmt.Println(user)
	fmt.Printf("%v+", req)
	ctx.String(http.StatusOK, "登录成功")
	return
}

func (u *UserHandler) SetJWTToken(ctx *gin.Context, uid int64) error {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
		Uid:       uid,
		UserAgent: ctx.Request.UserAgent(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PFX"))
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenStr)
	return err
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
	uc, ok := u.(*UserClaims) // 确保类型与中间件一致（指针/包名）
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
	uc, ok := ctx.MustGet("user").(*UserClaims)
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

type UserClaims struct {
	jwt.RegisteredClaims
	//声明你自己要放进token里面的数据
	Uid       int64
	UserAgent string
}
