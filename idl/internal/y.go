// Code generated by goyacc -l thrift.y. DO NOT EDIT.

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package internal

import (
	__yyfmt__ "fmt"

	"go.uber.org/thriftrw/ast"
)

type yySymType struct {
	yys int
	// Used to record line numbers when the line number at the start point is
	// required.
	line int

	docstring string

	// Holds the final AST for the file.
	prog *ast.Program

	// Other intermediate variables:

	bul bool
	str string
	i64 int64
	dub float64

	fieldType       ast.Type
	structType      ast.StructureType
	baseTypeID      ast.BaseTypeID
	fieldIdentifier int
	fieldRequired   ast.Requiredness

	field  *ast.Field
	fields []*ast.Field

	header  ast.Header
	headers []ast.Header

	function  *ast.Function
	functions []*ast.Function

	enumItem  *ast.EnumItem
	enumItems []*ast.EnumItem

	definition  ast.Definition
	definitions []ast.Definition

	typeAnnotations []*ast.Annotation

	constantValue    ast.ConstantValue
	constantValues   []ast.ConstantValue
	constantMapItems []ast.ConstantMapItem
}

const IDENTIFIER = 57346
const LITERAL = 57347
const INTCONSTANT = 57348
const DUBCONSTANT = 57349
const NAMESPACE = 57350
const INCLUDE = 57351
const CPP_INCLUDE = 57352
const VOID = 57353
const BOOL = 57354
const BYTE = 57355
const I8 = 57356
const I16 = 57357
const I32 = 57358
const I64 = 57359
const DOUBLE = 57360
const STRING = 57361
const BINARY = 57362
const MAP = 57363
const LIST = 57364
const SET = 57365
const ONEWAY = 57366
const TYPEDEF = 57367
const STRUCT = 57368
const UNION = 57369
const EXCEPTION = 57370
const EXTENDS = 57371
const THROWS = 57372
const SERVICE = 57373
const ENUM = 57374
const CONST = 57375
const REQUIRED = 57376
const OPTIONAL = 57377
const TRUE = 57378
const FALSE = 57379

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"LITERAL",
	"INTCONSTANT",
	"DUBCONSTANT",
	"NAMESPACE",
	"INCLUDE",
	"CPP_INCLUDE",
	"VOID",
	"BOOL",
	"BYTE",
	"I8",
	"I16",
	"I32",
	"I64",
	"DOUBLE",
	"STRING",
	"BINARY",
	"MAP",
	"LIST",
	"SET",
	"ONEWAY",
	"TYPEDEF",
	"STRUCT",
	"UNION",
	"EXCEPTION",
	"EXTENDS",
	"THROWS",
	"SERVICE",
	"ENUM",
	"CONST",
	"REQUIRED",
	"OPTIONAL",
	"TRUE",
	"FALSE",
	"'*'",
	"'='",
	"'{'",
	"'}'",
	"':'",
	"'('",
	"')'",
	"'<'",
	"','",
	"'>'",
	"'['",
	"']'",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	8, 73,
	9, 73,
	10, 73,
	-2, 9,
	-1, 3,
	1, 1,
	-2, 73,
}

const yyPrivate = 57344

const yyLast = 180

