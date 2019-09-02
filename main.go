package main

import (
	"fmt"
	"os"
	"p2g/lib"
	"p2g/utils"
)

func main() {
	// file := os.Args[1]
	file := "/Users/wangkun/localhost/test/test.php"
	if !utils.IsExist(file) {
		fmt.Println("File is not exists")
		os.Exit(0)
	}
	if !utils.IsFile(file) {
		fmt.Println("Does not support directory")
		os.Exit(0)
	}

	code, err := utils.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	astJson := utils.Parse(code)
	astJson = `[{"nodeType":"Stmt_Expression","expr":{"nodeType":"Expr_Assign","var":{"nodeType":"Expr_Variable","name":"a","attributes":{"startLine":3,"endLine":3}},"expr":{"nodeType":"Expr_FuncCall","name":{"nodeType":"Name","parts":["aa"],"attributes":{"startLine":3,"endLine":3}},"args":[{"nodeType":"Arg","value":{"nodeType":"Scalar_LNumber","value":1,"attributes":{"startLine":3,"endLine":3,"kind":10}},"byRef":false,"unpack":false,"attributes":{"startLine":3,"endLine":3}},{"nodeType":"Arg","value":{"nodeType":"Scalar_LNumber","value":2,"attributes":{"startLine":3,"endLine":3,"kind":10}},"byRef":false,"unpack":false,"attributes":{"startLine":3,"endLine":3}},{"nodeType":"Arg","value":{"nodeType":"Scalar_String","value":"a","attributes":{"startLine":3,"endLine":3,"kind":1}},"byRef":false,"unpack":false,"attributes":{"startLine":3,"endLine":3}},{"nodeType":"Arg","value":{"nodeType":"Expr_ConstFetch","name":{"nodeType":"Name","parts":["false"],"attributes":{"startLine":3,"endLine":3}},"attributes":{"startLine":3,"endLine":3}},"byRef":false,"unpack":false,"attributes":{"startLine":3,"endLine":3}}],"attributes":{"startLine":3,"endLine":3}},"attributes":{"startLine":3,"endLine":3}},"attributes":{"startLine":3,"endLine":3}}]`

	ast := lib.NewAst(astJson)
	gcode := ast.Convert()
	fmt.Printf("%+v\n", gcode)
}
