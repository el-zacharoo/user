# User gRPC microserverice #

This application uses `gRPC` and `proto` to send user data to a database (mongoDB). <br/>
This uses CRUD functions. <br />
For now I haven't added any securtity features, or protected my db connection as this is supposed to only showcase the very basics of CRUD functions in gRPC. 

## Get Started ##
- Install/update dependencies `go get -u`
- Initiate gRPC: `protoc --proto_path=proto user/v1/user.proto --go_out=. --go-grpc_out=.` This is the standard way of generating gRPC, however with this app I am using a makefile to manage my commands. So instead run `make proto`, this will also send the proto definitions to `https://buf.build` as a commit. 
- Run the application `go run .` or `make run`

## Set up Docker ## 
* To make a build on `docker` run `docker build --tag user .`
* To run on docker `docker run -d -p 8080:8080 user`

## Postman Setup ##
- Go to New, select gRPC Request
- Call `localhost:8080`
- Select the proto definitions file to put into the methods input
- Make sure server reflection is enabled in methods
- You can generate example JSON Messages (this is particularly useful when using the CREATE method, make sure the ID field isn't present as this is being     filled in automatically when an entry is created) 
- When making a GET Request make sure the JSON message only includes existing ID's <br />
 
```bash 
{ "id": "4942b3e9-3699-412a-9f16-01503113178f"}
```
