package models

type FetchIncidents struct {
	Incident Incident `json:"incident"`
	Status []QryIncidents `json:"statuses"`
}