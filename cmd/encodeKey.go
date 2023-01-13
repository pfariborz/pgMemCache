package cmd

import (
	"crypto/sha256"
	"encoding/hex"
)

var globalHash = map[string]map[string]string{}

func storeKeyPair() {
	entry := make(map[string]string)
	entry["key"] = "value"

	for k, v := range entry {
		sha256 := sha256.Sum256([]byte(entry[k]))

		hex := hex.EncodeToString(sha256[:])

		globalHash[hex] = map[string]string{}
		globalHash[hex][k] = v
	}
}
