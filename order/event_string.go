// Code generated by "stringer -type=Event"; DO NOT EDIT.

package order

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Created-1]
	_ = x[Invalid-2]
	_ = x[Completed-3]
	_ = x[Fulfilled-4]
}

const _Event_name = "CreatedInvalidCompletedFulfilled"

var _Event_index = [...]uint8{0, 7, 14, 23, 32}

func (i Event) String() string {
	i -= 1
	if i < 0 || i >= Event(len(_Event_index)-1) {
		return "Event(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Event_name[_Event_index[i]:_Event_index[i+1]]
}
