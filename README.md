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
$ minikube start --cpus 2 --memory 2096

$ eval $(minikube docker-env)
$ docker build --tag monolith:0.1.0 ./monolith --no-cache
$ docker build --tag favorites:0.1.0 ./favorites --no-cache

curl -L https://github.com/istio/istio/releases/download/1.12.0/istio-1.12.0-osx.tar.gz > istio-1.12.0-osx.tar.gz
gunzip istio-1.12.0-osx.tar.gz
tar xopf istio-1.12.0-osx.tar
export PATH=$PATH:$(pwd)/istio-1.12.0/bin

istioctl install
kubectl label namespace default istio-injection=enabled

$ kubectl apply -f k8s

$ kubectl port-forward service/istio-ingressgateway -n istio-system 8080:80
$ curl http://localhost:8080/products
[{"id":"1","name":"Potato"},{"id":"2","name":"Tomato"},{"id":"3","name":"Onion"},{"id":"4","name":"Carrot"}]
```
