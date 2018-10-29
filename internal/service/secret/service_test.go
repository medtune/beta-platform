package secret

import (
	"testing"
)

func clearSecrets() {
	secrets = make(map[string][]string, 0)
	secrets["signup"] = make([]string, 0, 0)
	secrets["auth"] = make([]string, 0, 0)
}

func TestRegister(t *testing.T) {
	type args struct {
		service string
		key     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"first registration", args{"auth", "abra"}, false},
		{"first registration", args{"signup", "cadabra"}, false},
		{"unknown service", args{"service 3", "abcdefgh"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Register(tt.args.service, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	clearSecrets()
}

func TestCheck(t *testing.T) {
	type args struct {
		service string
		key     string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"non existant", args{"auth", "abraaaaaa"}, false, true},
		{"already exist", args{"signup", "cadabra"}, true, false},
		{"unknown service", args{"service abc", "cadabra"}, false, true},
	}
	Register("auth", "abra")
	Register("signup", "cadabra")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Check(tt.args.service, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
	clearSecrets()
}
