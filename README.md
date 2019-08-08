# WebServer example in Go

This example shows how to build and deploy a containerized Go web server
application using [Kubernetes](https://kubernetes.io).

This directory contains:

- `main.go` contains the HTTP server implementation. It answers to all HTTP
requests with `Request Headers` in response.
- `Dockerfile` is used to build the Docker image for the application.
- `manifests` directory contains sample kubernetes manifests to deploy an application.

This application is available as a Docker image:

- `docker.io/iafeoktistov/go-webserver:1.0`
