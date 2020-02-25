/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SubscriptionData struct {
	AmfStatusUri string  `json:"amfStatusUri"`
	GuamiList    []Guami `json:"guamiList,omitempty"`
}
