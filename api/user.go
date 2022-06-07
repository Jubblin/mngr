package api

import (
	"github.com/gin-gonic/gin"
	"mngr/models"
	"mngr/ws"
	"net/http"
)

func RegisterUserEndpoints(router *gin.Engine, holders *ws.Holders) {
	rb := holders.Rb
	router.POST("/login", func(ctx *gin.Context) {
		var lu models.LoginUserViewModel
		if err := ctx.BindJSON(&lu); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		u, err := rb.UserRep.Login(&lu)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if u != nil {
			rb.Users[u.Token] = u
		}
		ctx.JSON(http.StatusOK, u)
	})

	router.POST("/registeruser", func(ctx *gin.Context) {
		var ru models.RegisterUserViewModel
		if err := ctx.BindJSON(&ru); err != nil {
			ctx.JSON(http.StatusBadRequest, false)
			return
		}
		if ru.Password != ru.RePassword {
			ctx.JSON(http.StatusBadRequest, false)
			return
		}
		u, err := rb.UserRep.Login(&models.LoginUserViewModel{Username: ru.Username, Password: ru.Password})
		if u != nil && err == nil {
			ctx.JSON(http.StatusNotFound, true)
			return
		}
		u, err = rb.UserRep.Register(&ru)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, false)
			return
		}
		if u != nil {
			rb.LoadUser()
			holders.Init()
		}
		ctx.JSON(http.StatusOK, true)
	})

	router.GET("/users", func(ctx *gin.Context) {
		services, err := rb.UserRep.GetUsers()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, services)
	})

	router.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		result, err := rb.UserRep.RemoveById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			rb.LoadUser()
			holders.Init()
			ctx.JSON(http.StatusOK, result)
		}
	})
}
