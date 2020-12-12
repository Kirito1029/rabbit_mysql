package main

import (
	"testing"
)

func Test_generateDSN(t *testing.T) {
	type args struct {
		config hotelOfferMgrConfig
	}
	tests := []struct {
		name    string
		args    args
		wantDsn string
	}{
		{
			name: "Generate DSN",
			args: args{
				config: hotelOfferMgrConfig{
					DbConfig: DbConfig{
						Address:      "127.0.0.1",
						Port:         "1534",
						DatabaseName: "testdatabase",
						Username:     "user",
						Password:     "test@111",
					},
				},
			},
			wantDsn: "user:test@111@tcp(127.0.0.1:1534)/testdatabase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDsn := generateDSN(tt.args.config); gotDsn != tt.wantDsn {
				t.Errorf("generateDSN() = %v, want %v", gotDsn, tt.wantDsn)
			}
		})
	}
}

func Test_sqlConnect(t *testing.T) {
	type args struct {
		conf DbConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sql Connect Default Config",
			args: args{
				conf: config.DbConfig,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sqlConnect(tt.args.conf)
			if (err != nil) != tt.wantErr {
				if isSqlInit {
					t.Errorf("sqlConnect() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
		})
	}
}

func Test_main(t *testing.T) {
	startConsumer = false
	tests := []struct {
		name string
	}{
		{
			name: "Test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
	startConsumer = true
}
