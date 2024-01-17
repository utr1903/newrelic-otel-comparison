# Golang

## OpenTelemetry

### Run the app

Get into the directory of the otel app:

```shell
cd ./golang/apps/otel
```

Run the code with:

```shell
export OTEL_SERVICE_NAME=comparison-golang-otel; export OTEL_EXPORTER_OTLP_ENDPOINT=<New Relic OTLP endpoint>; export OTEL_EXPORTER_OTLP_HEADERS=api-key=<License key>; go run main.go
```

where

- New Relic OTLP endpoint is https://otlp.nr-data.net:4317 for US and https://otlp.eu01.nr-data.net:4317 for EU.
- New Relic license key is your ingest key.

Generate traffic manually:

```shell
while true; do curl http://localhost:8080/api; sleep 1; done
```

### Monitor in New Relic

See your spans:

```
FROM Span SELECT * WHERE service.name = 'comparison-golang-otel'
```

See your span events:

```
FROM SpanEvent SELECT * WHERE service.name = 'comparison-golang-otel'
```

See your custom metric:

```
FROM Metric SELECT average(invocations) WHERE service.name = 'comparison-golang-otel' TIMESERIES
```

## New Relic

### Run the app

Get into the directory of the otel app:

```shell
cd ./golang/apps/newrelic
```

Run the code with:

```shell
export NEW_RELIC_APP_NAME=comparison-golang-newrelic; export NEW_RELIC_LICENSE_KEY=<License key>; go run main.go
```

Generate traffic manually:

```shell
while true; do curl http://localhost:8081/api; sleep 1; done
```

### Monitor in New Relic

See your transactions:

```
FROM Transaction SELECT * WHERE appName = 'comparison-golang-newrelic'
```

See your spans:

```
FROM Span SELECT * WHERE appName = 'comparison-golang-newrelic'
```

See your custom events:

```
FROM MyCustomEvent SELECT * WHERE appName = 'comparison-golang-newrelic'
```

See your logs

```
FROM Log SELECT * WHERE entity.name = 'comparison-golang-newrelic'
```
