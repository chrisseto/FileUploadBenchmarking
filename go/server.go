package main

import (
       "log"
       "net/http"
)

func UploadHandler(w http.ResponseWriter, req *http.Request) {
     switch req.Method {
            case "POST":
                 err2 := req.ParseMultipartForm(10000000)
                 if err2 != nil {
                     return
                 }
                 m := req.MultipartForm
                 files := m.File["data"]
                 log.Print(files[0].Filename)

            default:
                log.Print("Nope")
     }
}

func main() {
     http.HandleFunc("/", UploadHandler)

     err := http.ListenAndServe(":5555", nil)
     if err != nil {
        log.Fatal("Error: ", err)
     }
}
