package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	endpoint "simple-go-app/vote/pkg/endpoint"
	"simple-go-app/vote/pkg/logger"

	http1 "github.com/go-kit/kit/transport/http"
)

// makeAddHandler creates the handler logic
func makeAddHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	logger.Logger.Log("msg", "Making add handler")
	m.Handle("/add", http1.NewServer(endpoints.AddEndpoint, decodeAddRequest, encodeAddResponse, options...))
}

// decodeAddRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	logger.Logger.Log("msg", "decoding http body request")
	req := endpoint.AddRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetHandler creates the handler logic
func makeGetHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	logger.Logger.Log("msg", "Making get handler")
	m.Handle("/get", http1.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...))
}

// decodeGetRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	logger.Logger.Log("msg", "decoding http body request")
	req := endpoint.GetRequest{}
	// logger.Logger.Log("msg: body is ", r.Body)
	// err := json.NewDecoder(r.Body).Decode(&req) // there is no get body
	return req, nil
}

// encodeGetResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
