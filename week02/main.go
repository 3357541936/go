package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"week02.com/internal/middleware/login"
	"week02.com/internal/repository"
	"week02.com/internal/repository/dao"
	"week02.com/internal/service"
	"week02.com/internal/web"
)

func main() {
	db := initDB()
	server := initWebServer()

	initUserHandler(db, server)

	server.Run(":8081")
}

func initUserHandler(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	userRouter := web.NewUserHandler(us)
	userRouter.CreateRouter(server)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		ExposeHeaders:   []string{"x-jwt-token"},
		MaxAge:          12 * time.Hour,
	}), func(context *gin.Context) {
		//println("跨域功能开启!")
		// Web 治理: 熔断, 限流, 降级
		// 可观测性: 日志, metrics, tracing
		// 身份认证与鉴权等
	})

	//useSession(engine)
	useJWT(engine)

	return engine
}

func useSession(engine *gin.Engine) {
	loginCheck := &login.MiddlewareBuilder{}
	//store := cookie.NewStore([]byte("secret"))
	store, err := redis.NewStore(8, "tcp",
		"localhost:6379", "",
		[]byte("kEUuGqea3ADEk05zwBm5gOPd6x4SGyHe"),
		[]byte("pA7W4QQ6rkAThz2I5IeL7h5gLBgQFiFM"))
	if err != nil {
		panic("Redis仓库初始化错误!")
	}
	engine.Use(sessions.Sessions("ssid", store), loginCheck.CheckLogin())
}

func useJWT(engine *gin.Engine) {
	loginCheck := &login.MiddlewareJWTBuilder{}
	engine.Use(loginCheck.CheckLogin())
}
