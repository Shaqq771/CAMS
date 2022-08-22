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
    - Message Broker: Rabbitmq
- Rename .env.example to .env in root directory 
- Setup value in .env (fill in according to your settings)

Run:
# go run main.go
```

## Run with Docker
``` 
Step by step:
- install docker

Run:
# docker create network nabati
# docker-compose build --no-cache
# docker-compose up -d

Terminate:
# docker-compose down
```

## API Documentations (Postman)

```  
https://www.getpostman.com/collections/3c641355a056c17e7ed4
```

## Folder Tree Structure
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
main.go
```

## Example Log (ZAP)

```  
{"level":"info","msg":"SELECT * FROM product where id = $1 AND deleted_at IS NULL LIMIT 1","app_name":"Backend-Nabati","app_version":"0.1.0","log_type":"query","time":"2022-08-21T23:36:50+07:00","log":"zap"}
{"level":"info","msg":"product id not found0","app_name":"Backend-Nabati","app_version":"0.1.0","data":"error","log_type":"general error","time":"2022-08-21T23:36:50+07:00","log":"zap"}
{"level":"info","msg":"product id not found: 0","app_name":"Backend-Nabati","app_version":"0.1.0","data":"response","log_type":"general error","time":"2022-08-21T23:36:50+07:00","log":"zap"}
```

## Example Log (LOGRUS)

```  
{"app_name":"Backend-Nabati","app_version":"0.1.0","level":"info","log":"logrus","log_type":"query","msg":"SELECT * FROM product where id = $1 AND deleted_at IS NULL LIMIT 1","time":"2022-08-21T23:38:11+07:00"}
{"app_name":"Backend-Nabati","app_version":"0.1.0","error_type":"general error","level":"error","log":"logrus","log_type":"error","msg":"product id not found0","time":"2022-08-21T23:38:11+07:00"}
{"app_name":"Backend-Nabati","app_version":"0.1.0","error_type":"general error","level":"error","log":"logrus","log_type":"response","msg":"product id not found: 0","time":"2022-08-21T23:38:11+07:00"}
```

## Thank you. 