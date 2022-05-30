package models

type EthFlowInfo struct {
	DestMacAddr    string
	EthType        string
	FDescs         string
	FDir           []string
	SourceMacAddr  string
	VlanTags       []string
	SrcMacAddrEnd  string
	DestMacAddrEnd string
}
