# Metrics Collector Simulator

Metrics collector simulator can be used to setup multiple metrics collector in different namespaces in one managed cluster, to simulate thousands of managed clusters push metrics to ACM hub cluster for scale testing.

_Note:_ this simulator is for testing purpose only.

## Prereqs
You must meet the following requirements to setup metrics collector:

- ACM 2.1+ available
- `MultiClusterObservability` instance available and have following pods in `open-cluster-management-addon-observability` namespace:

	```
	$ oc get po -n open-cluster-management-addon-observability
	NAME                                               READY   STATUS    RESTARTS   AGE
	endpoint-observability-operator-7f8f949bc8-trwzh   2/2     Running   0          118m
	metrics-collector-deployment-74cbf5896f-jhg6v      1/1     Running   0          111m
	```

## Quick Start
### Setup metrics collector
You can run `manifestwork-vars.sh` to setup multiple metrics collector in each managed cluster.
Syntax:
```
# ./manifestwork-vars.sh <managed cluster starting index> <number of managed clusters> <hub name>
```

Each managed cluster's name begins with the prefix `k3s` and is followed by the index.

For example, setup 10 metrics collectors for managed clusters `k3s-3` to `k3s-5` with the following command:
```
# ./manifestwork-vars.sh 3 2 scale-hub
```

Check if the metrics collector running successfully in your managed cluster:
```
# kubectl get po -n open-cluster-management-addon-observability | grep metrics-collector-deployment
metrics-collector-deployment-7d69d9f897-xn8vz                    1/1     Running            0          22h
```

Check if the metrics collector manifestwork has successfully deployed to your managed cluster:
```
# oc get manifestwork -n <managed cluster namespace>

```

## Generate your own metrics data source
By default, `setup-metrics-collector.sh` is using metrics data defined in env `METRICS_IMAGE` as data source. You can build and push your own metrics data image with below command:
```
# METRICS_IMAGE=<example/metrics-data:latest> make all
```
## Setup metrics collector with your own metrics data source
Running below command to setup metrics collectors with your own data source:
```
# METRICS_IMAGE=<example/metrics-data:latest> ./setup-metrics-collector.sh 10
```
