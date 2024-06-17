FROM ubuntu:noble
WORKDIR /root
COPY server /root
COPY config.yaml /root
CMD ["./server"]
