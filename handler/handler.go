package handler

import (
	"net/http"
	"time"
	"web_adv/utils"
)

func RandomRequest(w http.ResponseWriter, r *http.Request) {
	t := utils.GetRandomTimeRequest()
	time.Sleep(time.Duration(t * int(time.Second)))

	if t%2 == 0 {
		utils.ResponseDone(w, []byte("success"))
	} else {
		utils.ResponseErr(w, "err", http.StatusBadRequest)
	}
}
