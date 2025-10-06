package bitbool

const (
	FeatureAFlag byte = 1 << iota // 00000001
	FeatureBFlag                  // 00000010
	FeatureCFlag                  // 00000100
	FeatureDFlag                  // 00001000
	FeatureEFlag                  // 00010000
	FeatureFFlag                  // 00100000
	FeatureGFlag                  // 01000000
	FeatureHFlag                  // 10000000
)

func CheckIfFeatureEnabled(features byte, flag byte) bool {
	return (features & flag) != 0
}

func RemoveFeature(features byte, flag byte) byte {
	return features & ^flag
}

func AddFeature(features byte, flag byte) byte {
	return features | flag
}
