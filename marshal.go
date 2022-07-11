package money

import (
	"encoding/json"
)

func (m Money) MarshalJSON() ([]byte, error) {
	dto := m.ExtractDTO()

	return json.Marshal(dto)
}

func (m *Money) UnmarshalJSON(data []byte) error {
	type Money2 Money

	dto := DTO{}
	if err := json.Unmarshal(data, &dto); err != nil {
		return json.Unmarshal(data, (*Money2)(m))
	}

	var err error
	*m, err = dto.ExtractMoney()

	return err
}
