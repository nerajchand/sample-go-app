# Sample Go Web App

This is a simple Go Web App that can be used for Training & Demo purposes.

There are multiple branches available with different colours, to help showcase the different image releases.

## Building & Push a New Container Image Releases to DockerHub

To create a new release the following steps will need to be performed:

```sh
docker buildx build --platform linux/amd64 . --tag narji/sample-go-app:0.1.0 --push
```

## Build & Tag a Local Image

```sh
docker buildx build --platform linux/amd64 . --tag sample-go-app
```

## Run the Container Locally

### Run From a Local Image

Once the image has been built locally the container can be run with docker as follows:

```sh
docker run -p 8080:8080 sample-go-app
```

### Run From a Pulled Imaged

Instead of building a local version you can also pull a remote image with the specified tag:

```sh
    docker image pull narji/sample-go-app:0.1.0
```

A quick look will show which images are locally available.

```sh
    docker image ls
```

These are the images you can use locally. So for the remote image we've downloaded to our
local computer this may look like this:

```sh
    docker container run -p 8080:8080 narji/sample-go-app:0.1.0
```
