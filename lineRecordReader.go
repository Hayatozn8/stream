package stream

import (
	"bufio"
	"os"
	"io"
)

type LineRecordReader struct{
	path string
	reader *bufio.Reader
	value string
	isEnd bool
}

func NewLineRecordReader(path string) (RecordReader, error){
	f, err := os.Open(path)

	if err != nil{
		return nil, err
	}

	reader := bufio.NewReader(f)

	return &LineRecordReader{
		path:path,
		reader:reader,
		isEnd:false,
	}, nil
}

func(this *LineRecordReader)Next() bool{
	if this.isEnd {
		return false
	}

	s, err:=this.reader.ReadString('\n')
	if err != nil{
		if err == io.EOF{
			this.isEnd = true
			return false
		}else {
			panic(err)
		}
	}

	this.value = s
	return true
}
func(this *LineRecordReader)GetCurrentValue() interface{}{
	return this.value
}

func(this *LineRecordReader)GetCurrentKey() interface{}{
	return ""
}
