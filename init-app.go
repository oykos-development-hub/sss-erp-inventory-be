package main

import (
	"log"
	"os"

	"gitlab.sudovi.me/erp/inventory-api/data"
	"gitlab.sudovi.me/erp/inventory-api/handlers"
	"gitlab.sudovi.me/erp/inventory-api/middleware"

	"github.com/oykos-development-hub/celeritas"
	"gitlab.sudovi.me/erp/inventory-api/services"
)

func initApplication() *celeritas.Celeritas {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "gitlab.sudovi.me/erp/inventory-api"

	models := data.New(cel.DB.Pool)

	RealEstateService := services.NewRealEstateServiceImpl(cel, models.RealEstate)
	RealEstateHandler := handlers.NewRealEstateHandler(cel, RealEstateService)

	ItemService := services.NewItemServiceImpl(cel, models.Item)
	ItemHandler := handlers.NewItemHandler(cel, ItemService)

	DispatchService := services.NewDispatchServiceImpl(cel, models.Dispatch)
	DispatchHandler := handlers.NewDispatchHandler(cel, DispatchService)

	DispatchItemService := services.NewDispatchItemServiceImpl(cel, models.DispatchItem, models.Item)
	DispatchItemHandler := handlers.NewDispatchItemHandler(cel, DispatchItemService)

	AssessmentService := services.NewAssessmentServiceImpl(cel, models.Assessment)
	AssessmentHandler := handlers.NewAssessmentHandler(cel, AssessmentService)

	myHandlers := &handlers.Handlers{
		RealEstateHandler:   RealEstateHandler,
		ItemHandler:         ItemHandler,
		AssessmentHandler:   AssessmentHandler,
		DispatchHandler:     DispatchHandler,
		DispatchItemHandler: DispatchItemHandler,
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	return cel
}
