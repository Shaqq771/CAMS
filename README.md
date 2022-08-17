# go-backend-boilerplate

Go backend api standard template

## Installation
``` 
# cd {your-go-path}/src 
# git clone https://gitito.nabatisnack.co.id/standardization/go-app-template.git
# cd go-app-template
```

## Branching Git Flow
``` 
Run development:
# git checkout develop

Run production:
# git checkout master
```

## Run with local machine
``` 
Step by step:
+ Setup local environment: 
    - Go version 1.8
    - Database: Postgres
    - Caching: Redis
    - Message Broker: Rabbitmq
- Rename .env.example to .env in root directory 

Run:
# go run cmd/main.go
```

## API Documentations (Postman)

```  
https://www.getpostman.com/collections/3c641355a056c17e7ed4
```

## Folder structure
``` 
- cmd 
- config 
+ delivery 
    - container 
    + http 
        - middleware
- docs
+ domain
    + health : 
        - constant
        - feature
        - model
        - repository
    + logistik
        - constant
        - feature
        - helper
        - model
        - repository
    + sales
        - constant
        - consumer
        - feature
        - model
        - repository
    + shared
        - constant
        - context
        - error
        - helper
        - query
        - response
+ infrastructure
    + broker
        - rabbitmq
    - database
    - jwt
    + logger
        - logrus
        - zap
    + shared
        - constant
- migration
```

## Thank you. 