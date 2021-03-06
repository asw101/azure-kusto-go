// Code generated by "stringer -type Kind"; DO NOT EDIT.

package errors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KOther-0]
	_ = x[KIO-1]
	_ = x[KInternal-2]
	_ = x[KDBNotExist-3]
	_ = x[KTimeout-4]
	_ = x[KLimitsExceeded-5]
	_ = x[KClientArgs-6]
	_ = x[KHTTPError-7]
	_ = x[KBlobstore-8]
	_ = x[KLocalFileSystem-9]
}

const _Kind_name = "KOtherKIOKInternalKDBNotExistKTimeoutKLimitsExceededKClientArgsKHTTPErrorKBlobstoreKLocalFileSystem"

var _Kind_index = [...]uint8{0, 6, 9, 18, 29, 37, 52, 63, 73, 83, 99}

func (i Kind) String() string {
	if i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
