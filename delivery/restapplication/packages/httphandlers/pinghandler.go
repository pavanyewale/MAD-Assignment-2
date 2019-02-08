package httphandlers

import (
	"fmt"
	"net/http"
	mthdroutr "pavan/gohttpexamples/sample4/delivery/restapplication/packages/mthdrouter"
	"pavan/gohttpexamples/sample4/delivery/restapplication/packages/resputl"
)

// PingHandler is a Basic ping utility for the service
type PingHandler struct {
	BaseHandler
}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

// Get function for PingHandler
func (p *PingHandler) Get(r *http.Request) resputl.SrvcRes {
	s := r.URL.Query()
	key, ok := s["key"]
	if ok {
		fmt.Println(key[0])
	}
	return resputl.ResponseNotImplemented(nil)
	return resputl.Response200OK("OK")
}
