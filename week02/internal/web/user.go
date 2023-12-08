package web

import (
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
	group.GET("/login", s.login)
	// 用户注册接口
	group.POST("/signup", s.signup)
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

	if err != nil {
		// 系统错误
		context.String(http.StatusOK, "系统不支持当前操作!")
		return
	}

	context.String(http.StatusOK, "申请被执行.")
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

	context.String(http.StatusOK, "This is Login Service")
}
