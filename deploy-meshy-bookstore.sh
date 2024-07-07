#!/bin/bash

set -aueo pipefail

for ns in bookbuyer bookstore bookthief bookwarehouse tcp-client tcp-echo-server; do
    kubectl create namespace $ns || true
    echo "Create the ${ns} service account and deployment:"
    kubectl apply -f ./manifests/${ns}.yaml
done


echo "Checkpoint"
for ns in bookbuyer bookstore bookthief bookwarehouse tcp-client tcp-echo-server; do
    kubectl get pods,deployments,serviceaccounts -n ${ns}
done
