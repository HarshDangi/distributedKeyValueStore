package consistentHashing

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"slices"
)

type hexNumber [16]byte

var ring []hexNumber
var hashToServer map[hexNumber]uint32
var ctx = context.Background()

type Servers struct {
	serversInfo []ServerInfo `json:"servers"`
}
type ServerInfo struct {
	id          uint32
	hexLocation string
}

type ServerInfoDecoded struct {
	id          uint32
	hexLocation hexNumber
}

// Loads the consistent hashing ring from redis into local memory for quick searching
func Initialise() error {
	data, err := os.ReadFile("../config/serverLocations.txt")
	if err != nil {
		log.Fatal("can't read the serverLocation config")
	}
	serversEncoded := Servers{}
	json.Unmarshal(data, &serversEncoded)

	serversDecoded := make([]ServerInfoDecoded, len(serversEncoded.serversInfo))

	for i, info := range serversEncoded.serversInfo {
		hexValue, _ := hex.DecodeString(info.hexLocation)
		serversDecoded[i] = ServerInfoDecoded{info.id, hexNumber(hexValue)}
	}

	slices.SortFunc(serversDecoded, func(a ServerInfoDecoded, b ServerInfoDecoded) int {
		return bytes.Compare(a.hexLocation[:], b.hexLocation[:])
	})

	for i, info := range serversDecoded {
		copy(ring[i][:], info.hexLocation[:])
		hashToServer[ring[i]] = info.id
	}

	return nil
}

func GetServer(key string) uint32 {
	serverHex := traverseClockwise(md5.Sum([]byte(key)))
	return hashToServer[serverHex]
}

func AddServer(serverId uint32, serverLocation hexNumber) error {
	data, err := os.ReadFile("../config/serverLocations.txt")
	if err != nil {
		log.Fatal("can't read the serverLocation config")
		return err
	}
	serversEncoded := Servers{}
	json.Unmarshal(data, &serversEncoded)

	serversEncoded.serversInfo = append(serversEncoded.serversInfo, ServerInfo{serverId, hex.EncodeToString(serverLocation[:])})

	marshalledValue, err := json.Marshal(serversEncoded)
	if err != nil {
		return err
	}
	os.WriteFile("../config/serverLocations.txt", marshalledValue, 0600)
	return nil
}

func getAffectedKeys() {

}

func RemoveServe() {

}
