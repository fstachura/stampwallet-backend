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

type PostAccountEmailConfirmationRequest struct {
	Token string `json:"token,omitempty" binding:"required"`
}
