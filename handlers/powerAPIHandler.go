package handlers

import(
    "log"
    "net/http"
    "sample.com/data"
    "sample.com/datasource"
    "github.com/gorilla/mux"
)

type Power struct{
	l *log.Logger
}

func NewPower(l *log.Logger) *Power{
	return &Power{l}
}

func (p *Power)  ServeHTTP(rw  http.ResponseWriter, r *http.Request){
	
	 if(r.Method == http.MethodGet){
	 	params := mux.Vars(r)
	 	name := params["name"]

	 	max_power := datasource.Get(name)
	 	if(max_power == 0){ // will add custom message
	 		return
	 		}else{
	 			sr := data.FmtPowerRes(name,max_power)
	 			err := sr.ToJSON(rw)
	 			if(err != nil ){
	 			http.Error(rw,"Internal Server Error",http.StatusInternalServerError)
	 		}
	 	
	 	}
	 }else{
	 	http.Error(rw,"The request type is not supported", http.StatusBadRequest)
	 }
	 

}
