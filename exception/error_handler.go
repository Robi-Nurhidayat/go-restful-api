package exception

import (
	"mymodule/helper"
	"mymodule/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if notFoundError(w,r,err) {
		return
	}

	if validationError(w,r,err) {
		return
	}

	internalServerError(w,r,err)
}


func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {

	notFoundStatus,ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusNotFound)
	
		WebResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "NOT FOUND",
			Data: notFoundStatus.Error,
		}
	
		helper.WriteToResponseBody(w,WebResponse)
		return true
	}else {
		return false
	}
	
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	badReq,ok := err.(validator.ValidationErrors)

	if ok {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusBadRequest)
	
		WebResponse := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: badReq.Error(),
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