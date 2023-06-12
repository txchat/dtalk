ARG version
## the lightweight scratch image we'll
FROM grafana/grafana:${version}
ADD provisioning/datasource.yml /etc/grafana/provisioning/datasources/
