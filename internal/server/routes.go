package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ostheperson/go-auth-service/internal/domain"
	"github.com/ostheperson/go-auth-service/internal/handlers"
)

func RegisterRoutes(s *domain.Server) http.Handler {
	r := gin.Default()
	// TODO: add role middleware
	// r.GET("/", s.HelloWorldHandler)
	// r.GET("/health", s.healthHandler)

	ah := handlers.NewAuthHandler(s)
	r.POST("/account/signup", ah.SignUp)
	r.POST("/account/signin", ah.SignIn)

	protected := r.Group("/v1")
	protected.Use(JwtAuthMiddleware(s.Env.AccessTokenSecret))
	uh := handlers.NewUsersHandler(s)

	protected.GET("/users", uh.GetUsers)
	protected.GET("/users/:id", uh.GetUser)
	protected.PATCH("/users/:id", uh.UpdateUser)
	protected.DELETE("/users/:id", uh.RemoveUser)
	protected.DELETE("/users/all", uh.GetUsers)

	return r
}

// func (s *Server) HelloWorldHandler(c *gin.Context) {
// 	resp := make(map[string]string)
// 	resp["message"] = "live"
//
// 	c.JSON(http.StatusOK, resp)
// }
//
// func (s *Server) healthHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, s.db.Health())
// }
