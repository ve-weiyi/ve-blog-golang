package routex

import "testing"

func TestNormalize(t *testing.T) {
	cases := []struct{ in, want string }{
		{"/api/v1/users", "/api/v1/users"},
		{"/api/v1/users/", "/api/v1/users"},
		{"api/v1/users", "/api/v1/users"},
		{"/api/v1/users?page=1", "/api/v1/users"},
		{"/api/v1/users/{id}", "/api/v1/users/:id"},
		{"/api/v1/users/{id}/posts/{post_id}", "/api/v1/users/:id/posts/:post_id"},
		{"/", "/"},
		{"", ""},
	}
	for _, c := range cases {
		if got := Normalize(c.in); got != c.want {
			t.Errorf("Normalize(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMatch(t *testing.T) {
	cases := []struct {
		pattern string
		path    string
		want    bool
	}{
		{"/api/v1/users", "/api/v1/users", true},
		{"/api/v1/users", "/api/v1/posts", false},
		{"/api/v1/users/:id", "/api/v1/users/123", true},
		{"/api/v1/users/{id}", "/api/v1/users/123", true},
		{"/api/v1/users/:id", "/api/v1/users", false},
		{"/api/v1/users/:id", "/api/v1/users/123/extra", false},
		{"/api/v1/*", "/api/v1/users", true},
		{"/api/v1/*", "/api/v1/users/123", false},
		{"/api/v1/**", "/api/v1/users/123/extra", true},
		{"/api/v1/**", "/api/v1", true},
		{"/api/v1/**/extra", "/api/v1/users/123/extra", true},
		{"/api/v1/users/", "/api/v1/users", true},
	}
	for _, c := range cases {
		if got := Match(c.pattern, c.path); got != c.want {
			t.Errorf("Match(%q, %q) = %v, want %v", c.pattern, c.path, got, c.want)
		}
	}
}
