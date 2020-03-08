package main

import (
	"log"
	"time"
)

func main() {
	log.Println("1")
	bigSlowOperation()
	log.Println("2")
}

func bigSlowOperation() {
	log.Println("3")
	defer trace("bigSlowOperation")()
	log.Println("4")
	time.Sleep(10 * time.Second)
	log.Println("5")
}

func trace(msg string) func() {
	log.Println("6")
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Println("7")
		log.Printf("exit %s (%s)", msg, time.Since(start))
		log.Println("8")
	}
}
