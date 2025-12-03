package stdlib_test

import (
	"os"
	"testing"

	"github.com/malivvan/cui/runtime"
	"github.com/malivvan/cui/runtime/require"
)

func TestReadFile(t *testing.T) {
	content := []byte("the quick brown fox jumps over the lazy dog")
	tf, err := os.CreateTemp("", "test")
	require.NoError(t, err)
	defer func() { _ = os.Remove(tf.Name()) }()

	_, err = tf.Write(content)
	require.NoError(t, err)
	_ = tf.Close()

	module(t, "os").call("read_file", tf.Name()).
		expect(&runtime.Bytes{Value: content})
}

func TestReadFileArgs(t *testing.T) {
	module(t, "os").call("read_file").expectError()
}
func TestFileStatArgs(t *testing.T) {
	module(t, "os").call("stat").expectError()
}

func TestFileStatFile(t *testing.T) {
	content := []byte("the quick brown fox jumps over the lazy dog")
	tf, err := os.CreateTemp("", "test")
	require.NoError(t, err)
	defer func() { _ = os.Remove(tf.Name()) }()

	_, err = tf.Write(content)
	require.NoError(t, err)
	_ = tf.Close()

	stat, err := os.Stat(tf.Name())
	if err != nil {
		t.Logf("could not get tmp file stat: %s", err)
		return
	}

	module(t, "os").call("stat", tf.Name()).expect(&runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			"name":      &runtime.String{Value: stat.Name()},
			"mtime":     &runtime.Time{Value: stat.ModTime()},
			"size":      &runtime.Int{Value: stat.Size()},
			"mode":      &runtime.Int{Value: int64(stat.Mode())},
			"directory": runtime.FalseValue,
		},
	})
}

func TestFileStatDir(t *testing.T) {
	td, err := os.MkdirTemp("", "test")
	require.NoError(t, err)
	defer func() { _ = os.RemoveAll(td) }()

	stat, err := os.Stat(td)
	require.NoError(t, err)

	module(t, "os").call("stat", td).expect(&runtime.ImmutableMap{
		Value: map[string]runtime.Object{
			"name":      &runtime.String{Value: stat.Name()},
			"mtime":     &runtime.Time{Value: stat.ModTime()},
			"size":      &runtime.Int{Value: stat.Size()},
			"mode":      &runtime.Int{Value: int64(stat.Mode())},
			"directory": runtime.TrueValue,
		},
	})
}

func TestOSExpandEnv(t *testing.T) {
	curMaxStringLen := runtime.MaxStringLen
	defer func() { runtime.MaxStringLen = curMaxStringLen }()
	runtime.MaxStringLen = 12

	_ = os.Setenv("VV", "FOO BAR")
	module(t, "os").call("expand_env", "$VV").expect("FOO BAR")

	_ = os.Setenv("VV", "FOO")
	module(t, "os").call("expand_env", "$VV $VV").expect("FOO FOO")

	_ = os.Setenv("VV", "123456789012")
	module(t, "os").call("expand_env", "$VV").expect("123456789012")

	_ = os.Setenv("VV", "1234567890123")
	module(t, "os").call("expand_env", "$VV").expectError()

	_ = os.Setenv("VV", "123456")
	module(t, "os").call("expand_env", "$VV$VV").expect("123456123456")

	_ = os.Setenv("VV", "123456")
	module(t, "os").call("expand_env", "${VV}${VV}").
		expect("123456123456")

	_ = os.Setenv("VV", "123456")
	module(t, "os").call("expand_env", "$VV $VV").expectError()

	_ = os.Setenv("VV", "123456")
	module(t, "os").call("expand_env", "${VV} ${VV}").expectError()
}
