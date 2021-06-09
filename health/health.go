package health

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"main/helper"
	"main/types"
	"strings"
)

func CheckHealth(client *redis.Client) {
	keys := client.Keys("*").Val()
	count := len(keys)

	for i := 0; i < count; i++ {
		subKeys := client.HKeys(keys[i]).Val()
		CheckSubKeyHealth(client, keys[i], subKeys)
		CheckSecondarySubKeyHealth(client, keys[i], subKeys)
	}
}

func CheckSubKeyHealth(client *redis.Client, key string, subKeys []string) {
	var recordType string
	//var subKey string
	var val string
	var previousHealthStatus bool
	var currentHealthStatus bool

	for i := 0; i < len(subKeys); i++ {
		val = client.HGet(key, subKeys[i]).Val()
		//subKey = subKeys[i]
		recordType = helper.GetRecordType(val)
		if strings.Contains(val, "SIMPLE"){
			continue
		}
		if recordType != "a"{
			continue
		}

		helper.Log("Stored Values From DB", val)

		var record types.Record
		err := json.Unmarshal([]byte(val), &record)
		if err != nil {
			fmt.Println("error during unmarshalling:", err)
		}

		if recordType == "a" {
			previousHealthStatus = record.A.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.A.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.A.Value.Secondary.IsHealthy = currentHealthStatus
		}else if recordType == "aaaa" {
			previousHealthStatus = record.AAAA.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.AAAA.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.AAAA.Value.Secondary.IsHealthy = currentHealthStatus
		}else if recordType == "txt" {
			previousHealthStatus = record.TXT.Value.Primary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.TXT.Value.Primary.HealthCheckConfig, previousHealthStatus)
			record.TXT.Value.Primary.IsHealthy = currentHealthStatus
		}else if recordType == "cname" {
			previousHealthStatus = record.CNAME.Value.Primary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.CNAME.Value.Primary.HealthCheckConfig, previousHealthStatus)
			record.CNAME.Value.Primary.IsHealthy = currentHealthStatus
		}else if recordType == "mx" {
			previousHealthStatus = record.MX.Value.Primary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.MX.Value.Primary.HealthCheckConfig, previousHealthStatus)
			record.MX.Value.Primary.IsHealthy = currentHealthStatus
		}else if recordType == "srv" {
			previousHealthStatus = record.SRV.Value.Primary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.SRV.Value.Primary.HealthCheckConfig, previousHealthStatus)
			record.SRV.Value.Primary.IsHealthy = currentHealthStatus
		}else if recordType == "caa" {
			previousHealthStatus = record.CAA.Value.Primary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.CAA.Value.Primary.HealthCheckConfig, previousHealthStatus)
			record.CAA.Value.Primary.IsHealthy = currentHealthStatus
		}else if recordType == "soa" {
			previousHealthStatus = record.SOA.Value.Primary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.SOA.Value.Primary.HealthCheckConfig, previousHealthStatus)
			record.SOA.Value.Primary.IsHealthy = currentHealthStatus
		}


		valueBytes, _ := json.Marshal(record)
		modifiedValueBytes := string(valueBytes)

		if previousHealthStatus != currentHealthStatus {
			//UpdateHealthStatus(client, key, subKeys[i], modifiedValueBytes)
			helper.Log("Modified Value", modifiedValueBytes)
		}else{
			helper.SimpleLog("No Changes Found!")
		}

		fmt.Println("test ",record.A.Value.Secondary.HealthCheckConfig.Port == "")
	}
}


