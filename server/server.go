package main

import (
	"errors"
	"fmt"
	"net"

	"go-grpc-example/config"
	"go-grpc-example/mockdata"
	"go-grpc-example/proto"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var log *zerolog.Logger

type server struct{}

func main() {

	// loads the config file
	config.LoadConfig()

	// port is initialized
	port := viper.GetString("server.port")

	// listener is created
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Error().Msgf("error in listening port "+port, err)
	}
	// for debugging
	fmt.Println("starting user management service server")

	// mockdata is generated
	mockdata.GenerateData()

	// new grpc server is created
	srv := grpc.NewServer()

	// new server is registered
	proto.RegisterUserManagementServiceServer(srv, &server{})
	reflection.Register(srv)

	// server starts listening
	if err := srv.Serve(listener); err != nil {
		log.Error().AnErr("unable to start server : ", err)
	}

}

// GetSingleUser - function fetches information about a single user
func (s *server) GetSingleUser(ctx context.Context, request *proto.RequestUser) (*proto.ResponseUser, error) {

	// variable decalarations
	var id []int
	response := &proto.ResponseUser{}

	// invalid input check
	if request.UserId == 0 {
		return &proto.ResponseUser{}, errors.New("INVALID USER ID")
	}

	// id is appended
	id = append(id, int(request.UserId))

	// GetData function is called with id
	userData, err := mockdata.GetData(id)
	if err != nil {
		return &proto.ResponseUser{}, err
	}

	// proto is populated
	for _, v := range userData {

		response.Id = int64(v.ID)
		response.City = v.City
		response.Fname = v.Fname
		response.Height = float32(v.Height)
		response.Married = v.Married
		response.Phone = v.Phone
	}

	// data is returned in proto format
	return response, nil

}

func (s *server) GetMultipleUsers(ctx context.Context, request *proto.RequestUsers) (*proto.ResponseUsers, error) {

	// variable decalarations
	var id []int
	var usersList []*proto.ResponseUser

	// invalid input check
	if len(request.UserId) == 0 {
		return &proto.ResponseUsers{}, status.Errorf(codes.InvalidArgument, "USER ID NOT PROVIDED")
	} 
	for _,v := range request.UserId {
		if v == 0 {
			return &proto.ResponseUsers{},status.Errorf(codes.InvalidArgument,"INVALID USER ID")
		}
       id = append(id,int(v))
		}
	

	// GetData function is called with id
	userData, err := mockdata.GetData(id)
	if err != nil {
		return &proto.ResponseUsers{}, err
	}

	// proto is populated
	for _, v := range userData {

		usersList = append(usersList,&proto.ResponseUser{
			Id: int64(v.ID),
			City: v.City,
			Fname: v.Fname,
			Phone: v.Phone,
			Married: v.Married,
			Height: float32(v.Height),
		})
	}

	// data is returned in proto format
	return &proto.ResponseUsers{Users: usersList}, nil
}
