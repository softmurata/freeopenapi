package models

type TestRequest struct {
	Id string `json:"id", yaml:"id", bson:"id", mapstructure:"Id"`
	Name string `json:"name", yaml:"name", bson:"name", mapstructure:"Name"`
	Age int32 `json:"age", yaml:"age", bson:"age", mapstructure:"Age"`
}
