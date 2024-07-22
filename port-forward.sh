#!/bin/sh

# forwards port 80 from gateway api to
c="kubectl get svc -n envoy-gateway-system --selector=gateway.envoyproxy.io/owning-gateway-namespace=default,gateway.envoyproxy.io/owning-gateway-name=eg -o jsonpath='{.items[0].metadata.name}'"
service=$(eval "$c")

kubectl -n envoy-gateway-system port-forward service/"$service" 8080:80