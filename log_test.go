package gu

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/bouk/monkey"
)

func TestLogFatalIfNotNil(t *testing.T) {
	var b bytes.Buffer
	log.SetOutput(&b)

	monkey.Patch(log.Fatal, func(v ...interface{}) {
		var std = log.New(os.Stderr, "", log.LstdFlags)
		std.Output(2, fmt.Sprint(v...))
		log.Print(v)
	})

	message := "this is test"
	LogFatalIfNotNil(nil, message)
	if b.Len() != 0 {
		t.Errorf("any stderr shoud not be output but %d length stderr output", b.Len())
	}

	LogFatalIfNotNil(errors.New("test error"), message)
	//if b.Len() == 0 || !strings.Contains(message, b.String()) {
	if b.Len() == 0 {
		t.Error("any stderr shoud be output but actual is not output at all")
	} else if !strings.Contains(b.String(), message) {
		t.Errorf("%q should be contained in %q but not", message, b.String())
	}
}
