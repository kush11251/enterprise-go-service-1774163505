package api

import (
    "enterprise-go-app/pkg/repository"
    "enterprise-go-app/pkg/service"
    "github.com/gin-gonic/gin"
    "net/http/httptest"
    "testing"
)

func TestListUsersRoute(t *testing.T) {
    repo := repository.NewUserRepository()
    svc := service.NewUserService(repo)
    r := gin.Default()
    RegisterRoutes(r, svc)

    req := httptest.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    if w.Code != 200 {
        t.Fatalf("expected 200 got %d", w.Code)
    }
}