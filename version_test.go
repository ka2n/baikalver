package baikalver

import "testing"

func TestVersionFromFWV(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"ok", args{"7500055"}, "1.7", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VersionFromFWV(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("VersionFromFWV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VersionFromFWV() = %v, want %v", got, tt.want)
			}
		})
	}
}
