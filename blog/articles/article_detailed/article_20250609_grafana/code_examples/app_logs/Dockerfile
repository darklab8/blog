FROM golang:1.24-bullseye AS build

RUN apt update
RUN apt install -y build-essential
RUN apt-get install ca-certificates -y
RUN gcc --version

WORKDIR /code

COPY go.mod main.go ./
RUN --mount=type=cache,target="/root/.cache/go-build" go build -v -o main main.go

FROM debian:11.11-slim AS runner
WORKDIR /code
COPY --from=build /code/main main
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT /code/main
