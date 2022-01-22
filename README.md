# Dns Health Check Service
This check whether an ip address or domain name is live or dead continuously after a interval. It is a cron service. 

Suppose an ip address for health check of a server are stored in database like below:

```
{
  "a": {
    "type": "FAIL_OVER",
    "value": {
      "primary": {
        "data": [
          {
            "ttl": 300,
            "ip": "1.2.2.1"
          },
          {
            "ttl": 300,
            "ip": "1.3.3.3"
          }
        ],
        "isHealthy": true,
        "healthCheckConfig": {
          "type": "TCP",
          "targetIPs": [
            "95.216.15.82",
            "95.216.20.168"
          ],
          "port": "80"
        }
      },
      "secondary": {
        "data": [
          {
            "ttl": 300,
            "ip": "2.1.2.2"
          },
          {
            "ttl": 300,
            "ip": "2.3.2.5"
          }
        ],
        "isHealthy": true,
        "healthCheckConfig": {
          "type": "TCP",
          "targetIPs": [
            "95.216.15.82",
            "95.216.20.168"
          ],
          "port": "80"
        }
      }
    }
  }
}
```

Now it will hit the ip adddress from the primary health check config. Then it will observe whether it get response or not. If get response, isHealthy status of primary portion will be true. Otherwise, it will be overwrote to false.
