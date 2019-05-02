package utils

func ToInt32(val *int) *int32 {
	mob := int32(*val)
	return &mob
}