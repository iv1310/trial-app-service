# Setup the prometheus stack.
1. Add the helm repo.
```sh
$ helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
$ helm repo update
```
2. Then, install it with custom values.
```sh
$ helm -n kube-prom-stack install kube-prometheus-stack prometheus-community/kube-prometheus-stack --create-namespace --values values-update.yaml
```
