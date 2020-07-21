# Kubernetes CRD + operator

Welcome! In this repo I'm hosting a very simple sample of a K8s Custom Resource Definition's operator written in Golang.

## Prepare your development environment

1. Clone https://github.com/kubernetes/code-generator.git
2. Step into code-generator/ root folder
3. Set GOPATH environment variable to the root of your Custom resource definition
4. run hack/update-codegen.sh

```bash
git clone https://github.com/kubernetes/code-generator.git
cd code-generator
export GOPATH=/opt/dev/kubernetes-custom-resource
hack/update-codegen.sh
```

```bash
vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/guille-mas/kubernetes-custom-resource/kubernetes-custom-resource/pkg/client \ github.com/guille-mas/kubernetes-custom-resource/kubernetes-custom-resource/pkg/apis \
github.com/guille-mas/kubernetes-custom-resource:v1
```
