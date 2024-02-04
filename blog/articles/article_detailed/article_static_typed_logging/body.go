package article_static_typed_logging

import (
	"strings"
	"time"
)

type Vars struct {
	StaticRoot string
	GoTypelog  string
	PyTypelog  string
}

var (
	Headline = "Typelog - type safe structured logging"
	Intro    = strings.ReplaceAll(`With modern logging systems able to parse JSON out of the box, we need defining easily jsonable logs.
Known solutions do not do it consistently and in a type safe way. Typelog comes to rescue.`, "\n", "")
	Date = time.Date(2024, time.January, 28, 0, 0, 0, 0, time.UTC)
)
