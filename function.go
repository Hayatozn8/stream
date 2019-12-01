package stream

type Predicate func(t interface{}) bool                        //test
type Function func(t interface{}) interface{}                  //apply
type Consumer func(t interface{})                              //accept
type Comparator func(o1 interface{}, o2 interface{}) int       //compare
type BiFunction func(t interface{}, u interface{}) interface{} //apply
