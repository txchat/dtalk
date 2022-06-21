ARG version
## the lightweight scratch image we'll
FROM nginx:${version}
COPY nginx.conf /etc/nginx/nginx.conf
