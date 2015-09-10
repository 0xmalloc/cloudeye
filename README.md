cloudeye
========
A couple of  Distributed Cloud Eyes on your cloud to help you know everything about your website/service 

## Dependency
cloudeye use influxdb (version > =0.9) as the mertic storeage, and use grafana to display the metric

## Why do This
we need a simple tools to monitor and analysis our logs, logstash + statsd +influxdb + grafana seems to be a perfect solution, but when comes to compex metric under any other tags (like see the latency under sepcific idc, specific action ,sepecific caller, sepecfic host, and all at the same time), it sucks. influxdb (version > 0.9) support tags and is easy to solve this problem but logstash and statsd can not do it,so we make the tool.

## How to Use
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

## The Result
what will you get by using this simple tools? look below:
![mahua](https://raw.githubusercontent.com/0xmalloc/cloudeye/master/doc/cloudeye-pic.png)

## How to Install
go get github.com/0xmalloc/cloudeye
