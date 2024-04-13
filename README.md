# memcached-operator

My adaptation of the Go Operator Tutorial from the official Operator SDK website

## Prerequisites

- Access to a [Kubernetes](https://kubernetes.io/) cluster
- [Go](https://go.dev/) 1.21
- [Operator SDK](https://sdk.operatorframework.io/) 1.34

If you do not already have a Kubernetes cluster, get started with Kubernetes in no time with [kind](https://kind.sigs.k8s.io/).

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

To clean up the resources:

```bash
kubectl delete -f config/samples/cache_v1alpha1_memcached.yaml
make undeploy
```

See also: [Go Operator Tutorial | Operator SDK](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/)

## License

[Apache 2.0](./LICENSE)
