package main

import (
	"fmt"
	"log"
	"net/http"
	"pcad/web"

	"github.com/nats-io/stan.go"
	"github.com/siuyin/dflt"
)

// 10 OMIT
func webStart(sc stan.Conn) {
	go func() {
		fmt.Println("web module starting up")

		rootHandler := &web.RootHandler{Msg: "Hello from Root"}
		http.Handle("/", rootHandler)

		orderHandler := &web.OrderHandler{Conn: sc, Subject: "orders"}
		http.Handle("/order", orderHandler)

		port := dflt.EnvString("WEB_PORT", ":8080")
		log.Fatal(http.ListenAndServe(port, nil))
	}()
}

// 20 OMIT
