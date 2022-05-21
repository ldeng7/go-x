package algo

import (
	"sort"
	"unsafe"

	"github.com/ldeng7/go-x/collectionx"
	"github.com/ldeng7/go-x/common"
)

type UndirectedGraph[V any, W common.Number] struct {
	nEdges int
	adj    map[*V]map[*V]W
}

func (g *UndirectedGraph[V, W]) Init(vertexes []*V, edges []GraphEdge[V, W]) *UndirectedGraph[V, W] {
	nEdges := 0
	adj := map[*V]map[*V]W{}
	for _, u := range vertexes {
		adj[u] = map[*V]W{}
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

func (g *UndirectedGraph[V, W]) ContainsVertex(u *V) bool {
	_, ok := g.adj[u]
	return ok
}

func (g *UndirectedGraph[V, W]) Vertexes() []*V {
	ret := make([]*V, 0, len(g.adj))
	for u := range g.adj {
		ret = append(ret, u)
	}
	return ret
}

func (g *UndirectedGraph[V, W]) Edges() []GraphEdge[V, W] {
	ret := make([]GraphEdge[V, W], 0, g.nEdges)
	for u, m := range g.adj {
		pu := uintptr(unsafe.Pointer(u))
		for v, w := range m {
			if pv := uintptr(unsafe.Pointer(v)); pu <= pv {
				ret = append(ret, GraphEdge[V, W]{u, v, w})
			}
		}
	}
	return ret
}

func (g *UndirectedGraph[V, W]) MinimumSpanningTreeEdgesByKruskal() ([]GraphEdge[V, W], bool) {
	nVertex := len(g.adj)
	indices := make(map[*V]int, nVertex)
	for u := range g.adj {
		indices[u] = len(indices)
	}
	edges := g.Edges()
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	au := (&ArrayUnion{}).Init(nVertex)
	edgesMin := make([]GraphEdge[V, W], nVertex-1)
	nEdgeMin := 0
	for _, e := range edges {
		u, v := e.VertexU, e.VertexV
		if !au.Merge(indices[u], indices[v]) {
			continue
		}
		edgesMin[nEdgeMin] = GraphEdge[V, W]{u, v, e.Weight}
		nEdgeMin++
		if nEdgeMin == nVertex-1 {
			break
		}
	}
	return edgesMin, nEdgeMin == nVertex-1
}

func (g *UndirectedGraph[V, W]) MinimumPathWeightSumsByDijkstra(src *V) map[*V]W {
	if !g.ContainsVertex(src) {
		return nil
	}

	nVertex := len(g.adj)
	indices := make(map[*V]int, nVertex)
	vertexes := make([]*V, nVertex)
	i := 0
	for u := range g.adj {
		indices[u], vertexes[i] = i, u
		i++
	}

	sums := make([]W, nVertex)
	visited := make([]bool, nVertex)
	iSrc := indices[src]
	sums[iSrc], visited[iSrc] = 0, true
	h := (&collectionx.BinaryHeap[int]{}).Init(make([]int, 0, nVertex*2), true, func(i, j int) bool {
		return sums[i] < sums[j]
	})
	h.Push(iSrc)

	for h.Len() != 0 {
		k, _ := h.Pop()
		sumK := sums[k]
		for u, w := range g.adj[vertexes[k]] {
			i := indices[u]
			if s := sumK + w; s < sums[i] || !visited[i] {
				sums[i], visited[i] = s, true
				h.Push(i)
			}
		}
	}

	ret := make(map[*V]W, nVertex-1)
	for i, s := range sums {
		if visited[i] && i != iSrc {
			ret[vertexes[i]] = s
		}
	}
	return ret
}
