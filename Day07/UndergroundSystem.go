package main

type UndergroundSystem struct {
	customer map[int]Node
	out      map[string][]Node2
}

type Node struct {
	place string
	time  int
}

type Node2 struct {
	place string
	time  int
	count int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{customer: make(map[int]Node), out: make(map[string][]Node2)}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.customer[id] = Node{place: stationName, time: t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	node := this.customer[id]
	_, ok := this.out[node.place]
	if !ok {
		this.out[node.place] = []Node2{}
	}

	node2 := this.getNode(node.place, stationName)
	node2.time += t - node.time
	node2.count++
}

func (this *UndergroundSystem) getNode(m, n string) *Node2 {
	for i := range this.out[m] {
		if this.out[m][i].place == n {
			return &this.out[m][i]
		}
	}

	node := Node2{place: n}
	this.out[m] = append(this.out[m], node)
	return &this.out[m][len(this.out[m])-1]
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	node := this.getNode(startStation, endStation)
	return float64(node.time) / float64(node.count)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
