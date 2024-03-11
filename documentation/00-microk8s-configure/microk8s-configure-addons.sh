#!/bin/bash

microk8s disable ha-cluster --force
microk8s enable dns
microk8s enable community
microk8s enable hostpath-storage
microk8s kubectl get pods -A
microk8s enable multus
microk8s enable ingress
microk8s enable dashboard
microk8s status
microk8s kubectl get pods -A