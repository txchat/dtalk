ARG version
## the lightweight scratch image we'll
FROM prom/prometheus:${version}
ADD prometheus.yml /etc/prometheus/
ADD service.json /etc/prometheus/
