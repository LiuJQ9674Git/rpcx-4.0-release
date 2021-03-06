package client

import (
	"fmt"
)

const _FailModeName = "FailoverFailfastFailtryFailbackup"

var _FailModeIndex = [...]uint8{0, 8, 16, 23, 33}

func (i FailMode) String() string {
	if i < 0 || i >= FailMode(len(_FailModeIndex)-1) {
		return fmt.Sprintf("FailMode(%d)", i)
	}
	return _FailModeName[_FailModeIndex[i]:_FailModeIndex[i+1]]
}

var _FailModeValues = []FailMode{0, 1, 2, 3}

var _FailModeNameToValueMap = map[string]FailMode{
	_FailModeName[0:8]:   0,
	_FailModeName[8:16]:  1,
	_FailModeName[16:23]: 2,
	_FailModeName[23:33]: 3,
}

// FailModeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FailModeString(s string) (FailMode, error) {
	if val, ok := _FailModeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FailMode values", s)
}

// FailModeValues returns all values of the enum
func FailModeValues() []FailMode {
	return _FailModeValues
}

// IsAFailMode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i FailMode) IsAFailMode() bool {
	for _, v := range _FailModeValues {
		if i == v {
			return true
		}
	}
	return false
}
