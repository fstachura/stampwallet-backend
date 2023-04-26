/*
 * StampWallet API Server
 *
 * StampWallet API Server REST Specification
 *
 * API version: 0.1.0
 * Contact: fbstachura@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"time"
)

type PostBusinessItemDefinitionRequest struct {
	PublicId string `json:"publicId,omitempty"`

	Name string `json:"name,omitempty"`

	Price int32 `json:"price,omitempty"`

	Description string `json:"description,omitempty"`

	ImageId string `json:"imageId,omitempty"`

	StartDate *time.Time `json:"startDate,omitempty"`

	EndDate *time.Time `json:"endDate,omitempty"`

	MaxAmount int32 `json:"maxAmount,omitempty"`

	Available bool `json:"available,omitempty"`
}
