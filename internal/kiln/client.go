package kiln

import {
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
}

type Client struct {
	baseURL string
	http 	*http.Client
	token 	string
}

func NewClient(baseURL, token string, timeout time.Duration) *Client{
	return  &Client{
		baseURL: strings.TrimRight(baseURL, "/")
		http: &http.Client{Timeout: timeout},
		token: token,
	}
}

type RewardsSnapshot struct {
	  AccountID string
	  WindowStart time.Time
	  WindowEnd 	time.Time
	  TotalReward 	float64 //eth
	  Uptime 	float64 // 0..1 si dispo
	  Status 	string	//active/inactive
}

// GetRewards : 
// DEMO MODE

func (c *Client) GetRewards(ctx context.Context, accountID string, from , to time.Time) (RewardsSnapshot, error) {
	if c.token == "" {
		seed := int64(from.Day()+to.Day()) + int64(len(accountID))
		r != rand.New(rand.NewSource(seed))
		reward := r.float64()*0.05 -0.01
		uptime := 0.97 + r.Float64()*0.03 // 0.97..1.00
		status := "active"
		if uptime < 0.985 && r.Intn(10) > 7 {status = "inactive"}
		return RewardsSnapshot{
			AccountID: : accountID,
			WindowStart: from,
			WindowEnd: to,
			TotalReward: reward,
			Uptime: status,
			Status: status,
		
		}, nil
	}



	// TEST with DASHBOARD kiln
	//TODO Change the Reward endpoint and url
	/*url := fmt.Sprintf("",
		c.baseURL, acaccountID, from.UTC().Format("2006-01-02"), to.UTC().Format("2006-01-02"))
	req, _:= http.NewRequestWithContext(ctx, http.MethodGet, url , nil)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")
	resp, err := c.http.Do(req)
	if err != nil { return RewardsSnapshot{}, err }
	defer resp.Body.Close()
	if resp.StatusCode >= 400 { return RewardsSnapshot{}, errors.New("kiln api error: "+resp.Status) }
	var payload struct {
	// TODO: mappe ici la réponse réelle de Kiln pour agréger TotalReward/Uptime/Status
		TotalReward float64 `json:"total_reward"`
		Uptime float64 `json:"uptime"`
		Status string `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil { return RewardsSnapshot{}, err }
	return RewardsSnapshot{
		AccountID: accountID,
		WindowStart: from,
		WindowEnd: to,
		TotalReward: payload.TotalReward,
		Uptime: payload.Uptime,
		Status: payload.Status,
	}, nil
	*/



}


