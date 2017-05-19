package constant

import "strconv"

// Gender type
// --------------------------------------------------
// (ISO 5218)
type Gender struct {
	basetype
}

// MarshalJSON implements json.Marshaler
func (gender Gender) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(gender.ID), 10)), nil
}

// GetGender returns matched area.
func GetGender(id uint8) Gender {
	for _, e := range GenderList {
		if e.ID == id {
			return e
		}
	}
	return Gender{}
}

// Male is gender male.
var Male = Gender{basetype{1, "男性"}}

// Female is gender female.
var Female = Gender{basetype{2, "女性"}}

// GenderList is list of Gender type
var GenderList = []Gender{
	{basetype{0, "未設定"}},
	Male,
	Female,
}
