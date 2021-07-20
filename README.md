# gojsonschema Docker Image

This repo is for "dockerizing" the `https://github.com/xeipuuv/gojsonschema`.

## Building
```sh
docker build --target gojsonschema -t gojsonschema .
```

## Running
```sh
docker run --rm -v "$(pwd):/opt" -w /opt gojsonschema -s test/testdata/schema.json test/testdata/document-valid.json
```
