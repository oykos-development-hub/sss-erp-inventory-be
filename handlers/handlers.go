package handlers

import (
	"net/http"
)

type Handlers struct {
	RealEstateHandler   RealEstateHandler
	ItemHandler         ItemHandler
	AssessmentHandler   AssessmentHandler
	DispatchHandler     DispatchHandler
	DispatchItemHandler DispatchItemHandler
}

type RealEstateHandler interface {
	CreateRealEstate(w http.ResponseWriter, r *http.Request)
	UpdateRealEstate(w http.ResponseWriter, r *http.Request)
	DeleteRealEstate(w http.ResponseWriter, r *http.Request)
	GetRealEstateById(w http.ResponseWriter, r *http.Request)
	GetRealEstatebyItemId(w http.ResponseWriter, r *http.Request)
	GetRealEstateList(w http.ResponseWriter, r *http.Request)
}

type ItemHandler interface {
	CreateItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
	GetItemById(w http.ResponseWriter, r *http.Request)
	GetItemList(w http.ResponseWriter, r *http.Request)
	GetItemListInOrganizationUnit(w http.ResponseWriter, r *http.Request)
}

type AssessmentHandler interface {
	CreateAssessment(w http.ResponseWriter, r *http.Request)
	UpdateAssessment(w http.ResponseWriter, r *http.Request)
	DeleteAssessment(w http.ResponseWriter, r *http.Request)
	GetAssessmentById(w http.ResponseWriter, r *http.Request)
	GetAssessmentList(w http.ResponseWriter, r *http.Request)
	GetAssessmentbyItemId(w http.ResponseWriter, r *http.Request)
}

type DispatchHandler interface {
	CreateDispatch(w http.ResponseWriter, r *http.Request)
	UpdateDispatch(w http.ResponseWriter, r *http.Request)
	DeleteDispatch(w http.ResponseWriter, r *http.Request)
	GetDispatchById(w http.ResponseWriter, r *http.Request)
	GetDispatchList(w http.ResponseWriter, r *http.Request)
}

type DispatchItemHandler interface {
	CreateDispatchItem(w http.ResponseWriter, r *http.Request)
	UpdateDispatchItem(w http.ResponseWriter, r *http.Request)
	DeleteDispatchItem(w http.ResponseWriter, r *http.Request)
	GetDispatchItemListByItemId(w http.ResponseWriter, r *http.Request)
	GetDispatchItemListByStatus(w http.ResponseWriter, r *http.Request)
}
