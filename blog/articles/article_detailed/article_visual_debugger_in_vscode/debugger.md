# For whom the article

The article is written for all those people using Pycharm but wishing to switch to Vscode for a reason. And also for beginners not knowing how to setup the visual debugger in vscode for python yet.

The article is written on request of multiple people who work with python through dev env setups with docker-compose, and needing connect to already existing containers and start dealing with python efficiently in it.

In this article we also cover ability to run visual debug for specific unit tests, and stress importance of doing it, because it helps us having comfortable entrypoints into our code with maximum visibility. That makes test driven development very plausable, or at least development when tests are written at the end of an atomic change made by pull request, for the reasons of having ability to see rapid feedback and trying to your code parts literally from within unit tests.

Article as main point guides people how to setup things for working with Python for Django *in already running* Docker containers, but we provide also explanation of doing it without docker, and also for Fastapi, and Flask for comparison.
Previous zero knowledge of vscode usage is assumed. Knowledge to operate `pip` and `venv` is assumed being present.

All examples provided in article are foundable in [examples](<https://github.com/darklab8/blog/tree/master/blog/articles/article_detailed/article_visual_debugger_in_vscode/examples>) folder. **We assume u will be opening vscode as working directory from within specific example folder** like
examples/simple_pyscript or examples/django_example, so that `.vscode` folder will be in the root of your file path.

P.S. In normal projects reopening workdir with `code -r .` is not required, as we open projects only once, but we will be often using it in this article to switch between multiple working directories when necessary.

# Python without docker

- We make assumption we utilize [flat python structure](https://packaging.python.org/en/latest/discussions/src-layout-vs-flat-layout/), because it is simple and makes project working maximum out of the box with minimum of extra wtfs.
- We also assume we will be utilizing unit testing :]
- As mentioned, as assume u will be opening as "Working Directory" specific example folders like examples/simple_pyscript.
    - `code -r .` shortcut allows quickly reopening vscode with changing working directory.
- We also assume you are using Linux as main OS, or at least using WSL2 with linux opened inside if u are developing at Windows or MacOS.
    - (the article material is tested to work on Kubuntu 22.04 LTS)

First assure u installed `Pytest Explorer` extension. It will autoinstall Python extension for intellisence along the way.
- pytest explorer id for search `@id:littlefoxteam.vscode-python-test-adapter`
- it will autoinstall Python extension along side: `@id:ms-python.python`
![]({{.StaticRoot}}article_visual_debugger_in_vscode/extension_pytest1.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/extension_pytest2.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/extension_pytest3.png)

If smth is working weird, like syntax highlighting by colors is not present, make sure to run Ctrl + Shift + P and write Reload Window. Select and apply, it quickly reloads vscode, including reloading extensions.
![]({{.StaticRoot}}article_visual_debugger_in_vscode/ctrl_shift_p_reloadwindow.png)

After that python simple file debugging is launchable as
- git clone/download https://github.com/darklab8/blog/tree/master/
- cd blog/articles/article_detailed/article_visual_debugger_in_vscode/examples
- cd simple_pyscript && code -r . # reopening to different workdir

just selecting Run -> Star Debugging (F5)

![]({{.StaticRoot}}article_visual_debugger_in_vscode/debug_python_script.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/debug_python_script2.png)

Now lets check it working for pytest test debug

- create venv `python3 -m venv .venv`
- `source .venv/bin/activate` (at windows can be `venv\Scripts\activate`)
- pip install -r requirements.txt

![]({{.StaticRoot}}article_visual_debugger_in_vscode/simple_pytest1.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/simple_pytest2.png)

# Basic functions of a visual debug in vscode

- on hover u see data. U can open nested data inside
- F10 step forward
- F11 step forward and deeper
- F5 continue
- use mouse to set red breakpoints around

# Python with settings for env vars

- cd ../simple_py_with_settings && code -r .

env vars are inputtable for different oses in settings.json

.vscode/settings.json
```json
{
    "python.testing.pytestArgs": [
        "."
    ],
    "python.testing.unittestEnabled": false,
    "python.testing.pytestEnabled": true,
    "terminal.integrated.env.linux": {
        "SOME_VAR": "abc"
    },
    "terminal.integrated.env.windows": {
        "SOME_VAR": "abc"
    },
    "terminal.integrated.env.osx": {
        "SOME_VAR": "abc"
    }
}
```
```
$ python3 file.py 
Example(foo=10, bar='abc', is_smth=False)
```

those env vars will be available for access from within launched web server and unit tests during visual debug usage too

# Python with running debug for django

cd ../django_example && code -r .

set breakpoint in a view at urls.py and launch Python Django debug, like in this picture below
![]({{.StaticRoot}}article_visual_debugger_in_vscode/red_breakpoint.png)
debug for django will work due to the present .vscode/launch.json containing
```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Python: Django",
            "type": "debugpy",
            "request": "launch",
            "program": "${workspaceFolder}/manage.py",
            "args": ["runserver"],
            "django": true,
            "justMyCode": false
          },
        {
            "name": "Python: Debug Tests", // KO with GUI (tests tab)
            "type": "debugpy",
            "request": "launch",
            "program": "${file}",
            "purpose": ["debug-test"],
            "console": "integratedTerminal",
            "justMyCode": false
        },
        {
            "name": "Debug specific tests", // OK with F5
            "type": "debugpy",
            "module": "pytest",
            "request": "launch",
            "purpose": ["debug-test"],
            "console": "integratedTerminal",
            "justMyCode": false,
            "args": ["test_sample.py::test_answer"]
        },
        {
            "name": "Python: Current File", // OK with F5
            "type": "debugpy",
            "request": "launch",
            "program": "${file}",
            "console": "internalConsole",
            "justMyCode": false
        }
    ]
}
```

out of all those settings, only Python: Django part is important to us at the moment for running debug via manage.py.
Take note that u can technically input arbitary arguments (under `args` key) to run it for execution of more different stuff 
![]({{.StaticRoot}}article_visual_debugger_in_vscode/django1.png)

take note of setting `"justMyCode": false`, it allows u navigating during visual debug third party libs too. Otherwise they will be skiped.
Crucial thing when debugging internal company libs, or when developing your own library that is doing smth with known third party libs.

![]({{.StaticRoot}}article_visual_debugger_in_vscode/django2.png)

one of the other settings written in launch.json will ensure u can utilize third party library code navigation `(justMyCode: false)` even when you use visual debug from unit test.
![]({{.StaticRoot}}article_visual_debugger_in_vscode/django3.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/django4.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/django5.png)

# Python with running debug for fastapi

- cd ../fastapi/example && code -r .

we create again venv, pip install -r requirements.txt, select new venv (at the bottom-right of the IDE. Select path to venv/bin/python3 manually if necessary there).
if smth glitches, we do `Ctrl + Shift + P -> Reload Window` trick

with small fixing to launch.json, we have it adapted for fastapi now
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Python: Fastapi",
            "type": "debugpy",
            "request": "launch",
            "program": "${workspaceFolder}/app.py",
            "justMyCode": false
          },
        {
            "name": "Python: Debug Tests", // KO with GUI (tests tab)
            "type": "debugpy",
            "request": "launch",
            "program": "${file}",
            "purpose": ["debug-test"],
            "console": "integratedTerminal",
            "justMyCode": false
        },
        {
            "name": "Debug specific tests", // OK with F5
            "type": "debugpy",
            "module": "pytest",
            "request": "launch",
            "purpose": ["debug-test"],
            "console": "integratedTerminal",
            "justMyCode": false,
            "args": ["test_sample.py::test_answer"]
        },
        {
            "name": "Python: Current File", // OK with F5
            "type": "debugpy",
            "request": "launch",
            "program": "${file}",
            "console": "internalConsole",
            "justMyCode": false
        }
    ]
}
```

![]({{.StaticRoot}}article_visual_debugger_in_vscode/fastapi1.png)

with written small code examples in app_test.py we have no trouble launching visual debug for unit tests of fastapi

![]({{.StaticRoot}}article_visual_debugger_in_vscode/fastapi2.png)

see folder `flask_example` for almost same example to try for Flask too.

# Python in already existing Docker

Ensure having installed Dev Containers app @id:ms-vscode-remote.remote-containers
![]({{.StaticRoot}}article_visual_debugger_in_vscode/docker_extension_install1.png)

- go to examples/django_example folder and run `docker compose build` and `docker compose run --service-ports shell`
    - `--service-ports` option ensures we forwarded 8000 port written in docker-compose in case of wishing to runserver and seeing it
- check the settings configured to make things working in `docker-compose.yml`.
    - We forwarded current volume insode to path /code, we made sure working directory in container /code too.

P.S. instead of usage of docker-compose, same is achievable with regular docker as `docker build --tag test .` and then `docker run -it -v $(pwd):/code -w /code --name shell --entrypoint=bash test`

Now we can enter already running container by using "Attach Visual Studio Code" option.
This option will open second instance of vscode, but already from within container.
![]({{.StaticRoot}}article_visual_debugger_in_vscode/docker_extension_install2.png)

Install Pytest Explorer `@id:littlefoxteam.vscode-python-test-adapter` (with autoinstalled Python inside)
and if necessary reload window by using `Ctrl + Shift + P -> Reload Window`

And go to working directory /code

if u did everything right, interface will be showing ability to run specific tests in debug mode like we did previously.
And u can also launch web server by using button from Debug menu on the left with selection of "Python: Django"
![]({{.StaticRoot}}article_visual_debugger_in_vscode/docker3.png)

to ensure that we are able to see dev server from the container, ensure we binded it to 0.0.0.0 instead of localhost!
that is doable with `python3 manage.py runserver 0.0.0.0:8000` for django, we added missing argument into launch.json.

![]({{.StaticRoot}}article_visual_debugger_in_vscode/docker4.png)
after setting breakpoint in view, and visiting `http://localhost:8000/polls` we get being stopped in breakpoint as desired.

Same steps are doable for FastAPI and Flask.
We added Dockerfile and docker-compose.yml for the same into fastapi_example and flask_example folders.

P.S. when u are in vscode from within container, opened terminal is opened automatically for within the container shell

# Ending

Congratulations, now u are able to use visual debug for python common web frameworks without and with docker for web server run itself and unit tests!

With showing examples how it is be done for two different python frameworks, where for fastapi we adjusted launch.json on a fly,
it is my intention u will be able to see how to do it similarily for any other existing python framework.

I encourage you trying to develop from within unit tests written with visual debug, since it is actually a comfortable way of having testing as part of a working code writing. Doing that will ensure most rapid feedback, and ensure u a writing better code in all kind of code internal characteristics.
Do check books like [TDD by kent beck]({{.SiteRoot}}favourite.html#TestDrivenDevelopmentByExample) and [Unit testing by Vladimir Khorikov]({{.SiteRoot}}favourite.html#UnitTestingPrinciplesPracticesandPatterns) in order to find out more about unit testing that enables writing maintainable code.

We showed examples with Docker, because modern web development could be having different C dependended libraries, that without docker is very struggling to install otherwise. Also docker makes universal way to setup dev envs for many people. Some companies just keep dev env straight in docker only for those reasons.
