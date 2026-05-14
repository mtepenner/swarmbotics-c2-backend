package authservice

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"
)

type CertificatePolicy struct {
	AllowedCommonNames []string
	MinimumValidity    time.Duration
}

func ValidateClientCertificate(pemBytes []byte, roots *x509.CertPool, policy CertificatePolicy, now time.Time) error {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return errors.New("certificate PEM block not found")
	}

	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("parse certificate: %w", err)
	}

	if now.Before(certificate.NotBefore) || now.After(certificate.NotAfter) {
		return errors.New("certificate is outside of its validity window")
	}

	if certificate.NotAfter.Sub(now) < policy.MinimumValidity {
		return errors.New("certificate validity is below the minimum policy window")
	}

	if len(policy.AllowedCommonNames) > 0 {
		allowed := false
		for _, commonName := range policy.AllowedCommonNames {
			if strings.EqualFold(commonName, certificate.Subject.CommonName) {
				allowed = true
				break
			}
		}
		if !allowed {
			return fmt.Errorf("common name %q is not approved", certificate.Subject.CommonName)
		}
	}

	if roots != nil {
		verifyOptions := x509.VerifyOptions{Roots: roots, CurrentTime: now}
		if _, err := certificate.Verify(verifyOptions); err != nil {
			return fmt.Errorf("verify certificate chain: %w", err)
		}
	}

	return nil
}
