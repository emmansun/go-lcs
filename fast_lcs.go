package lcs

import (
	"fmt"
	"log"
	"math"
	"reflect"
)

type Point struct {
	row    int
	column int
}

type Element struct {
	Point  //anonymous field
	lcsLen int
}

func newElement(lcsLen, row, column int) *Element {
	el := new(Element)
	el.row = row
	el.column = column
	el.lcsLen = lcsLen
	return el
}

func (pt Point) length() float64 {
	return math.Sqrt(math.Pow(float64(pt.row), float64(2)) + math.Pow(float64(pt.column), float64(2)))
}

func (pt Point) compareTo(otherPt Point) int {
	return int(pt.length() - otherPt.length())
}

func (pt Point) String() string {
	return fmt.Sprintf("(%d,%d)", pt.row, pt.column)
}

type LcsPair struct {
	modelIndexes  []int
	sampleIndexes []int
}

func (pair LcsPair) String() string {
	return fmt.Sprintf("%+v=%+v", pair.modelIndexes, pair.sampleIndexes)
}

var EMPTY_PAIR = &LcsPair{[]int{}, []int{}}

type LCSInterface interface {
	MaxLcsLen() int
	FindAllLcsPairs() []*LcsPair
}

type FastLCS struct {
	model               []interface{}
	sample              []interface{}
	lengthMatrix        [][]int
	pathDirectionMatrix [][]int
	rows                int
	columns             int
	dpReady             bool
}

func NewFastLCS(model, sample []interface{}) *FastLCS {
	log.SetFlags(log.Ldate | log.Lshortfile)
	fastLCS := new(FastLCS)
	fastLCS.model = model
	fastLCS.sample = sample
	fastLCS.init()
	return fastLCS
}

func convertString(s string) []interface{} {
	r := []rune(s)
	result := make([]interface{}, len(r))
	for i, v := range r {
		result[i] = v
	}
	return result
}

func NewFastLCSString(model, sample string) *FastLCS {
	return NewFastLCS(convertString(model), convertString(sample))
}

func (fastLCS *FastLCS) init() {
	fastLCS.dpReady = false
	fastLCS.rows = len(fastLCS.model) + 1
	fastLCS.columns = len(fastLCS.sample) + 1
	fastLCS.lengthMatrix = make([][]int, fastLCS.rows)
	for i := range fastLCS.lengthMatrix {
		fastLCS.lengthMatrix[i] = make([]int, fastLCS.columns)
	}

	fastLCS.pathDirectionMatrix = make([][]int, fastLCS.rows)
	for i := range fastLCS.pathDirectionMatrix {
		fastLCS.pathDirectionMatrix[i] = make([]int, fastLCS.columns)
	}
}

func equal(left, right interface{}) bool {
	return reflect.DeepEqual(left, right) //hope to check better solution
}

type DPCallback interface {
	WhenDiagonalEquals(row, column int)
	WhenLeftGreaterThanTop(row, column int)
	WhenTopGreaterThanLeft(row, column int)
	WhenTopEqualsLeft(row, column int)
}

func (lcs *FastLCS) WhenDiagonalEquals(row, column int) {
}

func (lcs *FastLCS) WhenLeftGreaterThanTop(row, column int) {
}

func (lcs *FastLCS) WhenTopGreaterThanLeft(row, column int) {
}

func (lcs *FastLCS) WhenTopEqualsLeft(row, column int) {
}

//ComputeDP Dynamic programming
func (lcs *FastLCS) ComputeDP(callback DPCallback) {
	if lcs.dpReady {
		return
	}
	lcs.dpReady = true
	for i := 1; i < lcs.rows; i++ {
		for j := 1; j < lcs.columns; j++ {
			switch {
			case (i < lcs.rows && j < lcs.columns && equal(lcs.model[i-1], lcs.sample[j-1])):
				lcs.lengthMatrix[i][j] = lcs.lengthMatrix[i-1][j-1] + 1
				lcs.pathDirectionMatrix[i][j] = 1
				callback.WhenDiagonalEquals(i, j)
				break
			case (lcs.lengthMatrix[i-1][j] > lcs.lengthMatrix[i][j-1]):
				lcs.lengthMatrix[i][j] = lcs.lengthMatrix[i-1][j]
				lcs.pathDirectionMatrix[i][j] = 2
				callback.WhenLeftGreaterThanTop(i, j)
				break
			case (lcs.lengthMatrix[i-1][j] < lcs.lengthMatrix[i][j-1]):
				lcs.lengthMatrix[i][j] = lcs.lengthMatrix[i][j-1]
				lcs.pathDirectionMatrix[i][j] = 3
				callback.WhenTopGreaterThanLeft(i, j)
				break
			default:
				lcs.lengthMatrix[i][j] = lcs.lengthMatrix[i][j-1]
				lcs.pathDirectionMatrix[i][j] = 4
				callback.WhenTopEqualsLeft(i, j)
				break
			}
		}
	}
	log.Printf("modelSize=%d, sampleSize=%d, LCSLen=%d\n", lcs.rows-1, lcs.columns-1, lcs.MaxLcsLen())
}

