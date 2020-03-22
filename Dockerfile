FROM golang:1.13 as builder

WORKDIR /root
ADD . .

RUN go env -w GOPROXY=https://goproxy.io,direct

RUN make

FROM golang:1.13
WORKDIR /root
COPY --from=builder /root/bin/fastgo /root/fastgo
RUN ln -s /root/fastgo /bin/fastgo

CMD ["fastgo", "server"]
