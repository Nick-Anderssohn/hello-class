// Copyright (C) 2017  Nicholas Anderssohn
package main

import (
	"context"
	"fmt"
	"hello-compsci/frontend-server/director"
	"hello-compsci/frontend-server/server"
	"net/http"
)

func main() {
	endpoints := []*server.Endpoint{
		{
			Path:        "/",
			HandlerFunc: http.FileServer(http.Dir("client/build/web")),
		},
	}

	server := director.NewHelloClassServer(context.Background(), endpoints, "0.0.0.0", "8080")
	if err := server.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
