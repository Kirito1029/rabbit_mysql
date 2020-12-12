package main

import (
	"testing"

	"github.com/streadway/amqp"
)

func Test_rmqConnect(t *testing.T) {
	type args struct {
		conf RMqConfig
	}
	tests := []struct {
		name     string
		args     args
		wantConn *amqp.Connection
		wantErr  bool
	}{
		{
			name: "Defaults",
			args: args{
				conf: config.RMqConfig,
			},
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

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Main",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
