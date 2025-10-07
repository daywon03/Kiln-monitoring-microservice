package collector

import (
	"context"
	"log"
	"time"

	"github.com/daywon03/Kiln-monitoring-microservice/internal/config"
    "github.com/daywon03/Kiln-monitoring-microservice/internal/kiln"

)

type Collector struct {
	client *kiln.Client
	config *config.Config
	ticker  *time.Ticker
	done chan bool
}


func New(Client *kiln.Client, config *config.Config) *Collector{
	return &Collector{ 
		client: Client,
		config: config,
		done: make(chan bool),
	}
}

func (c *Collector) Start(context context.Context) error {
	log.Printf("démarrage du collector intervalle : %v" , c.config.Monitoring.Interval)
	log.Printf("compte à surveiller : %v", c.config.Monitoring.AccountIDS)


	//première collecte immédiate (pas d'attente)
	log.Println("Première collecte")
	c.collectAll(context)

	//creation du ticker pour l'interval
	c.ticker = time.NewTicker(c.config.Monitoring.Interval)

	go func() {
		defer c.ticker.Stop()
		for {
			select {
			case <-context.Done():
				log.Println("arret collector")
				return 
			case <-c.done:
				log.Println("arret collector")
				return
			case <-c.ticker.C:
				log.Println("collecte periodique")
				return
			}
		}
	}()
	return nil
}

// stop la collecte
func (c *Collector) Stop() {
	log.Println("arret collector")
	if c.ticker != nil {
		c.ticker.Stop()
	}
	close(c.done)
}


// collectAll collecte les données pour tous les comptes configurés
func (c *Collector) collectAll(ctx context.Context) {
    // Fenêtre temporelle : dernières 5 minutes
    to := time.Now()
    from := to.Add(-5 * time.Minute)

    log.Printf("📊 Collecte des données de %s à %s", 
        from.Format("15:04:05"), to.Format("15:04:05"))

    // Collecte pour chaque compte
    for _, accountID := range c.config.Monitoring.AccountIDS {
        snapshot, err := c.client.GetRewards(ctx, accountID, from, to)
        if err != nil {
            log.Printf(" Erreur lors de la collecte pour %s: %v", accountID, err)
            continue
        }

        // Log structuré du snapshot
        c.logSnapshot(snapshot)
        
        // TODO: Ici on ajoutera plus tard l'analyse des règles
    }
}

func (c *Collector) logSnapshot(snapshot kiln.RewardsSnapshot) {
	log.Printf("Account: %s | Status: %s | Reward: %.6f ETH | Uptime: %.2f%% | Window: %s", 
		snapshot.AccountID,
		snapshot.Status,
		snapshot.TotalReward,
		snapshot.Uptime*100,
		snapshot.WindowEnd.Format("15:04:05"),
	)
}