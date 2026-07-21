FROM quay.io/prometheus/busybox:glibc

ARG TARGETPLATFORM

LABEL maintainer="Simon Schneider <dev@raynigon.com>"

COPY $TARGETPLATFORM/github_billing_exporter /bin/github_billing_exporter

EXPOSE      9776
USER        nobody
ENTRYPOINT  [ "/bin/github_billing_exporter" ]
