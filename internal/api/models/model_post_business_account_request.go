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

type PostBusinessAccountRequest struct {
	Name string `json:"name,omitempty" binding:"required"`

	Address string `json:"address,omitempty" binding:"required"`

	GpsCoordinates string `json:"gpsCoordinates,omitempty" binding:"required"`

	Description string `json:"description,omitempty" binding:"required"`

	Nip string `json:"nip,omitempty" binding:"required"`

	Krs string `json:"krs,omitempty" binding:"required"`

	Regon string `json:"regon,omitempty" binding:"required"`

	OwnerName string `json:"ownerName,omitempty" binding:"required"`
}
