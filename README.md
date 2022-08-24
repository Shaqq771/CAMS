# go-backend-boilerplate

Go backend api standard template

# Installation
## Clone Package from Git
``` 
# cd {your-go-path}/src 
# git clone https://gitito.nabatisnack.co.id/standardization/go-app-template.git
# cd go-app-template
```

## Branching Git Flow

Run development:
``` 
# git checkout develop
``` 
Run production:
``` 
# git checkout master
```

# Run Package 
## Run with local machine

Step by step:
+ Setup local environment: 
    - Go version 1.8
    - Database: Postgres
    - Message Broker: Rabbitmq
- Setup value in .env.example (fill in according to your settings)
- Run:
``` 
# make run_local
```

## Run with Docker

Step by step:
- install docker
- Run:
``` 
# make run_docker
``` 
- Terminate/stop docker container:
``` 
# make stop_docker
```

# Documentations
## API Documentations (Postman)

```  
Check in docs folder
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

# Logs
## Example Log (ZAP)
```  
{"level":"info","msg":"SELECT * FROM product where id = $1 AND deleted_at IS NULL LIMIT 1","app_name":"Backend-Nabati","app_version":"0.1.0","log_type":"query","time":"2022-08-21T23:36:50+07:00","log":"zap"}
{"level":"info","msg":"product id not found0","app_name":"Backend-Nabati","app_version":"0.1.0","data":"error","log_type":"general error","time":"2022-08-21T23:36:50+07:00","log":"zap"}
{"level":"info","msg":"product id not found: 0","app_name":"Backend-Nabati","app_version":"0.1.0","data":"response","log_type":"general error","time":"2022-08-21T23:36:50+07:00","log":"zap"}
```

## Example Log (LOGRUS) (DEFAULT)
Default using Logrus: Because logrus more fleksible to customize the logs structure and easy to read.
```  
{"app_name":"Backend-Nabati","app_version":"0.1.0","level":"info","log":"logrus","log_type":"query","msg":"SELECT * FROM product where id = $1 AND deleted_at IS NULL LIMIT 1","time":"2022-08-21T23:38:11+07:00"}
{"app_name":"Backend-Nabati","app_version":"0.1.0","error_type":"general error","level":"error","log":"logrus","log_type":"error","msg":"product id not found0","time":"2022-08-21T23:38:11+07:00"}
{"app_name":"Backend-Nabati","app_version":"0.1.0","error_type":"general error","level":"error","log":"logrus","log_type":"response","msg":"product id not found: 0","time":"2022-08-21T23:38:11+07:00"}
```

# Unit Test
## Create Mock Interface For Unit Test

Step by step:
- Install mockgen
```
go install github.com/golang/mock/mockgen@v1.6.0
```
- Open Makefile
- Add this code section mock:

```
mockgen -source="YOUR_GO_INTERFACE" -destination="YOUR_MOCK_DESTINATION"
```

Example:
```
mockgen -source="./domain/logistik/feature/feature.go" -destination="./domain/logistik/feature/mocks/feature_mock.go"
```
- Execute Makefile: 
```
# make mock 
```
- After complete, please check your mock in the destination folder/file.


## Running Unit Test 
- Open Makefile
- Add this code section test:
```
# go test -cover ./YOUR_GO_FOLDER
```
Example:
```
# go test -cover ./domain/logistik/feature
```
- Run Unit Test
```
make test
```
Results:
```
ok  backend-nabati/domain/logistik/feature  0.257s  coverage: 77.9% of statements
```

## Thank you. 