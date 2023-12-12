package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"week02.com/internal/domain"
	"week02.com/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (s *UserHandler) CreateRouter(server *gin.Engine) {
	group := server.Group("/user")
	// 用户登录接口
	group.POST("/login", s.login)
	// 用户注册接口
	group.POST("/signup", s.signup)
	// 用户编辑接口
	group.POST("/edit", s.edit)
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
		}
		context.String(http.StatusOK, "登录成功!")
	case service.ErrInvailidEmailOrPassword:
		context.String(http.StatusOK, "您的邮箱或密码无效!")
	default:
		context.String(http.StatusOK, "系统出错, 请联系管理员!")
	}
}

func (s *UserHandler) edit(context *gin.Context) {
	type EditData struct {
		Id            int64  `json:"id"`
		Username      string `json:"username"`
		Password      string `json:"password"`
		CheckPassword string `json:"checkPassword"`
	}
	var requestData EditData
	if err := context.Bind(&requestData); err != nil {
		return
	}
	if requestData.Password != requestData.CheckPassword {
		context.String(http.StatusOK, "两次密码输入不一致")
	}
	err := s.service.Edit(context, domain.User{
		Id:       requestData.Id,
		Email:    requestData.Username,
		Password: requestData.Password,
	})
	switch err {
	case nil:
		context.String(http.StatusOK, "个人信息修改成功")
	default:
		context.String(http.StatusOK, "系统出错,请联系管理员!")

	}
}
