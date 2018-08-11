package logrusbolt

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"strings"
)

func TestWrite(t *testing.T) {
	bufferOut := bytes.NewBufferString("")

	config := BoltHook{
		Bucket:    "test",
		Formatter: &logrus.JSONFormatter{},
		DBLoc:     "/tmp/test.db",
	}

	log := logrus.New()
	log.Out = bufferOut

	// Create boltdb
	hook, _ := NewHook(config)

	log.Hooks.Add(hook)

	// Create log
	log.Info("test info")

	if !strings.Contains(bufferOut.String(), `"test info"`) {
		t.Errorf("expected logrus message to have '%s', but got %#v", `"test info"`, bufferOut.String())
	}
}
