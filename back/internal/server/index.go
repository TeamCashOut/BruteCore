package server

import (
	"time"

	"api.brutecore/libs/lib_env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type ServerInterface interface {
	SetMiddleware()
	SetHandlers(hdlr HandlerInterface)
	Listen() error
}

type HandlerInterface interface {
	SetHandlers(app *fiber.App)
}

type Server struct {
	conf *lib_env.Config
	Srv  *fiber.App
}

const (
	loggerFormat = `
[${time}] [X-Request-ID: ${respHeader:x-request-id}] [${ip} -> ${host}]
${method} ${path} ${status} (${latency})
RequestHeaders: ${reqHeaders}
RequestBody: ${body}
ResponseHeaders: ${respHeader}
ResponseBody: ${resBody}
` + "\n"
	timeFormat = "02-Jan-2006 15:04:05"
)

func New(cf *lib_env.Config) ServerInterface {
	return &Server{
		Srv: fiber.New(fiber.Config{
			AppName:      cf.App.Name + " v" + cf.App.Version,
			BodyLimit:    15 * 1024 * 1024,
			WriteTimeout: time.Second * 120,
		}),
		conf: cf,
	}
}

func (s *Server) SetMiddleware() {
	s.Srv.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	s.Srv.Use(requestid.New())

	if s.conf.Logs.Http != nil {
		s.Srv.Use(logger.New(logger.Config{
			Format:     loggerFormat,
			TimeFormat: timeFormat,
			TimeZone:   "Asia/Almaty",
			Output:     s.conf.Logs.Http,
		}))
	}
}

func (s *Server) SetHandlers(hdlr HandlerInterface) {
	hdlr.SetHandlers(s.Srv)
}

func (s *Server) Listen() error {
	err := s.Srv.Listen(s.conf.Http.Port)
	if err != nil {
		return err
	}
	return nil
}
