package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func Date(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Current date and time: ", time.Now())
}
