# gojsonschema Docker Image

This repo is "dockerizing" the `https://github.com/xeipuuv/gojsonschema`.

# Usage

## Docker Hub

```sh
docker run --rm -v "$(pwd):/app" -w /app mingalevme/gojsonschema -s test/testdata/schema.json test/testdata/document-valid.json
```

## CI/CD (gitlab)

```sh
image: alpine
stages:
  - quality
gojsonschema:
  stage: quality
  needs:  []
  image:
    name: mingalevme/gojsonschema
    entrypoint: [""]
  script:
    - gojsonschema -s schema.json products.json
```

## Source

### Build from source

```sh
docker build --target gojsonschema -t gojsonschema .
```

### Run local image

```sh
docker run --rm -v "$(pwd):/app" -w /opt gojsonschema -s test/testdata/schema.json test/testdata/document-valid.json
```

# Dockerfile
```sh
https://github.com/mingalevme/gojsonschema/blob/master/Dockerfile
```
