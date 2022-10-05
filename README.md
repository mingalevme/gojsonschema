# gojsonschema Docker Image

This repo is "dockerizing" the `https://github.com/xeipuuv/gojsonschema`.

# Usage

## Docker Hub
```sh
docker run --rm -v "$(pwd):/opt" -w /opt mingalevme/gojsonschema -s test/testdata/schema.json test/testdata/document-valid.json
```

## Source

### Build
```sh
docker build --target gojsonschema -t gojsonschema .
```

### Run
```sh
docker run --rm -v "$(pwd):/opt" -w /opt gojsonschema -s test/testdata/schema.json test/testdata/document-valid.json
```

Dockerfile
```sh
https://github.com/mingalevme/gojsonschema/blob/master/Dockerfile
```
