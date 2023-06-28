package logger

type Format uint8

//go:generate stringer -type=Format -linecomment
const (
	Color Format = iota // color
	Text                // text
	Json                // json
)

func ParseFormat(fmt string) Format {
	switch fmt {
	case Text.String():
		return Text
	case Json.String():
		return Json
	default:
		return Color
	}
}
