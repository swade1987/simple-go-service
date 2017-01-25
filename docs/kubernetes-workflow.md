# Kubernetes CI/CD tutorial

## Prerequisites

In order for you to complete the Kubernetes CI/CD tutorial you will need to working Kubernetes cluster.

The cluster needs to have a dedicated Ingress node.

It is recommended that the [Kismatic Enterprise Toolkit](https://github.com/apprenda/kismatic) is used to setup the cluster.

An example of setting up an Ingress node with KET can be found [here](https://github.com/apprenda/kismatic/blob/master/docs/INGRESS.md)

Note: It is recommended that you create a bootstrap node and execute KET from that node.

## Workflow

Within the `kubernetes` directory you find three files: `deployment.yaml` , `ingress.yaml` and `service.yaml`.

These three files will need to be edited in order for you to deploy and verify the application is running successfully on the cluster.

### 1. Service.yaml

The following changes need to be made to this file.

- The service needs to be given a name:

```
metadata:
  name:
```

- The selector for the app needs to be provided:

```
selector:
    app:
```

- The port required to run our application needs to be provided, this was set earlier in the Docker CI/CD section of this tutorial.

```
ports:
  - port:
```

To learn more about how services work within Kubernetes the following link is a good place to start [https://kubernetes.io/docs/user-guide/services/](https://kubernetes.io/docs/user-guide/services/)

### 2. Deployment.yaml


### 3. Ingress.yaml


### 4. Execution

