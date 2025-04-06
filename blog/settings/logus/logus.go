package logus

import "github.com/darklab8/go-typelog/typelog"

var Log *typelog.Logger = typelog.NewLogger("darkblog", typelog.WithLogLevel(typelog.LEVEL_INFO))
