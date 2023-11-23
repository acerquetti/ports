package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/acerquetti/ports/app"
	"github.com/acerquetti/ports/domain"
)

var domainHTTPErrorMap = map[int]int{
	domain.ErrInvalidModel: http.StatusBadRequest,
}

type rest struct {
	service app.Service
	routes  []route
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler func(w http.ResponseWriter, req *http.Request)
}

func NewREST(service app.Service) *rest {
	r := &rest{
		service: service,
		routes: []route{
			{
				method: "POST",
				regex:  regexp.MustCompile(`^/ports$`),
			},
		},
	}
	r.routes[0].handler = r.Create
	return r
}

func (r *rest) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var routeMatched bool
	for _, route := range r.routes {
		if !route.regex.Match([]byte(req.RequestURI)) {
			continue
		}

		if route.method != req.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		routeMatched = true
		route.handler(w, req)
	}

	if !routeMatched {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (r *rest) Create(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var portRaw domain.PortRaw
	if err := json.Unmarshal(body, &portRaw); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	svcResult := make(chan error)
	go func() {
		svcResult <- r.service.Create(portRaw)
	}()

	select {
	// Capture result from app service called
	case err := <-svcResult:
		if err != nil {
			log.Print(err)
			w.WriteHeader(httpStatusFromDomainError(err))
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	// Timeout in case app service got stuck
	case <-time.After(10 * time.Second):
		w.WriteHeader(http.StatusInternalServerError)
	case <-req.Context().Done():
		log.Print(req.Context().Err())
	}
}

func httpStatusFromDomainError(err error) int {
	if portsErr, ok := err.(*domain.PortsError); ok {
		if httpErr, found := domainHTTPErrorMap[portsErr.Num()]; found {
			return httpErr
		}

		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}
