package helmet

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Helper function to create a test server with middleware
func createTestServer(config Config) *httptest.Server {
	middleware := New(config)
	mux := http.NewServeMux()
	mux.Handle("/", middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})))
	return httptest.NewServer(mux)
}

func Test_DefaultConfig(t *testing.T) {
	server := createTestServer(ConfigDefault)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	if resp.Header.Get("X-XSS-Protection") != "0" {
		t.Errorf("Expected X-XSS-Protection to be '0', got '%s'", resp.Header.Get("X-XSS-Protection"))
	}
	if resp.Header.Get("X-Content-Type-Options") != "nosniff" {
		t.Errorf("Expected X-Content-Type-Options to be 'nosniff', got '%s'", resp.Header.Get("X-Content-Type-Options"))
	}
	if resp.Header.Get("X-Frame-Options") != "SAMEORIGIN" {
		t.Errorf("Expected X-Frame-Options to be 'SAMEORIGIN', got '%s'", resp.Header.Get("X-Frame-Options"))
	}
	if resp.Header.Get("Referrer-Policy") != "no-referrer" {
		t.Errorf("Expected Referrer-Policy to be 'no-referrer', got '%s'", resp.Header.Get("Referrer-Policy"))
	}
}

func Test_CustomConfig(t *testing.T) {
	config := Config{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff-custom",
		XFrameOptions:         "DENY",
		ReferrerPolicy:        "origin",
		PermissionPolicy:      "geolocation=()",
		ContentSecurityPolicy: "default-src 'self'",
	}

	server := createTestServer(config)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	if resp.Header.Get("X-XSS-Protection") != "1; mode=block" {
		t.Errorf("Expected X-XSS-Protection to be '1; mode=block', got '%s'", resp.Header.Get("X-XSS-Protection"))
	}
	if resp.Header.Get("X-Content-Type-Options") != "nosniff-custom" {
		t.Errorf("Expected X-Content-Type-Options to be 'nosniff-custom', got '%s'", resp.Header.Get("X-Content-Type-Options"))
	}
	if resp.Header.Get("X-Frame-Options") != "DENY" {
		t.Errorf("Expected X-Frame-Options to be 'DENY', got '%s'", resp.Header.Get("X-Frame-Options"))
	}
	if resp.Header.Get("Referrer-Policy") != "origin" {
		t.Errorf("Expected Referrer-Policy to be 'origin', got '%s'", resp.Header.Get("Referrer-Policy"))
	}
	if resp.Header.Get("Content-Security-Policy") != "default-src 'self'" {
		t.Errorf("Expected Content-Security-Policy to be 'default-src 'self'', got '%s'", resp.Header.Get("Content-Security-Policy"))
	}
	if resp.Header.Get("Permissions-Policy") != "geolocation=()" {
		t.Errorf("Expected Permissions-Policy to be 'geolocation=()', got '%s'", resp.Header.Get("Permissions-Policy"))
	}
}
