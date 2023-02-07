package server

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	counter int64

	server *http.Server
	router *gin.Engine
}

func New() *Server {
	router := gin.Default()
	s := &Server{
		router:  router,
		counter: int64(0),
	}

	router.GET("/", s.CounterHandler)
	router.GET("/health_checks", s.HealthHandler)
	return s
}

func (s *Server) HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (s *Server) CounterHandler(ctx *gin.Context) {
	counter := atomic.AddInt64(&s.counter, 1)
	ctx.JSON(http.StatusOK, gin.H{"counter": counter})
}

func (s *Server) Start(addr string) error {
	s.server = &http.Server{
		Addr:        addr,
		Handler:     s.router,
		ReadTimeout: 10 * time.Second,
	}
	log.Printf("start server on %s", addr)
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	log.Println("stop server")
	if s.server != nil {
		return s.server.Close()
	}
	return nil
}
