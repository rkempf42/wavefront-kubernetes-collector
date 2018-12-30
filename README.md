# wavefront-kubernetes-collector

This collector enables monitoring Kubernetes clusters and sending metrics to [Wavefront](https://www.wavefront.com).

The collector scrapes the Kubelet summary API for Kubernetes metrics (based on [heapster](https://github.com/wavefronthq/wavefront-kubernetes-collector/tree/master/docs/heapster.md)). It additionally supports scraping Prometheus metrics format endpoints.

## Prerequisites
- Kubernetes 1.9+

## Configuration

The collector is plugin-driven and supports collecting metrics from multiple sources and writing metrics to Wavefront using a [Wavefront proxy](https://docs.wavefront.com/proxies.html) or via [direct ingestion](https://docs.wavefront.com/direct_ingestion.html).

See [configuration doc](https://github.com/wavefronthq/wavefront-kubernetes-collector/tree/master/docs/configuration.md) for detailed configuration information.

### Sources

Following sources are currently supported and can be configured using the `--source` flag:

1. Kubernetes source to collect performance metrics from the kubelet `/stats/summary` metrics API:
```
--source=kubernetes.summary_api:''
```
2. Prometheus source to scrape metrics from Prometheus metrics format endpoints such as kube state metrics:
```
--source=prometheus:''?url=http://kube-state-metrics.kube-system.svc.cluster.local:8080/metrics
```
Multiple prometheus sources can be added to scrape additional endpoints.

### Pod Auto Discovery
The collector can auto discover pods that export Prometheus format metrics. See the [discovery documentation](https://github.com/wavefronthq/wavefront-kubernetes-collector/tree/master/docs/discovery.md) for details.

### Sending metrics to Wavefront

#### Using Wavefront Proxy

```
--sink=wavefront:?proxyAddress=wavefront-proxy.default.svc.cluster.local:2878&clusterName=k8s-cluster&includeLabels=true
```

#### Using Direct Ingestion
```
--sink=wavefront:?server=https://<YOUR_INSTANCE>.wavefront.com&token=<YOUR_TOKEN>&clusterName=k8s-cluster&includeLabels=true
```

## Installation

1. Clone this repo.
2. Edit the `wavefront` sink in `deploy/kubernetes/4-collector-deployment.yaml`.
3. Edit or remove the `prometheus` sink in the above file.
4. Run `kubectl apply -f deploy/kubernetes`

To verify the installation, find the pod name of the deployed `wavefront-collector` and run:

```
kubectl logs -f COLLECTOR_POD_NAME -n wavefront-collector
```
