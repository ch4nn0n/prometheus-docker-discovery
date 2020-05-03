
# Prometheus Docker Discovery

Statically defining prometheus targets for docker containers is un-scalable and becomes a serious issue when operating n a dynamic environment. Defining targets through docker labels solves this problem, this application works by looking for docker labels and using the prometheus file-based service discovery.

The style of labeling takes inspiration from the Kubernetes service discovery mechanism and the Traefik docker configuration.

## Usage

### Prometheus Docker Discover
Configuration for this application is service
```yaml
prometheus.discovery.scrapeInterval: "15s"
prometheus.discovery.refreshInterval: "1m"

prometheus.discovery.logLevel: "info"
prometheus.discovery.label.<name>: "<value>"
```

### Docker labels
```yaml
prometheus.target.enabled: "true"
prometheus.target.job: "<name>" 
prometheus.target.port: "9090"
prometheus.target.path: "/metrics"

prometheus.target.label.<name>: "<value>"
```

### Prometheus configuration
Add the following snippet to the
```yaml
scrape_configs:
- job_name: 'node'
  file_sd_configs:
  - files:
    - 'docker-targets.json'
```

## Further reading
- https://prometheus.io/docs/guides/file-sd/
- https://prometheus.io/docs/prometheus/latest/configuration/configuration/#file_sd_config
