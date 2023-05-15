package io

import "encoding/json"

// define the vote
type Vote struct {
	Id      int    `json:"id"`
	Nominee string `json:"nominee"`
}

// a method for unmarshalling and returning the vote as a string (TODO: Might not need)
func (v Vote) String() string {
	b, err := json.Marshal(v)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
