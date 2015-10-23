package parser

import (
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/types"
	. "github.com/katydid/katydid/viper/ast"
	"github.com/katydid/katydid/viper/token"
	"strconv"
)

func newString(v interface{}) string {
	t := v.(*token.Token)
	return string(t.Lit)
}

func unquote(s1 string) string {
	s, err := strconv.Unquote(s1)
	if err != nil {
		return s1
	}
	return s
}

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Start	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Start : Rules	<< &Rules{X[0].([]*Rule), nil}, nil >>`,
		Id:         "Start",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Rules{X[0].([]*Rule), nil}, nil
		},
	},
	ProdTabEntry{
		String: `Start : Rules Space	<< &Rules{X[0].([]*Rule), X[1].(*expr.Space)}, nil >>`,
		Id:         "Start",
		NTType:     1,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Rules{X[0].([]*Rule), X[1].(*expr.Space)}, nil
		},
	},
	ProdTabEntry{
		String: `Rules : Rule	<< []*Rule{X[0].(*Rule)}, nil >>`,
		Id:         "Rules",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*Rule{X[0].(*Rule)}, nil
		},
	},
	ProdTabEntry{
		String: `Rules : Rules Rule	<< append(X[0].([]*Rule), X[1].(*Rule)), nil >>`,
		Id:         "Rules",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*Rule), X[1].(*Rule)), nil
		},
	},
	ProdTabEntry{
		String: `Rule : StartRule	<< &Rule{Start: X[0].(*Start)}, nil >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Rule{Start: X[0].(*Start)}, nil
		},
	},
	ProdTabEntry{
		String: `Rule : Final	<< &Rule{Final: X[0].(*Final)}, nil >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Rule{Final: X[0].(*Final)}, nil
		},
	},
	ProdTabEntry{
		String: `Rule : Internal	<< &Rule{Internal: X[0].(*Internal)}, nil >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Rule{Internal: X[0].(*Internal)}, nil
		},
	},
	ProdTabEntry{
		String: `Rule : Call	<< &Rule{Call: X[0].(*Call)}, nil >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Rule{Call: X[0].(*Call)}, nil
		},
	},
	ProdTabEntry{
		String: `Rule : Return	<< &Rule{Return: X[0].(*Return)}, nil >>`,
		Id:         "Rule",
		NTType:     3,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Rule{Return: X[0].(*Return)}, nil
		},
	},
	ProdTabEntry{
		String: `StartRule : Space "start" Equal State SemiColon	<< &Start{
      Before: X[0].(*expr.Space),
      Eq: X[2].(*expr.Keyword),
      State: X[3].(*State),
      SemiColon: X[4].(*expr.Keyword),
    }, nil >>`,
		Id:         "StartRule",
		NTType:     4,
		Index:      10,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Start{
				Before:    X[0].(*expr.Space),
				Eq:        X[2].(*expr.Keyword),
				State:     X[3].(*State),
				SemiColon: X[4].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `StartRule : "start" Equal State SemiColon	<< &Start{
      Eq: X[1].(*expr.Keyword),
      State: X[2].(*State),
      SemiColon: X[3].(*expr.Keyword),
    }, nil >>`,
		Id:         "StartRule",
		NTType:     4,
		Index:      11,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Start{
				Eq:        X[1].(*expr.Keyword),
				State:     X[2].(*State),
				SemiColon: X[3].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Final : Space "final" Equal State SemiColon	<< &Final{
      Before: X[0].(*expr.Space),
      Eq: X[2].(*expr.Keyword),
      State: X[3].(*State),
      SemiColon: X[4].(*expr.Keyword),
    }, nil >>`,
		Id:         "Final",
		NTType:     5,
		Index:      12,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Final{
				Before:    X[0].(*expr.Space),
				Eq:        X[2].(*expr.Keyword),
				State:     X[3].(*State),
				SemiColon: X[4].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Final : "final" Equal State SemiColon	<< &Final{
      Eq: X[1].(*expr.Keyword),
      State: X[2].(*State),
      SemiColon: X[3].(*expr.Keyword),
    }, nil >>`,
		Id:         "Final",
		NTType:     5,
		Index:      13,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Final{
				Eq:        X[1].(*expr.Keyword),
				State:     X[2].(*State),
				SemiColon: X[3].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Internal : Space "internal" State Expr State SemiColon	<< &Internal{
      Before: X[0].(*expr.Space),
      Src: X[2].(*State),
      Expr: X[3].(*expr.Expr),
      Dst: X[4].(*State),
      SemiColon: X[5].(*expr.Keyword),
    }, nil >>`,
		Id:         "Internal",
		NTType:     6,
		Index:      14,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Internal{
				Before:    X[0].(*expr.Space),
				Src:       X[2].(*State),
				Expr:      X[3].(*expr.Expr),
				Dst:       X[4].(*State),
				SemiColon: X[5].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Internal : "internal" State Expr State SemiColon	<< &Internal{
      Src: X[1].(*State),
      Expr: X[2].(*expr.Expr),
      Dst: X[3].(*State),
      SemiColon: X[4].(*expr.Keyword),
    }, nil >>`,
		Id:         "Internal",
		NTType:     6,
		Index:      15,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Internal{
				Src:       X[1].(*State),
				Expr:      X[2].(*expr.Expr),
				Dst:       X[3].(*State),
				SemiColon: X[4].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Call : Space "call" State Expr State State SemiColon	<< &Call{
      Before: X[0].(*expr.Space),
      Src: X[2].(*State),
      Expr: X[3].(*expr.Expr),
      ParentDst: X[4].(*State),
      ChildDst: X[5].(*State),
      SemiColon: X[6].(*expr.Keyword),
    }, nil >>`,
		Id:         "Call",
		NTType:     7,
		Index:      16,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Call{
				Before:    X[0].(*expr.Space),
				Src:       X[2].(*State),
				Expr:      X[3].(*expr.Expr),
				ParentDst: X[4].(*State),
				ChildDst:  X[5].(*State),
				SemiColon: X[6].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Call : "call" State Expr State State SemiColon	<< &Call{
      Src: X[1].(*State),
      Expr: X[2].(*expr.Expr),
      ParentDst: X[3].(*State),
      ChildDst: X[4].(*State),
      SemiColon: X[5].(*expr.Keyword),
    }, nil >>`,
		Id:         "Call",
		NTType:     7,
		Index:      17,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Call{
				Src:       X[1].(*State),
				Expr:      X[2].(*expr.Expr),
				ParentDst: X[3].(*State),
				ChildDst:  X[4].(*State),
				SemiColon: X[5].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Return : Space "return" State State Expr State SemiColon	<< &Return{
      Before: X[0].(*expr.Space),
      ParentSrc: X[2].(*State),
      ChildSrc: X[3].(*State),
      Expr: X[4].(*expr.Expr),
      Dst: X[5].(*State),
      SemiColon: X[6].(*expr.Keyword),
    }, nil >>`,
		Id:         "Return",
		NTType:     8,
		Index:      18,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Return{
				Before:    X[0].(*expr.Space),
				ParentSrc: X[2].(*State),
				ChildSrc:  X[3].(*State),
				Expr:      X[4].(*expr.Expr),
				Dst:       X[5].(*State),
				SemiColon: X[6].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `Return : "return" State State Expr State SemiColon	<< &Return{
      ParentSrc: X[1].(*State),
      ChildSrc: X[2].(*State),
      Expr: X[3].(*expr.Expr),
      Dst: X[4].(*State),
      SemiColon: X[5].(*expr.Keyword),
    }, nil >>`,
		Id:         "Return",
		NTType:     8,
		Index:      19,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &Return{
				ParentSrc: X[1].(*State),
				ChildSrc:  X[2].(*State),
				Expr:      X[3].(*expr.Expr),
				Dst:       X[4].(*State),
				SemiColon: X[5].(*expr.Keyword),
			}, nil
		},
	},
	ProdTabEntry{
		String: `State : Space id	<< &State{Before: X[0].(*expr.Space), Name: newString(X[1])}, nil >>`,
		Id:         "State",
		NTType:     9,
		Index:      20,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &State{Before: X[0].(*expr.Space), Name: newString(X[1])}, nil
		},
	},
	ProdTabEntry{
		String: `State : id	<< &State{Name: newString(X[0])}, nil >>`,
		Id:         "State",
		NTType:     9,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &State{Name: newString(X[0])}, nil
		},
	},
	ProdTabEntry{
		String: `State : Space string_lit	<< &State{Before: X[0].(*expr.Space), Name: newString(X[1])}, nil >>`,
		Id:         "State",
		NTType:     9,
		Index:      22,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &State{Before: X[0].(*expr.Space), Name: newString(X[1])}, nil
		},
	},
	ProdTabEntry{
		String: `State : string_lit	<< &State{Name: newString(X[0])}, nil >>`,
		Id:         "State",
		NTType:     9,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &State{Name: newString(X[0])}, nil
		},
	},
	ProdTabEntry{
		String: `Start : Expr Space	<<  >>`,
		Id:         "Start",
		NTType:     1,
		Index:      24,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Start : Expr	<<  >>`,
		Id:         "Start",
		NTType:     1,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Start : BuiltIn Space	<< &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil >>`,
		Id:         "Start",
		NTType:     1,
		Index:      26,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil
		},
	},
	ProdTabEntry{
		String: `Start : BuiltIn	<< &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil >>`,
		Id:         "Start",
		NTType:     1,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{BuiltIn: X[0].(*expr.BuiltIn)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : SpaceTerminal	<< &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Terminal: X[0].(*expr.Terminal)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : Function	<< &expr.Expr{Function: X[0].(*expr.Function)}, nil >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{Function: X[0].(*expr.Function)}, nil
		},
	},
	ProdTabEntry{
		String: `Expr : List	<< &expr.Expr{List: X[0].(*expr.List)}, nil >>`,
		Id:         "Expr",
		NTType:     10,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Expr{List: X[0].(*expr.List)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen Exprs CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     11,
		Index:      31,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), Params: X[3].([]*expr.Expr), CloseParen: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : Space id OpenParen CloseParen	<< &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     11,
		Index:      32,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Before: X[0].(*expr.Space), Name: newString(X[1]), OpenParen: X[2].(*expr.Keyword), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen Exprs CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     11,
		Index:      33,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), Params: X[2].([]*expr.Expr), CloseParen: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Function : id OpenParen CloseParen	<< &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "Function",
		NTType:     11,
		Index:      34,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Function{Name: newString(X[0]), OpenParen: X[1].(*expr.Keyword), CloseParen: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : EqualEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      35,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : ExclamationEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      36,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : LessThan Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      37,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : GreaterThan Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      38,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : GreaterEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      39,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : LessEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      40,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : TildeEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      41,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : StarEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      42,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : CaretEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      43,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : DollarEqual Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      44,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `BuiltIn : ColonColon Expr	<< &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil >>`,
		Id:         "BuiltIn",
		NTType:     12,
		Index:      45,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.BuiltIn{Symbol: X[0].(*expr.Keyword), Expr: X[1].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly Exprs CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     13,
		Index:      46,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), Elems: X[3].([]*expr.Expr), CloseCurly: X[4].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly Exprs CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     13,
		Index:      47,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), Elems: X[2].([]*expr.Expr), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : Space ListType OpenCurly CloseCurly	<< &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     13,
		Index:      48,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Before: X[0].(*expr.Space), Type: X[1].(types.Type), OpenCurly: X[2].(*expr.Keyword), CloseCurly: X[3].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `List : ListType OpenCurly CloseCurly	<< &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil >>`,
		Id:         "List",
		NTType:     13,
		Index:      49,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.List{Type: X[0].(types.Type), OpenCurly: X[1].(*expr.Keyword), CloseCurly: X[2].(*expr.Keyword)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Expr	<< []*expr.Expr{X[0].(*expr.Expr)}, nil >>`,
		Id:         "Exprs",
		NTType:     14,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []*expr.Expr{X[0].(*expr.Expr)}, nil
		},
	},
	ProdTabEntry{
		String: `Exprs : Exprs Comma Expr	<< append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil >>`,
		Id:         "Exprs",
		NTType:     14,
		Index:      51,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]*expr.Expr), expr.SetExprComma(X[2], X[1])), nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]bool"	<< types.LIST_BOOL, nil >>`,
		Id:         "ListType",
		NTType:     15,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BOOL, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]int"	<< types.LIST_INT, nil >>`,
		Id:         "ListType",
		NTType:     15,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_INT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]uint"	<< types.LIST_UINT, nil >>`,
		Id:         "ListType",
		NTType:     15,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_UINT, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]double"	<< types.LIST_DOUBLE, nil >>`,
		Id:         "ListType",
		NTType:     15,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_DOUBLE, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[]string"	<< types.LIST_STRING, nil >>`,
		Id:         "ListType",
		NTType:     15,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_STRING, nil
		},
	},
	ProdTabEntry{
		String: `ListType : "[][]byte"	<< types.LIST_BYTES, nil >>`,
		Id:         "ListType",
		NTType:     15,
		Index:      57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.LIST_BYTES, nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Terminal	<< X[0], nil >>`,
		Id:         "SpaceTerminal",
		NTType:     16,
		Index:      58,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SpaceTerminal : Space Terminal	<< expr.SetTerminalSpace(X[1], X[0]), nil >>`,
		Id:         "SpaceTerminal",
		NTType:     16,
		Index:      59,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.SetTerminalSpace(X[1], X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Literal : Bool	<< expr.NewBoolTerminal(X[0]), nil >>`,
		Id:         "Literal",
		NTType:     17,
		Index:      60,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBoolTerminal(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Literal : int_lit	<< expr.NewIntTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     17,
		Index:      61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewIntTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : uint_lit	<< expr.NewUintTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     17,
		Index:      62,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewUintTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : double_lit	<< expr.NewDoubleTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     17,
		Index:      63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewDoubleTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : string_lit	<< expr.NewStringTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     17,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewStringTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Literal : bytes_lit	<< expr.NewBytesTerminal(newString(X[0])) >>`,
		Id:         "Literal",
		NTType:     17,
		Index:      65,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewBytesTerminal(newString(X[0]))
		},
	},
	ProdTabEntry{
		String: `Terminal : Literal	<<  >>`,
		Id:         "Terminal",
		NTType:     18,
		Index:      66,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Terminal : bool_var	<< expr.NewVariableTerminal(types.SINGLE_BOOL) >>`,
		Id:         "Terminal",
		NTType:     18,
		Index:      67,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BOOL)
		},
	},
	ProdTabEntry{
		String: `Terminal : int_var	<< expr.NewVariableTerminal(types.SINGLE_INT) >>`,
		Id:         "Terminal",
		NTType:     18,
		Index:      68,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_INT)
		},
	},
	ProdTabEntry{
		String: `Terminal : uint_var	<< expr.NewVariableTerminal(types.SINGLE_UINT) >>`,
		Id:         "Terminal",
		NTType:     18,
		Index:      69,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_UINT)
		},
	},
	ProdTabEntry{
		String: `Terminal : double_var	<< expr.NewVariableTerminal(types.SINGLE_DOUBLE) >>`,
		Id:         "Terminal",
		NTType:     18,
		Index:      70,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_DOUBLE)
		},
	},
	ProdTabEntry{
		String: `Terminal : string_var	<< expr.NewVariableTerminal(types.SINGLE_STRING) >>`,
		Id:         "Terminal",
		NTType:     18,
		Index:      71,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_STRING)
		},
	},
	ProdTabEntry{
		String: `Terminal : bytes_var	<< expr.NewVariableTerminal(types.SINGLE_BYTES) >>`,
		Id:         "Terminal",
		NTType:     18,
		Index:      72,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.NewVariableTerminal(types.SINGLE_BYTES)
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< true, nil >>`,
		Id:         "Bool",
		NTType:     19,
		Index:      73,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< false, nil >>`,
		Id:         "Bool",
		NTType:     19,
		Index:      74,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `Equal : "="	<< &expr.Keyword{Value: "="}, nil >>`,
		Id:         "Equal",
		NTType:     20,
		Index:      75,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "="}, nil
		},
	},
	ProdTabEntry{
		String: `Equal : Space "="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "=",
    }, nil >>`,
		Id:         "Equal",
		NTType:     20,
		Index:      76,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `OpenParen : "("	<< &expr.Keyword{Value: "("}, nil >>`,
		Id:         "OpenParen",
		NTType:     21,
		Index:      77,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "("}, nil
		},
	},
	ProdTabEntry{
		String: `OpenParen : Space "("	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "(",
    }, nil >>`,
		Id:         "OpenParen",
		NTType:     21,
		Index:      78,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "(",
			}, nil
		},
	},
	ProdTabEntry{
		String: `CloseParen : ")"	<< &expr.Keyword{Value: ")"}, nil >>`,
		Id:         "CloseParen",
		NTType:     22,
		Index:      79,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ")"}, nil
		},
	},
	ProdTabEntry{
		String: `CloseParen : Space ")"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ")",
    }, nil >>`,
		Id:         "CloseParen",
		NTType:     22,
		Index:      80,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ")",
			}, nil
		},
	},
	ProdTabEntry{
		String: `OpenCurly : "{"	<< &expr.Keyword{Value: "{"}, nil >>`,
		Id:         "OpenCurly",
		NTType:     23,
		Index:      81,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "{"}, nil
		},
	},
	ProdTabEntry{
		String: `OpenCurly : Space "{"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "{",
    }, nil >>`,
		Id:         "OpenCurly",
		NTType:     23,
		Index:      82,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "{",
			}, nil
		},
	},
	ProdTabEntry{
		String: `CloseCurly : "}"	<< &expr.Keyword{Value: "}"}, nil >>`,
		Id:         "CloseCurly",
		NTType:     24,
		Index:      83,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "}"}, nil
		},
	},
	ProdTabEntry{
		String: `CloseCurly : Space "}"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "}",
    }, nil >>`,
		Id:         "CloseCurly",
		NTType:     24,
		Index:      84,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "}",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Comma : ","	<< &expr.Keyword{Value: ","}, nil >>`,
		Id:         "Comma",
		NTType:     25,
		Index:      85,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ","}, nil
		},
	},
	ProdTabEntry{
		String: `Comma : Space ","	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ",",
    }, nil >>`,
		Id:         "Comma",
		NTType:     25,
		Index:      86,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ",",
			}, nil
		},
	},
	ProdTabEntry{
		String: `SemiColon : ";"	<< &expr.Keyword{Value: ";"}, nil >>`,
		Id:         "SemiColon",
		NTType:     26,
		Index:      87,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ";"}, nil
		},
	},
	ProdTabEntry{
		String: `SemiColon : Space ";"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ";",
    }, nil >>`,
		Id:         "SemiColon",
		NTType:     26,
		Index:      88,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ";",
			}, nil
		},
	},
	ProdTabEntry{
		String: `HashTag : "#"	<< &expr.Keyword{Value: "#"}, nil >>`,
		Id:         "HashTag",
		NTType:     27,
		Index:      89,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "#"}, nil
		},
	},
	ProdTabEntry{
		String: `HashTag : Space "#"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "#",
    }, nil >>`,
		Id:         "HashTag",
		NTType:     27,
		Index:      90,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "#",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Ampersand : "&"	<< &expr.Keyword{Value: "&"}, nil >>`,
		Id:         "Ampersand",
		NTType:     28,
		Index:      91,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "&"}, nil
		},
	},
	ProdTabEntry{
		String: `Ampersand : Space "&"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "&",
    }, nil >>`,
		Id:         "Ampersand",
		NTType:     28,
		Index:      92,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "&",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Pipe : "|"	<< &expr.Keyword{Value: "|"}, nil >>`,
		Id:         "Pipe",
		NTType:     29,
		Index:      93,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "|"}, nil
		},
	},
	ProdTabEntry{
		String: `Pipe : Space "|"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "|",
    }, nil >>`,
		Id:         "Pipe",
		NTType:     29,
		Index:      94,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "|",
			}, nil
		},
	},
	ProdTabEntry{
		String: `OpenBracket : "["	<< &expr.Keyword{Value: "["}, nil >>`,
		Id:         "OpenBracket",
		NTType:     30,
		Index:      95,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "["}, nil
		},
	},
	ProdTabEntry{
		String: `OpenBracket : Space "["	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "[",
    }, nil >>`,
		Id:         "OpenBracket",
		NTType:     30,
		Index:      96,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "[",
			}, nil
		},
	},
	ProdTabEntry{
		String: `CloseBracket : "]"	<< &expr.Keyword{Value: "]"}, nil >>`,
		Id:         "CloseBracket",
		NTType:     31,
		Index:      97,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "]"}, nil
		},
	},
	ProdTabEntry{
		String: `CloseBracket : Space "]"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "]",
    }, nil >>`,
		Id:         "CloseBracket",
		NTType:     31,
		Index:      98,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "]",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Colon : ":"	<< &expr.Keyword{Value: ":"}, nil >>`,
		Id:         "Colon",
		NTType:     32,
		Index:      99,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ":"}, nil
		},
	},
	ProdTabEntry{
		String: `Colon : Space ":"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ":",
    }, nil >>`,
		Id:         "Colon",
		NTType:     32,
		Index:      100,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ":",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Exclamation : "!"	<< &expr.Keyword{Value: "!"}, nil >>`,
		Id:         "Exclamation",
		NTType:     33,
		Index:      101,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "!"}, nil
		},
	},
	ProdTabEntry{
		String: `Exclamation : Space "!"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "!",
    }, nil >>`,
		Id:         "Exclamation",
		NTType:     33,
		Index:      102,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "!",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Star : "*"	<< &expr.Keyword{Value: "*"}, nil >>`,
		Id:         "Star",
		NTType:     34,
		Index:      103,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "*"}, nil
		},
	},
	ProdTabEntry{
		String: `Star : Space "*"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "*",
    }, nil >>`,
		Id:         "Star",
		NTType:     34,
		Index:      104,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "*",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Underscore : "_"	<< &expr.Keyword{Value: "_"}, nil >>`,
		Id:         "Underscore",
		NTType:     35,
		Index:      105,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "_"}, nil
		},
	},
	ProdTabEntry{
		String: `Underscore : Space "_"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "_",
    }, nil >>`,
		Id:         "Underscore",
		NTType:     35,
		Index:      106,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "_",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Tilde : "~"	<< &expr.Keyword{Value: "~"}, nil >>`,
		Id:         "Tilde",
		NTType:     36,
		Index:      107,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "~"}, nil
		},
	},
	ProdTabEntry{
		String: `Tilde : Space "~"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "~",
    }, nil >>`,
		Id:         "Tilde",
		NTType:     36,
		Index:      108,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "~",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Dot : "."	<< &expr.Keyword{Value: "."}, nil >>`,
		Id:         "Dot",
		NTType:     37,
		Index:      109,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "."}, nil
		},
	},
	ProdTabEntry{
		String: `Dot : Space "."	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ".",
    }, nil >>`,
		Id:         "Dot",
		NTType:     37,
		Index:      110,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ".",
			}, nil
		},
	},
	ProdTabEntry{
		String: `At : "@"	<< &expr.Keyword{Value: "@"}, nil >>`,
		Id:         "At",
		NTType:     38,
		Index:      111,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "@"}, nil
		},
	},
	ProdTabEntry{
		String: `At : Space "@"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "@",
    }, nil >>`,
		Id:         "At",
		NTType:     38,
		Index:      112,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "@",
			}, nil
		},
	},
	ProdTabEntry{
		String: `RightArrow : "->"	<< &expr.Keyword{Value: "->"}, nil >>`,
		Id:         "RightArrow",
		NTType:     39,
		Index:      113,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "->"}, nil
		},
	},
	ProdTabEntry{
		String: `RightArrow : Space "->"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "->",
    }, nil >>`,
		Id:         "RightArrow",
		NTType:     39,
		Index:      114,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "->",
			}, nil
		},
	},
	ProdTabEntry{
		String: `EqualEqual : "=="	<< &expr.Keyword{Value: "=="}, nil >>`,
		Id:         "EqualEqual",
		NTType:     40,
		Index:      115,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "=="}, nil
		},
	},
	ProdTabEntry{
		String: `EqualEqual : Space "=="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "==",
    }, nil >>`,
		Id:         "EqualEqual",
		NTType:     40,
		Index:      116,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "==",
			}, nil
		},
	},
	ProdTabEntry{
		String: `ExclamationEqual : "!="	<< &expr.Keyword{Value: "!="}, nil >>`,
		Id:         "ExclamationEqual",
		NTType:     41,
		Index:      117,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "!="}, nil
		},
	},
	ProdTabEntry{
		String: `ExclamationEqual : Space "!="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "!=",
    }, nil >>`,
		Id:         "ExclamationEqual",
		NTType:     41,
		Index:      118,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "!=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `LessThan : "<"	<< &expr.Keyword{Value: "<"}, nil >>`,
		Id:         "LessThan",
		NTType:     42,
		Index:      119,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "<"}, nil
		},
	},
	ProdTabEntry{
		String: `LessThan : Space "<"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "<",
    }, nil >>`,
		Id:         "LessThan",
		NTType:     42,
		Index:      120,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "<",
			}, nil
		},
	},
	ProdTabEntry{
		String: `GreaterThan : ">"	<< &expr.Keyword{Value: ">"}, nil >>`,
		Id:         "GreaterThan",
		NTType:     43,
		Index:      121,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ">"}, nil
		},
	},
	ProdTabEntry{
		String: `GreaterThan : Space ">"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ">",
    }, nil >>`,
		Id:         "GreaterThan",
		NTType:     43,
		Index:      122,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ">",
			}, nil
		},
	},
	ProdTabEntry{
		String: `LessEqual : "<="	<< &expr.Keyword{Value: "<="}, nil >>`,
		Id:         "LessEqual",
		NTType:     44,
		Index:      123,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "<="}, nil
		},
	},
	ProdTabEntry{
		String: `LessEqual : Space "<="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "<=",
    }, nil >>`,
		Id:         "LessEqual",
		NTType:     44,
		Index:      124,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "<=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `GreaterEqual : ">="	<< &expr.Keyword{Value: ">="}, nil >>`,
		Id:         "GreaterEqual",
		NTType:     45,
		Index:      125,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: ">="}, nil
		},
	},
	ProdTabEntry{
		String: `GreaterEqual : Space ">="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: ">=",
    }, nil >>`,
		Id:         "GreaterEqual",
		NTType:     45,
		Index:      126,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  ">=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `TildeEqual : "~="	<< &expr.Keyword{Value: "~="}, nil >>`,
		Id:         "TildeEqual",
		NTType:     46,
		Index:      127,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "~="}, nil
		},
	},
	ProdTabEntry{
		String: `TildeEqual : Space "~="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "~=",
    }, nil >>`,
		Id:         "TildeEqual",
		NTType:     46,
		Index:      128,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "~=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `StarEqual : "*="	<< &expr.Keyword{Value: "*="}, nil >>`,
		Id:         "StarEqual",
		NTType:     47,
		Index:      129,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "*="}, nil
		},
	},
	ProdTabEntry{
		String: `StarEqual : Space "*="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "*=",
    }, nil >>`,
		Id:         "StarEqual",
		NTType:     47,
		Index:      130,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "*=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `CaretEqual : "^="	<< &expr.Keyword{Value: "^="}, nil >>`,
		Id:         "CaretEqual",
		NTType:     48,
		Index:      131,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "^="}, nil
		},
	},
	ProdTabEntry{
		String: `CaretEqual : Space "^="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "^=",
    }, nil >>`,
		Id:         "CaretEqual",
		NTType:     48,
		Index:      132,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "^=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `DollarEqual : "$="	<< &expr.Keyword{Value: "$="}, nil >>`,
		Id:         "DollarEqual",
		NTType:     49,
		Index:      133,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "$="}, nil
		},
	},
	ProdTabEntry{
		String: `DollarEqual : Space "$="	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "$=",
    }, nil >>`,
		Id:         "DollarEqual",
		NTType:     49,
		Index:      134,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "$=",
			}, nil
		},
	},
	ProdTabEntry{
		String: `ColonColon : "::"	<< &expr.Keyword{Value: "::"}, nil >>`,
		Id:         "ColonColon",
		NTType:     50,
		Index:      135,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{Value: "::"}, nil
		},
	},
	ProdTabEntry{
		String: `ColonColon : Space "::"	<< &expr.Keyword{
      Before: X[0].(*expr.Space),
      Value: "::",
    }, nil >>`,
		Id:         "ColonColon",
		NTType:     50,
		Index:      136,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Keyword{
				Before: X[0].(*expr.Space),
				Value:  "::",
			}, nil
		},
	},
	ProdTabEntry{
		String: `Space : Space space	<< expr.AppendSpace(X[0], newString(X[1])), nil >>`,
		Id:         "Space",
		NTType:     51,
		Index:      137,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return expr.AppendSpace(X[0], newString(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Space : space	<< &expr.Space{Space: []string{newString(X[0])}}, nil >>`,
		Id:         "Space",
		NTType:     51,
		Index:      138,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &expr.Space{Space: []string{newString(X[0])}}, nil
		},
	},
}
