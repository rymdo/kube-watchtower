FROM alpine:3.16
ENTRYPOINT ["/kube-watchtower"]
COPY kube-watchtower /