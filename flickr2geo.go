/*
    A tiny program to extract the latitude and longitude information from flickr photo sites
    Copyright (C) 2013  tpltnt

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
