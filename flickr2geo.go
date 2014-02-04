package main

import (
    "bytes"
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
     if !bytes.Contains(html, []byte("flickr_photos:location:latitude")) {
        fmt.Println("no latitude information found")
        os.Exit(11)
     }
     
     if !bytes.Contains(html, []byte("flickr_photos:location:longitude")) {
        fmt.Println("no longitude information found")
        os.Exit(12)
     }
     datachunks := bytes.Split(html, []byte("flickr_photos:location"))
     lat := bytes.Split(datachunks[1], []byte("\""))[2]
     lon := bytes.Split(datachunks[2], []byte("\""))[2]
     fmt.Printf("latitude = %s\n", lat)
	 fmt.Printf("longitude = %s\n", lon)
     fmt.Printf("for awk: %s %s\n", lat, lon)
}
