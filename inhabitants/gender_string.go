// Code generated by "stringer -type=Gender"; DO NOT EDIT.

package inhabitants

import "fmt"

const _Gender_name = "AsexualMaleFemale"

var _Gender_index = [...]uint8{0, 7, 11, 17}

func (i Gender) String() string {
	if i < 0 || i >= Gender(len(_Gender_index)-1) {
		return fmt.Sprintf("Gender(%d)", i)
	}
	return _Gender_name[_Gender_index[i]:_Gender_index[i+1]]
}