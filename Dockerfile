FROM scratch
COPY simple-smtp-cli /bin/simple-smtp-cli
ENTRYPOINT ["/bin/simple-smtp-cli"]
