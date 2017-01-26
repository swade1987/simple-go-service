# Deploying our application to the cluster

Now we have made the necessary changes to our Ingress, Deployment and Service definitions we are ready to deploy our application.

Using `kubectl` you need to apply the definitions to the cluster.

Hint: You will need to use the `kubectl apply` command. For more on this click [here](https://kubernetes.io/docs/user-guide/kubectl/kubectl_apply/)

## Scaling our application

One of the advantages to using Kubernetes is it allows us to scale up the number of instances of our application easily.

So lets scale up our application to have four instances running.

To make this possible you will need to make adjustments to the `deployment.yaml` file.

Then execute something like the command below:

```
$ kubectl apply ....
```

## Verfication

So we have deployed and scaled our application, now we need to verify its been successful.

For this we need to locate the hostname you chose to us in your Ingress definition.

Simple browse to http://<hostname> and you should see `Hello World!`.