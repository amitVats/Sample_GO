package main

import(
    "log"
    "net/http"
    "sample.com/handlers"
    "os"
    "github.com/gorilla/mux" 
    "time"
    "fmt"
)


func main(){
	l := log.New(os.Stdout,"power-api",log.LstdFlags)
	ph := handlers.NewPower(l)

    myRouter := mux.NewRouter()
    myRouter.Handle("/getPower/{name}", ph)

    //set up server
    server := &http.Server{
        Addr : ":9090",
        Handler: myRouter,
        ReadTimeout : 50 * time.Second
        WriteTimeout : 50 * time.Second
    }

    err := server.ListenAndServe()

    if(err != nil){
          fmt.Println(err)
    }

}
