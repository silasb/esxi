package parser

import (
	"io"

	"github.com/alecthomas/participle"
)

type Parser struct {
	parser *participle.Parser
}

func NewParser() (*Parser, error) {
	parser, err := participle.Build(&Config{})

	if err != nil {
		return nil, err
	}

	return &Parser{parser}, nil
}

func (p *Parser) Parse(io io.Reader) (*Config, error) {
	expr := &Config{}
	err := p.parser.Parse(io, expr)
	if err != nil {
		return expr, err
	}

	return expr, nil
}

type Bool bool

func (b *Bool) Capture(v []string) error { *b = v[0] == "true"; return nil }

type Value struct {
	Boolean    *Bool    `  @("true"|"false")`
	Identifier *string  `| @Ident { @"." @Ident }`
	String     *string  `| @(String|Char|RawString)`
	Number     *float64 `| @(Float|Int)`
	Minus      *Value   `| "-" @@`
	Null       *string  `| "<" "unset" ">"`
	// Array      []*Value `| "[" [ @@ { "," @@ } ] "]"`
	Array []*Entry `| "[" { @@ [ "," ] } "]" [ "," ]`
}

type Entry struct {
	Key   string `[ @Ident "=" ]`
	Type  string `[ "(" @Ident { @"." @Ident } ")" ]`
	Value *Value `( @@ [ "," ]`
	Block *Block `| @@ )`
}

type Block struct {
	// Parameters []*Value `{ @@ }`
	Entries []*Entry `"{" { @@ } "}" [ "," ]`
}

type Config struct {
	Config []*Entry `{ @@ }`
}
