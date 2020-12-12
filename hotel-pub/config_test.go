package main

import (
	"testing"
)

func Test_hotelOfferPubConfig_init(t *testing.T) {
	type fields struct {
		RMqConfig RMqConfig
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Test 1",
			fields:  fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &hotelOfferPubConfig{
				RMqConfig: tt.fields.RMqConfig,
			}
			if err := config.init(); (err != nil) != tt.wantErr {
				t.Errorf("hotelOfferPubConfig.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_hotelOfferPubConfig_setDefaults(t *testing.T) {
	type fields struct {
		RMqConfig RMqConfig
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Test 1",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &hotelOfferPubConfig{
				RMqConfig: tt.fields.RMqConfig,
			}
			config.setDefaults()
		})
	}
}
