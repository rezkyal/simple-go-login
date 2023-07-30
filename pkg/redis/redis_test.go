package redis

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisPkg_SetJSONTTL(t *testing.T) {
	type args struct {
		key   string
		value interface{}
		TTL   time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "failed to marshal",
			args: args{
				key:   "keyredis",
				value: make(chan int),
				TTL:   1,
			},

			wantErr: true,
		},
		{
			name: "success",
			args: args{
				key:   "keyredis",
				value: "",
				TTL:   1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := miniredis.RunT(t)
			r, _ := NewRedis(&redis.Options{
				Addr: s.Addr(),
			})

			if err := r.SetJSONTTL(context.Background(), tt.args.key, tt.args.value, tt.args.TTL); (err != nil) != tt.wantErr {
				t.Errorf("RedisPkg.SetJSONTTL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisPkg_GetJSON(t *testing.T) {
	type testData struct {
		Aaa string `json:"aaa"`
	}

	type args struct {
		key     string
		data    interface{}
		newData interface{}
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func(*miniredis.Miniredis)
		wantErr  bool
	}{
		{
			name: "failed when get json",
			args: args{
				key:     "testkey",
				data:    &testData{},
				newData: &testData{},
			},
			mockFunc: func(m *miniredis.Miniredis) {},
			wantErr:  true,
		},
		{
			name: "success",
			args: args{
				key:  "testkey1",
				data: &testData{},
				newData: &testData{
					Aaa: "123",
				},
			},
			mockFunc: func(m *miniredis.Miniredis) {
				m.Set("testkey1", `{"aaa":"123"}`)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := miniredis.RunT(t)
			r, _ := NewRedis(&redis.Options{
				Addr: s.Addr(),
			})

			tt.mockFunc(s)

			if err := r.GetJSON(context.Background(), tt.args.key, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("RedisPkg.GetJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.args.newData, tt.args.data)

		})
	}
}
