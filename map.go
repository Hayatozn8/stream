package stream

// MapPipeline
// extend ReferencePipeline implement PipelineStage
type MapPipeline struct {
	*ReferencePipeline
	mapper Function
}

func NewMapPipeline(previous PipelineStage, mapper Function) PipelineStage {
	p := &MapPipeline{
		ReferencePipeline: NewOpPipeline(previous),
		mapper:            mapper,
	}

	p.previousStage.SetNextStage(p)

	return p
}

func (this *MapPipeline) OpWrapSink(flags int, sink Sink) Sink {
	return NewMapSink(sink, this.mapper)
}

func (this *MapPipeline) OpIsStateful() bool {
	return false
}

// Map Sink
// implements Sink
type MapSink struct {
	*ChainedReference
	mapper Function
}

func NewMapSink(downstream Sink, mapper Function) Sink {
	return &MapSink{
		ChainedReference: NewChainedReference(downstream),
		mapper:           mapper,
	}
}

func (this *MapSink) Accept(u interface{}) {
	this.downstream.Accept(this.mapper(u))
}
