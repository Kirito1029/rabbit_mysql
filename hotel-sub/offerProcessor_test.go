package main

import (
	"testing"

	"gorm.io/gorm"
)

func Test_initDatabase(t *testing.T) {
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test DB Init",
			args: args{
				conn: sqlConn,
			},
			wantErr: false,
		},
		{
			name:    "Test DB Init",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initDatabase(tt.args.conn); (err != nil) != tt.wantErr {
				if isSqlInit {
					t.Errorf("initDatabase() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestOffer_StoreDB(t *testing.T) {
	type fields struct {
		OfferId   string
		HotelData Hotel
		RoomData  Room
		RatePlan  RatePlan
	}
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test Offer Store 1",
			fields: fields{
				OfferId: "of1",
				HotelData: Hotel{
					HotelId: "10",
					Name:    "test-1",
					Country: "India",
					Address: "Nil",
					Amenities: []string{
						"te",
					},
				},
				RoomData: Room{
					HotelId:     "10",
					RoomId:      "R1",
					Description: "test desc",
					Capacity: Capacity{
						Adult: 1,
					},
				},
				RatePlan: RatePlan{
					HotelId: "10",
					CancellationPolicy: []CancellationPolicy{
						CancellationPolicy{
							Type:         "testType",
							DaysToExpiry: 2,
						},
					},
					OtherConditions: []string{
						"Test Condition",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offer := &Offer{
				OfferId:   tt.fields.OfferId,
				HotelData: tt.fields.HotelData,
				RoomData:  tt.fields.RoomData,
				RatePlan:  tt.fields.RatePlan,
			}
			offer.StoreDB(tt.args.conn)
		})
	}
}

func TestOffer_Verify(t *testing.T) {
	type fields struct {
		OfferId   string
		HotelData Hotel
		RoomData  Room
		RatePlan  RatePlan
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Test Empty",
			wantErr: true,
		},
		{
			name: "Test Filled",
			fields: fields{
				OfferId: "of1",
				HotelData: Hotel{
					HotelId: "12",
					Name:    "test-1",
					Country: "India",
					Address: "Nil",
				},
				RoomData: Room{
					HotelId:     "12",
					RoomId:      "R1",
					Description: "test desc",
					Capacity: Capacity{
						Adult: 1,
					},
				},
				RatePlan: RatePlan{
					HotelId: "12",
					CancellationPolicy: []CancellationPolicy{
						CancellationPolicy{
							Type:         "testType",
							DaysToExpiry: 2,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offer := &Offer{
				OfferId:   tt.fields.OfferId,
				HotelData: tt.fields.HotelData,
				RoomData:  tt.fields.RoomData,
				RatePlan:  tt.fields.RatePlan,
			}
			if err := offer.Verify(); (err != nil) != tt.wantErr {
				t.Errorf("Offer.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHotelOffers_processHotelOffer(t *testing.T) {
	type fields struct {
		OfferList []Offer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test Empty",
		},
		{
			name: "Test Filled",
			fields: fields{
				OfferList: []Offer{
					Offer{

						OfferId: "of1",
						HotelData: Hotel{
							HotelId: "11",
							Name:    "test-1",
							Country: "India",
							Address: "Nil",
						},
						RoomData: Room{
							HotelId:     "11",
							RoomId:      "R1",
							Description: "test desc",
							Capacity: Capacity{
								Adult: 1,
							},
						},
						RatePlan: RatePlan{
							HotelId: "11",
							CancellationPolicy: []CancellationPolicy{
								CancellationPolicy{
									Type:         "testType",
									DaysToExpiry: 2,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Test Filled",
			fields: fields{
				OfferList: []Offer{
					Offer{

						OfferId: "of1",
						HotelData: Hotel{
							HotelId: "12",
							Name:    "test-1",
							Country: "India",
							Address: "Nil",
						},
						RoomData: Room{
							HotelId:     "11",
							RoomId:      "R1",
							Description: "test desc",
							Capacity: Capacity{
								Adult: 1,
							},
						},
						RatePlan: RatePlan{
							HotelId: "11",
							CancellationPolicy: []CancellationPolicy{
								CancellationPolicy{
									Type:         "testType",
									DaysToExpiry: 2,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offers := &HotelOffers{
				OfferList: tt.fields.OfferList,
			}
			offers.processHotelOffer()
		})
	}
}
