# user

This application uses `gRPC` and `proto` to send user data to mongoDB
This is a very basic application that uses CRUD functions, 

- Install/update dependencies `go get -u`
- Initiate gRPC: `protoc --proto_path=proto user/v1/user.proto --go_out=. --go-grpc_out=.`
- Run the application `go run .`
