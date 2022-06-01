## the lightweight scratch image we'll
## run our application within
FROM alpine:latest
## We have to copy the output from our
## builder stage to our production stage
ARG server_name
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn
ENV SERVERNAME=${server_name}
WORKDIR /usr/local/bin
COPY ${server_name} .
COPY ${server_name}.* /etc/txchat-${server_name}/config/
## we can then kick off our newly compiled
## binary exectuable!!
RUN apk update && apk add tzdata
CMD ["sh", "-c", "./${SERVERNAME} -conf /etc/txchat-${SERVERNAME}/config/${SERVERNAME}.toml"]
