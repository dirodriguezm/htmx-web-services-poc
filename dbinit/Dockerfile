FROM golang as build
WORKDIR /source
COPY ../go.work /source/go.work
COPY ../dbinit /source/dbinit
COPY ../lightcurve-service /source/lightcurve-service
RUN go build -v -o /app/dbinit alerce/dbinit

FROM golang
WORKDIR /app
COPY --from=build /app /app
CMD ["/app/dbinit"]
