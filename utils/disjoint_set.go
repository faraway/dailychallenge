package utils

type DisjointSet struct {
	Parent []int
	Rank   []int
}

func NewDisjointSet(size int) DisjointSet {
	ds := DisjointSet{
		Parent: make([]int, size),
		Rank:   make([]int, size),
	}
	for i,_ := range ds.Parent {
		ds.Parent[i] = i
		ds.Rank[i] = 1
	}
	return ds
}

/**
 Find with path compression
 (note the recursion)
 */
func (d DisjointSet) Find(x int) int {
	for d.Parent[x] == x {
		return x
	}
	d.Parent[x] = d.Find(d.Parent[x])
	return d.Parent[x]
}

/**
 Union with rank optimization
 */
func (d DisjointSet) Union(x int, y int) {
	rootX := d.Find(x)
	rootY := d.Find(y)
	if rootX != rootY {
		if d.Rank[rootX] > d.Rank[rootY] {
			d.Parent[rootY] = rootX
		} else if d.Rank[rootX] < d.Rank[rootY] {
			d.Parent[rootX] = rootY
		} else { //same height, pick one to be the child of the other, but need to update the rank
			d.Parent[rootY] = rootX
			d.Rank[rootX] += 1
		}
	}
}

func (d DisjointSet) Connected(x int, y int) bool {
	return d.Find(x) == d.Find(y)
}
