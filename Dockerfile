#
# https://github.com/xeipuuv/gojsonschema
# https://github.com/mingalevme/gojsonschema
# docker build --target gojsonschema -t mingalevme/gojsonschema .
#
FROM golang:alpine AS builder
WORKDIR /gojsonschema
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gojsonschema .

FROM alpine AS gojsonschema
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="mingalevme/gojsonschema" \
      org.label-schema.description="An implementation of JSON Schema for the Go programming language. Supports draft-04, draft-06 and draft-07." \
      org.label-schema.url="https://github.com/xeipuuv/gojsonschema" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/xeipuuv/gojsonschema" \
      org.label-schema.vendor="xeipuuv" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0" \
      org.label-schema.docker.cmd="docker run --rm -v \"$(pwd):/opt\" -w /opt gojsonschema -s test/testdata/schema.json test/testdata/document-valid.json"
COPY --from=builder /gojsonschema/gojsonschema /usr/local/bin/gojsonschema
ENTRYPOINT [ "/usr/local/bin/gojsonschema" ]
