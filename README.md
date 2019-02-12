# Logrus BoltDB Hook

With this hook logrus saves messages in the [BoltDB](https://github.com/coreos/bbolt)

## Install

```bash
$ go get github.com/trK54Ylmz/logrus-boltdb-hook
```

## Usage

```go
package main

import (
	"github.com/sirupsen/logrus"
	logrusbolt "github.com/trK54Ylmz/logrus-boltdb-hook"
)

func init() {
	config := logrusbolt.BoltHook{
		Bucket:    "test",
		Formatter: &logrus.JSONFormatter{},
		DBLoc:     "/tmp/test.db",
	}
	
	hook, err := logrusbolt.NewHook(config)
	
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Error(err)
	}
}


func main() {
	logrus.Info("test info")
}
```
