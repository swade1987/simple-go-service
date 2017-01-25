# Changes required to the Kubernetes ingress definition

The following changes need to be made to this file.

- The ingress definition needs to be given a name (ideally the same as the deployment and service definitions):

```
metadata:
  name:
```

- The service name (the same as the deployment and service) and port (the container port) must be specified:

```
backend:
  serviceName:
  servicePort:
```

## Setting up the hostname

In order for the hostname to work successfully the IP address of the ingress node needs to be obtained.

Once we have obtained the IP address execute the following command:

```
$
```


Finally the `host` must be specified, this is the domain name which will be used to browse to our application

## Additional documentation

To learn more about how ingress works within Kubernetes the following link is recommended [https://kubernetes.io/docs/user-guide/ingress/](https://kubernetes.io/docs/user-guide/ingress/)