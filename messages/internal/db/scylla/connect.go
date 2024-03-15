package scylla

import (
	"github.com/gocql/gocql"
)

func MustConnect(clusters ...string) *gocql.Session {
	cluster := gocql.NewCluster(clusters...)
	cluster.Keyspace = "messages"
	cluster.ProtoVersion = 4

	s, err := cluster.CreateSession()
	if err != nil {
		panic("failed to create session: " + err.Error())
	}

	return s
}
