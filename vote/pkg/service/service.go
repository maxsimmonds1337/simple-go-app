package service

import (
	"context"
	"fmt"
	"simple-go-app/vote/pkg/db"
	"simple-go-app/vote/pkg/io"

	"simple-go-app/vote/pkg/logger"

	_ "github.com/ibmdb/go_ibm_db"
)

// VoteService describes the service.
type VoteService interface {

	//adds a single vote
	Add(ctx context.Context, vote io.Vote) (v io.Vote, error error)
	//get the votes, returns a slice of votes
	Get(ctx context.Context) (v []io.Vote, error error)
}

type basicVoteService struct{}

func (b *basicVoteService) Add(ctx context.Context, vote io.Vote) (v io.Vote, error error) {
	session, err := db.GetDB2Connection()
	if err != nil {
		return v, err
	}
	defer session.Close()

	stmt, err := session.Prepare("INSERT INTO votes (Nominee, app) VALUES (?, ?)")
	if err != nil {
		return v, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(vote.Nominee, vote.App)
	if err != nil {
		return v, err
	}

	voteID, err := result.LastInsertId()
	if err != nil {
		return v, err
	}
	vote.Id = int(voteID)

	return vote, nil
}

func (b *basicVoteService) Get(ctx context.Context) (t []io.Vote, error error) {
	fmt.Print("getting stuff")
	logger.Logger.Log("err", error)
	session, err := db.GetDB2Connection()
	if err != nil {
		return t, err
	}
	defer session.Close()

	rows, err := session.Query("SELECT id, Nominee, app FROM votes")
	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var vote io.Vote
		err := rows.Scan(&vote.Id, &vote.Nominee, &vote.App)
		if err != nil {
			return t, err
		}
		t = append(t, vote)
	}

	if err := rows.Err(); err != nil {
		return t, err
	}

	return t, nil
}

// NewBasicVoteService returns a naive, stateless implementation of VoteService.
func NewBasicVoteService() VoteService {
	return &basicVoteService{}
}

// New returns a VoteService with all of the expected middleware wired in.
func New(middleware []Middleware) VoteService {
	var svc VoteService = NewBasicVoteService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
