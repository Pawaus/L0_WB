package nats

import (
	"github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

func Connect(clusterId, clientId, server string) (*stan.Conn, error) {
	nc, err := nats.Connect(server)
	if err != nil {
		return nil, err
	}
	sc, err := stan.Connect(clusterId, clientId, stan.NatsConn(nc))
	if err != nil {
		return nil, err
	}
	return &sc, nil
}
