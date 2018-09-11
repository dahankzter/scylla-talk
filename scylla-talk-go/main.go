package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/scylladb/gocqlx"

	"github.com/scylladb/gocqlx/qb"

	"github.com/gocql/gocql"
)

var (
	hosts string
)

func init() {
	flag.StringVar(&hosts, "hosts", "localhost", "Scylla hosts to connect to")
}

func main() {
	flag.Parse()

	c := gocql.NewCluster(hosts)
	c.CQLVersion = "3.0.0"
	c.Keyspace = "scylla_talk_ks"

	session, err := c.CreateSession()
	if err != nil {
		log.Printf("unable to create session, error=%s\n", err)
		return
	}

	defer session.Close()

	stmt, _ := qb.Select("user").Columns("id", "name", "email", "about", "properties").ToCql()

	q := session.Query(stmt)
	iter := gocqlx.Iter(q)

	var user User
	for iter.StructScan(&user) {
		fmt.Printf("USER: %s\n", user)
	}
}

type User struct {
	ID         gocql.UUID
	Name       string
	Email      string
	About      string
	Properties map[string]string
}

func (u *User) String() string {
	return fmt.Sprintf("ID=%s, Name=%s, Email=%s, About=%s, Properties=%s", u.ID, u.Name, u.Email, u.About, u.Properties)
}
