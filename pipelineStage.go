package stream

// PipelineStage
type PipelineStage interface {
	// Stream
	// getSourceShape() StreamShape
	// GetStreamAndOpFlags() int
	// exactOutputSizeIfKnown(Spliterator spliterator) long
	// S wrapAndCopyInto(S sink, Spliterator<P_IN> spliterator);
	// copyInto(Sink<P_IN> wrappedSink, Spliterator<P_IN> spliterator)
	// copyIntoWithCancel(wrappedSink Sink, Spliterator<P_IN> spliterator) bool

	// Spliterator<P_OUT> wrapSpliterator(Spliterator<P_IN> spliterator)
	// Node.Builder<P_OUT> makeNodeBuilder(long exactSizeIfKnown,IntFunction<P_OUT[]> generator)
	// Node<P_OUT> evaluate(Spliterator<P_IN> spliterator,boolean flatten,IntFunction<P_OUT[]> generator)
	OpIsStateful() bool
	GetNextStage() PipelineStage
	SetNextStage(next PipelineStage)
	GetPreviousStage() PipelineStage
	SetPreviousStage(previous PipelineStage)
	OpWrapSink(flags int, sink Sink) Sink
	// isParallel() bool
	GetDepth() int
	GetCombinedFlags() int
	// Evaluate(terminalOp TerminalOp) interface{}
}
