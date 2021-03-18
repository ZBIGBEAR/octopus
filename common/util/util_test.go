package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDir(t *testing.T) {
	path1 := "abc"
	b, err := IsDir(path1)
	assert.NotNil(t, err)
	assert.Equal(t, false, b)

	path2 := "F:\\workspace\\go\\src\\awesomeProject\\octopus\\common\\util"
	b, err = IsDir(path2)
	assert.Nil(t, err)
	assert.Equal(t, true, b)

	path3 := "F:\\workspace\\go\\src\\awesomeProject\\octopus\\common\\."
	b, err = IsDir(path3)
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}

func TestIsFile(t *testing.T) {
	path1 := "F:\\workspace\\go\\src\\awesomeProject\\octopus\\common\\util\\a.txt"
	b, err := IsFile(path1)
	assert.NotNil(t, err)
	assert.Equal(t, false, b)

	path2 := "F:\\workspace\\go\\src\\awesomeProject\\octopus\\common\\util\\util.go"
	b, err = IsFile(path2)
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}
