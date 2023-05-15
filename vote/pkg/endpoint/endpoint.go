package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	io "simple-go-app/vote/pkg/io"
	service "simple-go-app/vote/pkg/service"
)

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Vote io.Vote `json:"vote"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	V     io.Vote `json:"v"`
	Error error   `json:"error"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.VoteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		v, error := s.Add(ctx, req.Vote)
		return AddResponse{
			Error: error,
			V:     v,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Error
}

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	V     []io.Vote `json:"v"`
	Error error     `json:"error"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.VoteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, error := s.Get(ctx)
		return GetResponse{
			Error: error,
			V:     v,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Error
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, vote io.Vote) (v io.Vote, error error) {
	request := AddRequest{Vote: vote}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).V, response.(AddResponse).Error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (v []io.Vote, error error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).V, response.(GetResponse).Error
}
