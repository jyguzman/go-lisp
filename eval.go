package main

// import (
// 	"fmt"
// 	"os"
// )

// func evalList(exp []LispValue) any {
// 	firstArg := exp[0]
// 	t, val := firstArg.Type, firstArg.Val
// 	switch {
// 	case t == LOperator:
// 		return nil
// 	case t == LInteger || t == LFloat || t == LString || t == LBoolean || t == LNil || t == LSymbol:
// 		return val
// 	case t == LList:
// 		return eval(val)
// 	case t == LLambda:
// 		params, body :=
// 	default:
// 		return nil
// 	}
// }

// func eval(exp LispValue) any {
// 	return nil
// }

// // func plus(exp []interface{}) any {
// // 	firstArg := exp[1]
// // 	secondArg := exp[2]

// // }

// // func lambda(exp []interface{}) Lambda {
// // 	params = exp[0]
// // 	body = exp[1]
// // 	return Lambda{[]LispValue{}, []LispValue{}}
// // }

// func evalLambda(lambda Lambda, args []LispValue) any {
// 	params := lambda.Params
// 	body := lambda.Body

// 	if len(params) != len(args) {
// 		fmt.Println("Not enough arguments for function")
// 		os.Exit(1)
// 	}

// 	fnScope := make(map[interface{}]LispValue)
// 	for i := 0; i < len(args); i++ {
// 		fnScope[params[i].Val] = args[i]
// 	}

// 	for i := 0; i < len(args); i++ {
// 		arg, isParamInBody := fnScope[body[i].Val]
// 		if isParamInBody {
// 			body[i] = arg
// 		}
// 	}

// 	return nil
// }
