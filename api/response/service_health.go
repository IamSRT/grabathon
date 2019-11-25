package response

type HealthComponent struct {
	Type      string `json:"type"`
	IsHealthy bool   `json:"is_healthy"`
}

type ServiceHealth struct {
	IsHealthy        bool              `json:"is_healthy"`
	HealthComponents []HealthComponent `json:"components"`
}
