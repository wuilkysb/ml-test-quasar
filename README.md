![Mercado Libre](https://http2.mlstatic.com/frontend-assets/ui-navigation/5.15.0/mercadolibre/logo__large_plus@2x.png)

``` sh
$ git clone git@github.com:wuilkysb/ml-test-quasar.git $GOPATH/ml-test && cd $_
```

#  Operation Quasar Fire
##  Overview
API to decode an imperial message and position

![Screenshot](https://github.com/wuilkydb/ml-test-quasar/blob/master/readme_resource/packages.png?raw=true)

* router: entry points
* controllers: entry points between user requests and business logic
* services: business logic
* utils: useful features
* providers: dependency injector 

## Requirements

* go v1.15
* go module

## Build

* Install dependencies: 
```sh
$ go mod download
```

* OR
```sh
$ go get ./...
```

* Run test:
```sh 
$ go test ./... 
```

* Install swagger:
```sh 
$ go get -u github.com/swaggo/swag/cmd/swag
```

* Create swagger documentation
```sh 
$ swag init
```

## Environments
#### Required environment variables

* `PORT`: port for the server (example: 8080)

## Endpoints

### production endpoints

```
    curl --location --request POST 'http://ml-test-quasar.herokuapp.com/satellite/topsecret/' \
    --header 'accept: application/json' \
    --header 'Content-Type: application/json' \
    --data-raw '[
        {
            "distance": 100,
            "message": [
                "este",
                "",
                "",
                "mensaje",
                "",
                ""
            ],
            "name": "kenobi"
        },
        {
            "distance": 115.5,
            "message": [
                "",
                "es",
                "",
                "",
                "",
                "secreto"
            ],
            "name": "skywalker"
        },
        {
            "distance": 142.7,
            "message": [
                "este",
                "",
                "un",
                "",
                "ultra",
                ""
            ],
            "name": "sato"
        }
    ]'
```

### local endpoints
* for swagger documentation go to:
```
    http://localhost:8080/satellite/swagger/index.html
```

*  service get complete message and location:
``` 
  curl 'http://localhost:8080/satellite/topsecret/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  ----data-raw '[
    {
        "distance": 100,
        "message": [
            "este",
            "",
            "",
            "mensaje",
            "",
            ""
        ],
        "name": "kenobi"
    },
    {
        "distance": 115.5,
        "message": [
            "",
            "es",
            "",
            "",
            "",
            "secreto"
        ],
        "name": "skywalker"
    },
    {
        "distance": 142.7,
        "message": [
            "este",
            "",
            "un",
            "",
            "ultra",
            ""
        ],
        "name": "sato"
    }
    ]'
```