package models

type ScsAsDataContextsResponse struct {
	ScsAsDataContexts []ScsAsDataContext `json:"scsAsDataContexts,omitempty" yaml:"scsAsDataContexts" bson:"scsAsDataContexts" mapstructure:"ScsAsDataContexts"`
}
