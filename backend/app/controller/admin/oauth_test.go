package admin

import (
	"errors"
	"rustdesk-api-server-pro/internal/errcode"
	"testing"
)

func TestValidateOAuthRedirectURL(t *testing.T) {
	for _, valid := range []string{"", "https://desk.example.com/admin/auth/oauth/github/callback", "http://127.0.0.1/callback"} {
		if err := validateOAuthRedirectURL(valid); err != nil {
			t.Fatalf("expected %q to be valid: %v", valid, err)
		}
	}
	for _, invalid := range []string{"javascript:alert(1)", "/relative/callback", "ftp://desk.example.com/callback"} {
		if err := validateOAuthRedirectURL(invalid); err == nil {
			t.Fatalf("expected %q to be rejected", invalid)
		}
	}
}

func TestCleanOAuthValues(t *testing.T) {
	values := cleanOAuthValues([]string{" read:user ", "user:email", "read:user", ""})
	if len(values) != 2 || values[0] != "read:user" || values[1] != "user:email" {
		t.Fatalf("unexpected cleaned values: %#v", values)
	}
}

func TestOAuthCallbackErrorCode(t *testing.T) {
	tests := map[string]string{
		"no bindable oauth account":          errcode.ERR2208.Message,
		"context deadline exceeded":          errcode.ERR2209.Message,
		"state invalid or expired":           errcode.ERR2210.Message,
		"provider returned invalid response": errcode.ERR2212.Message,
	}
	for message, expected := range tests {
		if actual := oauthCallbackErrorCode(errors.New(message)); actual != expected {
			t.Fatalf("oauthCallbackErrorCode(%q) = %q, want %q", message, actual, expected)
		}
	}
	errcodeTests := map[errcode.Entry]string{
		errcode.ERR2023: errcode.ERR2208.Message,
		errcode.ERR2004: errcode.ERR2210.Message,
		errcode.ERR2030: errcode.ERR2209.Message,
	}
	for entry, expected := range errcodeTests {
		err := errcode.New(entry.Code, entry.Message)
		if actual := oauthCallbackErrorCode(err); actual != expected {
			t.Fatalf("oauthCallbackErrorCode(%q) = %q, want %q", err.Error(), actual, expected)
		}
	}
}

func TestMaskOAuthSecret(t *testing.T) {
	if actual := maskOAuthSecret("github-secret-a0f09583"); actual != "********a0f09583" {
		t.Fatalf("unexpected secret hint: %q", actual)
	}
	if actual := maskOAuthSecret("short"); actual != "********" {
		t.Fatalf("short secrets must not be exposed: %q", actual)
	}
	if actual := maskOAuthSecret(""); actual != "" {
		t.Fatalf("empty secret hint should stay empty: %q", actual)
	}
}
