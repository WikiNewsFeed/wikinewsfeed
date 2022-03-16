package metrics

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	bolt "go.etcd.io/bbolt"
)

var (
	subscriberCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "wnf_subscribers_total",
		Help: "The total number of feed subscribers",
	})

	hitsCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "wnf_hits_total",
		Help: "The total number of feed hits",
	})

	subscriberHitsCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "wnf_subscriber_hits_total",
		Help: "The total number of feed hits by subscribers",
	})
)

func SubscribeIfNotAlready(subscriberId string) {
	db, err := GetDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if len(subscriberId) == 0 {
		subscriberId = "anonymous"
	}

	writeError := db.Update(func(tx *bolt.Tx) error {
		currentTime := time.Now()
		subscriptions, _ := tx.CreateBucketIfNotExists([]byte("subscriptions"))
		subscriber := subscriptions.Get([]byte(subscriberId))
		if subscriber == nil {
			subscriptions.Put([]byte(subscriberId), []byte(currentTime.Format(time.RFC3339)))
			subscriberCount.Inc()
		}

		return nil
	})

	if writeError != nil {
		fmt.Println(writeError)
	}
}

// Experimental
func IncrementHits(subscriberId string) {
	if len(subscriberId) > 0 {
		subscriberHitsCount.Inc()
	}

	hitsCount.Inc()
}
