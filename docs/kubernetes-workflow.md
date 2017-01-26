# Continuous Delivery of our application to a Kubernetes cluster

## Prerequisites

In order for you to complete the Kubernetes CI/CD tutorial you will need to working Kubernetes cluster.

The cluster needs to have a dedicated Ingress node.

It is recommended that the [Kismatic Enterprise Toolkit](https://github.com/apprenda/kismatic) is used to setup the cluster.

An example of setting up an Ingress node with KET can be found [here](https://github.com/apprenda/kismatic/blob/master/docs/INGRESS.md)

Note: It is recommended that you create a bootstrap node and execute KET from that node.

## Workflow

Within the `kubernetes` directory you find three files: `deployment.yaml` , `ingress.yaml` and `service.yaml`.

These three files will need to be edited in order for you to deploy the application onto the Kubernetes cluster.

### 1. Service.yaml

The steps required to edit the Kubernetes service definition can be found [here](kubernetes-service.md)

### 2. Deployment.yaml

The steps required to edit the Kubernetes deployment definition can be found [here](kubernetes-deployment.md)

### 3. Ingress.yaml

The steps required to edit the Kubernetes ingress definition can be found [here](kubernetes-ingress.md)

### 4. Execution

The steps to deploy and scale our application can be found [here](kubernetes-execution.md)