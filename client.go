package cybersource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	cfg    Config
	client *http.Client
}

func NewClient(cfg Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &Client{
		cfg:    cfg,
		client: &http.Client{Timeout: 15 * time.Second},
	}, nil
}

func (c *Client) doRequestRaw(method, endpoint string, payload any) (int, http.Header, []byte, error) {
	var bodyBytes []byte
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return 0, nil, nil, fmt.Errorf("marshal error: %w", err)
		}
		bodyBytes = b
	} else {
		bodyBytes = []byte("{}")
	}

	u, err := url.Parse(c.cfg.BaseURL + endpoint)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("invalid url: %w", err)
	}
	host := u.Host

	vcDate := time.Now().UTC().Format(time.RFC1123)
	vcDate = strings.Replace(vcDate, "UTC", "GMT", 1)

	digest := buildDigest(bodyBytes)
	message := buildMessage(host, vcDate, method, u.Path, digest, c.cfg.MerchantID)

	signature, err := signPayload(c.cfg.SecretKey, message)
	if err != nil {
		return 0, nil, nil, err
	}
	sigHeader := buildSignatureHeader(c.cfg.ApiKey, signature)

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return 0, nil, nil, err
	}

	req.Header.Set("Host", host)
	req.Header.Set("v-c-date", vcDate)
	req.Header.Set("Digest", digest)
	req.Header.Set("v-c-merchant-id", c.cfg.MerchantID)
	req.Header.Set("Signature", sigHeader)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	requestID := res.Header.Get("v-c-correlation-id")
	if requestID == "" {
		requestID = res.Header.Get("x-request-id")
	}

	if res.StatusCode >= 300 {
		return res.StatusCode, res.Header, body, &HTTPError{
			StatusCode: res.StatusCode,
			Status:     res.Status,
			Body:       body,
			Headers:    res.Header,
			RequestID:  requestID,
		}
	}

	return res.StatusCode, res.Header, body, nil
}

func (c *Client) doRequest(method, endpoint string, payload any, result any) error {
	_, _, body, err := c.doRequestRaw(method, endpoint, payload)
	if err != nil {
		return err
	}

	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("decode error: %w", err)
		}
	}

	return nil
}
