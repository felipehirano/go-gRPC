## Projeto com o intuito de estudo sobre gRPC utilizando golang

### Protocol buffers

- Foi criado primeiramente uma pasta chamada /proto que será responsável por guardar os arquivos .proto
  para se trabalhar com o protocol buffers;
- Os arquivos .proto servem para gerar uma mensagem para cada estrutura de dados que será serializado e,
  para isso, é necessário especificar o nome e o tipo de cada mensagem;
- São eles que vão gerar as stubs para se trabalhar com o golang;
- Criar uma pasta chamada /pb que é ela que vai armazenar os arquivos stubs que serão gerados;
- O comando para gerar as stubs é (Sempre que alterar um .proto, é necessário rodar esse comando):
  `'protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.'`;
- Para mais detalhes, basta olhar a documentacao: https://developers.google.com/protocol-buffers/docs/gotutorial;

### Servidor

- Criar uma pasta chamada /server que irá armazenar o arquivo responsável por ser o nosso servidor;
- Assim que foi criado o o arquivo server.go, basta rodar o comando:
  `'go run server.go'`

### Clients

- Para rodar o client (evans) instale o evans e execute ele no terminal com o seguinte comando:
  `'evans [NAME_FILE].proto ou o evans -r repl --host localhost --port 50051'`
- Documentacao do evans: https://github.com/ktr0731/evans
- Temos os clients também em cmd/clients que pode chamar os servicos disponibilizados rodando o comando
  `'go run client.go'`;

### Servicos

- Os servicos são criados a partir das assinaturas que estão nos arquivos stubs gerados pelo protoc;
- Os servicos estão dentro da pasta /server, e para esse projeto, temos somente o user.go;
