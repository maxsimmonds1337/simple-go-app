package service

import (
	"context"

	"github.com/maxsimmonds1337/simple-go-app/vote/pkg/io"
)

// VoteService describes the service.
type VoteService interface {

	//adds a single vote
	Add(ctx context.Context, vote io.Vote) (v io.Vote, error error)
	//get the votes, returns a slice of votes
	Get(ctx context.Context) (v []io.Vote, error error)
}
