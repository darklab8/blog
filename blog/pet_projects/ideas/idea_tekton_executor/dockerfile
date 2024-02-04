FROM ubuntu:20.04

WORKDIR /install

# installing docker
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive TZ=Etc/UTC apt-get -y install tzdata
RUN apt-get -y install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
RUN mkdir -p /etc/apt/keyrings
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg
RUN echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
RUN apt-get update
RUN apt-get  -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin

# Get kubectl
RUN curl -LO https://dl.k8s.io/release/v1.24.0/bin/linux/amd64/kubectl
RUN install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# Minikube
RUN curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
RUN install minikube-linux-amd64 /usr/local/bin/minikube

#helm
RUN curl -LO https://get.helm.sh/helm-v3.9.0-linux-amd64.tar.gz
RUN tar -zxvf helm-v3.9.0-linux-amd64.tar.gz
RUN mv linux-amd64/helm /usr/local/bin/helm

# tekton CLI
RUN curl -LO https://github.com/tektoncd/cli/releases/download/v0.24.0/tkn_0.24.0_Linux_x86_64.tar.gz
RUN tar xvzf tkn_0.24.0_Linux_x86_64.tar.gz  -C /usr/local/bin/ tkn

WORKDIR /code

CMD service docker start && sleep 5 && minikube start --kubernetes-version=v1.24.0 --force && kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml && tail -F /dev/null