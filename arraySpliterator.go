package stream

// ArraySpliterator
type ArraySpliterator struct {
	slice           []interface{}
	index           int
	fence           int
	characteristics int
}

func NewArraySpliterator(slice []interface{}, index int, fence int, additionalCharacteristics int) Spliterator {
	requireNonNull(slice)
	return &ArraySpliterator{
		slice:           slice,
		index:           index,
		fence:           fence,
		characteristics: additionalCharacteristics | SIZED | SUBSIZED,
	}
}

func (this *ArraySpliterator) TryAdvance(action Consumer) bool {
	// requireNonNull(action)
	if this.index >= 0 && this.index < this.fence {
		action(this.slice[this.index])
		this.index++
		return true
	}
	return false
}

func (this *ArraySpliterator) ForEachRemaining(action Consumer) {
	// requireNonNull(action)

	s := this.slice
	hi := this.fence // hoist accesses and checks from loop
	i := this.index
	this.index = hi

	if len(s) >= hi && i >= 0 && i < hi {
		for j := i; j < hi; j++ {
			action(s[j])
		}
	}
}

func (this *ArraySpliterator) TrySplit() Spliterator {
	lo := this.index
	mid := (lo + this.fence) >> 1

	if lo >= mid {
		return nil
	} else {
		// 0 1 2 3 4 5 6 7 8 9 10  index=0, fence=10
		// this: 5 6 7 8 9 10
		// new: 0 1 2 3 4 5
		this.index = mid
		return NewArraySpliterator(this.slice, lo, mid, this.characteristics)
	}
}

func (this *ArraySpliterator) EstimateSize() int64 {
	return int64(this.fence - this.index)
}

func (this *ArraySpliterator) GetExactSizeIfKnown() int64 {
	if (this.characteristics & SIZED) == 0 {
		return -1
	} else {
		return this.EstimateSize()
	}
}

func (this *ArraySpliterator) Characteristics() int {
	return this.characteristics
}

func (this *ArraySpliterator) HasCharacteristics(characteristics int) bool {
	return (this.characteristics & characteristics) == characteristics
}

func (this *ArraySpliterator) GetComparator() Comparator {
	if this.HasCharacteristics(SORTED) {
		return nil
	}
	panic(IllegalStateException(""))
}
