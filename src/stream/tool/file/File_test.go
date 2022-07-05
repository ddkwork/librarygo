package file_test

import (
	"github.com/ddkwork/librarygo/src/stream/tool/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	body := `
-Xms256m
-Xmx2000m
-XX:ReservedCodeCacheSize=512m
-Xss2m
-XX:NewSize=128m
-XX:MaxNewSize=128m
-XX:+IgnoreUnrecognizedVMOptions
-XX:+UseG1GC
-XX:SoftRefLRUPolicyMSPerMB=50
-XX:CICompilerCount=2
-XX:+HeapDumpOnOutOfMemoryError
-XX:-OmitStackTraceInFastThrow
-ea
-Dsun.io.useCanonCaches=false
-Djdk.http.auth.tunneling.disabledSchemes=""
-Djdk.attach.allowAttachSelf=true
-Djdk.module.illegalAccess.silent=true
-Dkotlinx.coroutines.debug=off
-Dsun.tools.attach.tmp.only=true
`
	f := file.New()
	ok := f.WriteTruncate("clion64.vmoptions", body)
	assert.True(t, ok)
	return
	assert.True(t, f.WriteAppend("1.txt", "111"))
	assert.True(t, f.WriteAppend("k/1.txt", "222"))
	f.Copy("include/Common/EABase", "new/include")
}
