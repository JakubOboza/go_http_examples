package main

/*
  http://www.gorillatoolkit.org/pkg/mux
  is pretty much over fucking powered :OOOO
*/

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

type GoHandleFuncHandler func(w http.ResponseWriter, r *http.Request)
type WebHandler func(router *mux.Router, w http.ResponseWriter, r *http.Request)

func HandlerAwareOfRouter(router *mux.Router,  foo WebHandler) GoHandleFuncHandler {
  return (func (w http.ResponseWriter, req *http.Request){
    foo(router, w, req)
  })
}

func HomeHandler(router *mux.Router, w http.ResponseWriter, req *http.Request){
  url, _ := router.Get("service").URL("service_name", "trolling-and-molling")
  fmt.Fprintf(w, "Home handler! %s", url)
}

func ServicesHandler(router *mux.Router, w http.ResponseWriter, req *http.Request){
  url, _ := router.Get("services").URL()
  fmt.Fprintf(w, "Services List Page! Go -> %s", url)
}

func ServiceHandler(router *mux.Router, w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  service_name := params["service_name"]
  fmt.Fprintf(w, "Services handler! \n You selected %s", service_name)
}

func main(){

    router := mux.NewRouter()
    router.HandleFunc("/", HandlerAwareOfRouter(router, HomeHandler) )
    router.HandleFunc("/services", HandlerAwareOfRouter(router,ServicesHandler)).Name("services")
    router.HandleFunc("/services/{service_name}", HandlerAwareOfRouter(router,ServiceHandler)).Name("service")
    http.Handle("/", router)

    http.ListenAndServe(":8080", nil)

}
