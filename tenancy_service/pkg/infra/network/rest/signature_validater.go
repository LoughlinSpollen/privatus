package rest

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"net/http"
	"strconv"
)

type signatureValidater struct {
}

func NewSignatureValidater() *signatureValidater {
	log.Debug("NewSignatureValidater")

	return &signatureValidater{}
}

func (s *signatureValidater) ValidateJSON(w http.ResponseWriter, r *http.Request, body []byte) bool {
	log.Debug("signatureValidater Validate")

	header := r.Header
	if header.Get("Content-Type") != "application/vnd.api+json" {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "Missing Content-Type header")
		return false
	}

	if header.Get("Accept") != "application/vnd.api+json" {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "Missing Accept header")
		return false
	}

	if header.Get("Content-Length") != strconv.Itoa(len(body)) {
		w.WriteHeader(http.StatusLengthRequired)
		fmt.Fprintf(w, "Incorrect Content-Length header")
		return false
	}
	// Todo security domain and digest header validation

	return true
}
