package main

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (offers *HotelOffers) processHotelOffer() {
	for _, offer := range offers.OfferList {
		err := offer.Verify()
		if err != nil {
			log.Errorf("Error [%s] verifying incoming data [%v]", err.Error(), offer)
			continue
		}
		offer.StoreDB(sqlConn)
	}
	return
}

func (offer *Offer) Verify() (err error) {
	if offer.HotelData.HotelId != "" {
		if offer.HotelData.HotelId != offer.RoomData.HotelId ||
			offer.HotelData.HotelId != offer.RatePlan.HotelId {
			return errors.New("Inconsistent data")
		}
		return nil
	}
	return errors.New("Empty Hotel ID")
}

func (offer *Offer) StoreDB(conn *gorm.DB) {
	TableConn["hotel"].Create(&offer.HotelData)
	amenities := TableConn["amenities"]
	for _, amenity := range offer.HotelData.Amenities {
		amenities.Create(&HotelAmenitiesMapping{HotelId: offer.HotelData.HotelId, Amenity: amenity})
	}
	TableConn["room"].Create(&offer.RoomData)
	TableConn["ratePlan"].Create(&offer.RatePlan)
	ratePlanCancellation := TableConn["ratePlanCancellation"]
	for _, Cancellation := range offer.RatePlan.CancellationPolicy {
		ratePlanCancellation.Create(&RatePlanCancellationPolicyMapping{
			HotelId:            offer.RatePlan.HotelId,
			CancellationPolicy: Cancellation,
			RatePlanId:         offer.RatePlan.PlanId,
		})
	}
	ratePlanOtherCond := TableConn["ratePlanOtherCond"]
	for _, Conditions := range offer.RatePlan.OtherConditions {
		ratePlanOtherCond.Create(&RatePlanOtherConditionsMapping{
			HotelId:    offer.RatePlan.HotelId,
			Conditions: Conditions,
			RatePlanId: offer.RatePlan.PlanId,
		})
	}
	return
}

func initDatabase(conn *gorm.DB) (err error) {
	if conn == nil {
		log.Errorf("DB Connection not intialized")
		return errors.New("DB Connection not initialized")
	}
	var Tables []interface{}
	Tables = append(Tables, &Hotel{}, &Room{}, &RatePlan{}, &HotelAmenitiesMapping{},
		&RatePlanOtherConditionsMapping{}, &RatePlanCancellationPolicyMapping{})
	_ = Tables
	for _, table := range Tables {
		err = conn.AutoMigrate(table)
		if err != nil {
			return
		}
	}
	TableConn["hotel"], TableConn["room"], TableConn["ratePlan"] = conn.Table("hotels"), conn.Table("rooms"), conn.Table("rate_plans")
	TableConn["amenities"], TableConn["ratePlanCancellation"], TableConn["ratePlanOtherCond"] = conn.Table("hotel_amenities_mappings"),
		conn.Table("rate_plan_cancellation_policy_mappings"),
		conn.Table("rate_plan_other_conditions_mappings")
	return
}
