package go_call_name

import "strconv"

// go mod download "module name"
func CallName(name string, age int) interface{} {
	return "Hello " + name + " are you already " + strconv.Itoa(int(age)) + "?"
}