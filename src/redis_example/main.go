package main

import (
  "fmt"
  "github.com/garyburd/redigo/redis"
)


func main(){
  fmt.Println("Example of using Redis in Go!")

    c, err := redis.Dial("tcp", ":6379")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer c.Close()


    c.Send("SET", "foo", "bar")
    c.Send("GET", "foo")
    c.Flush()
    c.Receive() // reply from SET
    v, _ := c.Receive() // reply from GET

    fmt.Printf("Value of 'foo' is = %s \n", v)

}