package main

import (
	"gitlab.sudovi.me/erp/inventory-api/handlers"
	"gitlab.sudovi.me/erp/inventory-api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

func routes(app *celeritas.Celeritas, middleware *middleware.Middleware, handlers *handlers.Handlers) *chi.Mux {
	// middleware must come before any routes

	//api
	app.Routes.Route("/api", func(rt chi.Router) {

		rt.Post("/real-estates", handlers.RealEstateHandler.CreateRealEstate)
		rt.Get("/real-estates/{id}", handlers.RealEstateHandler.GetRealEstateById)
		rt.Get("/real-estates", handlers.RealEstateHandler.GetRealEstateList)
		rt.Put("/real-estates/{id}", handlers.RealEstateHandler.UpdateRealEstate)
		rt.Delete("/real-estates/{id}", handlers.RealEstateHandler.DeleteRealEstate)
		rt.Get("/item/{id}/real-estates", handlers.RealEstateHandler.GetRealEstatebyItemId)

		rt.Post("/items", handlers.ItemHandler.CreateItem)
		rt.Get("/items/{id}", handlers.ItemHandler.GetItemById)
		rt.Get("/items", handlers.ItemHandler.GetItemList)
		rt.Put("/items/{id}", handlers.ItemHandler.UpdateItem)
		rt.Delete("/items/{id}", handlers.ItemHandler.DeleteItem)

		rt.Post("/assessments", handlers.AssessmentHandler.CreateAssessment)
		rt.Get("/assessments/{id}", handlers.AssessmentHandler.GetAssessmentById)
		rt.Get("/assessments", handlers.AssessmentHandler.GetAssessmentList)
		rt.Put("/assessments/{id}", handlers.AssessmentHandler.UpdateAssessment)
		rt.Delete("/assessments/{id}", handlers.AssessmentHandler.DeleteAssessment)
		rt.Get("/assessments/{id}/item", handlers.AssessmentHandler.GetAssessmentbyItemId)

		rt.Post("/dispatches", handlers.DispatchHandler.CreateDispatch)
		rt.Get("/dispatches/{id}", handlers.DispatchHandler.GetDispatchById)
		rt.Get("/dispatches", handlers.DispatchHandler.GetDispatchList)
		rt.Put("/dispatches/{id}", handlers.DispatchHandler.UpdateDispatch)
		rt.Delete("/dispatches/{id}", handlers.DispatchHandler.DeleteDispatch)

		rt.Post("/dispatch-items", handlers.DispatchItemHandler.CreateDispatchItem)
		rt.Put("/dispatch-items/{id}", handlers.DispatchItemHandler.UpdateDispatchItem)
		rt.Delete("/dispatch-items/{id}", handlers.DispatchItemHandler.DeleteDispatchItem)
		rt.Get("/item/{id}/dispatch-items", handlers.DispatchItemHandler.GetDispatchItemListByItemId)
	})

	return app.Routes
}
