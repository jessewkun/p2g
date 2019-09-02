package lib

import "fmt"

func SClass(code string, node map[string]interface{}) string {
	nodeType, _ := node["nodeType"].(string)
	varr, _ := node["var"].(map[string]interface{})
	variable := ""
	switch nodeType {
	case "Expr_Assign":
		variable = varr["name"].(string)
	default:
		fmt.Printf("unknoen type %+v\n", nodeType)
		return code
	}

	exprExpr, _ := node["expr"].(map[string]interface{})
	switch exprExpr["nodeType"].(string) {
	case "Scalar_LNumber": // int
		fallthrough
	case "Scalar_DNumber": // float
		code += variable + " := " + exprExpr["value"].(string)
	case "Scalar_String": // string
		code += variable + " := \"" + exprExpr["value"].(string) + "\""
	case "Expr_ConstFetch": // true/false/null
		name := exprExpr["name"].(map[string]interface{})
		parts := name["parts"].([]string)
		code += variable + " := " + parts[0]
	case "Expr_New": // new class
	case "Expr_BinaryOp_BitwiseOr": // |
	case "Expr_BinaryOp_Plus": // +
	default:
		fmt.Printf("unknoen type %+v\n", nodeType)

	}
	return ""
}
