package propublica_api

import (
	"gopkg.in/resty.v0"
	"encoding/json"
)

type BillsResults struct {
	Status string
	Results []struct {
		Chamber string
		Congress string
		Num_results string
		Offset string
		Bills []Bill
	}
}

type Bill struct {
	Title string
	Number string
	Bill_uri string
	Introduced_date string
	Committees string
	Latest_major_action string
	Latest_major_action_date string
	Sponsor_uri string
	Cosponsors string
}

func (api *Client) GetLatestBills(congress string, chamber string, status string) (*BillsResults, error) {
	resp, _ := resty.R().
		SetHeader("X-API-Key", api.token).
		Get("https://api.propublica.org/congress/v1/" + congress + "/" + chamber + "/bills/" + status + ".json")

	data := &BillsResults{}

	if err := json.Unmarshal([]byte(resp.String()), &data); err != nil {
        return nil, err
    }

	return data, nil
}