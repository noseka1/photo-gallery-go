FROM registry.access.redhat.com/ubi8/ubi-minimal AS base
RUN microdnf install golang git

FROM base AS builder
ADD . /bin
WORKDIR /bin
RUN go build -o photo-gallery-go ./cmd/photo-gallery-go/main.go

FROM registry.access.redhat.com/ubi8/ubi-minimal
EXPOSE 8080
COPY --from=builder /bin/photo-gallery-go /
CMD ["/photo-gallery-go"]