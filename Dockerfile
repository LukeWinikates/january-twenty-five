FROM gcr.io/distroless/static-debian12:latest
COPY --chmod=555 build/housesitter-z2m-linux-amd64 .
USER       nobody
EXPOSE     6724
ENTRYPOINT [ "/housesitter-z2m-linux-amd64" ]