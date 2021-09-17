const { NodeTracerProvider } = require('@opentelemetry/sdk-trace-node');
const { SimpleSpanProcessor } = require('@opentelemetry/sdk-trace-base');
const { Resource } = require('@opentelemetry/resources');
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions');
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');
//const { CollectorTraceExporter } = require('@opentelemetry/exporter-collector-grpc')
const { AWSXRayPropagator } = require('@opentelemetry/propagator-aws-xray');
const { AWSXRayIdGenerator } = require('@opentelemetry/id-generator-aws-xray');
const { trace } = require('@opentelemetry/api');
const { HttpInstrumentation } = require('@opentelemetry/instrumentation-http');
const { ExpressInstrumentation } = require('@opentelemetry/instrumentation-express');
const { registerInstrumentations } = require('@opentelemetry/instrumentation');

module.exports = () => {

  const serviceName = process.env.SERVICE_NAME || "nodejs-service";
  const jaegerAgentHost = process.env.JAEGER_AGENT_HOST || 'localhost';
  const jaegerAgentPort = process.env.JAEGER_AGENT_PORT || '14250';
  

  const provider = new NodeTracerProvider({
    idGenerator: new AWSXRayIdGenerator(),
    resource: new Resource({
      [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
    }),
  });
  registerInstrumentations({
    instrumentations: [
      new HttpInstrumentation(),
      new ExpressInstrumentation()
    ]
  })

  //const otlpExporter = new CollectorTraceExporter({
  //  url: `${jaegerAgentHost}:${jaegerAgentPort}`
  //})

  //provider.addSpanProcessor( new SimpleSpanProcessor(otlpExporter))

  provider.addSpanProcessor(
    new SimpleSpanProcessor(
      new JaegerExporter({
          host: jaegerAgentHost,
          port: jaegerAgentPort
      })
    )
  );

  provider.register({
    propagator: new AWSXRayPropagator()
  });
  console.log("Tracing initialized");
  return trace.getTracer("microservice-demo");
}