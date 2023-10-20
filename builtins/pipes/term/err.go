//go:build !js
// +build !js

package term

import (
	"os"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/tty"
	"github.com/lmorg/murex/utils"
)

// Terminal: Standard Error

// Err is the Stderr interface for term
type Err struct {
	term
}

func (t *Err) File() *os.File {
	return os.Stderr
}

// SetDataType is a null method because the term interface is write-only
func (t *Err) SetDataType(string) {}

// Write is the io.Writer() interface for term
func (t *Err) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	t.mutex.Unlock()

	i, err = tty.Stderr.Write(b)
	if err != nil {
		tty.Stdout.WriteString(err.Error())
	}

	return
}

// Writeln writes an OS-specific terminated line to the stderr
func (t *Err) Writeln(b []byte) (int, error) {
	return t.Write(appendBytes(b, utils.NewLineByte...))
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *Err) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(t, dataType)
}
