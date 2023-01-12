FROM golang:latest
RUN mkdir -p /go/src/abassGarane/workouts
WORKDIR /go/src/abassGarane/workouts
ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build  -o bin/muscles -ldflags "-w -s" -a -installsuffix cgo cmd/web/*

FROM scratch
COPY --from=0 /go/src/abassGarane/workouts/bin/muscles .
CMD ["./muscles"]
EXPOSE 4000:80

