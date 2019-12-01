package stream

// FilterPipeline
// extend ReferencePipeline implement PipelineStage
type FilterPipeline struct {
	*ReferencePipeline
	predicate Predicate
}

func NewFilterPipeline(previous PipelineStage, predicate Predicate) PipelineStage {
	p := &FilterPipeline{
		ReferencePipeline: NewOpPipeline(previous),
		predicate:         predicate,
	}

	p.previousStage.SetNextStage(p)

	return p
}

func (this *FilterPipeline) OpWrapSink(flags int, sink Sink) Sink {
	return NewFilterSink(sink, this.predicate)
}

func (this *FilterPipeline) OpIsStateful() bool {
	return false
}

// FilterSink: implements Sink
type FilterSink struct {
	*ChainedReference
	predicate Predicate
}

func NewFilterSink(downstream Sink, predicate Predicate) Sink {
	return &FilterSink{
		ChainedReference: NewChainedReference(downstream),
		predicate:        predicate,
	}
}

func (this *FilterSink) Begin(size int64) {
	this.downstream.Begin(-1)
}

func (this *FilterSink) Accept(u interface{}) {
	if this.predicate(u) {
		this.downstream.Accept(u)
	}
}
