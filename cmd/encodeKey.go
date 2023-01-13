package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var globalHash = map[string]map[string]string{}

func storeKeyPair(key, value string) {
	entry := make(map[string]string)
	entry[key] = value

	for k, v := range entry {
		sha256 := sha256.Sum256([]byte(entry[k]))

		hex := hex.EncodeToString(sha256[:])

		globalHash[hex] = map[string]string{}
		globalHash[hex][k] = v
	}

	for k, v := range entry {
		fmt.Println(k)
		fmt.Println(v)
	}
}
