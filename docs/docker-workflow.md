# Docker CI/CD tutorial

## 1. Dockerfile and .dockerignore

You will find located in the root directory two files `Dockerfile` and `.dockerignore`

These files will need to be **edited** in order to get the application to run inside a Docker container.

Tip: Remember we only want the application binary to be inside the container.

Tip: The `.dockerignore` file should include all files and directories you don't want inside of the container.

For more information please reference [https://docs.docker.com/engine/reference/builder/](https://docs.docker.com/engine/reference/builder/)

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

Now we have created a docker image we need to validate its fit for purpose.

To validate this we will use the `docker-compose.yml` in the root directory.

To read more on Docker Compose click [here](https://docs.docker.com/compose/overview/)

Tip: Some key parts of the `docker-compose.yml` are missing and need to be added.

To verify the container is fit for purpose execute the following command from the root directory:

```
$ docker-compose up -d
```

Now in your browser browse to [http://localhost:5555](http://localhost:5555) you should see `Hello World`.

## 4. Uploading the container image to Docker Hub