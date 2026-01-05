package cybersource

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func buildDigest(body []byte) string {
	hash := sha256.Sum256(body)
	return "SHA-256=" + base64.StdEncoding.EncodeToString(hash[:])
}

func signPayload(secretKeyBase64, message string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(secretKeyBase64)
	if err != nil {
		return "", fmt.Errorf("secretKey no es Base64 válido: %w", err)
	}
	mac := hmac.New(sha256.New, keyBytes)
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

func buildMessage(host, vcDate, method, path, digest, merchantID string) string {
	requestTarget := fmt.Sprintf("request-target: %s %s", strings.ToLower(method), path)

	return fmt.Sprintf(
		"host: %s\n"+
			"v-c-date: %s\n"+
			"%s\n"+
			"digest: %s\n"+
			"v-c-merchant-id: %s",
		host, vcDate, requestTarget, digest, merchantID,
	)
}

func buildSignatureHeader(apiKey, signature string) string {
	return fmt.Sprintf(
		`keyid="%s", algorithm="HmacSHA256", headers="host v-c-date request-target digest v-c-merchant-id", signature="%s"`,
		apiKey, signature,
	)
}
