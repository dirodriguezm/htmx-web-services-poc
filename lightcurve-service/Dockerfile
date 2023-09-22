FROM golang as build
WORKDIR /source
COPY ./lightcurve-service /source
RUN go build -v -o /app/lightcurve .

FROM golang
WORKDIR /app
COPY --from=build /app /app
COPY --from=build /source/templates /app/templates
COPY --from=build /source/static /app/static
CMD ["/app/lightcurve"]
