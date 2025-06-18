FROM golang:tip-alpine3.22 AS build

WORKDIR /app

COPY . .

RUN go build -o /bin/app ./cmd/main.go


FROM build

COPY --from=build ./app .

CMD [ "/bin/app" ]