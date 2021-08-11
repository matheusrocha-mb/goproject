package server

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/matheusrocha-mb/goproject/internal/database"
	"google.golang.org/grpc"
)

const SERVER_LOCAL_PORT = 9002
const SERVER_CONTAINER_PORT = 9001

type Server struct {
	db    database.DatabaseRepo
	grpc  *grpc.Server
	gport string
}

func Run(gport string) {

	fmt.Printf("\nStarting server local on port: %d\nStarting server on container at port: %d\n", SERVER_LOCAL_PORT, SERVER_CONTAINER_PORT)
}

func grpcServer() {

	port:= os.Getenv("GRPC_SERVER_PORT")

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Info("Listener error: %v", err)
	}

	grpcsrv := grpc.NewServer()
	
	log.Info("gRPC server started on port: %v", port)
	if err := grpcsrv.Serve(listener); err != nil {
		log.Info("gRPC error: %v", err)
	}
}
