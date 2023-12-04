package function

import (
	"fmt"

	"handler/function/static/ausf"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ue-authentications" {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Hello World!")
		fmt.Println(ausf.StartAUSF())
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Invalid request type")
	}
}
