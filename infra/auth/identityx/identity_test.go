package identityx

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestKindOf(t *testing.T) {
	v7 := mustUUIDv7(t)
	v4 := uuid.New()

	tests := []struct {
		name string
		id   string
		want Kind
	}{
		{"uuidv7", v7.String(), KindUser},
		{"uuidv7 uppercase", strings.ToUpper(v7.String()), KindUnknown},
		{"uuidv4", v4.String(), KindUnknown},
		{"uuid compact (historical user shape)", strings.ReplaceAll(v7.String(), "-", ""), KindGuest},
		{"md5 guest", "d41d8cd98f00b204e9800998ecf8427e", KindGuest},
		{"uppercase hex", "D41D8CD98F00B204E9800998ECF8427E", KindUnknown},
		{"seed root", "root", KindUnknown},
		{"empty", "", KindUnknown},
		{"invalid uuid", "0190f3a1-7c2e-7c3a-9f4e-1a2b3c4d5e6g", KindUnknown},
		{"urn form", v7.URN(), KindUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KindOf(tt.id); got != tt.want {
				t.Fatalf("KindOf(%q) = %v, want %v", tt.id, got, tt.want)
			}
		})
	}
}

func TestNewUserID(t *testing.T) {
	id, err := NewUserID()
	if err != nil {
		t.Fatalf("NewUserID() error: %v", err)
	}

	u, err := uuid.Parse(id)
	if err != nil {
		t.Fatalf("NewUserID() returned unparsable id %q: %v", id, err)
	}
	if u.Version() != 7 {
		t.Fatalf("NewUserID() version = %d, want 7", u.Version())
	}
	if u.String() != id {
		t.Fatalf("NewUserID() = %q, not canonical lowercase form", id)
	}
	if len(id) != userIDLength {
		t.Fatalf("NewUserID() length = %d, want %d", len(id), userIDLength)
	}
	if KindOf(id) != KindUser {
		t.Fatalf("KindOf(NewUserID()) != KindUser")
	}
}

func TestNewGuestID(t *testing.T) {
	id, err := NewGuestID()
	if err != nil {
		t.Fatalf("NewGuestID() error: %v", err)
	}
	if len(id) != guestIDLength {
		t.Fatalf("NewGuestID() length = %d, want %d", len(id), guestIDLength)
	}
	if KindOf(id) != KindGuest {
		t.Fatalf("KindOf(NewGuestID()) = %v, want guest", KindOf(id))
	}
}

func mustUUIDv7(t *testing.T) uuid.UUID {
	t.Helper()
	id, err := uuid.NewV7()
	if err != nil {
		t.Fatalf("uuid.NewV7() error: %v", err)
	}
	return id
}
