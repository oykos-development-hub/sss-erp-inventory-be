package handlers

import (
	"context"
	"net/http"
	"strconv"

	"gitlab.sudovi.me/erp/inventory-api/contextutil"
	"gitlab.sudovi.me/erp/inventory-api/dto"
	"gitlab.sudovi.me/erp/inventory-api/errors"
	"gitlab.sudovi.me/erp/inventory-api/services"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

// DispatchHandler is a concrete type that implements DispatchHandler
type dispatchHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.DispatchService
}

// NewDispatchHandler initializes a new DispatchHandler with its dependencies
func NewDispatchHandler(app *celeritas.Celeritas, dispatchService services.DispatchService) DispatchHandler {
	return &dispatchHandlerImpl{
		App:     app,
		service: dispatchService,
	}
}

func (h *dispatchHandlerImpl) CreateDispatch(w http.ResponseWriter, r *http.Request) {
	var input dto.DispatchDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	userIDString := r.Header.Get("UserID")

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized)
		return
	}

	ctx := context.Background()
	ctx = contextutil.SetUserIDInContext(ctx, userID)

	res, err := h.service.CreateDispatch(ctx, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Dispatch created successfuly", res)
}

func (h *dispatchHandlerImpl) UpdateDispatch(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.DispatchDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	userIDString := r.Header.Get("UserID")

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized)
		return
	}

	ctx := context.Background()
	ctx = contextutil.SetUserIDInContext(ctx, userID)

	res, err := h.service.UpdateDispatch(ctx, id, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Dispatch updated successfuly", res)
}

func (h *dispatchHandlerImpl) DeleteDispatch(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	userIDString := r.Header.Get("UserID")

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized)
		return
	}

	ctx := context.Background()
	ctx = contextutil.SetUserIDInContext(ctx, userID)

	err = h.service.DeleteDispatch(ctx, id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Dispatch deleted successfuly")
}

func (h *dispatchHandlerImpl) GetDispatchById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetDispatch(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *dispatchHandlerImpl) GetDispatchList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetDispatchListInput
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	res, total, err := h.service.GetDispatchList(&input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}
