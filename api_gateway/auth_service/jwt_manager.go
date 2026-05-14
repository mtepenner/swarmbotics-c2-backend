package authservice

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Claims struct {
	Subject   string   `json:"sub"`
	Audience  string   `json:"aud,omitempty"`
	Issuer    string   `json:"iss,omitempty"`
	Roles     []string `json:"roles,omitempty"`
	IssuedAt  int64    `json:"iat"`
	ExpiresAt int64    `json:"exp"`
}

func IssueToken(secret []byte, claims Claims, now time.Time) (string, error) {
	if len(secret) == 0 {
		return "", errors.New("secret must not be empty")
	}
	if claims.Subject == "" {
		return "", errors.New("subject must not be empty")
	}

	if claims.ExpiresAt == 0 {
		claims.ExpiresAt = now.Add(15 * time.Minute).Unix()
	}
	claims.IssuedAt = now.UTC().Unix()

	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("marshal header: %w", err)
	}
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("marshal claims: %w", err)
	}

	encodedHeader := base64.RawURLEncoding.EncodeToString(headerJSON)
	encodedClaims := base64.RawURLEncoding.EncodeToString(claimsJSON)
	payload := encodedHeader + "." + encodedClaims
	signature := signToken(secret, payload)

	return payload + "." + signature, nil
}

func ValidateToken(token string, secret []byte, now time.Time) (Claims, error) {
	var claims Claims
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return claims, errors.New("token must contain header, claims, and signature")
	}

	signedPayload := parts[0] + "." + parts[1]
	expectedSignature := signToken(secret, signedPayload)
	if subtle.ConstantTimeCompare([]byte(expectedSignature), []byte(parts[2])) != 1 {
		return claims, errors.New("token signature is invalid")
	}

	claimsBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return claims, fmt.Errorf("decode claims: %w", err)
	}
	if err := json.Unmarshal(claimsBytes, &claims); err != nil {
		return claims, fmt.Errorf("unmarshal claims: %w", err)
	}

	if claims.ExpiresAt == 0 || now.UTC().Unix() >= claims.ExpiresAt {
		return claims, errors.New("token is expired")
	}

	return claims, nil
}

func signToken(secret []byte, payload string) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
