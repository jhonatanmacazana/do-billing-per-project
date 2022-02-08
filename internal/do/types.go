package do

import (
	"github.com/digitalocean/godo"
)

type ResourceWithInfo struct {
	Name         string  `json:"resourceName,omitempty"`
	Type         string  `json:"resourceType,omitempty"`
	PriceMonthly float64 `json:"priceMonthly,omitempty"`
	PriceHourly  float64 `json:"priceHourly,omitempty"`
	Months       int     `json:"months,omitempty"`
	Total        float64 `json:"total,omitempty"`

	Meta *godo.ProjectResource `json:"meta,omitempty"`
}

type ResourceType string

const (
	Database          ResourceType = "do:dbaas:"
	Domain            ResourceType = "do:domain:"
	Droplet           ResourceType = "do:droplet:"
	FloatingIP        ResourceType = "do:floatingip:"
	KubernetesCluster ResourceType = "do:kubernetes:"
	LoadBalancer      ResourceType = "do:loadbalancer:"
	Space             ResourceType = "do:space:"
	Volume            ResourceType = "do:volume:"
)
