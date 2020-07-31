FROM alpine:3.12
RUN apk add --no-cache tzdata
ADD ./bin/monbus /monbus
CMD ["/monbus", "--logtostderr", "-v", "2"]
EXPOSE 80