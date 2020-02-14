package throw

import (
	"fmt"
	"io"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type errType1 struct {
	m string
}

func (errType1) Error() string {
	return ""
}

type errType2 struct {
	m func() // incomparable
}

func (errType2) Error() string {
	return ""
}

func TestIsEqual(t *testing.T) {
	require.Panics(t, func() {
		var e1, e2 error = errType2{}, errType2{}
		runtime.KeepAlive(e1 == e2)
	})

	require.True(t, IsEqual(errType1{}, errType1{}))
	require.False(t, IsEqual(errType1{"A"}, errType1{"B"}))

	require.False(t, IsEqual(errType2{}, errType2{}))

	require.False(t, IsEqual(errType1{}, errType2{}))
	require.False(t, IsEqual(errType2{}, errType1{}))
}

type errBuilder struct {
	bottomErr error
}

func (v errBuilder) _err0() error {
	return WithDetails(v.bottomErr, Unsupported())
}

func (v errBuilder) _err1() error {
	return WithStackAndDetails(v._err0(), struct {
		msg string
		v0  int
	}{"err1Txt", 1})
}

func (v errBuilder) _err2() error {
	return WithStack(v._err1())
}

func (v errBuilder) _err3() error {
	panic(v._err2())
}

func (v errBuilder) _err4() (err error) {
	defer func() {
		err = RM(recover(), err, "panicCatch", struct{ position int }{7})
	}()
	return v._err3()
}

func newChain(bottom error) error {
	return errBuilder{bottom}._err4()
}

func TestWrapPanicExt(t *testing.T) {
	err := WrapPanicExt("test", 0)
	st := OutermostStack(err).StackTrace()
	s := st.StackTraceAsText()
	methodName := "github.com/insolar/assured-ledger/ledger-core/v2/vanilla/throw.TestWrapPanicExt"
	require.True(t, strings.HasPrefix(st.StackTraceAsText(), methodName), "missing method: %s", s)
}

func TestStackOf(t *testing.T) {
	errChain := newChain(io.EOF)
	fmt.Println(errChain)
}
