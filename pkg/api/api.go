package api

import (
    "enterprise-go-app/pkg/service"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func RegisterRoutes(r *gin.Engine, svc *service.UserService) {
    r.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, svc.ListUsers())
    })

    r.GET("/users/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
            return
        }
        if user, ok := svc.GetUser(id); ok {
            c.JSON(http.StatusOK, user)
        } else {
            c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        }
    })
}