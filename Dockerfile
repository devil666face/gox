FROM ubuntu:20.04
RUN DEBIAN_FRONTEND=noninteractive \
    apt-get update --quiet --quiet && \
    apt-get install --quiet --quiet --yes \
    --no-install-recommends --no-install-suggests \
    openssl ca-certificates \
    && apt-get --quiet --quiet clean \
    && rm --recursive --force /var/lib/apt/lists/* /tmp/* /var/tmp/*
WORKDIR /gox
COPY ./bin/gox .
CMD ["./gox","-config","./config/config.yaml"]
