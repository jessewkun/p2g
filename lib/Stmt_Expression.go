package lib

import (
	"fmt"
	"p2g/utils"
)

type StmtExpressionExpr struct {
	NodeType   string     `json:"nodeType,omitempty"`
	Evar       Evar       `json:"var,omitempty"`
	Expr       Expr       `json:"expr,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

const EXPR_ASSIGN = "Expr_Assign"

const SCALAR_LNUMBER = "Scalar_LNumber"
const SCALAR_DNUMBER = "Scalar_DNumber"
const SCALAR_STRING = "Scalar_String"
const EXPR_CONSTFETCH = "Expr_ConstFetch"
const EXPR_NEW = "Expr_New"
const EXPR_FUNCCALL = "Expr_FuncCall"

const EXPR_BINARYOP_PLUS = "Expr_BinaryOp_Plus"
const EXPR_BINARYOP_MINUS = "Expr_BinaryOp_Minus"
const EXPR_BINARYOP_MUL = "Expr_BinaryOp_Mul"
const EXPR_BINARYOP_DIV = "Expr_BinaryOp_Div"
const EXPR_BINARYOP_MOD = "Expr_BinaryOp_Mod"

const EXPR_ARG = "Arg"

const EXPR_BITWISENOT = "Expr_BitwiseNot"
const EXPR_BINARYOP_BITWISEOR = "Expr_BinaryOp_BitwiseOr"
const EXPR_BINARYOP_BITWISEAND = "Expr_BinaryOp_BitwiseAnd"
const EXPR_BINARYOP_BITWISEXOR = "Expr_BinaryOp_BitwiseXor"
const EXPR_BINARYOP_SHIFTLEFT = "Expr_BinaryOp_ShiftLeft"
const EXPR_BINARYOP_SHIFTRIGHT = "Expr_BinaryOp_ShiftRight"

type Evar struct {
	NodeType   string     `json:"nodeType,omitempty"`
	Name       string     `json:"name,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type Expr struct {
	NodeType   string      `json:"nodeType,omitempty"`
	Name       Name        `json:"name,omitempty"`
	Eexpr      Eexpr       `json:"expr,omitempty"`
	Args       []Arg       `json:"args,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Attributes Attributes  `json:"attributes,omitempty"`
	Left       Left        `json:"left,omitempty"`
	Right      Right       `json:"right,omitempty"`
}

type Eexpr struct {
	NodeType   string      `json:"nodeType,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Attributes Attributes  `json:"attributes,omitempty"`
}

type Arg struct {
	NodeType   string     `json:"nodeType,omitempty"`
	Value      Expr       `json:"value,omitempty"`
	ByRef      bool       `json:"byRef,omitempty"`
	Unpack     bool       `json:"unpack,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type ArgValue struct {
	NodeType   string      `json:"nodeType,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Attributes Attributes  `json:"attributes,omitempty"`
}

type Name struct {
	NodeType   string     `json:"nodeType,omitempty"`
	Parts      []string   `json:"parts,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type Left struct {
	NodeType   string      `json:"nodeType,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Attributes Attributes  `json:"attributes,omitempty"`
}

type Right struct {
	NodeType   string      `json:"nodeType,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Attributes Attributes  `json:"attributes,omitempty"`
}

func (se *StmtExpressionExpr) Index() string {
	code := ""
	variable := ""
	switch se.NodeType {
	case EXPR_ASSIGN:
		variable = se.Evar.Name
	default:
		fmt.Printf("unknoen type %+v\n", se.Evar)
		return code
	}
	s := variable + " := " + se.Expr.Format("")
	return s
}

func (expr Expr) Format(code string) string {
	switch expr.NodeType {
	case SCALAR_LNUMBER:
		fallthrough
	case SCALAR_DNUMBER:
		code += utils.ToString(expr.Value)
	case SCALAR_STRING:
		code += "\"" + utils.ToString(expr.Value) + "\""
	case EXPR_CONSTFETCH:
		code += expr.Name.Parts[0]
	// case EXPR_NEW: // new class
	case EXPR_FUNCCALL: // function call
		code += expr.Name.Parts[0] + "("
		for k, arg := range expr.Args {
			fmt.Printf("%+v\n", arg)
			if arg.ByRef {
				code += "*" + arg.Value.Format(code)
			} else {
				code += arg.Value.Format(code)
			}
			if k+1 < len(expr.Args) {
				code += ", "
			}
		}
		code += ")"
	case EXPR_ARG:
		if e, ok := expr.Value.(Eexpr); ok {
			code += utils.ToString(e.Value)
		}
	case EXPR_BITWISENOT:
		code += " := ~" + utils.ToString(expr.Eexpr.Value)
	case EXPR_BINARYOP_BITWISEOR: // |
		code += utils.ToString(expr.Left.Value) + " | " + utils.ToString(expr.Right.Value)
	case EXPR_BINARYOP_BITWISEAND: // &
		code += utils.ToString(expr.Left.Value) + " & " + utils.ToString(expr.Right.Value)
	case EXPR_BINARYOP_BITWISEXOR: // ^
		code += utils.ToString(expr.Left.Value) + " ^ " + utils.ToString(expr.Right.Value)
	case EXPR_BINARYOP_SHIFTLEFT: // <<
		code += utils.ToString(expr.Left.Value) + " << " + utils.ToString(expr.Right.Value)
	case EXPR_BINARYOP_SHIFTRIGHT: // >>
		code += utils.ToString(expr.Left.Value) + " >> " + utils.ToString(expr.Right.Value)
	case EXPR_BINARYOP_PLUS: // +
		code += utils.ToString(expr.Left.Value) + " + " + utils.ToString(expr.Right.Value)
	case EXPR_BINARYOP_MINUS: // -
		code += utils.ToString(expr.Left.Value) + " - " + utils.ToString(expr.Right.Value)
	case EXPR_BINARYOP_MUL: // *
		code += utils.ToString(expr.Left.Value) + " * " + utils.ToString(expr.Right.Value)
	case EXPR_BINARYOP_DIV: //  /
		code += utils.ToString(expr.Left.Value) + " / " + utils.ToString(expr.Right.Value)
	case EXPR_BINARYOP_MOD: // %
		code += utils.ToString(expr.Left.Value) + " % " + utils.ToString(expr.Right.Value)
	default:
		fmt.Printf("unknoen type %+v\n", expr.NodeType)

	}
	return code
}
