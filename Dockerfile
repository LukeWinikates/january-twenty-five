FROM gcr.io/distroless/static-debian12:latest
COPY --chmod=555 build/server-linux-amd64 .
USER       nobody
EXPOSE     6724
ENTRYPOINT [ "/server-linux-amd64" ]