package consistentHashing

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/harshdangi/distributedKeyValueStore/redisClient"
	"github.com/redis/go-redis/v9"
)

type hexNumber [16]byte

var ring []hexNumber
var hashToServer map[hexNumber]string
var ctx = context.Background()

// Loads the consistent hashing ring from redis into local memory for quick searching
func Initialise() error {
	vals, err := redisClient.Client.ZRangeByScore(ctx, "consistentHash", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err != nil {
		return err
	}
	ring = make([]hexNumber, len(vals)/2)
	for i, v := range vals {
		if i%2 != 0 {
			hexValue, _ := hex.DecodeString(v)
			copy(ring[i][:], hexValue)
			hashToServer[ring[i]] = vals[i-1]
		}
	}
	return nil
}

func GetServer(key string) string {
	serverHex := traverseClockwise(md5.Sum([]byte(key)))
	return hashToServer[serverHex]
}

func AddServer() {

}

func getAffectedKeys() {

}

func RemoveServe() {

}
