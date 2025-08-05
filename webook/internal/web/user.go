package web

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"regexp"
	"time"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) RegisterUser(server *gin.Engine) {
	ug := server.Group("/users")
	//ug.GET("/profile", u.Profile)
	ug.GET("/profile", u.ProfileJWT)
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.LoginJWT)
	//ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
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
	if err == service.ErrUserDuplicatedEmail {
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
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		},
		Uid:       user.Id,
		UserAgent: ctx.Request.UserAgent(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte("dopWHWvRXiyHULAkR90XQsR06Uvl7PFX"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	ctx.Header("x-jwt-token", tokenStr)

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

func (u *UserHandler) Edit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "编辑接口待实现")
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
func (u *UserHandler) ProfileJWT(ctx *gin.Context) {
	c, ok := ctx.Get("claims")
	//你可以断定必然有claims
	if !ok {
		//可以考虑监控住这里
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	//ok代表 是不是userClaims
	claims, ok := c.(*UserClaims)
	if !ok {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	println(claims.Uid)
	ctx.String(http.StatusOK, "你的profile")
	//这边就是补充profile其他代码
}

type UserClaims struct {
	jwt.RegisteredClaims
	//声明你自己要放进token里面的数据
	Uid       int64
	UserAgent string
}
