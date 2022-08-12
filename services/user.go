package services

import (
	"context"
	"fmt"
	"time"
	"github.com/felipehirano/go-gRPC/pb"
	"io"
	"log"
)

// Interfaces a serem implementadas: 

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// 	AddUserVerboose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerbooseClient, error)
// 	AddUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_AddUsersClient, error)
// }

type UserService struct {
	// Se adicionar um servico que nao existe no protocol buffers, ele não vai dar problema.
	pb.UnimplementedUserServiceServer
}

func NewUserService () *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Insert - Database	
	fmt.Println(req.Name);

	return &pb.User{
		Id:   "123",
		Name: req.GetName(),
		Email:  req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerboose(req *pb.User, stream pb.UserService_AddUserVerbooseServer) error {
	
	stream.Send(&pb.UserResultSteam{
		Status: "Init",
		User: &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultSteam{
		Status: "Inserting",
		User: &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultSteam{
		Status: "User has been inserted",
		User: &pb.User{
			Id:   "123",
			Name: req.GetName(),
			Email:  req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultSteam{
		Status: "Completed",
		User: &pb.User{
			Id:   "123",
			Name: req.GetName(),
			Email:  req.GetEmail(),
		},
	})

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	
	users := []*pb.User{}
	
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
		}
		users = append(users, &pb.User{
			Id:   "123",
			Name: req.GetName(),
			Email:  req.GetEmail(),
		})
		fmt.Println("Adding", req.GetName());
	}
}