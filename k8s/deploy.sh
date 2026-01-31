#!/bin/bash

kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0/standard-install.yaml
helm repo add kong https://charts.konghq.com

helm repo update

helm install dev ./url-short-cluster/

helm install kong kong/ingress

