FROM golang:1.10-alpine

RUN apk add --no-cache git
RUN mkdir -p src/github.com/MartinToruan/tax_calculator

RUN go get -u github.com/golang/dep/cmd/dep

COPY Gopkg.toml src/github.com/MartinToruan/tax_calculator/
COPY main.go src/github.com/MartinToruan/tax_calculator/
COPY helper src/github.com/MartinToruan/tax_calculator/helper
COPY logic src/github.com/MartinToruan/tax_calculator/logic
COPY model src/github.com/MartinToruan/tax_calculator/model
COPY persistence src/github.com/MartinToruan/tax_calculator/persistence
COPY svc src/github.com/MartinToruan/tax_calculator/svc

RUN cd src/github.com/MartinToruan/tax_calculator && dep ensure -v

WORKDIR src/github.com/MartinToruan/

CMD ["sh", "-c", "go run tax_calculator/main.go"]