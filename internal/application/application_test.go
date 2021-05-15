package application

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyRequest(t *testing.T, request *http.Request) (*http.Request, ed25519.PublicKey) {
	pubkey, privkey, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Errorf("error generating signing keypair: %s", err)
	}
	timestamp := "1608597133"

	request.Header.Set("X-Signature-Timestamp", timestamp)

	var msg bytes.Buffer
	msg.WriteString(timestamp)

	body, _ := request.GetBody()
	bodyBs, _ := ioutil.ReadAll(body)
	msg.Write(bodyBs)
	signature := ed25519.Sign(privkey, msg.Bytes())
	request.Header.Set("X-Signature-Ed25519", hex.EncodeToString(signature[:ed25519.SignatureSize]))

	return request, pubkey
}

func TestUnverifyRequest(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/interactions", strings.NewReader(`{ "type": 1 }`))
	req, pubKey := VerifyRequest(t, req)

	t.Log("modify the timestamp then this request becomes invalid")
	req.Header.Set("X-Signature-Timestamp", "1234")

	r := GetRouter(pubKey)

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusUnauthorized)
}

func TestVerifyRequest(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/interactions", strings.NewReader(`{ "type": 1 }`))
	req, pubKey := VerifyRequest(t, req)

	r := GetRouter(pubKey)

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func Test404(t *testing.T) {
	w := httptest.NewRecorder()
	r := GetRouter(ed25519.PublicKey{})

	req, _ := http.NewRequest("GET", "/unhandled", strings.NewReader(""))
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)
}
