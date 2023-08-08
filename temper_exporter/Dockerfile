FROM quay.io/prometheus/busybox:latest

COPY temper_exporter /bin/temper_exporter

EXPOSE 9897
USER nobody

ENTRYPOINT [ "/bin/temper_exporter" ]