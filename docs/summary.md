# Summary

* [Get started](start.md)

## Clients

* [Install](clients/install.md)
* [CLI commands](clients/cli.md)
* [Python API](clients/python.md)
* [Environments](clients/environments.md)
* [Uninstall](clients/uninstall.md)

## Workloads

* Realtime APIs
  * [Example](workloads/realtime/example.md)
  * [Predictor](workloads/realtime/predictors.md)
  * [Configuration](workloads/realtime/configuration.md)
  * [Models](workloads/realtime/models.md)
  * [Parallelism](workloads/realtime/parallelism.md)
  * [Server-side batching](workloads/realtime/server-side-batching.md)
  * [Autoscaling](workloads/realtime/autoscaling.md)
  * [Statuses](workloads/realtime/statuses.md)
  * [Metrics](workloads/realtime/metrics.md)
  * Multi-model
    * [Example](workloads/realtime/multi-model/example.md)
    * [Configuration](workloads/realtime/multi-model/configuration.md)
    * [Caching](workloads/realtime/multi-model/caching.md)
  * Traffic Splitter
    * [Example](workloads/realtime/traffic-splitter/example.md)
    * [Configuration](workloads/realtime/traffic-splitter/configuration.md)
  * [Troubleshooting](workloads/realtime/troubleshooting.md)
* [Async APIs](workloads/async/introduction.md)
  * [Example](workloads/async/example.md)
  * [Predictor](workloads/async/predictors.md)
  * [Configuration](workloads/async/configuration.md)
  * [Statuses](workloads/async/statuses.md)
  * [Webhooks](workloads/async/webhooks.md)
  * [Metrics](workloads/async/metrics.md)
* Batch APIs
  * [Example](workloads/batch/example.md)
  * [Predictor](workloads/batch/predictors.md)
  * [Configuration](workloads/batch/configuration.md)
  * [Jobs](workloads/batch/jobs.md)
  * [Statuses](workloads/batch/statuses.md)
  * [Metrics](workloads/batch/metrics.md)
* Task APIs
  * [Example](workloads/task/example.md)
  * [Definition](workloads/task/definitions.md)
  * [Configuration](workloads/task/configuration.md)
  * [Jobs](workloads/task/jobs.md)
  * [Statuses](workloads/task/statuses.md)
  * [Metrics](workloads/task/metrics.md)
* Dependencies
  * [Example](workloads/dependencies/example.md)
  * [Python packages](workloads/dependencies/python-packages.md)
  * [System packages](workloads/dependencies/system-packages.md)
  * [Custom images](workloads/dependencies/images.md)
* Observability
  * [Logging](workloads/observability/logging.md)
  * [Metrics](workloads/observability/metrics.md)

## Clusters

* AWS
  * [Install](clusters/aws/install.md)
  * [Update](clusters/aws/update.md)
  * [Auth](clusters/aws/auth.md)
  * [Security](clusters/aws/security.md)
  * [Multi-instance type](clusters/aws/multi-instance-type.md)
  * [Spot instances](clusters/aws/spot.md)
  * [Networking](clusters/aws/networking/index.md)
    * [Custom domain](clusters/aws/networking/custom-domain.md)
    * [HTTPS (via API Gateway)](clusters/aws/networking/https.md)
    * [VPC peering](clusters/aws/networking/vpc-peering.md)
  * [Setting up kubectl](clusters/aws/kubectl.md)
  * [Uninstall](clusters/aws/uninstall.md)
* GCP
  * [Install](clusters/gcp/install.md)
  * [Credentials](clusters/gcp/credentials.md)
  * [Multi-instance type](clusters/gcp/multi-instance-type.md)
  * [Setting up kubectl](clusters/gcp/kubectl.md)
  * [Uninstall](clusters/gcp/uninstall.md)
* [Private Docker registry](clusters/registry.md)
