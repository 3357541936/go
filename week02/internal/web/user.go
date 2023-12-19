package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"week02.com/internal/domain"
	"week02.com/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

type UserClaims struct {
	jwt.RegisteredClaims
	uid int64
}

var SignedKey = []byte("kEUuGqea3ADEk05zwBm5gOPd6x4SGyHO")

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (s *UserHandler) CreateRouter(server *gin.Engine) {
	group := server.Group("/user")
	// 用户登录接口
	group.POST("/login", s.loginJWT)
	// 用户注册接口
	group.POST("/signup", s.signup)
	// 用户补充基本信息接口
	group.POST("/edit", s.edit)
	// 用户修改基本信息接口
	group.POST("/profile", s.profile)
}

func (s *UserHandler) signup(context *gin.Context) {
	// 声明请求字段信息
	type SignupData struct {
		Email          string `json:"email"`
		PinNumber      string `json:"pin_number"`
		CheckPinNumber string `json:"check_pin_number"`
	}
	var req SignupData

	// 获取请求数据信息,传递字段需要与结构体定义字段一致.
	if err := context.Bind(&req); err != nil {
		return
	}

	// 调用service模块
	err := s.service.Signup(context, domain.User{
		Email:    req.Email,
		Password: req.PinNumber,
	})

	switch err {
	case nil:
		context.String(http.StatusOK, "注册成功.")
	case service.ErrDuplicateEmail:
		context.String(http.StatusOK, "注册邮箱冲突, 请更改!")
	default:
		context.String(http.StatusOK, "系统出错, 请联系管理员!")

	}
}

func (s *UserHandler) loginJWT(context *gin.Context) {
	type LoginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var requestData LoginData
	if err := context.Bind(&requestData); err != nil {
		return
	}
	u, err := s.service.Login(context, requestData.Username, requestData.Password)
	switch err {
	case nil:
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
			uid: u.Id,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			},
		})
		tokenStr, err := token.SignedString(SignedKey)
		if err != nil {
			context.String(http.StatusOK, "系统出错!")
			return
		}
		context.Header("x-jwt-token", tokenStr)
		// 登录成功
		type Obj struct {
			Id          int64  `json:"id"`
			Email       string `json:"email"`
			Name        string `json:"name"`
			Birth       int64  `json:"birth"`
			Description string `json:"description"`
		}
		context.JSON(http.StatusOK, &Obj{
			Id:          u.Id,
			Email:       u.Email,
			Name:        u.Name,
			Birth:       u.Birth,
			Description: u.Description,
		})
	case service.ErrInvailidEmailOrPassword:
		context.String(http.StatusOK, "您的邮箱或密码无效!")
	default:
		context.String(http.StatusOK, "系统出错, 请联系管理员!")
	}
}

func (s *UserHandler) login(context *gin.Context) {
	type LoginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var requestData LoginData
	if err := context.Bind(&requestData); err != nil {
		return
	}
	u, err := s.service.Login(context, requestData.Username, requestData.Password)
	switch err {
	case nil:
		sess := sessions.Default(context)
		sess.Set("userId", u.Id)
		sess.Options(sessions.Options{
			MaxAge: 60 * 15,
		})
		err := sess.Save()
		if err != nil {
			context.String(http.StatusOK, "系统出错, 请联系管理员!")
			return
		}
		type Obj struct {
			Id          int64  `json:"id"`
			Email       string `json:"email"`
			Name        string `json:"name"`
			Birth       int64  `json:"birth"`
			Description string `json:"description"`
		}
		context.JSON(http.StatusOK, &Obj{
			Id:          u.Id,
			Email:       u.Email,
			Name:        u.Name,
			Birth:       u.Birth,
			Description: u.Description,
		})
	case service.ErrInvailidEmailOrPassword:
		context.String(http.StatusOK, "您的邮箱或密码无效!")
	default:
		context.String(http.StatusOK, "系统出错, 请联系管理员!")
	}
}

func (s *UserHandler) edit(context *gin.Context) {
	type EditData struct {
		Id          int64  `json:"id"`
		Name        string `json:"name"`
		Birth       string `json:"birth"`
		Description string `json:"description"`
	}
	var requestData EditData
	if err := context.Bind(&requestData); err != nil {
		return
	}
	//sess := sessions.Default(context)
	//uid := sess.Get("userId")

	user := context.MustGet("user").(UserClaims)

	if requestData.Name == "" || requestData.Birth == "" || requestData.Description == "" {
		context.String(http.StatusOK, "提交信息中存在未填写项")
		return
	}
	timeParse, _ := time.Parse("2006-01-02 15:05:04", requestData.Birth)
	err := s.service.Edit(context, domain.User{
		Id:          user.uid,
		Name:        requestData.Name,
		Birth:       timeParse.UnixMilli(),
		Description: requestData.Description,
	})
	switch err {
	case nil:
		context.String(http.StatusOK, "个人信息修改成功")
	default:
		context.String(http.StatusOK, "系统出错,请联系管理员!")
	}
}

func (s *UserHandler) profile(context *gin.Context) {
	type ProfileData struct {
		Id          int64  `json:"id"`
		Name        string `json:"name"`
		Birth       string `json:"birth"`
		Description string `json:"description"`
	}
	var requestData ProfileData
	if err := context.Bind(&requestData); err != nil {
		return
	}

	//sess := sessions.Default(context)
	//uid := sess.Get("userId")
	user := context.MustGet("user").(UserClaims)
	if requestData.Name == "" || requestData.Birth == "" || requestData.Description == "" {
		context.String(http.StatusOK, "提交信息中存在未填写项")
		return
	}
	timeParse, _ := time.Parse("2006-01-02 15:05:04", requestData.Birth)
	err := s.service.Profile(context, domain.User{
		Id:          user.uid,
		Name:        requestData.Name,
		Birth:       timeParse.UnixMilli(),
		Description: requestData.Description,
	})
	switch err {
	case nil:
		context.String(http.StatusOK, "个人信息修改成功")
	default:
		context.String(http.StatusOK, "系统出错,请联系管理员!")

	}
}
