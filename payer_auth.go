package cybersource

import (
	"encoding/json"

	"github.com/hugochinchilla79/cybersource_sdk_go/models"
)

// SetupPayerAuth sets up payer authentication with Cardinal.
// This service must be called before the enrollment check.
// POST /risk/v1/authentication-setups
func (c *Client) SetupPayerAuth(req models.PayerAuthSetupRequest) (models.PayerAuthSetupAPIResponse, error) {
	status, _, body, err := c.doRequestRaw("POST", "/risk/v1/authentication-setups", req)

	var parsed models.PayerAuthSetupResponse
	_ = json.Unmarshal(body, &parsed)

	apiResp := models.PayerAuthSetupAPIResponse{
		HTTPStatus: status,
		Body:       body,
		Data:       parsed,
	}
	return apiResp, err
}

// CheckEnrollment verifies that a card is enrolled in a payer authentication program.
// This is used to determine if 3DS authentication is required.
// POST /risk/v1/authentications
func (c *Client) CheckEnrollment(req models.CheckEnrollmentRequest) (models.CheckEnrollmentAPIResponse, error) {
	status, _, body, err := c.doRequestRaw("POST", "/risk/v1/authentications", req)

	var parsed models.CheckEnrollmentResponse
	_ = json.Unmarshal(body, &parsed)

	apiResp := models.CheckEnrollmentAPIResponse{
		HTTPStatus: status,
		Body:       body,
		Data:       parsed,
	}
	return apiResp, err
}

// ValidateAuthenticationResults validates the authentication results after a challenge.
// Call this after the user completes the 3DS challenge and is redirected back to returnUrl.
// POST /risk/v1/authentication-results
func (c *Client) ValidateAuthenticationResults(req models.ValidateAuthResultsRequest) (models.CheckEnrollmentAPIResponse, error) {
	status, _, body, err := c.doRequestRaw("POST", "/risk/v1/authentication-results", req)

	var parsed models.CheckEnrollmentResponse
	_ = json.Unmarshal(body, &parsed)

	apiResp := models.CheckEnrollmentAPIResponse{
		HTTPStatus: status,
		Body:       body,
		Data:       parsed,
	}
	return apiResp, err
}
