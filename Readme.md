# Media streaming with gRPC and gRPC-web

This is a pet project using gRPC and gRPC-web to stream media (audio and video) over HTTP/2 using gRPC.

The project uses Go (golang) to implement the video streaming service and as an added feature, an authentication server, implemented in Julia (julialang) to authenticate requests to the video server.

The client uses gRPC-web through an envoy-proxy to perform auth operations on the auth server as well as stream media from the video server.

## Requirements
- Docker
- docker-compose
- A browser
- python3

## Installation
To run this application, just pull the source from this repo using the following command
```bash
git pull https://github.com/2HgO/grpc-streaming
```

To start up the services, run the following command from the root directory of the project
```bash
docker-compose up --build
```

*<b>Note</b>: It takes a while to instantiate the auth server because the julia package registry has to be downloaded and installed in the built image*

In a separate terminal window, run the following command to start up a static page where you can stream the video
```bash
make static-gateway
```