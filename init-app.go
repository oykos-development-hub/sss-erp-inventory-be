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

	ErrorLogService := services.NewErrorLogServiceImpl(cel, models.ErrorLog)
	ErrorLogHandler := handlers.NewErrorLogHandler(cel, ErrorLogService)

	RealEstateService := services.NewRealEstateServiceImpl(cel, models.RealEstate)
	RealEstateHandler := handlers.NewRealEstateHandler(cel, RealEstateService, ErrorLogService)

	ItemService := services.NewItemServiceImpl(cel, models.Item)
	ItemHandler := handlers.NewItemHandler(cel, ItemService, ErrorLogService)

	DispatchService := services.NewDispatchServiceImpl(cel, models.Dispatch)
	DispatchHandler := handlers.NewDispatchHandler(cel, DispatchService, ErrorLogService)

	DispatchItemService := services.NewDispatchItemServiceImpl(cel, models.DispatchItem, models.Item)
	DispatchItemHandler := handlers.NewDispatchItemHandler(cel, DispatchItemService, ErrorLogService)

	AssessmentService := services.NewAssessmentServiceImpl(cel, models.Assessment)
	AssessmentHandler := handlers.NewAssessmentHandler(cel, AssessmentService, ErrorLogService)

	LogService := services.NewLogServiceImpl(cel, models.Log)
	LogHandler := handlers.NewLogHandler(cel, LogService, ErrorLogService)

	myHandlers := &handlers.Handlers{
		RealEstateHandler:   RealEstateHandler,
		ItemHandler:         ItemHandler,
		AssessmentHandler:   AssessmentHandler,
		DispatchHandler:     DispatchHandler,
		DispatchItemHandler: DispatchItemHandler,
		LogHandler:          LogHandler,
		ErrorLogHandler:     ErrorLogHandler,
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	return cel
}
