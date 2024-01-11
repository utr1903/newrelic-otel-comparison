# Java

## OpenTelemetry

### Run the app

Get into the directory of the otel app:

```shell
cd ./java/otel
```

Download the OpenTelemetry auto-instrumentation JAR file:

```shell
wget https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v1.32.0/opentelemetry-javaagent.jar
```

Build the code with:

```shell
mvn -f pom.xml clean package
```

Run the code with:

```shell
export OTEL_SERVICE_NAME=comparison-java-otel; export OTEL_EXPORTER_OTLP_ENDPOINT=<New Relic OTLP endpoint>; export OTEL_EXPORTER_OTLP_HEADERS=api-key=<License key>; export OTEL_LOGS_EXPORTER=otlp; java -javaagent:./opentelemetry-javaagent.jar -jar ./target/app-0.0.1.jar
```

where

- New Relic OTLP endpoint is https://otlp.nr-data.net:4317 for US and https://otlp.eu01.nr-data.net:4317 for EU.
- New Relic license key is your ingest key.

Generate traffic:

```shell
while true; do curl http://localhost:8080/api; sleep 1; done
```

### Monitor in New Relic

See your spans:

```
FROM Span SELECT * WHERE service.name = 'comparison-java-otel'
```

See your span events:

```
FROM SpanEvent SELECT * WHERE service.name = 'comparison-java-otel'
```

See your custom metric:

```
FROM Metric SELECT average(invocations) WHERE service.name = 'comparison-java-otel' TIMESERIES
```

See your logs

```
FROM Log SELECT * WHERE service.name = 'comparison-java-otel'
```

## New Relic

### Run the app

Get into the directory of the otel app:

```shell
cd ./java/newrelic
```

Download the OpenTelemetry auto-instrumentation JAR file:

```shell
wget https://download.newrelic.com/newrelic/java-agent/newrelic-agent/8.8.0/newrelic-agent-8.8.0.jar
```

Build the code with:

```shell
mvn -f pom.xml clean package
```

Run the code with:

```shell
export NEW_RELIC_APP_NAME=comparison-java-newrelic; export NEW_RELIC_LICENSE_KEY=<License key>; java -javaagent:./newrelic-agent-8.8.0.jar -jar ./target/app-0.0.1.jar
```

Generate traffic:

```shell
while true; do curl http://localhost:8080/api; sleep 1; done
```

### Monitor in New Relic

See your transactions:

```
FROM Transaction SELECT * WHERE appName = 'comparison-java-newrelic'
```

See your spans:

```
FROM Span SELECT * WHERE appName = 'comparison-java-newrelic'
```

See your custom events:

```
FROM MyCustomEvent SELECT * WHERE appName = 'comparison-java-newrelic'
```

See your logs

```
FROM Log SELECT * WHERE entity.name = 'comparison-java-newrelic'
```
