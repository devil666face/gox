FROM ubuntu:20.04
RUN DEBIAN_FRONTEND=noninteractive \
    apt-get update && \
    apt-get install --yes \
    --no-install-recommends --no-install-suggests \
    openssl ca-certificates curl \
    && apt-get clean \
    && rm --recursive --force /var/lib/apt/lists/* /tmp/* /var/tmp/*
WORKDIR /gox
COPY ./bin/gox .
CMD ["./gox","-config","./config/config.yaml"]
