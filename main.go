package main

import "github.com/gocql/gocql"

func main() {
	cluster := gocql.NewCluster("localhost:9042")
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	session.Close()
}
