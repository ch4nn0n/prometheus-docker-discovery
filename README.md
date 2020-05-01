
# Prometheus Docker Discovery

Statically defining prometheus targets for docker containers is un-scalable and becomes a serious issue when operating n a dynamic environment. Defining targets through docker labels solves this problem, this application works by looking for docker labels and using the prometheus file-based service discovery.

The style of labeling takes inspiration from the Kubernetes service discovery mechanism and the Traefik docker configuration.

