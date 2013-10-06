package main

/*
  http://www.gorillatoolkit.org/pkg/mux
  is pretty much over fucking powered :OOOO
*/

import (
  "fmt"
  "github.com/gorilla/pat"
  "net/http"
)

type WebHandler func(router *pat.Router, w http.ResponseWriter, r *http.Request)

func HandlerAwareOfRouter(router *pat.Router,  foo WebHandler) http.HandlerFunc  {
  return (func (w http.ResponseWriter, req *http.Request){
    foo(router, w, req)
  })
}

func HomeHandler(router *pat.Router, w http.ResponseWriter, req *http.Request){
  url, _ := router.GetRoute("service").URL("service_name", "trolling-and-molling")
  fmt.Fprintf(w, "Home handler! %s", url)
}

func ServicesHandler(router *pat.Router, w http.ResponseWriter, req *http.Request){
  url, _ := router.GetRoute("services").URL()
  fmt.Fprintf(w, "Services List Page! Go -> %s", url)
}

func ServiceHandler(router *pat.Router, w http.ResponseWriter, req *http.Request){
  service_name := req.URL.Query().Get(":service_name")
  fmt.Fprintf(w, "Services handler! \n You selected %s", service_name)
}

func main(){

    router := pat.New()
    router.Get("/services/{service_name}", HandlerAwareOfRouter(router,ServiceHandler)).Name("service")
    router.Get("/services", HandlerAwareOfRouter(router,ServicesHandler)).Name("services")
    router.Get("/", HandlerAwareOfRouter(router,HomeHandler)).Name("home")
    http.Handle("/", router)

    http.ListenAndServe(":8080", nil)

}
