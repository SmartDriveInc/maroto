package border_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SmartDriveInc/maroto/v2/pkg/consts/border"
)

func TestType_IsValid(t *testing.T) {
	t.Run("When type is empty, should not be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Type("")

		// Act & Assert
		assert.False(t, borderType.IsValid())
	})
	t.Run("When type is full, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Full

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
	t.Run("When type is left, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Left

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
	t.Run("When type is top, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Top

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
	t.Run("When type is right, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Right

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
	t.Run("When type is bottom, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Bottom

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
}

func TestBorderConfig_ToGofpdfString(t *testing.T) {
	tests := []struct {
		name   string
		config border.BorderConfig
		want   string
	}{
		{
			name:   "No borders",
			config: border.BorderConfig{},
			want:   "",
		},
		{
			name:   "Full border",
			config: border.BorderConfig{Left: true, Top: true, Right: true, Bottom: true},
			want:   "1",
		},
		{
			name:   "Left and Top",
			config: border.BorderConfig{Left: true, Top: true},
			want:   "LT",
		},
		{
			name:   "Right and Bottom",
			config: border.BorderConfig{Right: true, Bottom: true},
			want:   "RB",
		},
		{
			name:   "Only Left",
			config: border.BorderConfig{Left: true},
			want:   "L",
		},
		{
			name:   "All except Bottom",
			config: border.BorderConfig{Left: true, Top: true, Right: true},
			want:   "LTR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.config.ToGofpdfString())
		})
	}
}

func TestBorderConfig_HasBorder(t *testing.T) {
	tests := []struct {
		name   string
		config border.BorderConfig
		want   bool
	}{
		{
			name:   "No borders",
			config: border.BorderConfig{},
			want:   false,
		},
		{
			name:   "Has Left border",
			config: border.BorderConfig{Left: true},
			want:   true,
		},
		{
			name:   "Has multiple borders",
			config: border.BorderConfig{Left: true, Top: true},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.config.HasBorder())
		})
	}
}

func TestFromType(t *testing.T) {
	tests := []struct {
		name string
		t    border.Type
		want border.BorderConfig
	}{
		{
			name: "Full type",
			t:    border.Full,
			want: border.BorderConfig{Left: true, Top: true, Right: true, Bottom: true},
		},
		{
			name: "Left type",
			t:    border.Left,
			want: border.BorderConfig{Left: true},
		},
		{
			name: "None type",
			t:    border.None,
			want: border.BorderConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, border.FromType(tt.t))
		})
	}
}

func TestNewConfig(t *testing.T) {
	config := border.NewConfig(true, false, true, false)
	expected := border.BorderConfig{Left: true, Top: false, Right: true, Bottom: false}
	assert.Equal(t, expected, config)
}
