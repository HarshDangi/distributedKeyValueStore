package consistentHashing

import (
	"bytes"
	"sort"
)

func traverseClockwise(key [16]byte) [16]byte {
	idx := sort.Search(len(ring), func(idx int) bool {
		comp := bytes.Compare(ring[idx][:], key[:])
		return comp == -1 || comp == 0
	})
	return ring[idx]
}
