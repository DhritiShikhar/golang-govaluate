package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	// simple expression without parameters
	exp1, _ := govaluate.NewEvaluableExpression("86 * 23 + 098 - 35 + 34")
	res, _ := exp1.Evaluate(nil)
	fmt.Println(res)

	// with parameters
	exp2, _ := govaluate.NewEvaluableExpression("a > 0")
	params1 := make(map[string]interface{}, 1)
	params1["a"] = -1
	res2, _ := exp2.Evaluate(params1)
	fmt.Println(res2)

	exp3, _ := govaluate.NewEvaluableExpression("(successful_requests * total_requests)/10 > 9")
	params2 := make(map[string]interface{}, 2)
	params2["successful_requests"] = 5
	params2["total_requests"] = 8
	res3, _ := exp3.Evaluate(params2)
	fmt.Println(res3)

	exp4, _ := govaluate.NewEvaluableExpression("http_response == 'OK'")
	params3 := make(map[string]interface{}, 1)
	params3["http_response"] = "OK"
	res4, _ := exp4.Evaluate(params3)
	fmt.Println(res4)
}
