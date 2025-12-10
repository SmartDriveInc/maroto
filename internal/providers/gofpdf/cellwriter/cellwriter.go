package cellwriter

import (
	"github.com/SmartDriveInc/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/SmartDriveInc/maroto/v2/pkg/consts/border"
	"github.com/SmartDriveInc/maroto/v2/pkg/core/entity"
	"github.com/SmartDriveInc/maroto/v2/pkg/props"
)

type CellWriter interface {
	SetNext(next CellWriter)
	GetNext() CellWriter
	GetName() string
	Apply(width, height float64, config *entity.Config, prop *props.Cell)
}

type cellWriter struct {
	stylerTemplate
	defaultColor *props.Color
}

func NewCellWriter(fpdf gofpdfwrapper.Fpdf) *cellWriter {
	return &cellWriter{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "cellWriter",
		},
		defaultColor: &props.BlackColor,
	}
}

func (c *cellWriter) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	if prop == nil {
		bd := border.None
		if config.Debug {
			bd = border.Full
		}

		c.fpdf.CellFormat(width, height, "", string(bd), 0, "C", false, 0, "")
		return
	}

	// Determine border string - BorderConfig takes precedence
	var borderStr string
	if config.Debug {
		borderStr = string(border.Full)
	} else if prop.BorderConfig != nil {
		borderStr = prop.BorderConfig.ToGofpdfString()
	} else {
		borderStr = string(prop.BorderType)
	}

	c.fpdf.CellFormat(width, height, "", borderStr, 0, "C", prop.BackgroundColor != nil, 0, "")
}
