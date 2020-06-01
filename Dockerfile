FROM alpine:3.12
ADD ./bin/monbus /monbus
CMD ["/monbus", "--logtostderr", "-v", "2"]