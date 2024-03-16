package scylla

import (
	"github.com/alserov/chatter/messages/internal/db/migrations/scylla"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func MustConnect(ks string, hosts ...string) gocqlx.Session {
	cluster := gocql.NewCluster(hosts...)
	cluster.ProtoVersion = 4
	cluster.Keyspace = ks

	s, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic("failed to init session: " + err.Error())
	}

	scylla.MustMigrate(s)

	return s
}
