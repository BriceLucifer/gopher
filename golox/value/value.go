package value

type Value uint64

type ValueArray struct {
	count    int
	capacity int
	values   *Value
}
