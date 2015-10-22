package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type Request struct {
	Name string
}

type Response struct {
	Greeting string
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func helloPost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    var request Request
    decoder := json.NewDecoder(req.Body)
	error := decoder.Decode(&request)
	if error != nil {
		panic("Error with JSON")
	}
	postResp := Response{Greeting: "Hello, " + request.Name + "!"}
	json.NewEncoder(rw).Encode(postResp) 
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello/", helloPost)
    
    server := http.Server{
            Addr: "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}