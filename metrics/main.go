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

	uniqueSubscriberHitsCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wnf_subscriber_unique_hits_total",
		Help: "The total number of unique hits",
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
		subscriptions, error := tx.CreateBucketIfNotExists([]byte("subscriptions"))
		subscriber := subscriptions.Get([]byte(subscriberId))
		if subscriber == nil {
			subscriptions.Put([]byte(subscriberId), []byte(currentTime.Format(time.RFC3339)))
			subscriberCount.Inc()
		}

		return error
	})

	if writeError != nil {
		fmt.Println(writeError)
	}
}

// Experimental
func IncrementHits(subscriberId string) {
	if len(subscriberId) > 0 {
		db, err := GetDb()
		if err != nil {
			panic(err)
		}
		defer db.Close()

		writeError := db.Update(func(tx *bolt.Tx) error {
			currentTime := time.Now()
			dailyHits, error := tx.CreateBucketIfNotExists([]byte(fmt.Sprintf("hits_%s", currentTime.Format("2006-01-02"))))
			dailyHits.Put([]byte(subscriberId), []byte(currentTime.Format(time.RFC3339)))
			uniqueSubscriberHitsCount.Set(float64(dailyHits.Stats().KeyN))
			return error
		})

		if writeError != nil {
			fmt.Println(writeError)
		}

		subscriberHitsCount.Inc()
	}

	hitsCount.Inc()
}
