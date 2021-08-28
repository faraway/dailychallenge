package utils

type DisjointSet []int

func NewDisjointSet(size int) DisjointSet {
	ds := make([]int, size)
	for i,_ := range ds {
		ds[i] = i
	}
	return ds
}

func (d DisjointSet) Find(x int) int {
	for d[x] != x {
		x = d[x]
	}
	return x
}

func (d DisjointSet) Union(x int, y int) {
	rootX := d.Find(x)
	rootY := d.Find(y)
	if rootX != rootY {
		d[rootY] = rootX
	}
}

func (d DisjointSet) Connected(x int, y int) bool {
	return d.Find(x) == d.Find(y)
}
