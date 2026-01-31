#!/bin/bash

kubectl port-forward dev-url-short-cluster-user-db-0 5001:5432 &

kubectl port-forward dev-url-short-cluster-url-db-0 5002:5432 &

wait



