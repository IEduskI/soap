package soap

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"time"
)

// generateSecurityHeader Set the values in the security structure.
func generateSecurityHeader(usr, pass string, ttl int64) Security {
	// Generate a random nonce
	nonce := generateNonce()

	// Generate the created timestamp in the required format
	created := time.Now().UTC().Format(time.RFC3339)

	// Concatenate the nonce, created timestamp, and password
	passwordDigest := generatePasswordDigest(string(nonce), created, pass)

	expirationDuration := time.Duration(ttl) * time.Minute

	expiration := time.Now().UTC().Add(expirationDuration).Format(time.RFC3339)

	// Create the Security struct
	header := Header{
		MustUnderstand: "1",
		Wsse:           "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
		Timestamp: Timestamp{
			ID:      "Timestamp-20046406",
			Wsu:     "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
			Created: created,
			Expires: expiration,
		},
		UsernameToken: UsernameToken{
			ID:       "UsernameToken-20914066",
			Wsu:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
			Username: usr,
			Password: Password{
				Type:           "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest",
				PasswordDigest: passwordDigest,
			},
			Nonce:   base64.StdEncoding.EncodeToString(nonce),
			Created: created,
		},
	}

	return Security{Header: header}
}

func generateNonce() []byte {
	nonce := make([]byte, 16)
	io.ReadFull(rand.Reader, nonce)
	return nonce
}

func generatePasswordDigest(nonce, created, password string) string {
	sha := sha1.New()
	sha.Write([]byte(nonce + created + password))
	return base64.StdEncoding.EncodeToString(sha.Sum(nil))
}
