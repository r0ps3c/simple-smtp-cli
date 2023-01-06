FROM golang as build
COPY . /go/src
RUN \
	cd /go/src && \
	make

FROM scratch
COPY --from=build /go/src/simple-smtp-cli /bin/simple-smtp-cli
ENTRYPOINT ["/bin/simple-smtp-cli"]
