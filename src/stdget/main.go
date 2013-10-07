package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
)

func FetchAndCheck(url string) bool {
  res, err := http.Get(url)
  if err != nil {
    return false
  }
  if (res.StatusCode == 200){
    return true
  }
  return false
}

func Worker(finish chan bool, url string){
  for {
    select {
      case exit := <- finish:
        if(exit){
          fmt.Println("Quiting life!")
          return
        }
      default:
        if(FetchAndCheck(url)){
          fmt.Printf("Url:(%s) is reachable\n", url)
        }else{
          fmt.Printf("ALARM! Url:(%s) is unreachable\n", url)
        }
        fmt.Printf("Sleeping for 5 sec\n")
        time.Sleep(5 * time.Second)
    } // end select
  } // end for
}

func Worker2(finish chan bool, fn func() ){
  for {
    select {
      case exit := <- finish:
        if(exit){
          fmt.Println("Quiting life!")
          return
        }
      default:
        fn()
        fmt.Printf("Sleeping for 5 sec\n")
        time.Sleep(5 * time.Second)
    } // end select
  } // end for
}

func main() {

  // for i := 0; i < len(os.Args); i++{
  //   fmt.Printf("Arg -> %s\n", os.Args[i])
  // }

  ch := make(chan bool)
  url := os.Args[1]

  go Worker(ch, url)
  go Worker2(ch, func(){
    if(FetchAndCheck(url)){
      fmt.Printf("Url:(%s) is reachable\n", url)
    }else{
      fmt.Printf("ALARM! Url:(%s) is unreachable\n", url)
    }
  })

  var str string

  time.Sleep(120 * time.Second)
  fmt.Println("Sending kill signal to worker")
  ch <- true
  fmt.Scanf("Hit enter to finish\n", str)
}