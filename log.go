package main

import (
	"log"
	"os"
)

var logError = log.New(os.Stderr, "", 0)
var logSuccess = log.New(os.Stdout, "", 0)
