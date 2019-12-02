package stream

type RecordReader interface{
	Next() bool
	GetCurrentValue() interface{}
	GetCurrentKey() interface{}
}
