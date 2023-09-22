FROM golang as build
WORKDIR /source
COPY ./aladin-service /source
RUN go build -v -o /app/aladinservice .

FROM golang
WORKDIR /app
COPY --from=build /app /app
COPY --from=build /source/templates /app/templates
COPY --from=build /source/static /app/static
CMD ["/app/aladinservice"]
