package baikalver

import (
	"testing"
)

func TestModelFromHWV(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Model
		wantErr bool
	}{
		{"GX10", args{"29297"}, GX10, false},
		{"G28", args{"29302"}, G28, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ModelFromHWV(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelFromHWV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ModelFromHWV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModel_Equal(t *testing.T) {
	type args struct {
		v Model
	}
	tests := []struct {
		name string
		i    Model
		args args
		want bool
	}{
		{"true", GX10, args{GX10}, true},
		{"false", GX10, args{Giant}, false},
		{"Cube", Cube, args{CubeRev2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Equal(tt.args.v); got != tt.want {
				t.Errorf("Model.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModel_String(t *testing.T) {
	tests := []struct {
		name string
		i    Model
		want string
	}{
		{"ok", GX10, "GX10"},
		{"unknown", Unknown, ""},
		{"unknown2", Model(0x99), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("Model.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
