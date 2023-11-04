package server

import (
	"app/common/interceptor"
	"app/common/logger"
	"app/common/public"
	"app/internal/model"
	"app/internal/router"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func New() *Server {
	new(model.User).InitSystemUser()
	handle := public.GetHandler()
	gin.SetMode(handle.Config.Server.Mode)
	return &Server{engine: gin.New(), handle: handle}
}

type Server struct {
	engine *gin.Engine
	server *http.Server
	handle *public.Handler
}

// Start api服务入口
func (srv *Server) Start() error {
	srv.server = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", srv.handle.Config.Server.Host, srv.handle.Config.Server.Port),
		Handler:        srv.engine,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	srv.engine.Use(
		interceptor.Cors(),
		interceptor.Context(),
		interceptor.Recover(),
	)
	router.Route(srv.engine)
	logger.System("server listen: http://%s", srv.server.Addr)
	return srv.listen()
}

func (srv *Server) listen() (err error) {
	errs := make(chan error, 1)
	go func() {
		if err = srv.server.ListenAndServe(); err != nil {
			if err != nil {
				errs <- err
			}
		}
	}()
	manager := make(chan os.Signal, 1)
	signal.Notify(manager, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err = <-errs:
		break
	case s := <-manager:
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			err = srv.server.Close()
		}
	}
	return err
}
