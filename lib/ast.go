package lib

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Ast struct {
	Lists []Node
	code  string
}

type Node struct {
	NodeType   string              `json:"nodeType,omitempty"`
	Expr       *StmtExpressionExpr `json:"expr,omitempty"`
	Attributes Attributes          `json:"attributes,omitempty"`
}

type Attributes struct {
	StartLine int
	EndLine   int
	Kind      int
}

const STMT_EXPRESSION = "Stmt_Expression"
const STMT_CLASS = "Stmt_Class"

func NewAst(code string) *Ast {
	var lists []Node
	json.Unmarshal([]byte(code), &lists)
	return &Ast{
		Lists: lists,
	}
}

func (ast Ast) Convert() string {
	for _, node := range ast.Lists {
		switch node.NodeType {
		case STMT_EXPRESSION:
			ast.code += node.Expr.Index()
		// case STMT_CLASS:
		// 	ast.code = SClass(ast.code, item["expr"].(map[string]interface{}))
		default:
			fmt.Println(reflect.TypeOf(node.NodeType))
		}
	}
	return ast.code
}
