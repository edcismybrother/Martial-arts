package code

import "encoding/json"

// Unmarshal unmarshal,byte[0] is type of msg from,there has four type: universe,planet,world and person
// in currently
func Unmarshal(data []byte) (*PersonInfo, error) {
	if len(data) < 1 {
		return nil, ErrWRONGFORMAT
	}
	if data[0] < 1 || data[0] > 4 {
		return nil, ErrTYPENOTEXSITED
	}
	info := &PersonInfo{}
	err := json.Unmarshal(data[1:], info)
	return info, err
}
