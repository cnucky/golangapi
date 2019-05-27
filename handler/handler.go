package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Todo struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Completed   bool    `json:"completed"`
	Due     time.Time   `json:"due"`
}

type Todos  []Todo

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request)  {
	todos := Todos{
		Todo{Name:"Write presentation"},
		Todo{Name:"Host meetup"},
		Todo{Id:22,Name:"Joel",Completed:true,Due:time.Now()},
	}
	if err := json.NewEncoder(w).Encode(todos); err != nil{
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

