![Mercado Libre](https://http2.mlstatic.com/frontend-assets/ui-navigation/5.15.0/mercadolibre/logo__large_plus@2x.png)

``` sh
$ git clone git@github.com:wuilkysb/ml-test-quasar.git $GOPATH/ml-test && cd $_
```


#  Overview
API to decode an imperial message and position

# Requirements

* go v1.15
* go module

# Build

* Install dependencies: 
```sh
$ go mod download
```

* Run test:
```sh 
$ go test ./... 
```

# Environments
#### Required environment variables

* `SERVER_PORT`: port for the server

# Endpoints

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
    ]' \
  --compressed
```