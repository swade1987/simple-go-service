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

**The Build task should have two commands, one for Golang and another for Docker**

To verify the build has been successful execute the following commands:

```
$ make build
$ docker images
```

You should see your image listed after executing the second command above.

## 3. Proving the container is fit for purpose

Now we have created a docker image we need to validate it is fit for purpose.

To validate this we will use the `docker-compose.yml` in the root directory.

To read more on Docker Compose click [here](https://docs.docker.com/compose/overview/)

**Some key parts of the `docker-compose.yml` are missing and need to be added.**

To verify the container is fit for purpose execute the following command from the root directory:

```
$ docker-compose up -d
```

Now in your browser browse to [http://localhost:5555](http://localhost:5555) you should see `Hello World`.

## 4. Uploading the container image to Docker Hub

As we have now verified the container is fit for purpose we should upload it to Docker Hub.

Firstly, we need to create a repository on Docker Hub itself for our images.

Browse to [https://hub.docker.com/add/repository/](https://hub.docker.com/add/repository/)

I would suggest `simple-go-app` as the repository name.

Once you have created the repository we need to execute the `docker login` command on our local machines to login to Docker Hub.

For help with the `docker login` command click [here](https://docs.docker.com/engine/reference/commandline/login/)

## Manual uploading

In order to push a newly created image to Docker Hub will need to perform the following steps:

1. Build a new image with the correct tag using the repo name and a specific version `<username>/simple-go-app:1.0`

2. Push the newly created image upto Docker Hub.

3. Removing the image from your local machine.

If you have successfully pushed the new image to Docker Hub so you should be able to browse to it from your browser.

## Automated uploading

We can go one step further and automate the process of creating and uploading new images.

To make this possible we need to update the `Makefile`

1. Adjust the `IMAGE_NAME` variable.

2. Update the `push` task to replicate the stages you executed during the manual execution.

Note: For this to work you will need to set the `IMAGE_VERSION` environment variable before each execution.

Once you think you have a working task update the `IMAGE_VERSION` to be `1.1` and then execute `make push`

If you are successful you should see v1.1 of the container in your Docker Hub repository.