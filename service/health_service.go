package service

import (
	"grabathon/api/request_response"
)

func isRedisHealthy() bool {
	return true
}

func GetServiceHealth() request_response.ServiceHealth {
	// Initialize service healthy
	isServiceHealthy := true

	// Check health for each component
	redisHealth := request_response.HealthComponent{Type: "redis", IsHealthy: isRedisHealthy()}
	if !redisHealth.IsHealthy {
		isServiceHealthy = false
	}

	// Define service health
	serviceHealth := request_response.ServiceHealth{IsHealthy: isServiceHealthy, HealthComponents: []request_response.HealthComponent{}}

	// Add health components
	serviceHealth.HealthComponents = append(serviceHealth.HealthComponents, redisHealth)

	return serviceHealth
}
