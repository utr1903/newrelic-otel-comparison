package server

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type Server struct {
	app        *newrelic.Application
	randomizer *rand.Rand
}

func New(app *newrelic.Application) *Server {
	// Instantiate randomizer
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Server{
		app:        app,
		randomizer: randomizer,
	}
}

// Server handler
func (s *Server) Handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	if s.randomizer.Intn(15) == 1 {

		// Get transaction from context
		txt := newrelic.FromContext(r.Context())

		// Record exception
		txt.NoticeError(errors.New("failed"))

		// Record custom event
		s.app.RecordCustomEvent(
			"MyCustomEvent",
			map[string]interface{}{
				"mykey": "myvalue",
			})

		// Set response
		s.createHttpResponse(w, http.StatusInternalServerError, []byte("NOT OK"))
	} else {

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
