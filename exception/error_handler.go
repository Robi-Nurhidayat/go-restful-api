package exception

import (
	"mymodule/helper"
	"mymodule/model/web"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if notFoundError(w,r,err) {
		return
	}

	internalServerError(w,r,err)
}


func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {

	exception,ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusNotFound)
	
		WebResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "NOT FOUND",
			Data: exception.Error,
		}
	
		helper.WriteToResponseBody(w,WebResponse)
		return true
	}else {
		return false
	}
	
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