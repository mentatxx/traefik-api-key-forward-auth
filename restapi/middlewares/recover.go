package middlewares

import (
	"log"
	"net/http"
)

type ApiPanicHandler struct {
	Next http.Handler
}

// ServeHTTP implements the http.Handler interface.
// It recovers from panics of all next handlers and logs them
func (h *ApiPanicHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("HTTP Handler error: %#v \n", r)
			wr.WriteHeader(http.StatusInternalServerError)
			if _, err := wr.Write([]byte("Internal error occurred")); err != nil {
				log.Println("Failed to write API response")
			}
		}
	}()
	h.Next.ServeHTTP(wr, req)
}
