#!/bin/bash

for ns in bookbuyer bookstore bookthief bookwarehouse tcp-client tcp-echo-server; do
    kubectl delete namespace $ns --wait=true
    ## kubectl patch namespace $ns -p '{"spec":{"finalizers":[]}}' --type=merge
done
