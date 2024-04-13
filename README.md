# memcached-operator

My adaptation of the Go Operator Tutorial from the official Operator SDK website

## Prerequisites

- Access to a [Kubernetes](https://kubernetes.io/) cluster with [cert-manager](https://cert-manager.io/) installed
- [Go](https://go.dev/) 1.21
- [Operator SDK](https://sdk.operatorframework.io/) 1.34

If you do not already have a Kubernetes cluster, get started with Kubernetes in no time with [kind](https://kind.sigs.k8s.io/).

An example of installing cert-manager on a fresh kind cluster using the official Helm chart:

```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm -n cert-manager install \
    cert-manager \
    jetstack/cert-manager \
    --version v1.14.4 \
    --set installCRDs=true \
    --create-namespace
```

## Developing

Fork and clone this repository, make it your working directory, then run:

```bash
make deploy
```

This runs the Memcached operator as a Kubernetes [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) under the [namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) `memcached-operator-system`:

```bash
kubectl -n memcached-operator-system get deploy
```

Sample output:

```text
NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE
memcached-operator-controller-manager   1/1     1            1           8m
```

Now deploy a sample Memcached [custom resource \(CR\)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) using the template available under `config/samples/cache_v1alpha1_memcached.yaml`:

```bash
kubectl apply -f config/samples/cache_v1alpha1_memcached.yaml
```

Observe that the Memcached CR is created along with its associated Deployment and Pods:

```bash
kubectl get memcached,deploy,pod
```

Sample output:

```text
NAME                                                  AGE
memcached.cache.donaldsebleung.com/memcached-sample   37s

NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/memcached-sample   3/3     3            3           22s

NAME                                   READY   STATUS    RESTARTS   AGE
pod/memcached-sample-df947c465-85rh4   1/1     Running   0          22s
pod/memcached-sample-df947c465-ck4ht   1/1     Running   0          22s
pod/memcached-sample-df947c465-jvzcf   1/1     Running   0          22s
```

Now edit the Memcached CR with `kubectl edit` and observe which modifications are allowed / rejected:

1. Increase `.spec.size` from `3` to `5`
1. Decrease `.spec.size` from `5` to `1`
1. Increase `.spec.size` from `1` to `6`
1. Decrease `.spec.size` from `6` to `0`
1. Decrease `.spec.size` from `1` to `-1`
1. Edit `.spec.containerPort` from `11211` to `0`
1. Edit `.spec.containerPort` from `11211` to `33221`

To clean up the resources:

```bash
kubectl delete -f config/samples/cache_v1alpha1_memcached.yaml
make undeploy
```

## See also

1. [Go Operator Tutorial | Operator SDK](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/)
1. [Admission Webhooks | Operator SDK](https://sdk.operatorframework.io/docs/building-operators/golang/webhook/)

## License

[Apache 2.0](./LICENSE)
