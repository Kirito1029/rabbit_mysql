package main

type (
	CancellationPolicy struct {
		Type         string `json:"type,omitempty"`
		DaysToExpiry int    `json:"expires_days_before,omitempty"`
	}
	Capacity struct {
		Adult         int `json:"max_adults,omitempty" gorm:"column:adultcapacity"`
		ExtraChildren int `json:"extra_children,omitempty" gorm:"column:extrachildren"`
	}
	Hotel struct {
		HotelId     string   `json:"hotel_id,omitempty" gorm:"column:id;size:10"`
		Name        string   `json:"name,omitempty" gorm:"column:name"`
		Country     string   `json:"country,omitempty" gorm:"column:country"`
		Address     string   `json:"address,omitempty" gorm:"column:address"`
		Latitude    float64  `json:"latitude,omitempty" gorm:"column:latitude"`
		Longitude   float64  `json:"longitude,omitempty" gorm:"column:longitude"`
		Telephone   string   `json:"telephone,omitempty" gorm:"column:telephone"`
		Amenities   []string `json:"amenities,omitempty" gorm:"-"`
		Description string   `json:"description,omitempty" gorm:"column:description"`
		RoomCount   int      `json:"room_count,omitempty" gorm:"column:roomcount"`
		Currency    string   `json:"currency,omitempty" gorm:"column:currency"`
	}
	Room struct {
		HotelId     string   `json:"hotel_id,omitempty" gorm:"column:hotelid;size:10"`
		RoomId      string   `json:"room_id,omitempty" gorm:"column:roomid"`
		Description string   `json:"description,omitempty" gorm:"column:description"`
		Name        string   `json:"name,omitempty" gorm:"column:name"`
		Capacity    Capacity `json:"capacity,omitempty" gorm:"embedded"`
		Hotel       Hotel    `gorm:"foreignKey:hotelid"`
	}

	RatePlan struct {
		HotelId            string               `json:"hotel_id,omitempty" gorm:"column:hotelid;size:10"`
		CancellationPolicy []CancellationPolicy `json:"cancellation_policy,omitempty" gorm:"-"`
		PlanId             string               `json:"rate_plan_id,omitempty" gorm:"column:planid;primarykey;size:10"`
		Name               string               `json:"name,omitempty" gorm:"column:name"`
		OtherConditions    []string             `json:"other_conditions,omitempty" gorm:"-"`
		MealPlan           string               `json:"meal_plan,omitempty" gorm:"column:mealplan"`
		Hotel              Hotel                `gorm:"foreignKey:hotelid"`
	}
	Offer struct {
		OfferId   string   `json:"cm_offer_id,omitempty"`
		HotelData Hotel    `json:"hotel,omitempty"`
		RoomData  Room     `json:"room,omitempty"`
		RatePlan  RatePlan `json:"rate_plan,omitempty"`
	}
	HotelOffers struct {
		OfferList []Offer `json:"offers"`
	}
	DbConfig struct {
		Address      string `yaml:"address"`
		Port         string `yaml:"port"`
		DatabaseName string `yaml:"databaseName"`
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
	}
	RMqConfig struct {
		Address string `yaml:"address"`
		Port    string `yaml:"port"`
	}
	hotelOfferMgrConfig struct {
		DbConfig  DbConfig  `yaml:"dbconfig"`
		RMqConfig RMqConfig `yaml:"rabbitmqconfig"`
	}
	HotelAmenitiesMapping struct {
		Amenity string `gorm:"column:amenity;index:HAM,unique"`
		HotelId string `gorm:"column:hotelid;size:10;index:HAM,unique"`
		Hotel   Hotel  `gorm:"foreignKey:hotelid"`
	}
	RatePlanCancellationPolicyMapping struct {
		CancellationPolicy CancellationPolicy `gorm:"embedded;index:CanHotelID,unique"`
		HotelId            string             `gorm:"column:hotelid;size:10;index:CanHotelID,unique"`
		RatePlanId         string             `gorm:"column:rateplanid;size:10;index:CanHotelID,unique"`
		Hotel              Hotel              `gorm:"foreignKey:hotelid"`
		RatePlan           RatePlan           `gorm:"foreignKey:rateplanid"`
	}
	RatePlanOtherConditionsMapping struct {
		Conditions string   `gorm:"index:ConHotelID,unique"`
		HotelId    string   `gorm:"column:hotelid;size:10;index:ConHotelID,unique"`
		RatePlanId string   `gorm:"column:rateplanid;size:10;index:ConHotelID,unique"`
		Hotel      Hotel    `gorm:"foreignKey:hotelid"`
		RatePlan   RatePlan `gorm:"foreignKey:rateplanid"`
	}
)
