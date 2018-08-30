package parser

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	_, err := NewParser()

	assert.Nil(t, err)
}

func TestNewParserWithFailure(t *testing.T) {
	_, err := NewParser()
	assert.Nil(t, err)
}

func TestParseIOWithEmptyString(t *testing.T) {
	p, _ := NewParser()

	b := bytes.NewBufferString("")
	config, err := p.Parse(b)
	assert.Nil(t, err)

	assert.Equal(t, len(config.Config), 0)
}

func TestParseIOSimpleConfig(t *testing.T) {
	p, _ := NewParser()

	b := bytes.NewBufferString(`var = (type) {
		hello = "world"
	}`)

	config, err := p.Parse(b)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, len(config.Config), 1)

	entry0 := config.Config[0]
	assert.Equal(t, entry0.Key, "var")
	assert.Equal(t, entry0.Type, "type")

	value := config.Config[0].Value
	assert.Nil(t, value)

	block := config.Config[0].Block
	assert.Equal(t, len(block.Entries), 1)

	entry := block.Entries[0]
	assert.Equal(t, entry.Key, "hello")

	s := *entry.Value.String
	assert.Equal(t, s, "world")
}

func TestBool(t *testing.T) {
	p, _ := NewParser()

	b := bytes.NewBufferString(`var = (type) {
		hello = "true"
	}`)

	config, err := p.Parse(b)
	assert.Nil(t, err)

	block := config.Config[0].Block
	assert.Equal(t, len(block.Entries), 1)

	entry := block.Entries[0]
	assert.Equal(t, entry.Key, "hello")

	s := *entry.Value.Boolean
	assert.Equal(t, s, Bool(true))
}
