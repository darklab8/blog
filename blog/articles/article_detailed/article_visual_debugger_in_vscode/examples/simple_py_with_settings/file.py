from dataclasses import dataclass
from copy import copy
import os

@dataclass
class Example:
    foo: int
    bar: str
    is_smth: bool = False

def func(e: Example) -> Example:
    e = copy(e)
    e.foo += 5
    return e

if __name__=="__main__":
    a = Example(5, os.environ["SOME_VAR"])
    b = func(a)
    print(b)
    assert a.bar == "abc"

