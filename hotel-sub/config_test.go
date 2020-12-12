package main

import (
	"testing"
)

func Test_hotelOfferMgrConfig_init(t *testing.T) {
	type fields struct {
		DbConfig  DbConfig
		RMqConfig RMqConfig
		cfgFile   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Config Init",
			fields: fields{
				cfgFile: "./dependency/hotelOfferMgrConfig.yaml",
			},
			wantErr: false,
		},
		{
			name: "Config Init Invalid",
			fields: fields{
				cfgFile: "./dependency/invalidConfig.yaml",
			},
			wantErr: true,
		},
		{
			name: "Config Init Invalid",
			fields: fields{
				cfgFile: "./dependency/invalidConfidg.yaml",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &hotelOfferMgrConfig{
				DbConfig:  tt.fields.DbConfig,
				RMqConfig: tt.fields.RMqConfig,
			}
			configFile = tt.fields.cfgFile
			if err := config.init(); (err != nil) != tt.wantErr {
				t.Errorf("hotelOfferMgrConfig.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_hotelOfferMgrConfig_setDefaults(t *testing.T) {
	configFile = "./hotelOfferMgrConfig.yaml"
	type fields struct {
		DbConfig  DbConfig
		RMqConfig RMqConfig
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Set Default",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &hotelOfferMgrConfig{
				DbConfig:  tt.fields.DbConfig,
				RMqConfig: tt.fields.RMqConfig,
			}
			config.setDefaults()
		})
	}
}
