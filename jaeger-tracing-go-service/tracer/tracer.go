package tracer

import (
	"log"
	"os"

	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
)

func InitTracer() {
	jaegerHost := os.Getenv("OTEL_JAEGER_ENDPOINT")
	jaegerServiceName := os.Getenv("OTEL_JAEGER_SERVICE_NAME")
	//jaegerExporter := os.Getenv("OTEL_EXPORTER")

	exporter, err := jaeger.NewRawExporter(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint((jaegerHost + "/api/traces"))))
	if err != nil {
		log.Fatalf("Error initializing Trace Provider: %v", err)
	}
	idg := xray.NewIDGenerator()
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.ServiceNameKey.String(jaegerServiceName),
		)),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithIDGenerator(idg),
	)
	otel.SetTracerProvider(tp)
	//otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTextMapPropagator(xray.Propagator{})
}
