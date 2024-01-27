package logus

import (
	"github.com/darklab8/darklab_goutils/goutils/logus_core"
)

var (
	Log *logus_core.Logger
)

func init() {
	Log = logus_core.NewLogger("darkblog")
}
