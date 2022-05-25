package models

import (
	"time"
)

type AccessAndMobilityData struct {
	Location             *UserLocation  `json:"location,omitempty" bson:"location"`
	LocationTs           *time.Time     `json:"locationTs,omitempty" bson:"locationTs"`
	TimeZone             string         `json:"timeZone,omitempty" bson:"timeZone"`
	TimeZoneTs           *time.Time     `json:"timeZoneTs,omitempty" bson:"timeZoneTs"`
	AccessType           AccessType     `json:"accessType,omitempty" bson:"accessType"`
	RegStates            []RmInfo       `json:"regStates,omitempty" bson:"regStates"`
	RegStatesTs          *time.Time     `json:"regStatesTs,omitempty" bson:"regStatesTs"`
	ConnStates           []CmInfo       `json:"connStates,omitempty" bson:"connStates"`
	ConnStatesTs         *time.Time     `json:"connStatesTs,omitempty" bson:"connStatesTs"`
	ReachabilityStatus   UeReachability `json:"reachabilityStatus,omitempty" bson:"reachabilityStatus"`
	ReachabilityStatusTs *time.Time     `json:"reachabilityStatusTs,omitempty" bson:"reachabilityStatusTs"`
	SmsOverNasStatus     SmsSupport     `json:"smsOverNasStatus,omitempty" bson:"smsOverNasStatus"`
	SmsOverNasStatusTs   *time.Time     `json:"smsOverNasStatusTs,omitempty" bson:"smsOverNasStatusTs"`
	// True  The serving PLMN of the UE is different from the HPLMN of the UE; False  The serving PLMN of the UE is the HPLMN of the UE.
	RoamingStatus   bool       `json:"roamingStatus,omitempty" bson:"roamingStatus"`
	RoamingStatusTs *time.Time `json:"roamingStatusTs,omitempty" bson:"roamingStatusTs"`
	CurrentPlmn     *PlmnId    `json:"currentPlmn,omitempty" bson:"currentPlmn"`
	CurrentPlmnTs   *time.Time `json:"currentPlmnTs,omitempty" bson:"currentPlmnTs"`
	RatType         []RatType  `json:"ratType,omitempty" bson:"ratType"`
	RatTypesTs      *time.Time `json:"ratTypesTs,omitempty" bson:"ratTypesTs"`
}
