package api

import (
	"encoding/json"
	"testing"
)

func TestAbPeerSameServerAcceptsOfficialClientShape(t *testing.T) {
	for _, tc := range []struct {
		name string
		body string
		want *bool
	}{
		{name: "true", body: `{"id":"10001","same_server":true}`, want: boolPtr(true)},
		{name: "false", body: `{"id":"10001","same_server":false}`, want: boolPtr(false)},
		{name: "null", body: `{"id":"10001","same_server":null}`, want: nil},
		{name: "missing", body: `{"id":"10001"}`, want: nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got AbPeer
			if err := json.Unmarshal([]byte(tc.body), &got); err != nil {
				t.Fatalf("unmarshal official peer payload: %v", err)
			}
			if tc.want == nil {
				if got.SameServer != nil {
					t.Fatalf("same_server = %v, want nil", *got.SameServer)
				}
				return
			}
			if got.SameServer == nil || *got.SameServer != *tc.want {
				t.Fatalf("same_server = %v, want %v", got.SameServer, *tc.want)
			}
		})
	}
}

func boolPtr(v bool) *bool { return &v }
