FROM openjdk:11.0.7-jre-slim

ARG JAR_FILE=build/libs/jaeger-tracing-java-service-0.0.1-SNAPSHOT.jar

COPY opentelemetry-javaagent-all.jar opentelemetry-javaagent-all.jar
COPY ${JAR_FILE} app.jar

#ADD https://github.com/aws-observability/aws-otel-java-instrumentation/releases/latest/download/aws-opentelemetry-agent.jar /opt/aws-opentelemetry-agent.jar
#ENV JAVA_TOOL_OPTIONS=-javaagent:/opt/aws-opentelemetry-agent.jar

# ENTRYPOINT ["java", "-jar","/app.jar"]
ENTRYPOINT ["java", "-javaagent:opentelemetry-javaagent-all.jar", "-jar","/app.jar"]
