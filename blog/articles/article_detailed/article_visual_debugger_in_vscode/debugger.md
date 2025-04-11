# For whom the article is

This article is written for all those people using Pycharm but wishing to switch to VSCode for a reason. And also for beginners not knowing how to setup the visual debugger in VSCode for python yet.

The article is written at the request of multiple people who work with python through dev env setups with docker-compose, and needing connect to already existing containers and start dealing with python efficiently in it.

In the article, we also cover the ability to run visual debug for specific unit tests, and stress the importance of doing it, because it helps us having comfortable entrypoints into our code with maximum visibility. That makes test-driven development very plausible, or at least development when tests are written at the end of an atomic change made by pull request, for the reasons of having the ability to see rapid feedback and trying to your code parts literally from within unit tests.

Article as main point guides people how to setup things for working with Python for Django *in already running* Docker containers, but we provide also explanation of doing it without docker, and also for Fastapi, and Flask for comparison.
Previous zero knowledge of VSCode usage is assumed. Basic knowledge of how to use `pip` and `venv` is assumed.

All examples in the article are in the [examples](https://github.com/darklab8/blog/tree/master/blog/articles/article_detailed/article_visual_debugger_in_vscode/examples) folder. 
**We assume you’re opening VSCode with the working directory set to a specific example folder** (e.g., `examples/simple_pyscript` or `examples/django_example`) so that the `.vscode` folder is at the project root.

P.S. In normal projects reopening workdir with `code -r .` is not required, as we open projects only once, but we will be often using it in this article to switch between multiple working directories when necessary.

# Python without docker

- We make the assumption we utilize [flat python structure](https://packaging.python.org/en/latest/discussions/src-layout-vs-flat-layout/), because it is simple and makes a project at maximum working out of the box with a minimum of extra wtf.
- We also assume we will be utilizing unit testing :]
- As mentioned, we assume you will be opening as "Working Directory" specific example folders like `examples/simple_pyscript`.
- `code -r .` shortcut allows quickly reopening VSCode with changing working directory.
In normal projects reopening workdir with `code -r .` is not required, as we open projects only once, but we will be often using it in this article to switch between multiple working directories when necessary.
- We also assume you are using Linux as a main OS, or at least using WSL2 with linux opened inside if you are developing at Windows or MacOS.
    - (the article material is tested to work on Kubuntu 22.04 LTS)

First, ensure you installed `Pytest Explorer` extension. It will autoinstall Python extension for intellisense along the way.
- pytest explorer id for search `@id:littlefoxteam.vscode-python-test-adapter`
- it will autoinstall Python extension along side: `@id:ms-python.python`
![]({{.StaticRoot}}article_visual_debugger_in_vscode/extension_pytest1.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/extension_pytest2.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/extension_pytest3.png)

If smth is working weirdly, like syntax highlighting by colors is not present, make sure to run Ctrl + Shift + P and write Reload Window. Select and apply, it quickly reloads VSCode, including reloading extensions.
![]({{.StaticRoot}}article_visual_debugger_in_vscode/ctrl_shift_p_reloadwindow.png)

After that python simple file debugging is launchable as
git clone/download https://github.com/darklab8/blog/tree/master/
```shell
git clone https://github.com/darklab8/blog.git
cd blog/blog/articles/article_detailed/article_visual_debugger_in_vscode/examples
cd simple_pyscript && code -r # will reopen to different workdir
```

In the left sidebar, click on `file.py` in the directory tree.
Select in the top menu: Run -> Star Debugging (F5)

![]({{.StaticRoot}}article_visual_debugger_in_vscode/debug_python_script.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/debug_python_script2.png)

Now let's check it is working for pytest test debug:
```shell
# create venv 
python3 -m venv .venv
# activate
source .venv/bin/activate # at windows can be `venv\Scripts\activate`
pip install -r requirements.txt
```

Select a new venv in VSCode using one of these methods:

1. Click the interpreter in the bottom-right corner and select `.venv/bin/python3` manually if needed.
2. On Linux: `Ctrl + Shift + P` > `Enter interpreter path`
3. On Mac: `CMD + P` > `Python: Select Interpreter`

If smth glitches, we do `Ctrl + Shift + P -> Reload Window` trick

![]({{.StaticRoot}}article_visual_debugger_in_vscode/simple_pytest1.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/simple_pytest2.png)

# Basic functions of a visual debug in VSCode

- Click to the left of the code to set a red breakpoint.
- F10 step over
- F11 step in
- F5 continue until next breakpoint or end of the program
- Hover to see data; you can expand nested data.

# Python with settings for env vars

```shell
cd ../simple_py_with_settings && code -r .
```

You can update `.vscode/settings.json` to set up your debug parameters.

```shell
cat .vscode/settings.json
```
will show content:
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

```shell
$ python3 file.py 
Example(foo=10, bar='abc', is_smth=False)
```

Those env vars will be available for access from within the launched web server and unit tests during visual debug usage too.

# Python with running debug for django

```shell
cd ../django_example && code -r .
```

set a breakpoint in a view at `urls.py` and launch Python Django debug, like in this picture below
![]({{.StaticRoot}}article_visual_debugger_in_vscode/red_breakpoint.png)

Debug for Django will work due to the present `.vscode/launch.json` containing
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
            "args": ["runserver", "0.0.0.0:8000"],
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


Of all the settings, only the `Python: Django` section matters for running the debugger via `manage.py` right now.
Take note that you can technically input arbitrary arguments (under `args` key) to run it for execution of more different stuff 
![]({{.StaticRoot}}article_visual_debugger_in_vscode/django1.png)

take note of setting `"justMyCode": false`, it allows you navigating during visual debug third party libs too. Otherwise, they will be skipped.
Crucial thing when debugging internal company libs, or when developing your own library that is doing smth with known third-party libs.

![]({{.StaticRoot}}article_visual_debugger_in_vscode/django2.png)

The `justMyCode: false` setting in `launch.json` allows navigation through third-party library code, even during visual debugging from unit tests.
![]({{.StaticRoot}}article_visual_debugger_in_vscode/django3.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/django4.png)
![]({{.StaticRoot}}article_visual_debugger_in_vscode/django5.png)

# Python with running debug for fastapi

```shell
cd ../fastapi_example && code -r .
```

Install packages in virtualenv again:
```shell
deactivate
python3 -m venv .venv
source .venv/bin/activate # at windows can be `venv\Scripts\activate`
pip install -r requirements.txt
```

with small fixes to `.vscode/launch.json`, we have it adapted for fastapi now
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

With written small code examples in `app_test.py` we have no trouble launching visual debug for unit tests of fastapi

![]({{.StaticRoot}}article_visual_debugger_in_vscode/fastapi2.png)

See folder `flask_example` for almost same example to try for Flask too.

# Python in already existing Docker

Ensure having installed Dev Containers app `@id:ms-vscode-remote.remote-containers`
![]({{.StaticRoot}}article_visual_debugger_in_vscode/docker_extension_install1.png)

```shell
cd ../examples/django_example
docker compose build
docker compose run --service-ports shell
```

Instead of docker-compose usage, same is achievable with regular `docker` command:
```shell
docker build --tag test 
docker run -it -v $(pwd):/code -w /code --name shell --entrypoint=bash test
```

Now we can enter the already running container by using **Attach Visual Studio Code** option.
This option opens a second VSCode instance, running from within the container.
![]({{.StaticRoot}}article_visual_debugger_in_vscode/docker_extension_install2.png)

Install Pytest Explorer `@id:littlefoxteam.VSCode-python-test-adapter` (with autoinstalled Python inside).

Since the extension is installed inside a running container, you’ll need to repeat the installation each time a new container is created or started.

In  the left sidebar go to the working directory `/code`

If everything is set up correctly, you’ll see green arrows next to test functions, allowing you to run them in debug mode. (Green arrow at the left of a line code)

You can also launch the web server from the Debug menu by selecting `Python: Django`.

![]({{.StaticRoot}}article_visual_debugger_in_vscode/docker3.png)

To ensure that we are able to see dev server from the container, ensure we binded it to `0.0.0.0` instead of `localhost`!
That is doable with `python3 manage.py runserver 0.0.0.0:8000` for django, we added missing argument into `.vscode/launch.json`.

![]({{.StaticRoot}}article_visual_debugger_in_vscode/docker4.png)
after setting breakpoint in view, and visiting `http://localhost:8000/polls` we get being stopped in breakpoint as desired.

The same steps are doable for FastAPI and Flask.
We added Dockerfile and docker-compose.yml for the same into fastapi_example and flask_example folders.

P.S. When you are in VSCode from within container, opened terminal is opened automatically for within the container shell

# Ending

Congratulations, now you are able to use visual debug for python common web frameworks without and with docker for web server run itself and unit tests!

With showing examples how it is done for two different python frameworks, where for fastapi we adjusted launch.json on a fly,
it is my intention that you will be able to see how to do it similarly for any other existing python framework.

I encourage you to try development from within unit tests written with visual debug, since it is actually a comfortable way of having testing as part of a working code writing. Doing that will ensure most rapid feedback, and ensure you a writing better code in all kind of code internal characteristics.
Do check books like [TDD by kent beck]({{.SiteRoot}}favourite.html#TestDrivenDevelopmentByExample) and [Unit testing by Vladimir Khorikov]({{.SiteRoot}}favourite.html#UnitTestingPrinciplesPracticesandPatterns) in order to find out more about unit testing that enables writing maintainable code.

We showed examples with Docker, because modern web development could be having different C depended libraries, that without docker is very struggling to install otherwise. Also, docker makes universal way to setup dev envs for many people and documents as a code how your application is buildable. Some companies just keep dev env straight in docker only for those reasons.