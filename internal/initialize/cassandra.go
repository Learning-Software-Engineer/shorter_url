package initialize

import (
	"github.com/a6slab/shorter_url/global"
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

func InitCassandra() *gocql.Session {
	r := global.Config.Cassandra
	cluster := gocql.NewCluster(r.Url)
	cluster.Keyspace = r.Keyspace

	session, err := cluster.CreateSession()

	if err != nil {
		logrus.Fatal(err)
	}

	return session
}
