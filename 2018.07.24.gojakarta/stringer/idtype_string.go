// Code generated by "stringer -type IDType"; DO NOT EDIT.

package identity

import "strconv"

const _IDType_name = "KTPSIMPassport"

var _IDType_index = [...]uint8{0, 3, 6, 14}

func (i IDType) String() string {
	if i < 0 || i >= IDType(len(_IDType_index)-1) {
		return "IDType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IDType_name[_IDType_index[i]:_IDType_index[i+1]]
}
