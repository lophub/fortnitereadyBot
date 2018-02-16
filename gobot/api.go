package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(name string) (*PlayerStats, error) {
	// https://api.fortnitetracker.com/v1/profile/pc/LopDropFlop
	url := fmt.Sprintf("https://api.fortnitetracker.com/v1/profile/pc/%s", name)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("TRN-Api-Key", header)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	stats := new(PlayerStats)
	err = json.Unmarshal(data, stats)

	return stats, err
}
