package main

import (
	"testing"
)

func Test_rmqConnect(t *testing.T) {
	type args struct {
		conf RMqConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test With Localhost Default Port",
			args: args{
				conf: RMqConfig{
					Address: "127.0.0.1",
					Port:    "5672",
				},
			},
			wantErr: false,
		},
		{
			name: "Test invalid Port",
			args: args{
				conf: RMqConfig{
					Address: "0.0.0.0",
					Port:    "8035",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := rmqConnect(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("rmqConnect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_startHotelStatsConsumer(t *testing.T) {
	startConsumer = false
	tests := []struct {
		name string
	}{
		{
			name: "Test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startHotelStatsConsumer()
		})
	}
	startConsumer = true
}
