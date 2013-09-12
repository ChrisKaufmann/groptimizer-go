package main

import (
	"github.com/msbranco/goconfig"
	"net/http"
	"fmt"
)

var (
	port  string
	myurl string
)

func init() {
	var err error
	c, err := goconfig.ReadConfigFile("config")
	if err != nil {
		err.Error()
	}
	port, err = c.GetString("Web", "port")
	if err != nil {
		err.Error()
	}
	myurl, err = c.GetString("Web", "url")
	if err != nil {
		err.Error()
	}
}

func main() {
	http.HandleFunc("/", handleRoot)
	print("Listening on 127.0.0.1:" + port + "\n")
	http.ListenAndServe("127.0.0.1:"+port, nil)
}
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "O Hai")
}
