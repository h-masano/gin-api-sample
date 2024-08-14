# gin-api-sample

This is a sample project of gin API Server with api-codegen running on Docker container.

## How to run

``` sh
docker-compose build
docker-compose up -d
```

## How to modify API definition and generate gin code

1. modify `openapi.yaml`
1. go into the container

    ``` sh
    docker exec -it gin-api-sample-api-1 sh
    ```

1. generate gin code

    ``` sh
    make generate
    ```
