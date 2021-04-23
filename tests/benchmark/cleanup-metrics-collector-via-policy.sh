#!/bin/bash

# Managed cluser name
mc=$1

start=$2
end=$3

for i in $(seq $start $end); do
  file=policy-${mc}-${i}-delete.yaml
  echo "creating $file"
  cp policy-scale26-disable.yaml ${file}
  sed -i "s~__CLUSTER__~${mc}~" ${file}
  sed -i "s~__NUMBER__~${i}~" ${file}
  #sed -i "s~__CLUSTERID__~$(cat /proc/sys/kernel/random/uuid)~" ${file}
  echo "applying delete policy $file"
  oc apply -f ${file}
  # TODO: delete $file
done

