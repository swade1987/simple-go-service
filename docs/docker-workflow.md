# Docker CI/CD tutorial

## 1. Dockerfile and .dockerignore

You will find located in the root directory two files `Dockerfile` and `.dockerignore`

These files will need to be **edited** in order to get the application to run inside a Docker container.

Tip: Remember we only want the application binary to be inside the container.

Tip: The `.dockerignore` file should include all files and directories you don't want inside of the container.

For more information please reference [https://docs.docker.com/engine/reference/builder/#/from](https://docs.docker.com/engine/reference/builder/#/from)

The blue navigation menu on the right hand side of the page (for the link above) is very informative.

## 2. Building the Docker container

To build the Docker container you will need to edit the "build" task within the `Makefile` in the root directory

Tip: The [docker build](https://docs.docker.com/engine/reference/commandline/build/) link is a good place to start.

To verify the build has been successful execute the following commands:

```
$ make build
$ docker images
```

You should see your image listed after executing the second command above.

## 3. Proving the container is fit for purpose


## 4. Uploading the container image to Docker Hub