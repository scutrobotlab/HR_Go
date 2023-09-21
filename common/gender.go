package common

const (
	GenderUnknown     = 0
	GenderMale        = 1
	GenderFemale      = 2
	GenderUnspecified = 9
)

func GenderIntToString(gender int32) string {
	switch gender {
	case GenderUnknown:
		return "unknown"
	case GenderMale:
		return "male"
	case GenderFemale:
		return "female"
	case GenderUnspecified:
		return "unspecified"
	}
	return ""
}

func GenderStringToInt(gender string) int32 {
	switch gender {
	case "unknown":
		return GenderUnknown
	case "male":
		return GenderMale
	case "female":
		return GenderFemale
	case "unspecified":
		return GenderUnspecified
	}
	return GenderUnknown
}
