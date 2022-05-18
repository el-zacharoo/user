# user

This application uses `gRPC` and `proto` to send user data to mongoDB. <br/>
This is a very basic application that uses CRUD functions. 

- Install/update dependencies `go get -u`
- Initiate gRPC: `protoc --proto_path=proto user/v1/user.proto --go_out=. --go-grpc_out=.`
- Run the application `go run .`

## Postman Setup 
- Go to New, select gRPC Request
- Call `localhost:8080`
- Select the proto definitions file
- Make sure server reflection is enabled in methods
- You can generate example JSON Messages (this is particularly useful when using the CREATE method, make sure the ID field isn't present as this is being     filled in automatically when an entry is created) 
- When making a GET Request make sure the JSON message only includes existing ID's <br />
 
```bash 
   { "id": "4942b3e9-3699-412a-9f16-01503113178f"}
```
