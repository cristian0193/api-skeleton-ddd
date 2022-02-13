package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/context"
)

const rootPrefix = "/golang-service"

type Server struct {
	server  *echo.Echo
	startAt time.Time
	port    int
}

type healthResponse struct {
	Status  string    `json:"status"`
	StartAt time.Time `json:"start_at"`
}

func NewServer(port int) *Server {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())

	path := e.Group(rootPrefix)

	server := &Server{server: e, startAt: time.Now(), port: port}

	path.GET("/health", server.health)

	return server
}

func (s *Server) health(c echo.Context) error {
	r := &healthResponse{Status: "OK", StartAt: s.startAt}
	return c.JSON(http.StatusOK, r)
}

func (s *Server) Start() error {
	port := fmt.Sprintf(":%v", s.port)
	if err := s.server.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server.Start: %w", err)
	}
	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server.Shutdown: %w", err)
	}
	return nil
}
