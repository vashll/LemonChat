package msgServer

import "encoding/json"

type request struct {
	ClientId      string
	RequestTime   int64
	RequestHandle string
	Params        string
}

func (r *request) ToString() string {
	b, err := json.Marshal(r)
	if err != nil {
		return "json maishal failed!"
	}

	return string(b)
}

