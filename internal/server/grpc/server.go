package grpc

import (
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"todoApp/pkg/service"
)

type Server struct {
	services *service.Service
	server   *grpc.Server
}

func NewServer(services *service.Service, logger *logrus.Logger) *Server {
	logr := logrus.NewEntry(logger)
	return &Server{
		services: services,
		server: grpc.NewServer(
			grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
				grpcLogrus.StreamServerInterceptor(logr),
				grpcRecovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
				grpcLogrus.UnaryServerInterceptor(logr),
				grpcRecovery.UnaryServerInterceptor(),
			)),
		),
	}
}

func (s *Server) Run(port string) error {
	addr := fmt.Sprintf(":%s", port)

	RegisterServiceServer(s.server, s)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return s.server.Serve(l)
}

func (s *Server) ShutDown() {
	s.server.GracefulStop()
}

func (s *Server) mustEmbedUnimplementedServiceServer() {}
