package handler

import "github.com/gin-gonic/gin"

func (obj *handler) LogIn(c *gin.Context) {
	ctx := c.Request.Context()

	obj.uc.RegisterUser(ctx)
}

func (obj *handler) LogUp(c *gin.Context) {
	ctx := c.Request.Context()

	obj.uc.RegisterUser()
}
