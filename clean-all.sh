#!/bin/bash

NAMESPACES=("bookbuyer" "bookstore" "bookthief" "bookwarehouse" "tcp-client" "tcp-echo-server")

for ns in "${NAMESPACES[@]}"; do
    kubectl delete namespace $ns --wait=true
    ## kubectl patch namespace $ns -p '{"spec":{"finalizers":[]}}' --type=merge
done
