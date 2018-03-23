package lcs

import (
	"log"
	"sort"
)

var ZERO_POINT = &Point{0, 0}

type NormalLCS struct {
	FastLCS
	lcsEdges          [][][2]*Point
	equivalentEdges   [][][2]*Point
	commonVerticeRows []int
	commonVerticeCols []int
	commonRowSet      map[int]bool
	commonColSet      map[int]bool
	commonVertices    []*Point
}

func NewNormalLCS(model, sample []interface{}) *NormalLCS {
	log.SetFlags(log.Ldate | log.Lshortfile)
	normalLCS := new(NormalLCS)
	normalLCS.model = model
	normalLCS.sample = sample
	normalLCS.init()
	normalLCS.lcsEdges = make([][][2]*Point, normalLCS.rows)
	for i := range normalLCS.lcsEdges {
		normalLCS.lcsEdges[i] = make([][2]*Point, normalLCS.columns)
	}
	normalLCS.equivalentEdges = make([][][2]*Point, normalLCS.rows)
	for i := range normalLCS.equivalentEdges {
		normalLCS.equivalentEdges[i] = make([][2]*Point, normalLCS.columns)
	}
	normalLCS.commonRowSet = map[int]bool{}
	normalLCS.commonColSet = map[int]bool{}
	return normalLCS
}

func NewNormalLCSString(model, sample string) *NormalLCS {
	return NewNormalLCS(convertString(model), convertString(sample))
}

func (lcs *NormalLCS) WhenDiagonalEquals(row, column int) {
	top := lcs.lengthMatrix[row][column-1]
	left := lcs.lengthMatrix[row-1][column]
	lcs.commonRowSet[row] = true
	lcs.commonColSet[column] = true
	lcs.lcsEdges[row][column][0] = &Point{row - 1, column - 1}
	lcs.commonVertices = append(lcs.commonVertices, &Point{row, column})
	if lcs.lengthMatrix[row][column] == top {
		lcs.equivalentEdges[row][column][0] = &Point{row, column - 1}
	}
	if lcs.lengthMatrix[row][column] == left {
		lcs.equivalentEdges[row][column][1] = &Point{row - 1, column}
	}
}

func (lcs *NormalLCS) WhenLeftGreaterThanTop(row, column int) {
	lcs.lcsEdges[row][column][0] = &Point{row - 1, column}
}

func (lcs *NormalLCS) WhenTopGreaterThanLeft(row, column int) {
	lcs.lcsEdges[row][column][0] = &Point{row, column - 1}
}

func (lcs *NormalLCS) WhenTopEqualsLeft(row, column int) {
	lcs.lcsEdges[row][column][0] = &Point{row, column - 1}
	lcs.lcsEdges[row][column][1] = &Point{row - 1, column}
}

func findEndIndex(slice []int, start, end int) int {

	if !sort.IntsAreSorted(slice) {
		sort.Ints(slice)
	}
	endIdx := sort.SearchInts(slice, end)
	if endIdx == 0 && end < slice[0] {
		return -1
	}
	if endIdx == len(slice) {
		endIdx = endIdx - 1
	}
	if slice[endIdx] >= start {
		return slice[endIdx]
	}
	return -1
}

func (lcs *NormalLCS) findNextVertexRowIndex(startPoint, endPoint *Point) int {
	return findEndIndex(lcs.commonVerticeRows, startPoint.row,
		endPoint.row)
}

func (lcs *NormalLCS) findNextVertexColIndex(startPoint, endPoint *Point) int {
	return findEndIndex(lcs.commonVerticeCols, startPoint.column,
		endPoint.column)
}

func (this *NormalLCS) hasEdges(vertex *Point) bool {
	return (this.lcsEdges[vertex.row][vertex.column][0] != nil || this.lcsEdges[vertex.row][vertex.column][1] != nil)
}

func (this *NormalLCS) traverseEdges(points [2]*Point) map[string]*LcsPair {
	var candidates = map[string]*LcsPair{}
	for _, p := range points {
		if p != nil {
			nextVertex := this.findNextVertex(ZERO_POINT, p)
			for k, p1 := range this.findAllLcsPairs(nextVertex) {
				candidates[k] = p1
			}
		}
	}
	return candidates
}

func (lcs *NormalLCS) findNextVertex(startVertex, endVertex *Point) *Point {
	nextPoint := &Point{-1, -1}
	if lcs.pathDirectionMatrix[endVertex.row][endVertex.column] == 1 {
		return endVertex
	}
	if endVertex.row > endVertex.column {
		nextPoint.column = lcs.findNextVertexColIndex(startVertex, endVertex)
		if nextPoint.column < 0 {
			return nil
		}
		nextPoint.row = lcs.findNextVertexRowIndex(startVertex, endVertex)
	} else {
		nextPoint.row = lcs.findNextVertexRowIndex(startVertex, endVertex)
		if nextPoint.row < 0 {
			return nil
		}
		nextPoint.column = lcs.findNextVertexColIndex(startVertex, endVertex)
	}
	if nextPoint.row < 0 || nextPoint.column < 0 {
		return nil
	}
	return nextPoint
}

func (this *NormalLCS) findAllLcsPairs(vertex *Point) map[string]*LcsPair {
	if vertex == nil || !this.hasEdges(vertex) {
		return map[string]*LcsPair{EMPTY_PAIR.String(): &LcsPair{[]int{}, []int{}}}
	}

	// traverse all normal LCS edges to build LCS
	candidates := this.traverseEdges(this.lcsEdges[vertex.row][vertex.column])
	if this.pathDirectionMatrix[vertex.row][vertex.column] == 1 {
		candidatesClone := map[string]*LcsPair{}
		for _, candidate := range candidates {
			candidate.ModelIndexes = append(candidate.ModelIndexes, vertex.row-1)
			candidate.SampleIndexes = append(candidate.SampleIndexes, vertex.column-1)
			candidatesClone[candidate.String()] = candidate
		}
		candidates = candidatesClone
		// traverse all equivalent edges to get all candidates that are as
		// good as the current LCS
		for k, c := range this.traverseEdges(this.equivalentEdges[vertex.row][vertex.column]) {
			candidates[k] = c
		}
	}
	return candidates
}

func (this *NormalLCS) FindAllLcsPairs() []*LcsPair {
	if !this.dpReady {
		this.ComputeDP(this)
	}
	if this.lengthMatrix[this.rows-1][this.columns-1] == 0 {
		return []*LcsPair{EMPTY_PAIR}
	}
	this.commonVerticeRows = make([]int, 0, len(this.commonRowSet))
	for k := range this.commonRowSet {
		this.commonVerticeRows = append(this.commonVerticeRows, k)
	}
	this.commonVerticeCols = make([]int, 0, len(this.commonColSet))
	sort.Ints(this.commonVerticeRows)
	for k := range this.commonColSet {
		this.commonVerticeCols = append(this.commonVerticeCols, k)
	}
	sort.Ints(this.commonVerticeCols)
	candidates := this.findAllLcsPairs(this.findNextVertex(ZERO_POINT, &Point{this.rows - 1, this.columns - 1}))
	results := make([]*LcsPair, 0, len(candidates))
	for _, candidate := range candidates {
		results = append(results, candidate)
	}
	return results
}
