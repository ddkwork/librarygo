package clang_test

import (
	"github.com/ddkwork/librarygo/src/clang"
	"github.com/ddkwork/librarygo/src/mycheck"
	"testing"
)

func TestClangFormat(t *testing.T) {
	c := clang.New()
	assert := mycheck.Assert(t)
	assert.True(c.WriteClangFormatBody("D:\\codespace\\workspace\\src\\cppkit\\hyperdbgui"))
	assert.True(c.Format("D:\\codespace\\workspace\\src\\cppkit\\hyperdbgui\\IopXxxControlFile.ds"))
}
