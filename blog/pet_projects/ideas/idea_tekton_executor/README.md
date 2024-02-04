# Description:

While checking tekton, I discovered that it is really promising tool which by design will allow me to have

- CI code not dependended on any CI provider
- CI code easy to run locally

# Goals:
As first goal I wish to receive tekton executor, which would be integrated with Github Actions
as second goal I wish to have docker in docker availability to my pipelines

end result will be available docker image

# base deps inside docker:
- minikube
- kubectl
- helm
- docker in docker inside minikube
- tekton inside minikube

# note to remember

launch docker container as privileged
docker run image -it --privileged