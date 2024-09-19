package cuda

type AlphaCompTypes int

const (
	AlphaCompTypeOver AlphaCompTypes = iota
	AlphaCompTypeIn
	AlphaCompTypeOut
	AlphaCompTypeAtop
	AlphaCompTypeXor
	AlphaCompTypePlus
	AlphaCompTypeOverPremul
	AlphaCompTypeInPremul
	AlphaCompTypeOutPremul
	AlphaCompTypeAtopPremul
	AlphaCompTypeXorPremul
	AlphaCompTypePlusPremul
	AlphaCompTypePremul
)
