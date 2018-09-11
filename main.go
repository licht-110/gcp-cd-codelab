package main

import (
  "io"
  "net/http"
  "log"
  "os"
)

var elog *log.Logger
var ilog *log.Logger
func init(){
  elog = log.New(os.Stdout, "[ERROR]", log.LstdFlags|log.LUTC)
  ilog = log.New(os.Stdout, "[INFO]", log.LstdFlags|log.LUTC)
}

func hello(w http.ResponseWriter, r *http.Request) {
  ilog.Print("Hello, world!")
  io.WriteString(w, "Hello World!")
}

func fail(w http.ResponseWriter, r *http.Request) {
  elog.Print("Error!")
  w.WriteHeader(500)
  io.WriteString(w, "Error!")
}

func main() {
  http.HandleFunc("/fail", fail)
  http.HandleFunc("/", hello)
  http.ListenAndServe(":80", nil)
}
