package validator

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"Naufal", "naufal@example.com", false},
		{"", "naufal@example.com", true},
		{"Naufal", "not-an-email", true},
	}

	for _, tt := range tests {
		err := ValidateUser(tt.name, tt.email)
		if (err != nil) != tt.wantErr {
			t.Fatalf("Validate(%q, %q) error = %v, wantErr = %v",
				tt.name, tt.email, err, tt.wantErr)
		}
	}
}
