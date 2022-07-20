package color

// Enum represents a color used by the game.
//
// This captures the valid general colors that the game data may set. These
// color constants should correlate to a similar color when the a driver
// renders them, but no guarantees about matching the value the color reports
// should be expected (e.g. a terminal driver may map down to 16 colors).
type Enum uint8

const (
	Black Enum = iota
	DarkGray
	Gray
	BrightGray
	White
	BrightWhite
	DarkRed
	Red
	BrightRed
	DarkGreen
	Green
	BrightGreen
	DarkBlue
	Blue
	BrightBlue
	DarkYellow
	Yellow
	BrightYellow
	DarkPurple
	Purple
	BrightPurple
	DarkCyan
	Cyan
	BrightCyan
	DarkBrown
	Brown
	BrightBrown
	DarkPink
	Pink
	BrightPink
	DarkOrange
	Orange
	BrightOrange

	count
)

// Value returns the general color value set for the color.
func (c Enum) Value() uint32 {
	return values[c]
}

// Name returns the name of the color.
func (c Enum) Name() string {
	return names[c]
}

// SetValue sets the value the color reports.
func (c Enum) SetValue(v uint32) {
	values[c] = 0xFFFFFF & v
}

// Reset resets the value of the color to the default.
func (c Enum) Reset() {
	values[c] = defaults[c]
}

// ResetAll resets all colors to their default value.
func ResetAll() {
	for c := Enum(0); c < count; c++ {
		values[c] = defaults[c]
	}
}

var (
	names = [count]string{
		"black", "dark gray", "gray", "bright gray", "white", "bright white",
		"dark red", "red", "bright red", "dark green", "green", "bright green",
		"dark blue", "blue", "bright blue", "dark yellow", "yellow", "bright yellow",
		"dark purple", "purple", "bright purple", "dark cyan", "cyan", "bright cyan",
		"dark brown", "brown", "bright brown", "dark pink", "pink", "bright pink",
		"dark orange", "orange", "bright orange",
	}
	values = [count]uint32{
		0x000000, 0x333333, 0x666666, 0x999999, 0xCCCCCC, 0xFFFFFF,
		0x220000, 0x880000, 0xff0000, 0x002200, 0x008800, 0x00FF00,
		0x000022, 0x000088, 0x0000ff, 0x222200, 0x888800, 0xFFFF00,
		0x220022, 0x880088, 0xff00ff, 0x002222, 0x008888, 0x00FFFF,
		0x221100, 0x551F00, 0x883300, 0x8E1592, 0xCA56CB, 0xFF88FF,
		0xCC3311, 0xCD5421, 0xFF7711,
	}
	defaults = [count]uint32{
		0x000000, 0x333333, 0x666666, 0x999999, 0xCCCCCC, 0xFFFFFF,
		0x220000, 0x880000, 0xff0000, 0x002200, 0x008800, 0x00FF00,
		0x000022, 0x000088, 0x0000ff, 0x222200, 0x888800, 0xFFFF00,
		0x220022, 0x880088, 0xff00ff, 0x002222, 0x008888, 0x00FFFF,
		0x221100, 0x551F00, 0x883300, 0x8E1592, 0xCA56CB, 0xFF88FF,
		0xCC3311, 0xCD5421, 0xFF7711,
	}
)
