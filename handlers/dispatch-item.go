package handlers

import (
	"net/http"
	"strconv"

	"gitlab.sudovi.me/erp/inventory-api/dto"
	"gitlab.sudovi.me/erp/inventory-api/errors"
	"gitlab.sudovi.me/erp/inventory-api/services"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

// DispatchItemHandler is a concrete type that implements DispatchItemHandler
type dispatchitemHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.DispatchItemService
}

// NewDispatchItemHandler initializes a new DispatchItemHandler with its dependencies
func NewDispatchItemHandler(app *celeritas.Celeritas, dispatchitemService services.DispatchItemService) DispatchItemHandler {
	return &dispatchitemHandlerImpl{
		App:     app,
		service: dispatchitemService,
	}
}

func (h *dispatchitemHandlerImpl) CreateDispatchItem(w http.ResponseWriter, r *http.Request) {
	var input dto.DispatchItemDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateDispatchItem(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "DispatchItem created successfuly", res)
}

func (h *dispatchitemHandlerImpl) UpdateDispatchItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.DispatchItemDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateDispatchItem(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "DispatchItem updated successfuly", res)
}

func (h *dispatchitemHandlerImpl) DeleteDispatchItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteDispatchItem(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "DispatchItem deleted successfuly")
}

func (h *dispatchitemHandlerImpl) GetDispatchItemListByItemId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetDispatchItemList(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *dispatchitemHandlerImpl) GetDispatchItemListByStatus(w http.ResponseWriter, r *http.Request) {
	var input dto.DispatchItemStatus
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.GetDispatchItemListbyStatus(input.Type, input.DispatchID)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "DispatchItem fetch successfuly", res)
}
