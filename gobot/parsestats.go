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

			if since < 36000 {
				_, ok := p.ReportedMatches[match.ID]
				if !ok {
					p.ReportedMatches[match.ID] = struct{}{}
					b.Send(TestChannel, FormatMatch(n, match))
					b.Send(TestChannel, FormatPlayerStats(n, s))
				}
			}
		}

	}
}

func FormatPlayerStats(name string, s *PlayerStats) string {
	stats := s.GetLifeTimeStats()

	str := "```\n"
	str += fmt.Sprintf("%s LifetimeStats \n", name)
	str += line("Wins", fmt.Sprintf("%s", stats["Wins"])) + "\n"
	str += line("Kills", fmt.Sprintf("%s", stats["Kills"])) + "\n"
	str += line("Matches", fmt.Sprintf("%s", stats["Matches Played"])) + "\n"

	str += "```"
	return str
}

func FormatMatch(name string, match Match) string {
	str := "```\n"
	str += fmt.Sprintf("%s Just Finished a game\n", name)
	str += line("Kills", fmt.Sprintf("%d", match.Kills)) + "\n"
	str += line("Duration", fmt.Sprintf("%d mins", match.MinutesPlayed)) + "\n"
	str += line("Matches", fmt.Sprintf("%d", match.Matches)) + "\n"

	str += "```"
	return str
}

func line(label string, value string) string {
	return fmt.Sprintf("%-15s: %s", label, value)
}
