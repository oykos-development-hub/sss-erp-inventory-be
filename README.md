# golang-laravel

go test --coverprofile=coverage.out . --tags integration
go test --coverprofile=coverage.out && go tool cover --html=coverage.out
go tool cover --html=coverage.out
go test -cover . --tags integration --count=1
go test -v . --tags integration --count=1

## Application structure

-   data (contains your database models)
-   handlers (handles HTTP requests and responses)
-   logs (handles application logging)
-   mail (handles email sending)
-   middleware (contains HTTP middleware functions)
-   migrations (contains database migration scripts)
-   public (contains static files served publicly)
-   services (contains business logic)
-   storage (handles file storage)
-   views (contains templates for HTML views)
-   errors (contains custom errors and error handling logic)
