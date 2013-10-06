package main

import (
  "fmt"
  "github.com/gorilla/http"
  "os"
  "bytes"
)

func main(){

  fmt.Println("Getching google page")
  if _, err := http.Get(os.Stdout, "http://www.google.pl/"); err != nil {
    fmt.Printf("could not fetch: %v", err)
  }

  // Lets read it to memory! ha biatches
  fmt.Printf("\n\nNow to buffer\n\n")

  var buf *bytes.Buffer = new(bytes.Buffer)
  if _, err := http.Get(buf, "http://www.google.pl"); err != nil {
    fmt.Printf("could not fetch: %v", err)
  }else{
    fmt.Printf("%s", buf.String())
  }

}