# Description

Ci tool to create pipelines purely run locally
Intended for cloud CI agnostic solution and being executed from other graphical CI tools

# Version1: language agnostic version in yaml

Create CI tool like Tekton, but with Shell / Docker executors (optionally k8s executor later)
With syntax like in Gitlab CI and Docker CI
With inserting values like in Helm / Docker-compose
Must be made in Golang

# Version2: it should work at least in go. Optionally adding support for other languages later.

- https://docs.dagger.io/cookbook like dagger, but free, executable locally / easily embedding into Github Actions/Gitlab / self hosted optionally with its own runners
- Those docs and libs may help on this path
    - https://blog.gitea.com/creating-go-actions/ gitea added support for native golang actions
    - https://github.com/sethvargo/go-githubactions - Github Actions
