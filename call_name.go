package go_call_name

import "strconv"

// v2.0.1
func CallName(name string, age int, response string) interface{} {
	return "Hello " + name + " are you already " + strconv.Itoa(int(age)) + "?" + response + " I guess."
}