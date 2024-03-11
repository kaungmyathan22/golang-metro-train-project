# Golang trains station project

### Trains
#### Create Train
    curl -X POST \
         http://localhost:8000/v1/trains \
        -H 'cache-control: no-cache' \
        -H 'content-type: application/json' \
        -d '{"driverName": "Menaka", "operatingStatus": true}'
#### Get train by id
    CURL -X GET "http://localhost:8000/v1/trains/1"
#### Delete train
    CURL -X DELETE "http://localhost:8000/v1/trains/1"