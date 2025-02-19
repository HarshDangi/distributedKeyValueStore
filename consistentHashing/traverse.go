package consistentHashing

import (
	"bytes"
	"sort"
)

func traverseClockwise(key [16]byte) hexNumber {
	if bytes.Compare(ring[len(ring)-1][:], key[:]) == 1 { // if key is bigger than the largest then obviously first server would be its target i.e. looping around
		return ring[0]
	}

	idx := sort.Search(len(ring), func(idx int) bool {
		comp := bytes.Compare(ring[idx][:], key[:])
		return comp == 1 || comp == 0
	})
	return ring[idx]
}

func traverseAntiClockwise(key [16]byte) hexNumber {
	if bytes.Compare(ring[0][:], key[:]) == -1 { // if key is smaller than the first server, then the last server would be the target i.e. looping around
		return ring[len(ring)-1]
	}

	idx := sort.Search(len(ring), func(idx int) bool {
		comp := bytes.Compare(ring[idx][:], key[:])
		return comp == -1 || comp == 0
	})
	return ring[idx]
}
