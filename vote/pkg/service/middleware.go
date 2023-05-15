package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	io "simple-go-app/vote/pkg/io"
)

// Middleware describes a service middleware.
type Middleware func(VoteService) VoteService

type loggingMiddleware struct {
	logger log.Logger
	next   VoteService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a VoteService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next VoteService) VoteService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Add(ctx context.Context, vote io.Vote) (v io.Vote, error error) {
	defer func() {
		l.logger.Log("method", "Add", "vote", vote, "v", v, "error", error)
	}()
	return l.next.Add(ctx, vote)
}
func (l loggingMiddleware) Get(ctx context.Context) (v []io.Vote, error error) {
	defer func() {
		l.logger.Log("method", "Get", "v", v, "error", error)
	}()
	return l.next.Get(ctx)
}
