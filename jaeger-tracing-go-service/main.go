package main

import (
	"jaeger-tracing-go-service/config"
	"jaeger-tracing-go-service/routes"
	"jaeger-tracing-go-service/tracer"
	"log"

	"github.com/gin-gonic/gin"

	gintrace "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

var tr = otel.Tracer("go-service")

func main() {
	tracer.InitTracer()

	// Set client options
	config.Connect()

	// Init Router
	router := gin.Default()
	router.Use(gintrace.Middleware("go-service"))

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8091"))
}
