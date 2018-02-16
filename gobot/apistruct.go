package main

import (
	"encoding/json"
)

func (p *PlayerStats) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}

type PlayerStats struct {
	AccountID        string `json:"accountId"`
	PlatformID       int    `json:"platformId"`
	PlatformName     string `json:"platformName"`
	PlatformNameLong string `json:"platformNameLong"`
	EpicUserHandle   string `json:"epicUserHandle"`
	Stats            struct {
		P2 struct {
			TrnRating struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"trnRating"`
			Score struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"score"`
			Top1 struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"top1"`
			Top3 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top3"`
			Top5 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top5"`
			Top6 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top6"`
			Top10 struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"top10"`
			Top12 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top12"`
			Top25 struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"top25"`
			Kd struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kd"`
			WinRatio struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"winRatio"`
			Matches struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"matches"`
			Kills struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kills"`
			MinutesPlayed struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"minutesPlayed"`
			Kpm struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kpm"`
			Kpg struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kpg"`
			AvgTimePlayed struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"avgTimePlayed"`
			ScorePerMatch struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"scorePerMatch"`
			ScorePerMin struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"scorePerMin"`
		} `json:"p2"`
		P10 struct {
			TrnRating struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"trnRating"`
			Score struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"score"`
			Top1 struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"top1"`
			Top3 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top3"`
			Top5 struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"top5"`
			Top6 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top6"`
			Top10 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top10"`
			Top12 struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"top12"`
			Top25 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top25"`
			Kd struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kd"`
			WinRatio struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"winRatio"`
			Matches struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"matches"`
			Kills struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kills"`
			MinutesPlayed struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"minutesPlayed"`
			Kpm struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kpm"`
			Kpg struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kpg"`
			AvgTimePlayed struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"avgTimePlayed"`
			ScorePerMatch struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"scorePerMatch"`
			ScorePerMin struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"scorePerMin"`
		} `json:"p10"`
		P9 struct {
			TrnRating struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"trnRating"`
			Score struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"score"`
			Top1 struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"top1"`
			Top3 struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"top3"`
			Top5 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top5"`
			Top6 struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"top6"`
			Top10 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top10"`
			Top12 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top12"`
			Top25 struct {
				Label        string `json:"label"`
				Field        string `json:"field"`
				Category     string `json:"category"`
				ValueInt     int    `json:"valueInt"`
				Value        string `json:"value"`
				Rank         int    `json:"rank"`
				DisplayValue string `json:"displayValue"`
			} `json:"top25"`
			Kd struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kd"`
			WinRatio struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"winRatio"`
			Matches struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"matches"`
			Kills struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kills"`
			MinutesPlayed struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueInt     int     `json:"valueInt"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"minutesPlayed"`
			Kpm struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kpm"`
			Kpg struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"kpg"`
			AvgTimePlayed struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"avgTimePlayed"`
			ScorePerMatch struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"scorePerMatch"`
			ScorePerMin struct {
				Label        string  `json:"label"`
				Field        string  `json:"field"`
				Category     string  `json:"category"`
				ValueDec     float64 `json:"valueDec"`
				Value        string  `json:"value"`
				Rank         int     `json:"rank"`
				Percentile   float64 `json:"percentile"`
				DisplayValue string  `json:"displayValue"`
			} `json:"scorePerMin"`
		} `json:"p9"`
	} `json:"stats"`
	LifeTimeStats []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"lifeTimeStats"`
	RecentMatches []struct {
		ID              int     `json:"id"`
		AccountID       string  `json:"accountId"`
		Playlist        string  `json:"playlist"`
		Kills           int     `json:"kills"`
		MinutesPlayed   int     `json:"minutesPlayed"`
		Top1            int     `json:"top1"`
		Top5            int     `json:"top5"`
		Top6            int     `json:"top6"`
		Top10           int     `json:"top10"`
		Top12           int     `json:"top12"`
		Top25           int     `json:"top25"`
		Matches         int     `json:"matches"`
		Top3            int     `json:"top3"`
		DateCollected   string  `json:"dateCollected"`
		Score           int     `json:"score"`
		Platform        int     `json:"platform"`
		TrnRating       float64 `json:"trnRating,omitempty"`
		TrnRatingChange float64 `json:"trnRatingChange,omitempty"`
	} `json:"recentMatches"`
}
