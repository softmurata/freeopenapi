package models

type ScsAsDataContextsResponse struct {
	ScsAsDataContexts []ScsAsDataContextResponse `json:"scsAsDataContexts,omitempty" yaml:"scsAsDataContexts" bson:"scsAsDataContexts" mapstructure:"ScsAsDataContexts"`
}
