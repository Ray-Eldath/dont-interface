package main

func test(inter1 interface{}, inter2 ...interface{}) (interface{}, string) {
	var inter3 interface{}
	var s string

	return inter3, s
}

func test2(str string, int int32, inter4 []interface{}) []interface{} { return nil }

type test3 struct {
	i      int
	inter5 interface{}
	inter6 []interface{}
}

type I interface {
	test4(f float64, inter7 interface{}, inter8 ...interface{}) (iii interface{}, sss string)
}
