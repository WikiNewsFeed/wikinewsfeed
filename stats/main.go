package stats

import (
	"fmt"
	"math/big"
	"time"

	bolt "go.etcd.io/bbolt"
)

type GeneratedStats struct {
	SubscribersTotal int `json:"subscribers_total"`
}

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
		}

		return nil
	})

	if writeError != nil {
		fmt.Println(writeError)
	}
}

// Experimental
func IncrementHits(subscriberId string) {
	db, err := GetDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	writeError := db.Update(func(tx *bolt.Tx) error {
		currentTime := time.Now()
		hits, _ := tx.CreateBucketIfNotExists([]byte("hits"))
		currentFormatted := currentTime.Format("2006-01-02")
		currentCounter := hits.Get([]byte(currentFormatted))
		if currentCounter == nil {
			hits.Put([]byte(currentFormatted), big.NewInt(0).Bytes())
		}

		counterAmount := hits.Get([]byte(currentFormatted))
		counter := new(big.Int).SetBytes(counterAmount)
		updatedCounter := counter.Add(counter, big.NewInt(1))
		hits.Put([]byte(currentFormatted), updatedCounter.Bytes())
		return nil
	})

	if writeError != nil {
		fmt.Println(writeError)
	}
}

func GetStats() (*GeneratedStats, error) {
	db, err := GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var stats = GeneratedStats{}

	readError := db.View(func(tx *bolt.Tx) error {
		subscriptions := tx.Bucket([]byte("subscriptions"))
		subscribersTotal := subscriptions.Stats().KeyN
		stats.SubscribersTotal = subscribersTotal
		return nil
	})

	if readError != nil {
		return nil, err
	}

	return &stats, nil
}
