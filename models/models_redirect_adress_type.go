package models

type RedirectAddressType string

// List of RedirectAddressType
const (
	IPV4_ADDRRedirectAddressType RedirectAddressType = "IPV4_ADDR"
	IPV6_ADDRRedirectAddressType RedirectAddressType = "IPV6_ADDR"
	URLRedirectAddressType       RedirectAddressType = "URL"
	SIP_URIRedirectAddressType   RedirectAddressType = "SIP_URI"
)
