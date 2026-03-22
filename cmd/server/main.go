package main

import (
    "enterprise-go-app/pkg/api"
    "enterprise-go-app/pkg/repository"
    "enterprise-go-app/pkg/service"
    "github.com/gin-gonic/gin"
)

func main() {
    repo := repository.NewUserRepository()
    svc := service.NewUserService(repo)
    r := gin.Default()
    api.RegisterRoutes(r, svc)
    r.Run(":8080")
}