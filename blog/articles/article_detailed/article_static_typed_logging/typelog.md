**Keeping logs clean, rich and useful**

## Intro

With modern logging systems able to parse JSON out of the box, we need to define easily jsonable logs. The alternative to define regex for parsing logs is a rather dreadful fragile experience limited in its nature.
Known solutions do not do it consistently and in a type safe way. It is easy to shoot into your own legs with them.

Structured logging became part of std library [for golang](https://go.dev/blog/slog) ,
as well part of official documentation [for python](https://docs.python.org/3/howto/logging-cookbook.html#implementing-structured-logging). Golang managed to reach a rather big comfort out of the box, but it has some parts missing... like type safety. U can insert Any objects into logging msgs easily. [Python std way](https://docs.python.org/3/howto/logging-cookbook.html#implementing-structured-logging) is in even more raw state. And they all lack a way to do it in a type safe way.

For reading {{.LinkTypeSafety}}

## Solution

<p align="center">
  <img src="{{.StaticRoot}}typelog/typelog.png" style="width: 200px; height: 200px;"/>
</p>

Typelog comes to bring best parts out of different solutions, and having added typing safety against shooting into your own legs! Presence of type **Any** is as maximum as possible removed from its public interface parts.

### [Go version]({{.GoTypelog}})

if u will try inserting invalid to log objects, u will get errors before runtime!

```go
package types

type TaskID string

type WorkerID int
```
<span></span>

```go
package typedlogs

import (
	"log/slog"

	"github.com/darklab8/go-typelog/examples/types"
	"github.com/darklab8/go-typelog/typelog"
)

func TaskID(value types.TaskID) typelog.LogType { return typelog.String("task_id", string(value)) }

func WorkerID(value types.WorkerID) typelog.LogType {
	return typelog.Int("worker_id", int(value))
}
```
<span></span>

```go
package examples

var logger *typelog.Logger = typelog.NewLogger("worker")

func TestTypedLogs(t *testing.T) {
	worker_id := types.WorkerID(5)
	logger.Info("Worker was started", typedlogs.WorkerID(worker_id))

  logger.Info("Worker was started", 123123, "asdasd") // TYPING ERROR

	logger := logger.Log.WithFields(typedlogs.WorkerID(worker_id), typedlogs.TaskID("abc"))
	logger.Info("Worker started task")
	logger.Info("Worker finished task")
}
```

[link to go-typelog repository]({{.GoTypelog}})

### [Python version]({{.PyTypelog}})

In order to use py-typelog in python to its full capacity u will need to turn on mypy or pyright in [preferably strict mod](https://careers.wolt.com/en/blog/tech/professional-grade-mypy-configuration).
py-typelog comes with mypy typelog-stubs as part of its package.

```py
# types.py
from dataclasses import dataclass
from typing import NewType

TaskID = NewType("TaskID", int)


@dataclass(frozen=True)
class Task:
    smth: str
    b: int
```
<span></span>

```python
# logtypes.py
from typing import Any, Dict

from typelog import LogType

from . import types


def TaskID(value: types.TaskID) -> LogType:
    def wrapper(params: Dict[str, Any]) -> None:
        params["task_id"] = str(value)

    return wrapper


def Task(value: types.Task) -> LogType:
    def wrapper(params: Dict[str, Any]) -> None:
        params.update(value.__dict__)

    return wrapper
```
<span></span>

```python
import logging
import unittest

import typelog
from typelog import LogConfig, Loggers, get_logger
from typelog.types import LibName, LogLevel, RootLogLevel

from . import logtypes, types

logger = get_logger(__name__)

class TestExamples(unittest.TestCase):
    def setUp(self) -> None:
        Loggers(
            RootLogLevel(logging.DEBUG),
            LogConfig(LibName("examples"), LogLevel(logging.DEBUG)),
            add_time=True,
        ).configure()

    def test_basic(self) -> None:
        logger.warn("Writing something", logtypes.TaskID(types.TaskID(123)))

    def test_another_one(self) -> None:
        task = types.Task(smth="abc", b=4)
        logger.warn("Writing something", logtypes.Task(task))

        logger.warn("Writing something", 123123, "324234") # TYPING ERROR
```

[link to py-typelog repository]({{.PyTypelog}})

## What brings typelog?

- You are type safety protected against mistakes in logging!
- Parsing entities to logging records is decoupled from your working code
- You keep consistent same keys for same variables
- Easier to refactor logs, you can rename them, change keys or parsing rules in one click or few code lines across application
- Easier refactoring, brings better Domain language defined across your application
- You will be able to query your JSON logs in modern logging solutions, with capturing more useful logs in a query because of necessary identifiers present consistently.

## Concluding words

Once you tap into the power of context rich logging, you gain another powerful tool for easier program debugging on runtime. The power of it should not be understimated. Some problems are detectable only at runtime, only in production. It is univetable. It is very nice to have useful logs as source of data to investigate when such problems appear. Also, better logging brings easier time debugging unit tests as well, especially if turning third-party libraries to a warning level and turning on your logging level to debug during test runs.
