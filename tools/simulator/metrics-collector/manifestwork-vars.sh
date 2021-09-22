#!/bin/bash
# envsubst variable for generating a ManifestWork to deploy on the HUB
# sample values extracted from cchun's hub

# Managed cluster name
managed_cluster_prefix="k3s"
#export CLUSTER=k3s-deca4f1a # name of the managed cluster, must match the namespace used on the hub

# Hub Thanos URL to write to
# oc get deployment metrics-collector-deployment -n open-cluster-management-addon-observability -o jsonpath='{.spec.template.spec.containers[0].env[1].value}'
thanos_url=$(oc get deployment metrics-collector-deployment -n open-cluster-management-addon-observability -o jsonpath='{.spec.template.spec.containers[0].env[1].value}')

# Certificates 
managed_cluster_ca_cert=$(oc get secret observability-managed-cluster-certs -o jsonpath='{.data.ca\.crt}' -n open-cluster-management-addon-observability)
tls_cert=$(oc get secret observability-controller-open-cluster-management.io-observability-signer-client-cert -o jsonpath='{.data.tls\.crt}' -n open-cluster-management-addon-observability)
tls_key=$(oc get secret observability-controller-open-cluster-management.io-observability-signer-client-cert -o jsonpath='{.data.tls\.key}' -n open-cluster-management-addon-observability)
# TODO: Other possible values to customize
# 1. Namespace to deploy on managed cluster. Defaults to open-cluster-management-addon-observability
# 2. metrics collector image to use

cp manifestwork-template.yaml manifestwork.yaml

sed -e "s,THANOS_URL,$thanos_url," \
    -e "s/MANAGED_CLUSTER_CA_CERT/$managed_cluster_ca_cert/" \
    -e "s/TLS_CERT/$tls_cert/" \
    -e "s/TLS_KEY/$tls_key/" manifestwork-template.yaml > manifestwork.yaml

START=${1:-$START}
NUM_MC=${2:-$NUM_MC}
END=$((START + NUM_MC - 1))
echo "START: $START
END: $END
TOTAL Managed Clusters: $NUM_MC"

hub=${3:-$hub}

folder=$hub/manifestworks
mkdir -p $folder
# generates all manifestworks
for i in $(seq $START $END) ; do
    cluster_name="$managed_cluster_prefix-$i"
    sed -e "s/CLUSTER_NAME/$cluster_name/" manifestwork.yaml > $folder/$cluster_name.yaml
done

# apply all manifestworks
oc apply -f $folder
