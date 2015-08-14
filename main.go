package hello

import (
//    "fmt"
    "net/http"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("./")))
}

