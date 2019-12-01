package stream

type Sink interface {
	Accept(t interface{})
	Begin(size int64)
	End()
	CancellationRequested() bool
}

type TerminalSink interface {
	Sink
	GetResult() interface{}
}

type AccumulatingSink interface {
	Combine(other interface{})
}

// DefaultSink
type DefaultSink struct {
}

// abstract ChainedReference class implements Sink
type ChainedReference struct {
	downstream Sink
}

// if downstream is nil, panic
func NewChainedReference(downstream Sink) *ChainedReference {
	requireNonNull(downstream)

	return &ChainedReference{
		downstream: downstream,
	}
}

// func (this *ChainedReference) accept(t interface{}) {
// 	panic(IllegalStateException("called wrong accept method"))
// }

func (this *ChainedReference) Begin(size int64) {
	this.downstream.Begin(size)
}
func (this *ChainedReference) End() {
	this.downstream.End()
}
func (this *ChainedReference) CancellationRequested() bool {
	return this.downstream.CancellationRequested()
}
