package handlers

import (
	"net/http"
	"time"

	"github.com/crhino/grace/routes"
	"code.cloudfoundry.org/lager"
	"github.com/tedsuo/rata"
)

func New(logger lager.Logger) rata.Handlers {
	var sharedCounter uint64

	handlers := rata.Handlers{
		routes.Ping:             &Ping{},
		routes.InstanceIndex:    &InstanceIndex{},
		routes.StartedAt:        &StartedAt{Time: time.Now()},
		routes.Env:              &Env{},
		routes.Cwd:              &Cwd{},
		routes.ListExperiments:  &ListExperiments{},
		routes.Experiments:      &Experiments{},
		routes.Hello:            &Hello{},
		routes.Exit:             &Exit{},
		routes.Curl:             &Curl{},
		routes.MakeTmpFile:      &MakeTmpFile{},
		routes.DeleteTmpFile:    &DeleteTmpFile{},
		routes.Stick:            &Stick{},
		routes.Unstick:          &Unstick{},
		routes.MountFUSEFS:      &MountFUSEFS{},
		routes.ListFUSEFS:       &ListFUSEFS{},
		routes.IncrementCounter: &IncrementCounter{&sharedCounter},
		routes.ReadCounter:      &ReadCounter{&sharedCounter},
	}

	for route, handler := range handlers {
		handlers[route] = &LoggingHandler{
			Route:   route,
			Handler: handler,
			Logger:  logger,
		}
	}

	return handlers
}

type LoggingHandler struct {
	Route   string
	Handler http.Handler
	Logger  lager.Logger
}

func (h *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.Logger.Session(h.Route)
	session.Info("request.begin")
	h.Handler.ServeHTTP(w, r)
	session.Info("request.end")
}
