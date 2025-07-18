package v1

import (
	"awesomeProject666/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Init() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/items", h.getAllItems)
	v1.POST("/auth/login", h.login)
	v1.POST("/user", h.createUser)

	authGroup := v1.Group("/")
	authGroup.Use(middleware.JWTAuth(h.JWT))
	{
		authGroup.GET("/users", h.getAllUsers)
		authGroup.GET("/user/search", h.getUser)
	}

	return r
}
