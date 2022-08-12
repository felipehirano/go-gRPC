package main 

import (
	"github.com/felipehirano/go-gRPC/pb"
	"google.golang.org/grpc"
	"context"
	"log"
	"io"
	"fmt"
	"time"
)


func main() {
	// Cria uma conexão gRPC
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	// Encerra a conexão gRPC quando ninguém estiver utilizando esse connection; 
	defer connection.Close()

	user := pb.NewUserServiceClient(connection)

	// Inicia comunicação com o servidor;
	// AddUser(user)

	// Inicia comunicacao via stream a partir do server;
	// AddUserVerboose(user)

	// Inicia comunicacao via stream a partir do client;
	// AddUsers(user)

	// Inicia comunicacao bidirecional;
	AddUserStreamBoth(user)
}

func AddUser(user pb.UserServiceClient) {
	
	req := &pb.User{
		Id: "0",
		Name: "Thiago",
		Email: "thiago@gmail.com",
	}

	res, err := user.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("failed to add user: %v", err)
	}

	log.Println("User added:", res)
}

func AddUserVerboose(user pb.UserServiceClient) {
	req := &pb.User{
		Id: "0",
		Name: "Thiago",
		Email: "thiago@gmail.com",
	}

	responseStream, err := user.AddUserVerboose(context.Background(), req)

	if err != nil {
		log.Fatalf("failed to add user: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive stream: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(user pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id: "f1",
			Name: "Felipe1",
			Email: "fkenhirano4@gmail.com1",
		},
		&pb.User{
			Id: "f2",
			Name: "Felipe2",
			Email: "fkenhirano4@gmail.com2",
		},
		&pb.User{
			Id: "f3",
			Name: "Felipe3",
			Email: "fkenhirano4@gmail.com3",
		},
		&pb.User{
			Id: "f4",
			Name: "Felipe4",
			Email: "fkenhirano4@gmail.com4",
		},
		&pb.User{
			Id: "f5",
			Name: "Felipe5",
			Email: "fkenhirano4@gmail.com5",
		},
	}

	// O Background controla o fluxo de mensagens, garantindo que se a mensagem não chegar ele já para as requisicoes;
	stream, err := user.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("failed to add users: %v", err)
	}


	// o _ significa o indice;
	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receiving response: %v", err)
	}

	fmt.Println("Users added:", res)
}

func AddUserStreamBoth(user pb.UserServiceClient) {
	stream, err := user.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("failed to creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id: "f1",
			Name: "Felipe1",
			Email: "fkenhirano4@gmail.com1",
		},
		&pb.User{
			Id: "f2",
			Name: "Felipe2",
			Email: "fkenhirano4@gmail.com2",
		},
		&pb.User{
			Id: "f3",
			Name: "Felipe3",
			Email: "fkenhirano4@gmail.com3",
		},
		&pb.User{
			Id: "f4",
			Name: "Felipe4",
			Email: "fkenhirano4@gmail.com4",
		},
		&pb.User{
			Id: "f5",
			Name: "Felipe5",
			Email: "fkenhirano4@gmail.com5",
		},
	}

	// Para ficar enviando e recebendo dados ao mesmo tempo, é necessário um assincronismo;
	// O go utiliza go rountines como threads para criar esse assincronismo;
	// Então criamos uma go rountine para enviar e outra para receber as requisicoes;

	// go rountine para enviar as requisições;
	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user:", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	// Como utilizamos go rountine para crair o assincronismo, assim que cada go rountine termina, o programa também termina;
	// para tratar isso, é necessário criar um channel para ter o controle do encerramento do programa; 
	wait := make(chan struct{})
	
	// go rountine para receber as requisições;
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("failed to receive stream: %v", err)
				break
			}
			fmt.Printf("Recebendo user %v com status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()
	
	// Enquanto esse channel estiver rodando, o client não vai morrer;
	<-wait

}