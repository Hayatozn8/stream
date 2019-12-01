package stream

type Stream interface {
	Filter(predicate Predicate) Stream
	Map(mapper Function) Stream
	Reduce(accumulator BiFunction) interface{} //TODO
}

//skip spliterator
//TODO: parallel bool
func NewStream() Stream {
	return &StreamHelper{}
}

type StreamHelper struct {
	head        PipelineStage
	operators   PipelineStage //like LinkedList not ArrayList
	spliterator Spliterator
}

func (this *StreamHelper) Filter(predicate Predicate) Stream {
	// requireNonNull(predicate)

	this.operators = NewFilterPipeline(this.operators, predicate)

	return this
}
func (this *StreamHelper) Map(mapper Function) Stream {
	// requireNonNull(mapper)

	this.operators = NewMapPipeline(this.operators, mapper)

	return this
}
func (this *StreamHelper) Reduce(accumulator BiFunction) interface{} {
	//TODO
	// terminalSink := NewReducingSink(accumulator)
	// this.wrapAndCopyInto(terminalSink, this.spliterator)
	// return terminalSink.GetResult()
	return this.evaluate(NewReducingSink(accumulator))
}

func (this *StreamHelper) wrapSink(sink Sink) Sink {
	requireNonNull(sink)

	for p := this.operators; p.GetDepth() > 0; p = p.GetPreviousStage() {
		sink = p.OpWrapSink(p.GetPreviousStage().GetCombinedFlags(), sink)
	}
	return sink
}

// TODO param
// TODO if check
func (this *StreamHelper) copyInto(wrappedSink Sink, spliterator Spliterator) {
	requireNonNull(wrappedSink)

	wrappedSink.Begin(spliterator.GetExactSizeIfKnown())
	spliterator.ForEachRemaining(wrappedSink.Accept)
	wrappedSink.End()
}

func (this *StreamHelper) wrapAndCopyInto(sink Sink, spliterator Spliterator) {
	requireNonNull(sink)
	this.copyInto(this.wrapSink(sink), spliterator)
}

func (this *StreamHelper) evaluate(terminalSink TerminalSink) interface{} {
	this.wrapAndCopyInto(terminalSink, this.spliterator)
	return terminalSink.GetResult()
}

// IntSliceStream
// type IntSliceStream struct {
// }

func OfArray(array []interface{}) Stream {
	s := &StreamHelper{
		spliterator: NewArraySpliterator(array, 0, len(array), 0),
		head:        NewHeadPipeline(),
	}
	s.operators = s.head
	return s
}
