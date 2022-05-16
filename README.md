# user

This application uses `gRPC` and `proto` to send user data to mongoDB

Initiate gRPC: `protoc --proto_path=proto user/v1/user.proto --go_out=. --go-grpc_out=.`
