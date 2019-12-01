package stream

const (
	MSG_STREAM_LINKED = "stream has already been operated upon or closed"
	MSG_CONSUMED      = "source already consumed or closed"
)

//abstract class referencePipeline implement PipelineStage
// Subclass must implement:
// 	OpWrapSink
// 	OpIsStateful
type ReferencePipeline struct {
	previousStage   PipelineStage
	nextStage       PipelineStage
	sourceOrOpFlags int
	depth           int
	combinedFlags   int
	// Supplier<? extends Spliterator<?>> sourceSupplier
	linkedOrConsumed  bool
	sourceAnyStateful bool
	// sourceCloseAction Runnable
	parallel bool
}

func NewOpPipeline(previousStage PipelineStage) *ReferencePipeline {
	p := &ReferencePipeline{
		previousStage: previousStage,
	}

	p.depth = p.previousStage.GetDepth() + 1
	return p
}

func NewHeadPipeline() PipelineStage {
	return &ReferencePipeline{}
}

func (this *ReferencePipeline) GetNextStage() PipelineStage {
	return this.nextStage
}

func (this *ReferencePipeline) SetNextStage(next PipelineStage) {
	this.nextStage = next
}

func (this *ReferencePipeline) GetPreviousStage() PipelineStage {
	return this.previousStage
}

func (this *ReferencePipeline) SetPreviousStage(previous PipelineStage) {
	this.previousStage = previous
}

// TODO
// func (this *ReferencePipeline) isParallel() bool {
// 	return this.parallel
// }

func (this *ReferencePipeline) GetDepth() int {
	return this.depth
}

func (this *ReferencePipeline) GetCombinedFlags() int {
	return this.combinedFlags
}

func (this *ReferencePipeline) OpWrapSink(flags int, sink Sink) Sink {
	panic(UnsupportedOperationException(""))
}

func (this *ReferencePipeline) OpIsStateful() bool {
	panic(UnsupportedOperationException(""))
}

// func (this *ReferencePipeline) Evaluate(terminalOp TerminalOp) interface{} {
// 	// assert getOutputShape() == terminalOp.inputShape();
// 	if this.linkedOrConsumed {
// 		panic(IllegalStateException(MSG_STREAM_LINKED))
// 	}

// 	this.linkedOrConsumed = true

// 	// return isParallel()
// 	// 	   ? terminalOp.evaluateParallel(this, sourceSpliterator(terminalOp.getOpFlags()))
// 	// 	   : terminalOp.evaluateSequential(this, sourceSpliterator(terminalOp.getOpFlags()));
// }
