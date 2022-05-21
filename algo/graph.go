package algo

import "github.com/ldeng7/go-x/common"

type GraphEdge[V any, W common.Number] struct {
	VertexU *V
	VertexV *V
	Weight  W
}
