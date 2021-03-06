/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UdmInfo struct {
	GroupId string `json:"groupId,omitempty" bson:"groupId"`

	SupiRanges *[]SupiRange `json:"supiRanges,omitempty" bson:"supiRanges"`

	GpsiRanges *[]IdentityRange `json:"gpsiRanges,omitempty" bson:"gpsiRanges"`

	ExternalGroupIdentifiersRanges *[]IdentityRange `json:"externalGroupIdentifiersRanges,omitempty" bson:"externalGroupIdentifiersRanges"`

	RoutingIndicators []string `json:"routingIndicators,omitempty" bson:"routingIndicators"`
}
