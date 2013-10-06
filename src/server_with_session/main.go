package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/sessions"
    "github.com/gorilla/mux"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func MyHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    // Set some session values.
    session.Values["foo"] = "bar"
    session.Values[42] = 4

    // Set defaults for luls
    if(session.Values["iter"] == nil){
      session.Values["iter"] = 0
    }

    iter := session.Values["iter"].(int)
    session.Values["iter"] = iter + 1
    session.Save(r, w)

    fmt.Fprintf(w, "Session 'iter' value => %d", session.Values["iter"])
}

func main(){

    router := mux.NewRouter()
    router.HandleFunc("/", MyHandler)
    http.Handle("/", router)

    http.ListenAndServe(":8080", nil)

}