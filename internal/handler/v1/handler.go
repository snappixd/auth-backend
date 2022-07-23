package v1

import (
	"auth-back/internal/service"
	"auth-back/pkg/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)
	}
}

// func parseIdFromPath(c *gin.Context, param string) (primitive.ObjectID, error) {
// 	idParam := c.Param(param)
// 	if idParam == "" {
// 		return primitive.ObjectID{}, errors.New("empty id param")
// 	}

// 	id, err := primitive.ObjectIDFromHex(idParam)
// 	if err != nil {
// 		return primitive.ObjectID{}, errors.New("invalid id param")
// 	}

// 	return id, nil
// }
