package middlewares

import (
	"testing"

	"github.com/rezkyal/simple-go-login/entity/config"
)

func TestJwtAuthMiddleware(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			name:    "success",
			args:    args{},
			wantNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JwtAuthMiddleware(tt.args.cfg); (got == nil) != tt.wantNil {
				t.Errorf("JwtAuthMiddleware() = %v, wantNil %v", got, tt.wantNil)
			}
		})
	}
}
