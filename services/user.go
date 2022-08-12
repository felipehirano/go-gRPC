package services

import "github.com/felipehirano/go-gRPC/pb"

// Interface a ser implementada: 

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
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
		id:   "123"
		name: req.GetName(),
		email:  req.GetEmail(),
	}, nil
}