package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// testapi()
}

func testapi() {
	// resp, err := Get("LopDropFlop")
	// panicerror(err)
	// fmt.Println(resp.String())
	// fmt.Println(len(resp.RecentMatches))
}

func startbot() {
	b := NewBot(token)
	err := b.Connect()
	panicerror(err)

	// time.Sleep(10 * time.Second)
	b.Send(TestChannel, "Test")

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
