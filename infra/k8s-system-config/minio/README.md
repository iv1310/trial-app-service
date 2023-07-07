# Setup the minio object storage.
1. Add the helm repo.
```sh
$ helm repo add minio https://helm.min.io
$ helm repo update
```
2. Then, install it with custom values.
```sh
$ helm -n minio install minio --create-namespace --values values-update.yaml
```
