#!/bin/bash

# Name of the managed cluster
mc=$1

# Set up unique metrics collectors in the range start-end
start=$2
end=$3

# TODO: check start <= end

# Generate a policy via a template file
for i in $(seq $start $end); do
  file=policy-${mc}-${i}.yaml
  echo "creating $file"
  cp policy-scale26.yaml ${file}
  sed -i "s~__CLUSTER__~${mc}~" ${file}
  sed -i "s~__NUMBER__~${i}~" ${file}
  sed -i "s~__CLUSTERID__~$(cat /proc/sys/kernel/random/uuid)~" ${file}
  echo "applying policy $file"
  oc apply -f ${file}
  # TODO: cleanup $file
done
