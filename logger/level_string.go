// Code generated by "stringer -type=Level -linecomment"; DO NOT EDIT.

package logger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DebugLevel-0]
	_ = x[InfoLevel-1]
	_ = x[WarnLevel-2]
	_ = x[ErrorLevel-3]
	_ = x[FatalLevel-4]
	_ = x[PanicLevel-5]
	_ = x[NoLevel-6]
	_ = x[Disabled-7]
	_ = x[TraceLevel - -1]
}

const _Level_name = "tracedebuginfowarnerrorfatalpanicnodisabled"

var _Level_index = [...]uint8{0, 5, 10, 14, 18, 23, 28, 33, 35, 43}

func (i Level) String() string {
	i -= -1
	if i < 0 || i >= Level(len(_Level_index)-1) {
		return "Level(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _Level_name[_Level_index[i]:_Level_index[i+1]]
}