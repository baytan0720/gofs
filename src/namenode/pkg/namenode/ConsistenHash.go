package namenode

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type uints []uint32

func (x uints) Len() int           { return len(x) }
func (x uints) Less(i, j int) bool { return x[i] < x[j] }
func (x uints) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ConsistentHash struct {
	circle           map[uint32]string // 保存环
	members          map[string]bool   // 成员标记
	sortedHashes     uints             // 环上的虚拟节点，有序的slice
	NumberOfReplicas int               // 副本数量
	count            int64             // 节点数量
}

func MakeConsistentHash() *ConsistentHash {
	c := &ConsistentHash{
		NumberOfReplicas: 1024,
		circle:           make(map[uint32]string),
		members:          make(map[string]bool),
	}
	return c
}

func (c *ConsistentHash) FormatKey(elt string, idx int) string {
	return strconv.Itoa(idx) + elt
}

func (c *ConsistentHash) Add(elt string) {
	for i := 0; i < c.NumberOfReplicas; i++ {
		c.circle[c.hashKey(c.FormatKey(elt, i))] = elt
	}
	c.members[elt] = true
	c.updateSortedHashes()
	c.count++
}

func (c *ConsistentHash) Remove(elt string) {
	for i := 0; i < c.NumberOfReplicas; i++ {
		delete(c.circle, c.hashKey(c.FormatKey(elt, i)))
	}
	delete(c.members, elt)
	c.updateSortedHashes()
	c.count--
}

func (c *ConsistentHash) Get(name string) string {
	key := c.hashKey(name)
	i := c.search(key)
	return c.circle[c.sortedHashes[i]]
}

func (c *ConsistentHash) search(key uint32) (i int) {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	i = sort.Search(len(c.sortedHashes), f)
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return
}

func (c *ConsistentHash) GetThree(name string) []string {
	var (
		key   = c.hashKey(name)
		i     = c.search(key)
		start = i
		res   = make([]string, 0, 3)
		elem  = c.circle[c.sortedHashes[i]]
	)
	res = append(res, elem)
	for i = start + 1; i != start; i++ {
		if i >= len(c.sortedHashes) {
			i = 0
		}
		elem = c.circle[c.sortedHashes[i]]
		if !sliceContainsMember(res, elem) {
			res = append(res, elem)
		}
		if len(res) == 3 {
			break
		}
	}

	return res
}

func (c *ConsistentHash) GetN(name string, n int) []string {
	if c.count < int64(n) {
		n = int(c.count)
	}

	var (
		key   = c.hashKey(name)
		i     = c.search(key)
		start = i
		res   = make([]string, 0, n)
		elem  = c.circle[c.sortedHashes[i]]
	)

	res = append(res, elem)

	if len(res) == n {
		return res
	}

	for i = start + 1; i != start; i++ {
		if i >= len(c.sortedHashes) {
			i = 0
		}
		elem = c.circle[c.sortedHashes[i]]
		if !sliceContainsMember(res, elem) {
			res = append(res, elem)
		}
		if len(res) == n {
			break
		}
	}

	return res
}

func (c *ConsistentHash) hashKey(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *ConsistentHash) updateSortedHashes() {
	hashes := c.sortedHashes[:0]
	if cap(c.sortedHashes)/(c.NumberOfReplicas*4) > len(c.circle) {
		hashes = nil
	}
	for k := range c.circle {
		hashes = append(hashes, k)
	}
	sort.Sort(hashes)
	c.sortedHashes = hashes
}

func sliceContainsMember(set []string, member string) bool {
	for _, m := range set {
		if m == member {
			return true
		}
	}
	return false
}
