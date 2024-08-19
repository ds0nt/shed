package api

import (
	"github.com/ds0nt/shed/pkg/log"
	"github.com/ds0nt/shed/pkg/storage"
	"github.com/ds0nt/shed/pkg/storage/leveldb_storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Service struct {
	// Echo is the main router object
	Echo  *echo.Echo
	Store storage.Storer
}

// NewService creates a new instance of the Service
func NewService() *Service {
	store, err := leveldb_storage.NewLevelDBStorage("data")
	if err != nil {
		panic(err)
	}
	return &Service{
		Echo:  echo.New(),
		Store: store,
	}
}

// StartServer starts the server
func (s *Service) StartServer() {
	log.Info("Starting Server")

	// Middleware
	s.Echo.Use(middleware.Logger())
	s.Echo.Use(middleware.Recover())
	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	s.Echo.POST("/register", s.createUserHandler)
	s.Echo.POST("/login", s.loginUserHandler)

	s.Echo.GET("/conversations", s.listConversationsHandler)
	s.Echo.POST("/conversations", s.createConversationHandler)
	s.Echo.GET("/conversations/:id", s.getConversationHandler)
	s.Echo.POST("/send-message/:id", s.sendMessageHandler)

	// Serve static assets
	s.Echo.Static("/", "ui")

	// Start the server
	s.Echo.Start(":8080")
}
