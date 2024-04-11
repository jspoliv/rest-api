package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, name, surname, pw string) *Account {
	acc, err := NewAccount(name, surname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account =>", acc.Number)
	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "alice", "bob", "hunter2")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		seedAccounts(store)
	}

	server := NewApiServer(":3000", store)
	server.Run()
}
