package service

import (
	"grabathon/api/response"
)

func isRedisHealthy() bool {
	return true
}

func GetServiceHealth() response.ServiceHealth {
	// Initialize service healthy
	isServiceHealthy := true

	// Check health for each component
	redisHealth := response.HealthComponent{Type: "redis", IsHealthy: isRedisHealthy()}
	if !redisHealth.IsHealthy {
		isServiceHealthy = false
	}

	// Define service health
	serviceHealth := response.ServiceHealth{IsHealthy: isServiceHealthy, HealthComponents: []response.HealthComponent{}}

	// Add health components
	serviceHealth.HealthComponents = append(serviceHealth.HealthComponents, redisHealth)

	return serviceHealth
}
