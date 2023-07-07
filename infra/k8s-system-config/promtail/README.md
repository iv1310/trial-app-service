# Setup the promtail.
1. Add the helm repo(but no needed if we already added grafana repo).
```sh
$ helm repo add grafana https://grafana.github.io/helm-charts
$ helm repo update
```
2. Then, install it with custom values.
```sh
$ helm -n loki-promtail install loki-promtail grafana/promtail --create-namespace --values values-update.yaml
```
