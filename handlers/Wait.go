package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Wait(w http.ResponseWriter, r *http.Request){
	delay_str, param := r.URL.Query()["delay"]
	print(param)
	if param {
		delay_num, err := strconv.Atoi(delay_str[0])
		if err == nil{
			time.Sleep(time.Duration(delay_num) * time.Millisecond)
		} else {
			fmt.Fprint(w, "Uncorrected value DELAY")
		}

	} else {
		fmt.Fprint(w, "Uncorrected get request")
	}


}
