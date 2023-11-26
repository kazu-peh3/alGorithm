package constraints

// Signed は、符号付き整数型を表すインターフェースです。
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned は、符号なし整数型を表すインターフェースです。
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Float は、浮動小数点型を表すインターフェースです。
type Float interface {
	~float32 | ~float64
}

// Integer は、符号付きおよび符号なしの整数型を表すインターフェースです。
type Integer interface {
	Signed | Unsigned
}

// Number は、整数型および浮動小数点型を表すインターフェースです。
type Number interface {
	Integer | Float
}

// Ordered は、順序付け可能な型を表すインターフェースです。
type Ordered interface {
	Number | ~string
}
