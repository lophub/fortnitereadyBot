package main

import (
	"fmt"
	"time"
)

type APIPoller struct {
	Stats map[string]*PlayerStats

	NameList []string
}

func NewAPIPoller(namelist []string) *APIPoller {
	a := new(APIPoller)
	a.Stats = make(map[string]*PlayerStats)
	a.NameList = namelist

	return a
}

func (a *APIPoller) GetName(name string) *PlayerStats {
	if v, ok := a.Stats[name]; ok {
		return v
	}
	return nil
}

func (a *APIPoller) Run() {
	ticker := time.NewTicker(2000 * time.Millisecond)
	i := 0
	for _ = range ticker.C {
		person := a.NameList[i]
		s, err := Get(person)
		if err != nil {
			fmt.Println(err)
		}
		a.Stats[person] = s
		i++
		i = i % len(a.NameList)
	}
}
