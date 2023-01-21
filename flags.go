package main

import (
	"flag"
)

var (
	initDB bool = false
)

func flags() {
	flag.BoolVar(&initDB, "init", false, "Delete the existing DB and rebuild a new blank one before running.")
	flag.Parse()
}

