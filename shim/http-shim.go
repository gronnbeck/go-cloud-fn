package shim

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

type HTTPRequest struct {
	Body       string            `json:"body"`
	Header     map[string]string `json:"headers"`
	Method     string            `json:"method"`
	RemoteAddr string            `json:"remote_addr"`
	URL        string            `json:"url"`
}

type HTTPResponse struct {
	Body       string            `json:"body"`
	Header     map[string]string `json:"headers"`
	StatusCode int               `json:"status_code"`
}

type RequestHandler func(http.ResponseWriter, *http.Request)

func ServeHTTP(h RequestHandler) {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var httpRequest HTTPRequest
	err = json.Unmarshal(stdin, &httpRequest)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest(httpRequest.Method, httpRequest.URL, bytes.NewBufferString(httpRequest.Body))

	for k, v := range httpRequest.Header {
		r.Header.Add(k, v)
	}

	r.RemoteAddr = httpRequest.RemoteAddr

	w := httptest.NewRecorder()

	h(w, r)

	resp := w.Result()

	header := make(map[string]string)
	for k, v := range resp.Header {
		header[k] = strings.Join(v, ",")
	}
	httpResponse := HTTPResponse{
		Body:       w.Body.String(),
		Header:     header,
		StatusCode: resp.StatusCode,
	}

	out, err := json.Marshal(&httpResponse)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(out)
}
