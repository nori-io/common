package config

type (
	//go:generate mockgen -destination=../mocks/registry/config_config.go -package=mocks github.com/nori-io/common/v4/pkg/domain/config Config
	Config interface {
		Bool(key string, desc string) Bool
		Float(key string, desc string) Float
		Int(key string, desc string) Int
		Int8(key string, desc string) Int8
		Int32(key string, desc string) Int32
		Int64(key string, desc string) Int64
		UInt(key string, desc string) UInt
		UInt32(key string, desc string) UInt32
		UInt64(key string, desc string) UInt64
		Slice(key string, desc string) Slice
		SliceInt(key string, desc string) SliceInt
		SliceString(key string, desc string) SliceString
		String(key string, desc string) String
		StringMap(key string, desc string) StringMap
		StringMapInt(key string, desc string) StringMapInt
		StringMapSliceString(key string, desc string) StringMapSliceString
		StringMapString(key string, desc string) StringMapString

		SetDefault(key string, val interface{})
	}

	Bool                 func() bool
	Float                func() float64
	Int                  func() int
	Int8                 func() int8
	Int32                func() int32
	Int64                func() int64
	UInt                 func() uint
	UInt32               func() uint32
	UInt64               func() uint64
	Slice                func() []interface{}
	SliceInt             func() []int
	SliceString          func() []string
	String               func() string
	StringMap            func() map[string]interface{}
	StringMapInt         func() map[string]int
	StringMapSliceString func() map[string][]string
	StringMapString      func() map[string]string
)
