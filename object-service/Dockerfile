FROM golang as build
WORKDIR /source
COPY ./object-service /source
RUN go build -v -o /app/objectservice .

FROM golang
EXPOSE 8001
WORKDIR /app
COPY --from=build /app /app
COPY --from=build /source/templates /app/templates
COPY --from=build /source/static /app/static
CMD ["/app/objectservice"]
