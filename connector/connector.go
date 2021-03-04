package connector

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Connect establishes connection between server and client
func Connect() *grpc.ClientConn {

	// connection to server port
	conn, err := grpc.Dial("localhost"+viper.GetString("client.port"), grpc.WithInsecure())
	if err != nil {
		log.Panic().AnErr("unable to connect to server from client : ", err)
	}

	// connection is returned
	return conn

}
