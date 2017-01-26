# Deploying our application to the cluster

Now we have made the necessary changes to our Ingress, Deployment and Service definitions we are ready to deploy our application.

Using `kubectl` you need to apply the definitions to the cluster.

Hint: You will need to use the `kubectl apply` command. For more on this click [here](https://kubernetes.io/docs/user-guide/kubectl/kubectl_apply/)

## Scaling our application

One of the advantages to using Kubernetes is it allows us to scale up the number of instances of our application easily.

So lets scale up our application to have four instances running.

To make this possible you will need to make adjustments to the `deployment.yaml` file and then execute `kubectl apply`.

## Upgrading our application