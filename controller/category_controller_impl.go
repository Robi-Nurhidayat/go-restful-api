package controller

import (
	"mymodule/helper"
	"mymodule/model/web"
	"mymodule/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryReponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	webResponses := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryReponse,
	}

	// writer.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(writer)
	// err := encoder.Encode(webResponses)
	// helper.PanicIfError(err)

	helper.WriteToResponseBody(writer, webResponses)

}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryReponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	webResponses := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryReponse,
	}

	helper.WriteToResponseBody(writer, webResponses)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	webResponses := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponses)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	webResponses := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponses)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controller.CategoryService.FindAll(request.Context())

	webResponses := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponses)
}
