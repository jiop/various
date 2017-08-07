package main

import (
	"log"
	"time"
)

func main() {
	log.Println("# start")
	channel := make(chan string)

	go func(c chan string) {
		log.Println("## start")
		c <- "started"
		for {
			log.Println("## loop")
			c <- "message"
			select {
			case v := <-c:
				if v == "stop" {
					log.Println("## stopping")
					c <- "stopping"
					return
				}
			}
		}
	}(channel)

	for {
		log.Println("# loop")
		select {
		case <-time.After(5 * time.Second):
			log.Println("# send stop")
			channel <- "stop"
		case v := <-channel:
			log.Println(v)
			if v == "stopped" {
				return
			}
		}
	}
}
