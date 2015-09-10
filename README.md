cloudeye
========
A couple of  Distributed Cloud Eyes on your cloud to help you know everything about your website/service 

## Key concept

## How to use
until now the cloudeye only support the log in json format, like these below:
```javascript
{"t":1441077637,"action":"getSession","cost": 213, "db_cost": 23,"redis_cost": 12,"ret": 0, "idc":"yg", "request":{"seqid":"661330e611a6459cba5f5637280202ec","sid":"61443WVUGANICCEMAA4A"},"response":{"code":0}}
{"t":1441077638,"action":"getSession","cost": 113, "db_cost": 20,"redis_cost": 8,"ret": 1, "idc":"yg", "request":{"seqid":"661330e611a6459cba5f5637280202ec","sid":"61443WVUGANICCEMAA4A"},"response":{"code":1}}
```

suppose your logs are json format and it all output to the files 'ssn-sys.log' (log4j log4g log4x and syslog scribe flume etc  will be easy to do this ), and you have n metric are very important (like cost, db_cost, redis_cost and so on)
, let's face a more complex situation, you wana see these metric under specific action and under 
specific idc, and under specific action and specific idc at the same time,you can try these config below:

```javascript
    "metrics": [
            {
                "metricname": "session.cost",
                "value" : "cost",
                "type" : "counting",
                "time" : "t",
                "tags" : ["action", "idc"]
            },
            {
                "metricname": "session.dbcost",
                "value" : "db_cost",
                "type" : "counting",
                "time" : "t",
                "tags" : ["action", "idc"]
            },
            {
                "metricname": "session.rediscost",
                "value" : "redis_cost",
                "type" : "counting",
                "time" : "t",
                "tags" : ["action", "idc"]
            }
        ],
```
and these start a influxdb, and fullfill the backend_influxdb config:

```javascript
    "backend_influxdb" : {
        "host" : "localhost",
        "port" : 8086,
        "database" : "test02",
        "user" : "root",
        "pwd" : "root"
    }
```

the whole config file will be :

```javascript
{
    "filepath" : "./ssn-sys.log",
    "metrics": [
            {
                "metricname": "session.cost",
                "value" : "cost",
                "type" : "counting",
                "time" : "t",
                "tags" : ["action", "idc"]
            },
            {
                "metricname": "session.dbcost",
                "value" : "db_cost",
                "type" : "counting",
                "time" : "t",
                "tags" : ["action", "idc"]
            },
            {
                "metricname": "session.rediscost",
                "value" : "redis_cost",
                "type" : "counting",
                "time" : "t",
                "tags" : ["action", "idc"]
            }
        ],
    "backend_influxdb" : {
        "host" : "localhost",
        "port" : 8086,
        "database" : "test02",
        "user" : "root",
        "pwd" : "root"
    }
}
```
## How to install
go get github.com/0xmalloc/cloudeye


