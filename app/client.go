package app

import (
	"github.com/area3001/goira/comm"
)

func NewClient(opts *comm.NatsClientOpts) (*Client, error) {
	nc, err := comm.Dial(opts)
	if err != nil {
		return nil, err
	}

	return &Client{nc}, nil
}

type Client struct {
	nc *comm.NatsClient
}
