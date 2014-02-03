package main

import (
	"fmt"
	"io/ioutil"
	"log"
    "os"
	"net/http"
)

func main() {
     if 2 != len(os.Args) {
        fmt.Println("usage: " + os.Args[0] + " flickr-url")
        os.Exit(1)
     }
	 res, err := http.Get(os.Args[1])
	 if err != nil {
	 	log.Fatal(err)
	 }
	 html, err := ioutil.ReadAll(res.Body)
	 res.Body.Close()
	 if err != nil {
		log.Fatal(err)
	 }
	 fmt.Printf("%s\n", html)
}
