#!/bin/bash

set -aueo pipefail

# Source: https://release-v1-2.docs.openservicemesh.io/docs/getting_started/install_apps/
kubectl create namespace bookstore || true
kubectl create namespace bookbuyer || true
kubectl create namespace bookthief || true
kubectl create namespace bookwarehouse || true


echo "Create the bookbuyer service account and deployment:"
kubectl apply -f https://raw.githubusercontent.com/draychev/meshy-bookstore/main/manifests/bookbuyer.yaml

echo "Create the bookthief service account and deployment:"
kubectl apply -f https://raw.githubusercontent.com/draychev/meshy-bookstore/main/manifests/bookthief.yaml

echo "Create the bookstore service account, service, and deployment:"
kubectl apply -f https://raw.githubusercontent.com/draychev/meshy-bookstore/main/manifests/bookstore.yaml

echo "Create the bookwarehouse service account, service, and deployment:"
kubectl apply -f https://raw.githubusercontent.com/draychev/meshy-bookstore/main/manifests/bookwarehouse.yaml

echo "Create the mysql service account, service, and stateful set:"
kubectl apply -f https://raw.githubusercontent.com/draychev/meshy-bookstore/main/manifests/mysql.yaml


echo "Checkpoint"
kubectl get pods,deployments,serviceaccounts -n bookbuyer
kubectl get pods,deployments,serviceaccounts -n bookthief

kubectl get pods,deployments,serviceaccounts,services,endpoints -n bookstore
kubectl get pods,deployments,serviceaccounts,services,endpoints -n bookwarehouse
