FROM scratch
COPY bin/photo-gallery-go /
CMD ["/photo-gallery-go"]
EXPOSE 8080
