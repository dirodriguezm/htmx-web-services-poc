FROM golang as build
WORKDIR /source
COPY . /source
RUN go build -v -o /app/lightcurve .

FROM golang
WORKDIR /app
COPY --from=build /app /app
COPY --from=build /source/templates /app/templates
CMD ["/app/lightcurve"]
