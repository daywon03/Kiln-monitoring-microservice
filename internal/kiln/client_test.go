package kiln

import (
	"context"
	"time"
    "testing"
)

func TestClient_GetRewards(t *testing.T) {
    //Create client
    client := NewClient("https://api.kiln.fi", "", 10*time.Second)
    
    //Define period
    from := time.Now().Add(-1 * time.Hour)
    to := time.Now()
    
    //Testing
    snapshot, err := client.GetRewards(context.Background(), "acc_demo", from, to)

    //verification
    if err != nil {
        t.Fatalf("erreur %v", err)
    }

    if snapshot.AccountID != "acc_demo" {
        t.Errorf("AccountID must be : acc_demo, recieved %s", snapshot.AccountID)
    }

    if snapshot.Uptime < 0 || snapshot.Uptime > 1 {
        t.Errorf("Uptime must be between 0 and 1 , recieved : %f", snapshot.Uptime)
    }

    t.Logf("Test succed %+v", snapshot)
    
}