func CheckSecondarySubKeyHealth(client *redis.Client, key string, subKeys []string) {
	var recordType string
	//var subKey string
	var val string
	var previousHealthStatus bool
	var currentHealthStatus bool

	for i := 0; i < len(subKeys); i++ {
		val = client.HGet(key, subKeys[i]).Val()
		//subKey = subKeys[i]
		recordType = helper.GetRecordType(val)
		if strings.Contains(val, "SIMPLE"){
			continue
		}
		if recordType != "a"{
			continue
		}

		helper.Log("Stored Values From DB", val)

		var record types.Record
		err := json.Unmarshal([]byte(val), &record)
		if err != nil {
			fmt.Println("error during unmarshalling:", err)
		}

		if recordType == "a" {
			if record.A.Value.Secondary.HealthCheckConfig.Type == ""{
				continue
			}
			previousHealthStatus = record.A.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.A.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.A.Value.Secondary.IsHealthy = currentHealthStatus
		}else if recordType == "aaaa" {
			if record.AAAA.Value.Secondary.HealthCheckConfig.Type == ""{
				continue
			}
			previousHealthStatus = record.AAAA.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.AAAA.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.AAAA.Value.Secondary.IsHealthy = currentHealthStatus
		}else if recordType == "txt" {
			if record.TXT.Value.Secondary.HealthCheckConfig.Type == ""{
				continue
			}
			previousHealthStatus = record.TXT.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.TXT.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.TXT.Value.Secondary.IsHealthy = currentHealthStatus
		}else if recordType == "cname" {
			if record.CNAME.Value.Secondary.HealthCheckConfig.Type == ""{
				continue
			}
			previousHealthStatus = record.CNAME.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.CNAME.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.CNAME.Value.Secondary.IsHealthy = currentHealthStatus
		}else if recordType == "mx" {
			if record.MX.Value.Secondary.HealthCheckConfig.Type == ""{
				continue
			}
			previousHealthStatus = record.MX.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.MX.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.MX.Value.Secondary.IsHealthy = currentHealthStatus
		}else if recordType == "srv" {
			if record.SRV.Value.Secondary.HealthCheckConfig.Type == ""{
				continue
			}
			previousHealthStatus = record.SRV.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.SRV.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.SRV.Value.Secondary.IsHealthy = currentHealthStatus
		}else if recordType == "caa" {
			if record.CAA.Value.Secondary.HealthCheckConfig.Type == ""{
				continue
			}
			previousHealthStatus = record.CAA.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.CAA.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.CAA.Value.Secondary.IsHealthy = currentHealthStatus
		}else if recordType == "soa" {
			if record.SOA.Value.Secondary.HealthCheckConfig.Type == ""{
				continue
			}
			previousHealthStatus = record.SOA.Value.Secondary.IsHealthy
			currentHealthStatus = GetCurrentHealthStatus(record.SOA.Value.Secondary.HealthCheckConfig, previousHealthStatus)
			record.SOA.Value.Secondary.IsHealthy = currentHealthStatus
		}


		valueBytes, _ := json.Marshal(record)
		modifiedValueBytes := string(valueBytes)

		if previousHealthStatus != currentHealthStatus {
			UpdateHealthStatus(client, key, subKeys[i], modifiedValueBytes)
			helper.Log("Modified Value", modifiedValueBytes)
		}else{
			helper.SimpleLog("No Changes Found!")
		}
	}
}



func GetCurrentHealthStatus(healthCheckConfig types.FailOverHealthCheckConfig, previousHealthStatus bool) bool {
	var firstAttempt bool
	var secondAttempt bool
	var thirdAttempt bool

	if healthCheckConfig.Type == "TCP" {
		firstAttempt = helper.TcpConnect(healthCheckConfig.TargetIPs, healthCheckConfig.Port)
		if firstAttempt == previousHealthStatus {
			return previousHealthStatus
		}
		secondAttempt = helper.TcpConnect(healthCheckConfig.TargetIPs, healthCheckConfig.Port)
		thirdAttempt = helper.TcpConnect(healthCheckConfig.TargetIPs, healthCheckConfig.Port)
	} else {
		firstAttempt = helper.Curl(healthCheckConfig.TargetUrl)
		if firstAttempt == previousHealthStatus {
			return previousHealthStatus
		}
		secondAttempt = helper.Curl(healthCheckConfig.TargetUrl)
		thirdAttempt = helper.Curl(healthCheckConfig.TargetUrl)
	}

	fmt.Println("first attempt: ", firstAttempt)
	fmt.Println("second attempt: ", secondAttempt)
	fmt.Println("third attempt: ", thirdAttempt)

	if firstAttempt == secondAttempt && secondAttempt == thirdAttempt && thirdAttempt == firstAttempt {
		if firstAttempt != previousHealthStatus {
			return firstAttempt
		}
	}

	return previousHealthStatus
}

func UpdateHealthStatus(client *redis.Client, key string, subKey string, val string) {
	client.HSet(key, subKey, val)
}
