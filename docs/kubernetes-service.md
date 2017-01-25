# Changes required to the Kubernetes service definition

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