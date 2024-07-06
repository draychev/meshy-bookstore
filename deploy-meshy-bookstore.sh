#!/bin/bash

set -aueo pipefail

# Source: https://release-v1-2.docs.openservicemesh.io/docs/getting_started/install_apps/

for ns in bookbuyer bookstore bookthief bookwarehouse bookwatcher tcp-client tcp-echo-server tcp-demo; do
    kubectl create namespace $ns || true
    echo "Create the ${ns} service account and deployment:"
    kubectl apply -f https://raw.githubusercontent.com/draychev/meshy-bookstore/main/manifests/${ns}.yaml
done    


echo "Checkpoint"
for ns in bookbuyer bookstore bookthief bookwarehouse bookwatcher tcp-client tcp-echo-server tcp-demo; do
    kubectl get pods,deployments,serviceaccounts -n ${ns}
done
