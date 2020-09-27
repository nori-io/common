package meta

type LicenseType uint8

const (
	Unlicensed LicenseType = iota
	Proprietary
	Custom
	GPLv3
	GPLv2
	LGPLv3
	LGPLv2_1
	AGPLv3_0
	Apache2_0
	MPL_2_0
	PublicDomain
)
