package ast

import (

)

type Statement interface {
	statement()
}

type Expr interface {
	expr()
}

