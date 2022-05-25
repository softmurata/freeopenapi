package models

type QosCharacteristics struct {
	Var5qi            int32           `json:"5qi" yaml:"5qi" bson:"5qi" mapstructure:"Var5qi"`
	ResourceType      QosResourceType `json:"resourceType" yaml:"resourceType" bson:"resourceType" mapstructure:"ResourceType"`
	PriorityLevel     int32           `json:"priorityLevel" yaml:"priorityLevel" bson:"priorityLevel" mapstructure:"PriorityLevel"`
	PacketDelayBudget int32           `json:"packetDelayBudget" yaml:"packetDelayBudget" bson:"packetDelayBudget" mapstructure:"PacketDelayBudget"`
	PacketErrorRate   string          `json:"packetErrorRate" yaml:"packetErrorRate" bson:"packetErrorRate" mapstructure:"PacketErrorRate"`
	AveragingWindow   int32           `json:"averagingWindow,omitempty" yaml:"averagingWindow" bson:"averagingWindow" mapstructure:"AveragingWindow"`
	MaxDataBurstVol   int32           `json:"maxDataBurstVol,omitempty" yaml:"maxDataBurstVol" bson:"maxDataBurstVol" mapstructure:"MaxDataBurstVol"`
}
