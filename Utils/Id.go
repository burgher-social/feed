package Utils

import (
	"crypto/rand"
	"math/big"
	"time"
)

var TABLE string = ".0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
var maxBigInt = new(big.Int).SetUint64(18446744073709551615)
var lastTime int64
var lastRandom uint64

func GenerateId() string {
	var now int64 = time.Now().UnixMilli()
	var delta uint64 = 0
	randomNumber, _ := rand.Int(rand.Reader, maxBigInt)
	randomNumberWithDelta := randomNumber.Uint64()
	if lastTime == now && lastRandom == randomNumberWithDelta {
		delta = 1
		randomNumberWithDelta = randomNumberWithDelta + delta
	}
	lastTime = now
	lastRandom = randomNumberWithDelta
	kSortedId := make([]byte, 19)
	for i := 18; i >= 8; i-- {
		kSortedId[i] = TABLE[randomNumberWithDelta%64]
		randomNumberWithDelta = randomNumberWithDelta >> 6
	}
	for i := 7; i >= 0; i-- {
		kSortedId[i] = TABLE[now%64]
		now = now >> 6
	}
	return string(kSortedId)
}
