package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
)

const (
	// control
	C_RESET      = "\033[0m"
	C_BRIGHT     = "\033[1m"
	C_DIM        = "\033[2m"
	C_UNDERSCORE = "\033[4m"
	C_BLINK      = "\033[5m"
	C_REVERSE    = "\033[7m"
	C_HIDDEN     = "\033[8m"

	C_NULL = ""

	// foreground
	F_BLACK   = "\033[30m"
	F_RED     = "\033[31m"
	F_GREEN   = "\033[32m"
	F_YELLOW  = "\033[33m"
	F_BLUE    = "\033[34m"
	F_MAGENTA = "\033[35m"
	F_CYAN    = "\033[36m"
	F_WHITE   = "\033[37m"

	// background
	B_BLACK   = "\033[40m"
	B_RED     = "\033[41m"
	B_GREEN   = "\033[42m"
	B_YELLOW  = "\033[43m"
	B_BLUE    = "\033[44m"
	B_MAGENTA = "\033[45m"
	B_CYAN    = "\033[46m"
	B_WHITE   = "\033[47m"
)

var Config = map[token.Token]string{
	token.ILLEGAL: C_NULL,
	token.EOF:     C_NULL,
	token.COMMENT: F_BLUE,

	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	token.IDENT:  C_NULL, // main
	token.INT:    F_RED, // 12345
	token.FLOAT:  F_RED, // 123.45
	token.IMAG:   F_RED, // 123.45i
	token.CHAR:   C_BRIGHT + F_MAGENTA, // 'a'
	token.STRING: F_MAGENTA, // "abc"

	// Operators and delimiters
	token.ADD: C_NULL, // +
	token.SUB: C_NULL, // -
	token.MUL: C_NULL, // *
	token.QUO: C_NULL, // /
	token.REM: C_NULL, // %

	token.AND:     C_NULL, // &
	token.OR:      C_NULL, // |
	token.XOR:     C_NULL, // ^
	token.SHL:     C_NULL, // <<
	token.SHR:     C_NULL, // >>
	token.AND_NOT: C_NULL, // &^

	token.ADD_ASSIGN: C_NULL, // +=
	token.SUB_ASSIGN: C_NULL, // -=
	token.MUL_ASSIGN: C_NULL, // *=
	token.QUO_ASSIGN: C_NULL, // /=
	token.REM_ASSIGN: C_NULL, // %=

	token.AND_ASSIGN:     C_NULL, // &=
	token.OR_ASSIGN:      C_NULL, // |=
	token.XOR_ASSIGN:     C_NULL, // ^=
	token.SHL_ASSIGN:     C_NULL, // <<=
	token.SHR_ASSIGN:     C_NULL, // >>=
	token.AND_NOT_ASSIGN: C_NULL, // &^=

	token.LAND:  C_NULL, // &&
	token.LOR:   C_NULL, // ||
	token.ARROW: C_NULL, // <-
	token.INC:   C_NULL, // ++
	token.DEC:   C_NULL, // --

	token.EQL:    C_NULL, // ==
	token.LSS:    C_NULL, // <
	token.GTR:    C_NULL, // >
	token.ASSIGN: C_NULL, // =
	token.NOT:    C_NULL, // !

	token.NEQ:      C_NULL, // !=
	token.LEQ:      C_NULL, // <=
	token.GEQ:      C_NULL, // >=
	token.DEFINE:   C_NULL, // :=
	token.ELLIPSIS: C_NULL, // ...

	token.LPAREN: C_NULL, // (
	token.LBRACK: C_NULL, // [
	token.LBRACE: C_NULL, // {
	token.COMMA:  C_NULL, // ,
	token.PERIOD: C_NULL, // .

	token.RPAREN:    C_NULL, // )
	token.RBRACK:    C_NULL, // ]
	token.RBRACE:    C_NULL, // }
	//token.SEMICOLON: C_NULL, // ;
	token.COLON:     C_NULL, // :

	// Keywords
	token.BREAK:    F_YELLOW,
	token.CASE:     F_YELLOW,
	token.CHAN:     F_YELLOW,
	token.CONST:    F_YELLOW,
	token.CONTINUE: F_YELLOW,

	token.DEFAULT:     F_YELLOW,
	token.DEFER:       F_YELLOW,
	token.ELSE:        F_YELLOW,
	token.FALLTHROUGH: F_YELLOW,
	token.FOR:         F_YELLOW,

	token.FUNC:   C_BRIGHT + F_YELLOW,
	token.GO:     C_BRIGHT + F_YELLOW,
	token.GOTO:   F_YELLOW,
	token.IF:     F_YELLOW,
	token.IMPORT: F_YELLOW,

	token.INTERFACE: F_YELLOW,
	token.MAP:       F_YELLOW,
	token.PACKAGE:   C_BRIGHT + F_YELLOW,
	token.RANGE:     F_YELLOW,
	token.RETURN:    C_BRIGHT + F_YELLOW,

	token.SELECT: F_YELLOW,
	token.STRUCT: F_YELLOW,
	token.SWITCH: F_YELLOW,
	token.TYPE:   C_BRIGHT + F_YELLOW,
	token.VAR:    F_YELLOW,
}

func main() {
	var src []byte
	var err error
	if len(os.Args) > 1 {
		src, err = ioutil.ReadFile(os.Args[1])
	} else {
		src, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)
	index := 0
	out := ""
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF { break }
		c, ok := Config[tok]
		if !ok { continue }
		if lit == "" { lit = tok.String() }
		n := file.Offset(pos)
		begin, end := n, n + len(lit)
		out += string(src[index:begin]) + c + string(src[begin:end]) + C_RESET
		index = end
	}
	fmt.Println(out)
}
