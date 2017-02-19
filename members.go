package propublica_api

import (
	"gopkg.in/resty.v0"
	"encoding/json"
)

type MembersResults struct {
	Results []struct {
		Chamber string
		Congress string
		Num_results string
		Offset string
		Members []Member
	}
}

type Member struct {
	Id string
	Thomas_id string
	Api_uri string
	First_name string
	Middle_name string
	Last_name string
	Gender string
	Party string
	Twitter_account string
	Facebook_account string
	Facebook_id string
	Youtube_account string
	Rss_url string
	Current_party string
	Url string
	Domain string
	Dw_nominate string
	Ideal_point string
	Senority string
	Next_election string
	Total_votes string
	Missed_votes string
	Total_present string
	State string
	Missed_votes_pct string
	Votes_with_arty_pct string
	Times_topic_url string
	Times_tag string
	Govtrack_id string
	Cspan_id string
	Icpsr_id string
	Most_recent_vote string
	Roles []struct {
		Congress string
		Chamber string
		Title string
		State string
		Party string
		Seniority string
		District string
		Start_date string
		End_date string
		Bills_sponsored string
		Bills_cosponsored string
		Missed_votes_pct string
		Votes_with_part_pct string
		Committees []struct {
			Name string
			Code string
			Api_uri string
			Rank_in_party string
			Begin_date string
			End_date string
		}
	}
}

func (api *Client) GetMembers(congress string, chamber string) (*MembersResults, error) {
	resp, _ := resty.R().
		SetHeader("X-API-Key", api.token).
		Get("https://api.propublica.org/congress/v1/" + congress + "/" + chamber + "/members.json")

	data := &MembersResults{}

	if err := json.Unmarshal([]byte(resp.String()), &data); err != nil {
        return nil, err
    }

	return data, nil
}