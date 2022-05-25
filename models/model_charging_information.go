package models

type ChargingInformation struct {
	PrimaryChfAddress   string `json:"primaryChfAddress" yaml:"primaryChfAddress" bson:"primaryChfAddress" mapstructure:"PrimaryChfAddress"`
	SecondaryChfAddress string `json:"secondaryChfAddress" yaml:"secondaryChfAddress" bson:"secondaryChfAddress" mapstructure:"SecondaryChfAddress"`
}
