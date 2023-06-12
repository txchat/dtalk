ARG version
## the lightweight scratch image we'll
FROM mysql:${version}
COPY dtalk_biz.sql /docker-entrypoint-initdb.d/
COPY dtalk_record.sql /docker-entrypoint-initdb.d/
