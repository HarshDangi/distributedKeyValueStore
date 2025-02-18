package consistentHashing

import (
	"bytes"
	"sort"
)

func traverseClockwise(key [16]byte) [16]byte {
	if bytes.Compare(ring[len(ring)-1][:], key[:]) == 1 { // if key is bigger than the largest then obviously first server would be its target i.e. looping around
		return ring[0]
	}

	idx := sort.Search(len(ring), func(idx int) bool {
		comp := bytes.Compare(ring[idx][:], key[:])
		return comp == -1 || comp == 0
	})
	return ring[idx]
}
