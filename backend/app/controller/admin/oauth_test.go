package admin

import (
	"errors"
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
		"no bindable oauth account":          "oauth_account_not_bound",
		"context deadline exceeded":          "oauth_provider_unreachable",
		"state invalid or expired":           "oauth_state_expired",
		"provider returned invalid response": "oauth_auth_failed",
	}
	for message, expected := range tests {
		if actual := oauthCallbackErrorCode(errors.New(message)); actual != expected {
			t.Fatalf("oauthCallbackErrorCode(%q) = %q, want %q", message, actual, expected)
		}
	}
}
