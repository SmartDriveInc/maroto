// Package border contains all border types.
package border

// Type represents a border type.
type Type string

const (
	// None is the default border type.
	None Type = ""
	// Full is a border type that borders all sides.
	Full Type = "1"
	// Left is a border type that borders the left side.
	Left Type = "L"
	// Top is a border type that borders the top side.
	Top Type = "T"
	// Right is a border type that borders the right side.
	Right Type = "R"
	// Bottom is a border type that borders the bottom side.
	Bottom Type = "B"
)

// IsValid checks if the border type is valid.
func (t Type) IsValid() bool {
	return t == Full || t == Left || t == Top || t == Right || t == Bottom
}

// BorderConfig represents a border configuration with individual side controls.
type BorderConfig struct {
	Left   bool
	Top    bool
	Right  bool
	Bottom bool
}

// ToGofpdfString converts BorderConfig to gofpdf border string format.
func (b BorderConfig) ToGofpdfString() string {
	if !b.Left && !b.Top && !b.Right && !b.Bottom {
		return ""
	}

	if b.Left && b.Top && b.Right && b.Bottom {
		return "1" // Full border
	}

	var result string
	if b.Left {
		result += "L"
	}
	if b.Top {
		result += "T"
	}
	if b.Right {
		result += "R"
	}
	if b.Bottom {
		result += "B"
	}

	return result
}

// HasBorder returns true if any border side is enabled.
func (b BorderConfig) HasBorder() bool {
	return b.Left || b.Top || b.Right || b.Bottom
}

// FromType creates a BorderConfig from the legacy Type for backward compatibility.
func FromType(t Type) BorderConfig {
	switch t {
	case Full:
		return BorderConfig{Left: true, Top: true, Right: true, Bottom: true}
	case Left:
		return BorderConfig{Left: true}
	case Top:
		return BorderConfig{Top: true}
	case Right:
		return BorderConfig{Right: true}
	case Bottom:
		return BorderConfig{Bottom: true}
	default:
		return BorderConfig{}
	}
}

// NewConfig creates a BorderConfig with specified sides enabled.
func NewConfig(left, top, right, bottom bool) BorderConfig {
	return BorderConfig{
		Left:   left,
		Top:    top,
		Right:  right,
		Bottom: bottom,
	}
}
