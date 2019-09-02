package utils

func Parse(code string) string {
	code = `[{"nodeType":"Stmt_Expression","expr":{"nodeType":"Expr_Assign","var":{"nodeType":"Expr_Variable","name":"a","attributes":{"startLine":3,"endLine":3}},"expr":{"nodeType":"Expr_BinaryOp_Plus","left":{"nodeType":"Scalar_LNumber","value":1,"attributes":{"startLine":3,"endLine":3,"kind":10}},"right":{"nodeType":"Scalar_LNumber","value":2,"attributes":{"startLine":3,"endLine":3,"kind":10}},"attributes":{"startLine":3,"endLine":3}},"attributes":{"startLine":3,"endLine":3}},"attributes":{"startLine":3,"endLine":3}}]`
	return code
	// code := "package main"
	// for _, item := range tokenArr {
	// 	switch item["nodeType"] {
	// 	case "Stmt_Expression":
	// 		code = Stmt_Expression(code, item["expr"].(map[string]interface{}))
	// 	// case "Stmt_Class":
	// 	default:
	// 		fmt.Println(reflect.TypeOf(item["nodeType"]))
	// 	}
	// 	fmt.Println(reflect.TypeOf(item["nodeType"]))
	// 	// switch temp := item.(type) {
	// 	// case []interface{}:
	// 	// 	fmt.Printf("[]interface %+v\n", temp)
	// 	// case string:
	// 	// 	fmt.Printf("string %+v\n", temp)
	// 	// case float64:
	// 	// 	fmt.Printf("float64 %+v\n", temp)
	// 	// default:
	// 	// 	fmt.Printf("default %+v\n", temp)
	// 	// }
	// }
}
