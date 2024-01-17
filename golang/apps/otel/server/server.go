package server

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	randomizer  *rand.Rand
	invocations metric.Int64Counter
}

func New() *Server {

	// Instantiate randomizer
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create custom metric
	invocations, err := otel.GetMeterProvider().Meter("server").
		Int64Counter(
			"invocations",
			metric.WithDescription("Measures the number of method invocations."),
		)
	if err != nil {
		panic(err)
	}

	return &Server{
		randomizer:  randomizer,
		invocations: invocations,
	}
}

// Server handler
func (s *Server) Handler(
	w http.ResponseWriter,
	r *http.Request,
) {

	if s.randomizer.Intn(15) == 1 {
		// Increment to counter
		s.invocations.Add(r.Context(), 1,
			metric.WithAttributes(attribute.Bool("succeeded", false)))

		// Get span from context
		span := trace.SpanFromContext(r.Context())
		if span.IsRecording() {
			// Set OTel status code & description
			span.SetStatus(codes.Error, "failed")

			// Record exception
			span.RecordError(errors.New("failed"))

			// Record custom event
			span.AddEvent(
				"MyCustomEvent",
				trace.EventOption(
					trace.WithAttributes(
						attribute.String("mykey", "myvalue"),
					),
				),
			)
		}

		// Set response
		s.createHttpResponse(w, http.StatusInternalServerError, []byte("NOT OK"))
	} else {
		// Increment to counter
		s.invocations.Add(r.Context(), 1,
			metric.WithAttributes(attribute.Bool("succeeded", true)))

		// Set response
		s.createHttpResponse(w, http.StatusOK, []byte("OK"))
	}
}

func (s *Server) createHttpResponse(
	w http.ResponseWriter,
	statusCode int,
	body []byte,
) {
	w.WriteHeader(statusCode)
	w.Write(body)
}
