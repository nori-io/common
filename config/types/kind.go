package types

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int32
	Int64
	UInt
	UInt32
	UInt64
	Float
	Map
	Slice
	String
)
