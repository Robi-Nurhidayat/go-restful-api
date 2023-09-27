package exception

import (
	"mymodule/helper"
	"mymodule/model/web"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	internalServerError(w,r,err)
}


func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusInternalServerError)

	WebResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}

	helper.WriteToResponseBody(w,WebResponse)
}