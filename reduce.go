package stream

// implement AccumulatingSink, TerminalSink
type ReducingSink struct {
	empty    bool
	state    interface{}
	operator BiFunction
}

func NewReducingSink(operator BiFunction) *ReducingSink {
	return &ReducingSink{
		operator: operator,
	}
}

func (this *ReducingSink) Accept(t interface{}) {
	if this.empty {
		this.empty = false
		this.state = t
	} else {
		this.state = this.operator(this.state, t)
	}
}

func (this *ReducingSink) Begin(size int64) {
	this.empty = true
	this.state = nil
}

func (this *ReducingSink) End() {
	// nothing
}

func (this *ReducingSink) CancellationRequested() bool {
	return false
}

func (this *ReducingSink) GetResult() interface{} {
	return this.state
}

func (this *ReducingSink) Combine(other interface{}) {
	a, ok := other.(*ReducingSink)
	if !ok {
		panic(ClassCastException(""))
	}

	if !a.empty {
		this.Accept(a.state)
	}
}
