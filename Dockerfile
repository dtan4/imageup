FROM alpine:3.6

COPY bin/imageup /imageup

EXPOSE 8000

CMD ["/imageup"]
