package ints

import (
	"sort"
	"unsafe"

	"github.com/ldeng7/go-x/structs"
)

type UndirectedGraph struct {
	nEdges int
	adj    map[*GraphVertex]map[*GraphVertex]geWeightType
}

func (g *UndirectedGraph) Init(vertexes []*GraphVertex, edges []GraphEdge) *UndirectedGraph {
	nEdges := 0
	adj := map[*GraphVertex]map[*GraphVertex]geWeightType{}
	for _, u := range vertexes {
		adj[u] = map[*GraphVertex]geWeightType{}
	}

	for _, e := range edges {
		u, v, w := e.VertexU, e.VertexV, e.Weight
		mu, mv := adj[u], adj[v]
		if nil == mu || nil == mv {
			continue
		}
		if _, ok := mu[v]; !ok {
			nEdges++
		}
		//mu[v], mv[u] = w, w
		mu[v] = w
	}
	g.nEdges = nEdges
	g.adj = adj
	return g
}

func (g *UndirectedGraph) ContainsVertex(u *GraphVertex) bool {
	_, ok := g.adj[u]
	return ok
}

func (g *UndirectedGraph) Vertexes() []*GraphVertex {
	ret := make([]*GraphVertex, 0, len(g.adj))
	for u := range g.adj {
		ret = append(ret, u)
	}
	return ret
}

func (g *UndirectedGraph) Edges() []GraphEdge {
	ret := make([]GraphEdge, 0, g.nEdges)
	for u, m := range g.adj {
		pu := uintptr(unsafe.Pointer(u))
		for v, w := range m {
			if pv := uintptr(unsafe.Pointer(v)); pu <= pv {
				ret = append(ret, GraphEdge{u, v, w})
			}
		}
	}
	return ret
}

func (g *UndirectedGraph) MinimumSpanningTreeEdgesByKruskal() ([]GraphEdge, bool) {
	nVertex := len(g.adj)
	indices := make(map[*GraphVertex]int, nVertex)
	for u := range g.adj {
		indices[u] = len(indices)
	}
	edges := g.Edges()
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	au := (&structs.ArrayUnion{}).Init(nVertex)
	edgesMin := make([]GraphEdge, nVertex-1)
	nEdgeMin := 0
	for _, e := range edges {
		u, v := e.VertexU, e.VertexV
		if !au.Merge(indices[u], indices[v]) {
			continue
		}
		edgesMin[nEdgeMin] = GraphEdge{u, v, e.Weight}
		nEdgeMin++
		if nEdgeMin == nVertex-1 {
			break
		}
	}
	return edgesMin, nEdgeMin == nVertex-1
}

func (g *UndirectedGraph) MinimumPathWeightSumsByDijkstra(src *GraphVertex) map[*GraphVertex]int {
	if !g.ContainsVertex(src) {
		return nil
	}

	nVertex := len(g.adj)
	indices := make(map[*GraphVertex]int, nVertex)
	vertexes := make([]*GraphVertex, nVertex)
	i := 0
	for u := range g.adj {
		indices[u], vertexes[i] = i, u
		i++
	}

	sums := make([]int, nVertex)
	for i := 0; i < nVertex; i++ {
		sums[i] = -1
	}
	iSrc := indices[src]
	sums[iSrc] = 0
	q := (&PriorityQueue{}).Init(make([]int, 0, nVertex*2), true, func(i, j int) bool {
		return sums[i] < sums[j]
	})
	q.Push(iSrc)

	for 0 != q.Len() {
		k := *(q.Pop())
		sumK := sums[k]
		for u, w := range g.adj[vertexes[k]] {
			i := indices[u]
			if s, sumI := sumK+w, sums[i]; s < sumI || sumI == -1 {
				sums[i] = s
				q.Push(i)
			}
		}
	}

	ret := make(map[*GraphVertex]int, nVertex-1)
	for i, s := range sums {
		if s != -1 && i != iSrc {
			ret[vertexes[i]] = s
		}
	}
	return ret
}
