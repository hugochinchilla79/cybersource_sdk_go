package cybersource

import (
	"encoding/json"
	"fmt"

	"github.com/hugochinchilla79/cybersource_sdk_go/models"
)

func (c *Client) Authorize(req models.PaymentRequest) (models.PaymentAPIResponse, error) {
	status, _, body, err := c.doRequestRaw("POST", "/pts/v2/payments", req)

	var parsed models.PaymentResponse
	_ = json.Unmarshal(body, &parsed)

	apiResp := models.PaymentAPIResponse{
		HTTPStatus: status,
		Body:       body,
		Data:       parsed,
	}
	return apiResp, err
}

func (c *Client) AuthorizePayment(req models.PaymentRequest) (models.PaymentAPIResponse, error) {
	status, _, body, err := c.doRequestRaw("POST", "/pts/v2/payments", req)

	var parsed models.PaymentResponse
	_ = json.Unmarshal(body, &parsed)

	apiResp := models.PaymentAPIResponse{
		HTTPStatus: status,
		Body:       body,
		Data:       parsed,
	}

	return apiResp, err
}

func (c *Client) Capture(paymentID string, req models.CaptureRequest) (models.PaymentAPIResponse, error) {
	endpoint := fmt.Sprintf("/pts/v2/payments/%s/captures", paymentID)

	status, _, body, err := c.doRequestRaw("POST", endpoint, req)

	var parsed models.PaymentResponse
	_ = json.Unmarshal(body, &parsed)

	apiResp := models.PaymentAPIResponse{
		HTTPStatus: status,
		Body:       body,
		Data:       parsed,
	}

	return apiResp, err
}

func (c *Client) ReversePayment(paymentID string, req models.ReversalRequest) (models.PaymentAPIResponse, error) {
	endpoint := fmt.Sprintf("/pts/v2/payments/%s/reversals", paymentID)

	status, _, body, err := c.doRequestRaw("POST", endpoint, req)

	var parsed models.PaymentResponse
	_ = json.Unmarshal(body, &parsed)

	apiResp := models.PaymentAPIResponse{
		HTTPStatus: status,
		Body:       body,
		Data:       parsed,
	}

	return apiResp, err
}

func (c *Client) RefundCapture(captureID string, req models.RefundRequest) (models.PaymentAPIResponse, error) {
	endpoint := fmt.Sprintf("/pts/v2/captures/%s/refunds", captureID)

	status, _, body, err := c.doRequestRaw("POST", endpoint, req)

	var parsed models.PaymentResponse
	_ = json.Unmarshal(body, &parsed)

	apiResp := models.PaymentAPIResponse{
		HTTPStatus: status,
		Body:       body,
		Data:       parsed,
	}

	return apiResp, err
}
