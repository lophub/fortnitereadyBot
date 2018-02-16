package main

import (
	"fmt"
	"sort"
	"time"
)

type Parser struct {
	Players         map[string]*PlayerStats
	ReportedMatches map[int]struct{}
}

func NewParser() *Parser {
	p := new(Parser)
	p.Players = make(map[string]*PlayerStats)
	p.ReportedMatches = make(map[int]struct{})

	return p
}

func (b *Bot) Run() {
	go b.APIPoller.Run()
	p := b.Parser
	ticker := time.NewTicker(2000 * time.Millisecond)
	for _ = range ticker.C {
		for n, s := range b.APIPoller.Stats {
			// If player call is the same
			st, ok := p.Players[n]
			if ok {
				if fmt.Sprintf("%p", s) == fmt.Sprintf("%p", st) {
					continue
				}
			}
			p.Players[n] = s

			// If player is different
			sort.Sort(MatchList(s.RecentMatches))

			if len(s.RecentMatches) == 0 {
				continue
			}
			match := s.RecentMatches[0]
			now := time.Now().UTC()

			since := now.Sub(match.GetDate()).Seconds()
			if since < 3600 {
				_, ok := p.ReportedMatches[match.ID]
				if !ok {
					p.ReportedMatches[match.ID] = struct{}{}
					b.Send(TestChannel, FormatMatch(n, match))
				}
			}
		}

	}
}

func FormatMatch(name string, match Match) string {
	str := ""
	str += fmt.Sprintf("%s Just Finished a game\n")
	str += line("Kills", fmt.Sprintf("%d", match.Kills))
	str += line("Duration", fmt.Sprintf("%d mins", match.MinutesPlayed))
	str += line("Matches", fmt.Sprintf("%d", match.Matches))

	return str
}

func line(label string, value string) string {
	return fmt.Sprintf("%-15s: %s", line, value)
}
