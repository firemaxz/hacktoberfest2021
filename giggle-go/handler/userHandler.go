package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewHandlerUsers(service user.Service) *userHandler {
	return &userHandler{service}
}

// GET 
func (h *userHandler) Getusers(c *gin.Context) {
	// your logic here
	
	
	// response := helper.APIResponse("Success to get users", http.StatusOK, "success", user.FormatUsers(users))
	// c.JSON(http.StatusOK, response)
}


	