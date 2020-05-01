
# Prometheus Docker Discovery

Statically defining prometheus targets for docker containers is un-scalable and becomes a serious issue when operating n a dynamic environment. Defining targets through docker labels solves this problem, this application works by looking for docker labels and using the prometheus file-based service discovery.

The style of labeling takes inspiration from the Kubernetes service discovery mechanism and the Traefik docker configuration.

## Usage

### Prometheus Docker Discover
Configuration for this application is service
```yaml
prometheus.docker.discovery.enabled: "true"
prometheus.docker.discovery.logLevel: "info"
prometheus.docker.discovery.network: "<name>"
prometheus.docker.discovery.label.<name>: "<value>"
prometheus.docker.discovery.targetsFile: "<path_to_file>"
```

### Docker labels
```yaml
prometheus.docker.<service>.enabled: "true"
prometheus.docker.<service>.target.port: "6060"
prometheus.docker.<service>.target.path: "/metrics"
prometheus.docker.<service>.target.interval: "15s"
prometheus.docker.<service>.label.<name>: "<value>"
prometheus.docker.<service>.network: "<name>"
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
