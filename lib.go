package stream

func requireNonNull(obj interface{}) {
	if obj == nil {
		panic(NullPointerException(""))
	}
}
