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

// AssessmentHandler is a concrete type that implements AssessmentHandler
type assessmentHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.AssessmentService
}

// NewAssessmentHandler initializes a new AssessmentHandler with its dependencies
func NewAssessmentHandler(app *celeritas.Celeritas, assessmentService services.AssessmentService) AssessmentHandler {
	return &assessmentHandlerImpl{
		App:     app,
		service: assessmentService,
	}
}

func (h *assessmentHandlerImpl) CreateAssessment(w http.ResponseWriter, r *http.Request) {
	var input dto.AssessmentDTO
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

	res, err := h.service.CreateAssessment(ctx, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Assessment created successfuly", res)
}

func (h *assessmentHandlerImpl) UpdateAssessment(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.AssessmentDTO
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

	res, err := h.service.UpdateAssessment(ctx, id, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Assessment updated successfuly", res)
}

func (h *assessmentHandlerImpl) DeleteAssessment(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.DeleteAssessment(ctx, id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Assessment deleted successfuly")
}

func (h *assessmentHandlerImpl) GetAssessmentById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetAssessment(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *assessmentHandlerImpl) GetAssessmentList(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAssessmentList()
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *assessmentHandlerImpl) GetAssessmentbyItemId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, total, err := h.service.GetAssessmentbyItemId(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}
