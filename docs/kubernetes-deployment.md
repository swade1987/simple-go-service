# Changes required to the Kubernetes deployment definition

The `deployment.yaml` file contains the majority of the logic used to deploy our application onto the cluster.

The following changes need to be made to the `deployment.yaml` file within the `kubernetes` directory.

- The metadata for the deployment needs to be specified (this should reference the service definition name):

```
metadata:
  name:
  labels:
    name:
```

- The `app` name and container image `version` must be specified:

```
labels:
  app:
  version: ""
```

- The image name needs to be specified, this needs to be the `name` and the `version` of the image stored in Docker Hub.

```
containers:
  - image:  ""
```

- The `port` for the liveness and readiness probes need to be set. For more information on these probes see [here](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/)

- Finally the container port and name need to be specified:

```
- containerPort:
  name:
```

For a more detailed understanding of Kubernetes deployments the following link is recommended: [https://kubernetes.io/docs/user-guide/deployments/](https://kubernetes.io/docs/user-guide/deployments/)