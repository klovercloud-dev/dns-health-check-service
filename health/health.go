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
	}
}

func CheckSubKeyHealth(client *redis.Client, key string, subKeys []string) {
	var recordType string
	var record types.Record
	var subKey string
	var val string

	for i := 0; i < len(subKeys); i++ {
		val = client.HGet(key, subKeys[i]).Val()
		subKey = subKeys[i]
		recordType = helper.GetRecordType(val)
		if strings.Contains(val, "SIMPLE"){
			continue
		}

		if recordType == "a" && subKey != "www" && subKey != "@" && subKey != "verify"{
			helper.Log("Stored Values From DB", val)

			err := json.Unmarshal([]byte(val), &record)
			if err != nil {
				fmt.Println("error:", err)
			}

			previousHealthStatus := record.A.Value.Primary.IsHealthy
			currentHealthStatus := GetCurrentHealthStatus(record.A.Value.Primary.HealthCheckConfig, previousHealthStatus)

			record.A.Value.Primary.IsHealthy = currentHealthStatus
			valueBytes, _ := json.Marshal(record)

			//modifiedValueBytes := "{\"a\":" + string(valueBytes) + "}"
			modifiedValueBytes := string(valueBytes)
			if previousHealthStatus != currentHealthStatus {
				UpdateHealthStatus(client, key, subKeys[i], modifiedValueBytes)
			}

			helper.Log("Modified Value", modifiedValueBytes)
		}
	//	if recordType == "cname" && subKey != "www" && subKey != "@" && subKey != "verify" {
	//		helper.Log("Stored Values From DB", val)
	//		var record types.Record
	//		err := json.Unmarshal([]byte(val), &record)
	//		if err != nil {
	//			fmt.Println("error:", err)
	//		}
	//
	//		previousHealthStatus := record.CNAME.Value.Primary.IsHealthy
	//		currentHealthStatus := GetCurrentHealthStatus(record.CNAME.Value.Primary.HealthCheckConfig, previousHealthStatus)
	//		fmt.Println("current health status: ", currentHealthStatus)
	//		record.CNAME.Value.Primary.IsHealthy = currentHealthStatus
	//		valueBytes, _ := json.Marshal(record.CNAME)
	//
	//		modifiedValueBytes := "{\"cname\":" + string(valueBytes) + "}"
	//		if previousHealthStatus != currentHealthStatus {
	//			UpdateHealthStatus(client, key, subKeys[i], modifiedValueBytes)
	//		}
	//
	//		helper.Log("Modified Value", modifiedValueBytes)
	//	}
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
