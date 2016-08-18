// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	// "flag"
	"log"
	"net/http"
	// "text/template"

	// "github.com/gorilla/websocket"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

)

// var addr = flag.String("addr", ":8080", "http service address")
// var homeTemplate = template.Must(template.ParseFiles("home.html"))

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL)
// 	if r.URL.Path != "/" {
// 		http.Error(w, "Not found", 404)
// 		return
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "Method not allowed", 405)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	homeTemplate.Execute(w, r.Host)
// }


// var hub = newHub()

// func hello() http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		conn, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
// 		client.hub.register <- client
// 		go client.writePump()
// 		client.readPump()

// 	}
// }


// func hello() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		serveWs(hub, w, r)
// 	}
// }


func wsHub() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			hub *Hub
			err error
		)

		hubId := r.URL.Query().Get("hubId")

		if (hubId == "") {
			hubId = "hub000"
		}

		log.Println("wsHub,  hub id:")
		log.Println(hubId)
		log.Println("end wsHub.")

		hub = getHub(hubId)

		if (hub == nil) {
			hub, err = initializeHub(hubId)

			if err != nil {
				log.Println(err)
				return
			}

			if (hub == nil) {

				log.Println("hub nil")
				return
			}

		}

		serveWs(hub, w, r)
	}
}


func test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		hubId := r.URL.Query().Get("hubId")

		log.Println("test,  hub id:")
		log.Println(hubId)
		log.Println("end wsHub.")
	}
}


func test2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")

		log.Println("test2,  hub id:")
		log.Println(id)
		log.Println("end wsHub.")
	}
}


func main() {
	// flag.Parse()
	// hub := newHub()
	// go hub.run()
	// http.HandleFunc("/", serveHome)
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	serveWs(hub, w, r)
	// })
	// err := http.ListenAndServe(*addr, nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }


	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("./public"))
	//
	// e.GET("/ws", standard.WrapHandler(http.HandlerFunc(hello())))
	// e.GET("/ws/:hubId", standard.WrapHandler(http.HandlerFunc(wsHub())))
	e.GET("/ws", standard.WrapHandler(http.HandlerFunc(wsHub())))

	e.GET("/test", standard.WrapHandler(http.HandlerFunc(test())))
	e.GET("/test2/:id", standard.WrapHandler(http.HandlerFunc(test2())))


	e.Run(standard.New(":1323"))

}









