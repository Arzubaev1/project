package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	handler := handler.NewHandler(cfg, storage, logger)

	r.Use(customCORSMiddleware())

	// Login Api
	r.POST("/login", handler.Login)

	// Register Api
	r.POST("/register", handler.Register)

	// User Api
	r.POST("/user", handler.AuthMiddleware(), handler.CreateUser)
	r.GET("/user/:id", handler.AuthMiddleware(), handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.AuthMiddleware(), handler.UpdateUser)
	r.DELETE("/user/:id", handler.AuthMiddleware(), handler.DeleteUser)

	r.POST("/driver", handler.AuthMiddleware(), handler.CreateDriver)
	r.GET("/driver/:id", handler.AuthMiddleware(), handler.GetByIdDriver)
	r.GET("/driver", handler.GetListDriver)
	r.PUT("/driver/:id", handler.AuthMiddleware(), handler.UpdateDriver)
	r.DELETE("/driver/:id", handler.AuthMiddleware(), handler.DeleteDriver)

	r.POST("/car", handler.AuthMiddleware(), handler.CreateCar)
	r.GET("/car/:id", handler.AuthMiddleware(), handler.GetByIdCar)
	r.GET("/car", handler.GetListCar)
	r.PUT("/car/:id", handler.AuthMiddleware(), handler.UpdateCar)
	r.DELETE("/car/:id", handler.AuthMiddleware(), handler.DeleteCar)

	r.POST("/order", handler.AuthMiddleware(), handler.CreateOrder)
	r.GET("/order/:id", handler.AuthMiddleware(), handler.GetByIdOrder)
	r.GET("/order", handler.GetListOrder)
	r.PUT("/order/:id", handler.AuthMiddleware(), handler.UpdateOrder)
	r.DELETE("/order/:id", handler.AuthMiddleware(), handler.DeleteOrder)

	r.POST("/branch", handler.AuthMiddleware(), handler.CreateBranch)
	r.GET("/branch/:id", handler.AuthMiddleware(), handler.GetByIdBranch)
	r.GET("/branch", handler.GetListBranch)
	r.PUT("/branch/:id", handler.AuthMiddleware(), handler.UpdateBranch)
	r.DELETE("/branch/:id", handler.AuthMiddleware(), handler.DeleteBranch)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accesp-Encoding, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
