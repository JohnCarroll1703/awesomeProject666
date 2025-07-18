package v1

import (
	"awesomeProject666/app/internal/auth"
	"awesomeProject666/app/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	Logger   *logrus.Logger
	Services *service.Services
	JWT      *auth.JWTManager
}

func NewHandler(logger *logrus.Logger,
	services *service.Services,
	jwtManager *auth.JWTManager) *Handler {
	return &Handler{
		Logger:   logger,
		Services: services,
		JWT:      jwtManager,
	}
}

func (h *Handler) Health(c *gin.Context) {
	h.Logger.Info("Health check called")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
