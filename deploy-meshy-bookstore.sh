#!/bin/bash

set -aueo pipefail

NAMESPACES=("bookbuyer" "bookstore" "bookthief" "bookwarehouse" "bookstore-tcp-client" "bookstore-tcp-echo-server")

for ns in "${NAMESPACES[@]}"; do
    kubectl create namespace $ns || true
    echo "Create the ${ns} service account and deployment:"
    kubectl apply -f ./manifests/${ns}.yaml
done


echo "Checkpoint"
for ns in "${NAMESPACES[@]}"; do
    kubectl get pods,deployments,serviceaccounts -n ${ns}
done
