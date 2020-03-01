package mapsumtrie

type MapSum struct {
	t *trieMap
}

type trieMap struct {
	child map[byte]*trieMap
	sums  map[string]int
}

func newTrieMap(val int, key string) *trieMap {
	t := &trieMap{
		child: make(map[byte]*trieMap),
		sums:  make(map[string]int),
	}
	if key != "" {
		t.sums[key] = val
	}
	return t
}

func (t *trieMap) add(key string, val int) {
	curr := t
	for ix := range key {
		char := key[ix]
		if node, contains := curr.child[char]; contains {
			node.sums[key] = val
			curr = node
		} else {
			tmp := newTrieMap(val, key)
			curr.child[char] = tmp
			curr = tmp
		}
	}
}

func (t *trieMap) find(prefix string) (sum int) {
	curr := t

	var sums map[string]int
	for ix := range prefix {
		char := prefix[ix]
		if _, contains := curr.child[char]; !contains {
			return 0
		}

		sums = curr.child[char].sums
		curr = curr.child[char]
	}

	for _, val := range sums {
		sum += val
	}
	return
}

func Constructor() MapSum {
	return MapSum{
		t: newTrieMap(0, ""),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.t.add(key, val)
}

func (this *MapSum) Sum(prefix string) int {
	return this.t.find(prefix)
}
