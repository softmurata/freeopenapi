package models

type RouteToLocation struct {
	Dnai        string            `json:"dnai" yaml:"dnai" bson:"dnai" mapstructure:"Dnai"`
	RouteInfo   *RouteInformation `json:"routeInfo,omitempty" yaml:"routeInfo" bson:"routeInfo" mapstructure:"RouteInfo"`
	RouteProfId string            `json:"routeProfId,omitempty" yaml:"routeProfId" bson:"routeProfId" mapstructure:"RouteProfId"`
}
