package utils

import (
	"regexp"

	utils_logus1 "github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_logus"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

func InitRegexExpression(regex **regexp.Regexp, expression utils_types.RegExp) {
	var err error

	*regex, err = regexp.Compile(string(expression))
	utils_logus.Log.CheckFatal(err, "failed to init regex", utils_logus.Regex(expression), utils_logus1.FilePath(GetCurrentFile()))
}
