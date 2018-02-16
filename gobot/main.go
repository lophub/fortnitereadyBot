package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var _ = time.Now

var Names = []string{"Emyrks", "LopDropFlop", "r0bd0g364", "GuyWhoDoesThings"}

func main() {
	startbot()
}

func testapi() {
	resp, err := Get("LopDropFlop")
	panicerror(err)
	fmt.Println(resp.String())
	fmt.Println(len(resp.RecentMatches))
}

func startbot() {
	b := NewBot(token, Names)
	err := b.Connect()
	panicerror(err)

	go b.Run()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	b.Close()
}

func panicerror(err error) {
	if err != nil {
		panic(err)
	}
}
