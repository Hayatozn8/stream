package stream

const (
	ORDERED    = 0x00000010
	DISTINCT   = 0x00000001
	SORTED     = 0x00000004
	SIZED      = 0x00000040
	NONNULL    = 0x00000100
	IMMUTABLE  = 0x00000400
	CONCURRENT = 0x00001000
	SUBSIZED   = 0x00004000
)

type Spliterator interface {
	TryAdvance(action Consumer) bool
	ForEachRemaining(action Consumer)
	TrySplit() Spliterator
	EstimateSize() int64
	GetExactSizeIfKnown() int64
	Characteristics() int
	HasCharacteristics(characteristics int) bool
	GetComparator() Comparator
}
