# Setup the loki-distributed as a logging system.
1. Add the helm repo.
```sh
$ helm repo add grafana https://grafana.github.io/helm-charts
$ helm repo update
```
2. Then, install it with custom values.
```sh
$ helm -n loki-stack install loki-stack grafana/loki-distributed --create-namespace --values values-update.yaml
```
