FROM alpine:3.20 as base

RUN apk update && \
	apk add --no-cache \
	ghostscript \
	make \
	perl \
	wget

RUN apk add fontconfig --no-cache --repository https://www.freedesktop.org/wiki/Software/fontconfig --allow-untrusted

RUN wget -qO- "https://yihui.name/gh/tinytex/tools/install-unx.sh" | sh

ENV PATH="${PATH}:/root/bin"

# RUN tlmgr install xetex
RUN fmtutil-sys --all

RUN apk add texlive

RUN tlmgr install hyphenat tracklang xltxtra realscripts tcolorbox wrapfig roboto lipsum


FROM golang:1.23 as build-stage

WORKDIR /

COPY . ./
COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go texTemplate.go

FROM base

COPY --from=build-stage /main /app/main
COPY ./images /app/images

CMD ["/app/main"]