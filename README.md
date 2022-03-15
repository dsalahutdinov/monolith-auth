# monolith-auth
Example of how to enable monolith authenticated traffic into micro-services

## Local setup

Run monolith locally:
```sh
# builds docker images and runs server
$ docker-compose -f monolith/docker-compose.yml up --build server

# calls public products endpoint
$ curl http://localhost:9292/products
[{"id":"1","name":"Potato"},{"id":"2","name":"Tomato"},{"id":"3","name":"Onion"},{"id":"4","name":"Carrot"}]

# calls favorites endpoint with api token:
$ curl http://localhost:9292/favorites -H "Authorization: Token=qwerty"
[{"product_id":"1"},{"product_id":"2"}]
```

Run favorites service locally:
```sh
$ docker-compose -f favorites/docker-compose.yml up --build server

$ curl http://localhost:8383/favorites -H "X-Auth-Identity: 234"
[{"product_id":"1"},{"product_id":"3"}]
```

## Minikube setup
```
$ brew install minikube
$ minikube start --cpus 2 --memory 2048

$ eval $(minikube docker-env)
$ minikube -p minikube docker-env | source

$ docker build --tag monolith:0.1.0 ./monolith --no-cache
$ docker build --tag favorites:0.1.0 ./favorites --no-cache

#curl -L https://istio.io/downloadIstio | sh -
curl -L https://github.com/istio/istio/releases/download/1.12.0/istio-1.12.0-osx.tar.gz > istio-1.12.0-osx.tar.gz
gunzip istio-1.12.0-osx.tar.gz
tar xopf istio-1.12.0-osx.tar
export PATH=$PATH:$(pwd)/istio-1.12.0/bin

istioctl install
kubectl label namespace default istio-injection=enabled

$ kubectl apply -f k8s/monolith.yaml -f k8s/gateway.yaml

$ minikube tunnel
$ kubectl port-forward service/istio-ingressgateway -n istio-system 8080:80
$ curl http://localhost:8080/products
[{"id":"1","name":"Potato"},{"id":"2","name":"Tomato"},{"id":"3","name":"Onion"},{"id":"4","name":"Carrot"}]
```

## Canary
```shell
$ curl http://127.0.0.1/products/ #Version by weights

$ curl http://127.0.0.1/products/ -H "Canary: rails_next" #Force Canary Version
[{"id":"1","name":"Potato"},{"id":"2","name":"Tomato"},{"id":"3","name":"Onion"},{"id":"4","name":"Carrot"},{"id":"5","name":"Canary"}]
```

## Ingress Gateway Logs

```shell
kubectl logs $(kubectl get pods -n istio-system -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}' | grep istio-ingressgateway | head -1) -n istio-system -f
```


## WASM

### Setup (create ConfigMap)

```shell
cd k8s
tinygo build -o go_filter.wasm -scheduler=none -target=wasi ./golang-wasm-filter-envoy.go
kubectl create configmap wasm-binary --from-file=go_filter.wasm
kubectl label namespace default istio-injection=enabled --overwrite=true
```

### Logs

```shell
kubectl logs $(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}' | grep favorites | head -1) -c istio-proxy -f
```

### Apply WASM changes to EnvoyFilter

```shell
cd k8s
tinygo build -o go_filter.wasm -scheduler=none -target=wasi ./golang-wasm-filter-envoy.go
kubectl delete configmap wasm-binary && kubectl create configmap wasm-binary --from-file=go_filter.wasm
kubectl delete pods $(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}' | grep favorites)
```
