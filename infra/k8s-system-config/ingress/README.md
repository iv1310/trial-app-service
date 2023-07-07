# Setup the nginx-ingress controller.
1. Add the helm repo.
```sh
$ helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
$ helm repo update
```
2. Then, install the ingress with custom values.
```sh
$ helm -n ingress-nginx install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --values values-update.yaml
```
