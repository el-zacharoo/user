package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/el-zacharoo/user/handler"
	pbcnn "github.com/el-zacharoo/user/internal/gen/user/v1/userv1connect"
	"github.com/el-zacharoo/user/store"
	"github.com/rs/cors"
)

const port = "localhost:8081"

func main() {
	s := store.Connect()
	svc := &handler.UserServer{
		Store: s,
	}

	c := setCORS()
	mux := http.NewServeMux()
	path, h := pbcnn.NewUserServiceHandler(svc)
	mux.Handle(path, h)
	hndlr := c.Handler(mux)

	http.ListenAndServe(
		port,
		h2c.NewHandler(hndlr, &http2.Server{}),
	)
}

func setCORS() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type"},
		AllowedMethods:   []string{"POST"},
		AllowCredentials: true,
	})
	return c
}
