package ints

type gvValType = int
type geWeightType = int

type GraphVertex struct {
	Value gvValType
}

type GraphEdge struct {
	VertexU *GraphVertex
	VertexV *GraphVertex
	Weight  geWeightType
}
