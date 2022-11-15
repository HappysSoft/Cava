package model

type DataType int

const (
	Int DataType = iota
	Long
	Float
	Double
	String
	Boolean
	Identifier
)

type node struct {
	isEmpty  bool
	nodeType DataType
	value    interface {
		int | int64 | float64 | float32 | bool | string
	}
	next       *node
	friendNode *node
}
