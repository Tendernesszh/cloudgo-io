package service

import (
	"net/http"
	"github.com/unrolled/render"
)

type data struct {
	ID  string
	Password string
}

func recordHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		  r.ParseForm();
			dataput := data{
			    ID: r.Form["id"][0],
					Password: r.Form["password"][0],
			}
			formatter.HTML(w,http.StatusOK,"userinfo", dataput)
		}
}
