# YData custom k8s resource

## Install

`go get github.com/guille-mas/kubernetes-custom-resource`

## Build

`go build -o ydata .`

## Howto run it

```bash
./ydata -kubeconfig=$HOME/.kube/config
kubectl create -f yaml/crd.yml
kubectl create -f yaml/my-ydata.yml
kubectl get ydatas
```

## Backlog

Fix vanity import path guille.cloud/kubernetes-custom-resource. \
I think the root of the problem is that the meta tag is injected into the DOM at runtime at https://guille.cloud/kubernetes-custom-resource, so go can't see it when it checks that uri.

I've put whatever come to my mind as YData spec and status attributes. More meaningful attributes would be fine to have.
