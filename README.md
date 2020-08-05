# go-gin-mysql Example

PreRequisties 
 * Go( version 1.1 or more) should be installed.
 * Add Running MySQL server details(host,port,db name,etc) in config.yaml - No  need to create tables. 
 
   
Run Following commands

* go mod download
* go build
* go run main.go


The REST server serves Below ENDPOINTS, 


CREATE CUSTOMER     |    POST    |   http://localhost:8080/api/customer                                |    {"name":"david", "active":false}       
_________________________________________________________________________________________________________________________________________________________

ADD IPBLACKLIST     |    POST    |   http://localhost:8080/api/ipblacklist                             |    {"ip":"2130706433"}                       
_________________________________________________________________________________________________________________________________________________________

ADD USER AGENT      |    POST    |   http://localhost:8080/api/userblacklist                           |    {"user-agent":"Alien"}                    
_________________________________________________________________________________________________________________________________________________________

POST REQUEST        |    POST    |   http://localhost:8080/api/request                                 |   {"customerID":28,"tagID":2,"userID":"sri",
                                                                                                            "remoteIP":"123.234.56.78",                
                                                                                                           "timestamp":1596490607}                    
_________________________________________________________________________________________________________________________________________________________

UPDATE USER         |   PUT      |   http://localhost:8080/api/customer/6                              |    {"name":"sri", "active":false}
_________________________________________________________________________________________________________________________________________________________

DELETE IPBLACKLIST  |   DELETE   |  http://localhost:8080/api/ipblacklist?ip=2130706423                |    
__________________________________________________________________________________________________________________________________________________________

GET STATS        |    GET     |   http://localhost:8080/api/stats/?customerID=7&date=2020-08-04        |   
   
{
   "Customer": "7",
    "HourlyCounts": [
        {
            "RequestCount": 1,
            "InvalidCount": 2,
            "Time": "2020-08-04T19:00:00+02:00"
        },
        {
            "RequestCount": 1,
            "InvalidCount": 2,
            "Time": "2020-08-04T13:00:00+02:00"
        }
    ]
}
          