var yyAct = [...]int{

	32, 58, 65, 5, 7, 68, 87, 31, 11, 66,
	125, 73, 69, 70, 14, 12, 89, 12, 94, 13,
	127, 13, 96, 95, 62, 61, 60, 162, 155, 33,
	151, 129, 92, 160, 59, 59, 59, 148, 144, 133,
	135, 123, 71, 72, 85, 82, 73, 69, 70, 79,
	56, 107, 91, 54, 63, 121, 53, 67, 74, 57,
	90, 55, 19, 131, 132, 81, 84, 158, 76, 77,
	78, 106, 10, 8, 9, 140, 93, 71, 72, 118,
	116, 97, 16, 15, 100, 138, 28, 103, 99, 98,
	17, 102, 101, 147, 142, 104, 18, 112, 113, 157,
	114, 110, 111, 88, 52, 37, 36, 35, 34, 74,
	124, 30, 29, 119, 122, 117, 128, 120, 134, 126,
	21, 25, 26, 27, 105, 74, 24, 22, 20, 139,
	137, 136, 75, 109, 108, 3, 6, 143, 141, 64,
	80, 86, 146, 2, 4, 74, 83, 145, 23, 150,
	149, 152, 74, 84, 130, 115, 156, 154, 153, 159,
	42, 38, 84, 161, 1, 0, 0, 0, 43, 44,
	45, 46, 47, 48, 49, 50, 51, 39, 40, 41,
}
var yyPact = [...]int{

	-1000, -1000, -1000, -1000, -1000, 64, -31, -1000, 78, 85,
	58, -1000, -1000, -1000, 95, -1000, 81, -1000, 108, 107,
	-1000, -1000, 104, 103, 102, -1000, -1000, -1000, -1000, -1000,
	-1000, 101, 156, 100, 16, 13, 21, 20, -7, -19,
	-20, -21, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -7, -1000, -1000, -1000, -1000, 41, -1000, -1000,
	-1000, -1000, -1000, -1000, 8, 4, 3, 99, -1000, -1000,
	-1000, -1000, -1000, -1000, 12, -12, -28, -24, -25, -7,
	-31, -1000, -7, -31, -1000, -7, -31, 47, 11, -1000,
	-1000, -1000, -1000, 97, -1000, -7, -7, -1000, -1000, 96,
	-1000, -1000, 74, -1000, -1000, 68, -1000, -1000, 6, 0,
	-29, -27, -1000, -1000, -8, 29, -3, -1000, -1000, -1000,
	-1, -1000, -31, -1000, 41, 80, -1000, -7, -1000, 69,
	-1000, -1000, -1000, -1000, 90, -7, -1000, -4, -31, -1000,
	-7, 89, -6, -1000, 41, -1000, -1000, -9, -1000, -31,
	-1000, 41, -16, -1000, -7, 37, -1000, -7, -10, -1000,
	-1000, -17, -1000,
}
var yyPgo = [...]int{

	0, 0, 6, 164, 7, 161, 155, 154, 148, 146,
	2, 144, 143, 141, 9, 140, 139, 136, 135, 5,
	134, 133, 132, 1, 8, 124, 115, 99,
}
var yyR1 = [...]int{

	0, 3, 12, 12, 11, 11, 11, 11, 11, 18,
	18, 17, 17, 17, 17, 17, 17, 8, 8, 8,
	16, 16, 15, 15, 10, 10, 9, 9, 6, 6,
	7, 7, 7, 14, 14, 13, 25, 25, 26, 26,
	27, 27, 4, 4, 4, 4, 4, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 19, 19, 19, 19,
	19, 19, 19, 19, 20, 20, 21, 21, 23, 23,
	22, 22, 22, 1, 2, 24, 24, 24,
}
var yyR2 = [...]int{

	0, 2, 0, 2, 3, 4, 3, 4, 4, 0,
	3, 7, 6, 8, 8, 8, 11, 1, 1, 1,
	0, 3, 4, 6, 0, 3, 7, 9, 2, 0,
	1, 1, 0, 0, 3, 10, 1, 0, 1, 1,
	0, 4, 3, 8, 6, 6, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 4, 4, 0, 3, 0, 6, 0, 3,
	0, 6, 4, 0, 0, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -3, -12, -18, -11, -1, -17, -1, 9, 10,
	8, -24, 46, 50, -2, 5, 4, 5, 38, 4,
	33, 25, 32, -8, 31, 26, 27, 28, 5, 4,
	4, -4, -1, -4, 4, 4, 4, 4, -5, 21,
	22, 23, 4, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 4, 40, 40, 40, 29, 39, -23, 43,
	45, 45, 45, -23, -16, -10, -14, -1, -19, 6,
	7, 36, 37, 5, -1, -22, -4, -4, -4, 41,
	-15, -1, 41, -9, -1, 41, -13, -2, 4, 4,
	48, 40, 44, -1, 46, 47, 47, -23, -24, -2,
	-23, -24, -2, -23, -24, -25, 24, 40, -20, -21,
	4, -4, -23, -23, 4, -6, 6, -26, 11, -4,
	-14, 49, -19, 41, -1, 39, -24, 47, -23, 39,
	-7, 34, 35, 42, -1, 41, -24, -19, 5, -23,
	6, -4, 4, -23, 42, -24, -23, 4, 43, -19,
	-23, 39, -10, -24, -19, 44, -23, -27, 30, -23,
	43, -10, 44,
}
var yyDef = [...]int{

	2, -2, -2, -2, 3, 0, 77, 74, 0, 0,
	0, 10, 75, 76, 0, 4, 0, 6, 0, 0,
	73, 73, 0, 0, 0, 17, 18, 19, 5, 7,
	8, 0, 0, 0, 0, 0, 0, 0, 68, 0,
	0, 0, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 68, 20, 24, 33, 73, 73, 42, 70,
	73, 73, 73, 12, 73, 73, 74, 0, 11, 56,
	57, 58, 59, 60, 0, 73, 0, 0, 0, 68,
	77, 74, 68, 77, 74, 68, 77, 37, 0, 61,
	64, 66, 69, 0, 73, 68, 68, 13, 21, 0,
	14, 25, 29, 15, 34, 73, 36, 33, 73, 73,
	77, 0, 44, 45, 68, 32, 0, 73, 38, 39,
	74, 62, 77, 63, 73, 0, 72, 68, 22, 0,
	73, 30, 31, 28, 0, 68, 65, 0, 77, 43,
	68, 0, 0, 16, 73, 71, 23, 68, 24, 77,
	26, 73, 73, 67, 68, 40, 27, 68, 0, 35,
	24, 73, 41,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	43, 44, 38, 3, 46, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 42, 50,
	45, 39, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 48, 3, 49, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 40, 3, 41,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.prog = &ast.Program{Headers: yyDollar[1].headers, Definitions: yyDollar[2].definitions}
			yylex.(*lexer).program = yyVAL.prog
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.headers = nil
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.headers = append(yyDollar[1].headers, yyDollar[2].header)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.header = &ast.Include{
				Path: yyDollar[3].str,
				Line: yyDollar[1].line,
			}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.header = &ast.Include{
				Name: yyDollar[3].str,
				Path: yyDollar[4].str,
				Line: yyDollar[1].line,
			}
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.header = &ast.CppInclude{
				Path: yyDollar[3].str,
				Line: yyDollar[1].line,
			}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.header = &ast.Namespace{
				Scope: "*",
				Name:  yyDollar[4].str,
				Line:  yyDollar[1].line,
			}
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.header = &ast.Namespace{
				Scope: yyDollar[3].str,
				Name:  yyDollar[4].str,
				Line:  yyDollar[1].line,
			}
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.definitions = nil
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.definitions = append(yyDollar[1].definitions, yyDollar[2].definition)
		}
	case 11:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.definition = &ast.Constant{
				Name:  yyDollar[5].str,
				Type:  yyDollar[4].fieldType,
				Value: yyDollar[7].constantValue,
				Line:  yyDollar[1].line,
				Doc:   ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 12:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.definition = &ast.Typedef{
				Name:        yyDollar[5].str,
				Type:        yyDollar[4].fieldType,
				Annotations: yyDollar[6].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 13:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.definition = &ast.Enum{
				Name:        yyDollar[4].str,
				Items:       yyDollar[6].enumItems,
				Annotations: yyDollar[8].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 14:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.definition = &ast.Struct{
				Name:        yyDollar[4].str,
				Type:        yyDollar[3].structType,
				Fields:      yyDollar[6].fields,
				Annotations: yyDollar[8].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 15:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.definition = &ast.Service{
				Name:        yyDollar[4].str,
				Functions:   yyDollar[6].functions,
				Annotations: yyDollar[8].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 16:
		yyDollar = yyS[yypt-11 : yypt+1]
		{
			parent := &ast.ServiceReference{
				Name: yyDollar[7].str,
				Line: yyDollar[6].line,
			}

			yyVAL.definition = &ast.Service{
				Name:        yyDollar[4].str,
				Functions:   yyDollar[9].functions,
				Parent:      parent,
				Annotations: yyDollar[11].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.structType = ast.StructType
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.structType = ast.UnionType
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.structType = ast.ExceptionType
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.enumItems = nil
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.enumItems = append(yyDollar[1].enumItems, yyDollar[2].enumItem)
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.enumItem = &ast.EnumItem{
				Name:        yyDollar[3].str,
				Annotations: yyDollar[4].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			value := int(yyDollar[5].i64)
			yyVAL.enumItem = &ast.EnumItem{
				Name:        yyDollar[3].str,
				Value:       &value,
				Annotations: yyDollar[6].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.fields = nil
			yylex.(*lexer).autoFieldIndex = -1
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.fields = append(yyDollar[1].fields, yyDollar[2].field)
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.field = &ast.Field{
				ID:           yyDollar[3].fieldIdentifier,
				Name:         yyDollar[6].str,
				Type:         yyDollar[5].fieldType,
				Requiredness: yyDollar[4].fieldRequired,
				Annotations:  yyDollar[7].typeAnnotations,
				Line:         yyDollar[1].line,
				Doc:          ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 27:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.field = &ast.Field{
				ID:           yyDollar[3].fieldIdentifier,
				Name:         yyDollar[6].str,
				Type:         yyDollar[5].fieldType,
				Requiredness: yyDollar[4].fieldRequired,
				Default:      yyDollar[8].constantValue,
				Annotations:  yyDollar[9].typeAnnotations,
				Line:         yyDollar[1].line,
				Doc:          ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.fieldIdentifier = int(yyDollar[1].i64)

			// Use a negative value as the new basis for auto-assigned fields.
			if yyVAL.fieldIdentifier < 0 {
				yylex.(*lexer).autoFieldIndex = yyVAL.fieldIdentifier - 1
			}
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.fieldIdentifier = yylex.(*lexer).autoFieldIndex
			yylex.(*lexer).autoFieldIndex -= 1
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fieldRequired = ast.Required
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fieldRequired = ast.Optional
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.fieldRequired = ast.Unspecified
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.functions = nil
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.functions = append(yyDollar[1].functions, yyDollar[2].function)
		}
	case 35:
		yyDollar = yyS[yypt-10 : yypt+1]
		{
			yyVAL.function = &ast.Function{
				Name:        yyDollar[5].str,
				Parameters:  yyDollar[7].fields,
				ReturnType:  yyDollar[3].fieldType,
				Exceptions:  yyDollar[9].fields,
				OneWay:      yyDollar[2].bul,
				Annotations: yyDollar[10].typeAnnotations,
				Line:        yyDollar[4].line,
				Doc:         ParseDocstring(yyDollar[1].docstring),
			}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.bul = true
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.bul = false
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fieldType = nil
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fieldType = yyDollar[1].fieldType
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.fields = nil
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.fields = yyDollar[3].fields
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.fieldType = ast.BaseType{ID: yyDollar[2].baseTypeID, Annotations: yyDollar[3].typeAnnotations, Line: yyDollar[1].line}
		}
	case 43:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.fieldType = ast.MapType{KeyType: yyDollar[4].fieldType, ValueType: yyDollar[6].fieldType, Annotations: yyDollar[8].typeAnnotations, Line: yyDollar[1].line}
		}
	case 44:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.fieldType = ast.ListType{ValueType: yyDollar[4].fieldType, Annotations: yyDollar[6].typeAnnotations, Line: yyDollar[1].line}
		}
	case 45:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.fieldType = ast.SetType{ValueType: yyDollar[4].fieldType, Annotations: yyDollar[6].typeAnnotations, Line: yyDollar[1].line}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.fieldType = ast.TypeReference{Name: yyDollar[2].str, Line: yyDollar[1].line}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.BoolTypeID
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I8TypeID
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I8TypeID
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I16TypeID
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I32TypeID
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I64TypeID
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.DoubleTypeID
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.StringTypeID
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.BinaryTypeID
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantInteger(yyDollar[1].i64)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantDouble(yyDollar[1].dub)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantBoolean(true)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantBoolean(false)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantString(yyDollar[1].str)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantReference{Name: yyDollar[2].str, Line: yyDollar[1].line}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantList{Items: yyDollar[3].constantValues, Line: yyDollar[1].line}
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantMap{Items: yyDollar[3].constantMapItems, Line: yyDollar[1].line}
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.constantValues = nil
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.constantValues = append(yyDollar[1].constantValues, yyDollar[2].constantValue)
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.constantMapItems = nil
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.constantMapItems = append(yyDollar[1].constantMapItems, ast.ConstantMapItem{Key: yyDollar[3].constantValue, Value: yyDollar[5].constantValue, Line: yyDollar[2].line})
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.typeAnnotations = nil
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.typeAnnotations = yyDollar[2].typeAnnotations
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.typeAnnotations = nil
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.typeAnnotations = append(yyDollar[1].typeAnnotations, &ast.Annotation{Name: yyDollar[3].str, Value: yyDollar[5].str, Line: yyDollar[2].line})
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.typeAnnotations = append(yyDollar[1].typeAnnotations, &ast.Annotation{Name: yyDollar[3].str, Line: yyDollar[2].line})
		}
	case 73:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.line = yylex.(*lexer).line
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.docstring = yylex.(*lexer).LastDocstring()
		}
	}
	goto yystack /* stack new state and value */
}
