package main

import (
	"kubernetes-postgres/docs"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

// Main entry point of the service listening on 8080
func main() {
	//fmt.Println("starting...")
	/*logger := servicelog.GetInstance()
	logger.Println(time.Now().UTC(), "Starting service")
	// routes defined in routes.go
	router := NewRouter()
	http.ListenAndServe(":8080", router)*/
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"

	router := NewRouter()

	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", router)

}