func pop(stack []*Element) (*Element, []*Element) {
	element := stack[len(stack)-1]
	return element, stack[:len(stack)-1]
}

func reverse(arr []int) []int {
	temp := make([]int, len(arr))
	for i := range arr {
		temp[i] = arr[len(arr)-i-1]
	}
	return temp
}

func (this *FastLCS) createLcsPair(print []*Element) *LcsPair {
	pair := new(LcsPair)
	pair.modelIndexes = make([]int, 0, len(print)-1)
	pair.sampleIndexes = make([]int, 0, len(print)-1)
	for _, v := range print {
		if v.row > 0 && v.row < this.rows {
			pair.modelIndexes = append(pair.modelIndexes, v.row-1)
		}
		if v.column > 0 && v.column < this.columns {
			pair.sampleIndexes = append(pair.sampleIndexes, v.column-1)
		}
	}
	pair.modelIndexes = reverse(pair.modelIndexes)
	pair.sampleIndexes = reverse(pair.sampleIndexes)
	return pair
}

func (this *FastLCS) updateToMostTop(point *Point) {
	k := point.column - 1
	pointCol := point.column
	for {
		if this.lengthMatrix[point.row][k] == this.lengthMatrix[point.row][point.column] {
			if this.pathDirectionMatrix[point.row][k] == 1 {
				pointCol = k
			}
			k--
		} else {
			break
		}
	}
	point.column = pointCol
}

func (this *FastLCS) updateToMostLeft(point *Point) {
	k := point.row - 1
	pointRow := point.row
	for {
		if this.lengthMatrix[k][point.column] == this.lengthMatrix[point.row][point.column] {
			if this.pathDirectionMatrix[k][point.column] == 1 {
				pointRow = k
			}
			k--
		} else {
			break
		}
	}
	point.row = pointRow
}

func (this *FastLCS) searchElement(pt *Point, row, column, stype int) {
	switch this.pathDirectionMatrix[row][column] {
	case 1:
		pt.column = column
		pt.row = row
		// Emman, if there are no below process, we can NOT get complete candidate set
		if row > 0 && column > 0 {
			if stype == 0 {
				this.updateToMostLeft(pt)
			} else {
				this.updateToMostTop(pt)
			}
		}
		break
	case 2:
		this.searchElement(pt, row-1, column, stype)
		break
	case 3:
		this.searchElement(pt, row, column-1, stype)
		break
	case 4:
		if stype == 0 {
			this.searchElement(pt, row-1, column, stype)
		} else {
			this.searchElement(pt, row, column-1, stype)
		}
		break
	}
}

func (this *FastLCS) addAllJumpElements(store []*Element, leftBottomPoint, rightTopPoint *Point) []*Element {
	currentLcsLen := 0
	for i := rightTopPoint.row; i >= leftBottomPoint.row; i-- {
		for j := leftBottomPoint.column; j >= rightTopPoint.column; j-- {
			// Emman, check lcs len to avoid mix different lcs len cases.
			if this.pathDirectionMatrix[i][j] <= 1 && this.lengthMatrix[i][j] >= currentLcsLen {
				e := newElement(this.lengthMatrix[i][j], i, j)
				store = append(store, e)
				currentLcsLen = this.lengthMatrix[i][j]
			}
		}
	}
	return store
}

func (this *FastLCS) MaxLcsLen() int {
	if !this.dpReady {
		this.ComputeDP(this)
	}
	return this.lengthMatrix[this.rows-1][this.columns-1]
}

func (this *FastLCS) FindAllLcsPairs() []*LcsPair {
	if !this.dpReady {
		this.ComputeDP(this)
	}
	if this.lengthMatrix[this.rows-1][this.columns-1] == 0 {
		return []*LcsPair{EMPTY_PAIR}
	}
	var store, print []*Element
	print = make([]*Element, 0, this.lengthMatrix[this.rows-1][this.columns-1]+1)
	store = make([]*Element, 0, this.rows+this.columns)
	var storeTop *Element
	var results []*LcsPair
	virtualNode := newElement(this.lengthMatrix[this.rows-1][this.columns-1]+1, this.rows, this.columns)
	store = append(store, virtualNode)
	for len(store) > 0 {
		storeTop, store = pop(store)
		if storeTop.row <= 1 || storeTop.column <= 1 {
			if storeTop.row > 0 && storeTop.column > 0 {
				print = append(print, storeTop)
			}
			if len(print) == this.lengthMatrix[this.rows-1][this.columns-1]+1 {
				results = append(results, this.createLcsPair(print))
			}
			if len(store) > 0 {
				e := store[len(store)-1]
				for len(print) > 0 {
					ep := print[len(print)-1]
					if ep.lcsLen <= e.lcsLen {
						print = print[:len(print)-1]
					} else {
						break
					}
				}
			}
		} else {
			print = append(print, storeTop)
			leftBottomPoint := new(Point)
			this.searchElement(leftBottomPoint, storeTop.row-1, storeTop.column-1, 0)
			rightTopPoint := new(Point)
			this.searchElement(rightTopPoint, storeTop.row-1, storeTop.column-1, 1)
			store = this.addAllJumpElements(store, leftBottomPoint, rightTopPoint)
		}
	}
	return results
}
