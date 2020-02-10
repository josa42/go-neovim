package neovim

import (
	"fmt"

	"github.com/neovim/go-client/nvim"
)

const (
	WindowOptionFixWidth       BoolOption   = "winfixwidth"    // bool
	WindowOptionNumber         BoolOption   = "number"         // bool
	WindowOptionRelativeNumber BoolOption   = "relativenumber" // bool
	WindowOptionFoldColumn     IntOption    = "foldcolumn"     // "0"
	WindowOptionFoldMethod     StringOption = "foldmethod"     // "manual"
	WindowOptionFoldEnable     BoolOption   = "foldenable"     // bool
	WindowOptionWrap           BoolOption   = "wrap"           // bool
	WindowOptionCursorLine     BoolOption   = "cursorline"     // bool
	WindowOptionCursorColumn   BoolOption   = "cursorcolumn"   // bool
	WindowOptionSignColumn     StringOption = "signcolumn"     // string
	WindowOptionList           BoolOption   = "list"           // bool
	WindowOptionWidth          IntOption    = "winwidth"       // int
	WindowOptionColorColumn    StringOption = "colorcolumn"    // string
	WindowOptionSpell          BoolOption   = "spell"          // bool
)

// WindowOptionWidth

type WindowOptions struct {
	api      *Api
	windowID nvim.Window
}

func (o *WindowOptions) Number() bool {
	return o.getBool(WindowOptionNumber)
}

func (o *WindowOptions) SetNumber(value bool) {
	o.setBool(WindowOptionNumber, value)
}

func (o *WindowOptions) RelativeNumber() bool {
	return o.getBool(WindowOptionRelativeNumber)
}

func (o *WindowOptions) SetRelativeNumber(value bool) {
	o.setBool(WindowOptionRelativeNumber, value)
}

func (o *WindowOptions) Width() int {
	return o.getInt(WindowOptionWidth)
}

func (o *WindowOptions) SetWidth(value int) {
	o.setInt(WindowOptionWidth, value)
}

func (o *WindowOptions) ColorColumn() string {
	return o.getString(WindowOptionColorColumn)
}

func (o *WindowOptions) SetCursorLine(value bool) {
	o.setBool(WindowOptionCursorLine, value)
}

func (o *WindowOptions) CursorLine() bool {
	return o.getBool(WindowOptionCursorLine)
}

func (o *WindowOptions) SetCursorColumn(value bool) {
	o.setBool(WindowOptionCursorColumn, value)
}

func (o *WindowOptions) CursorColumn() bool {
	return o.getBool(WindowOptionCursorColumn)
}

type WindowSignColumnValue string

const (
	WindowSignColumnAuto WindowSignColumnValue = "auto"
	WindowSignColumnNo   WindowSignColumnValue = "no"
	WindowSignColumnYes  WindowSignColumnValue = "yes"
)

// When and how to draw the signcolumn.
func (o *WindowOptions) SignColumn() WindowSignColumnValue {
	return WindowSignColumnValue(o.getString(WindowOptionSignColumn))
}

// When and how to draw the signcolumn.
func (o *WindowOptions) SetSignColumn(value WindowSignColumnValue, number ...int) {

	v := string(value)
	if value != WindowSignColumnNo && len(number) > 0 {
		v = fmt.Sprintf("%s:%d", v, number[0])
	}

	o.setString(WindowOptionSignColumn, v)
}

func (o *WindowOptions) SetFixWidth(value bool) {
	o.setBool(WindowOptionFixWidth, value)
}

func (o *WindowOptions) FixWidth() bool {
	return o.getBool(WindowOptionFixWidth)
}

func (o *WindowOptions) Spell() bool {
	return o.getBool(WindowOptionSpell)
}

func (o *WindowOptions) SetSpell(value bool) {
	o.setBool(WindowOptionSpell, value)
}

// foldcolumn

func (o *WindowOptions) FoldColumn() int {
	return o.getInt(WindowOptionFoldColumn)
}

func (o *WindowOptions) SetFoldColumn(value int) {
	if value > 12 {
		value = 12
	}
	o.setInt(WindowOptionFoldColumn, value)
}

// foldmethod

type WindowFoldMethodValue string

const (
	// Folds are created manually.
	WindowFoldMethodManual WindowFoldMethodValue = "manual"

	// Lines with equal indent form a fold.
	WindowFoldMethodIndent WindowFoldMethodValue = "indent"

	// 'foldexpr' gives the fold level of a line.
	WindowFoldMethodExpr WindowFoldMethodValue = "expr"

	// Markers are used to specify folds.
	WindowFoldMethodMarker WindowFoldMethodValue = "marker"

	// Syntax highlighting items specify folds.
	WindowFoldMethodSyntax WindowFoldMethodValue = "syntax"

	// Fold text that is not changed.
	WindowFoldMethodDiff WindowFoldMethodValue = "diff"
)

func (o *WindowOptions) FoldMethod() WindowFoldMethodValue {
	return WindowFoldMethodValue(o.getString(WindowOptionFoldMethod))
}

func (o *WindowOptions) SetFoldMethod(value WindowFoldMethodValue) {
	o.setString(WindowOptionFoldMethod, string(value))
}

func (o *WindowOptions) FoldEnable() bool {
	return o.getBool(WindowOptionFoldEnable)
}

func (o *WindowOptions) SetFoldEnable(value bool) {
	o.setBool(WindowOptionFoldEnable, value)
}

func (o *WindowOptions) List() bool {
	return o.getBool(WindowOptionList)
}

func (o *WindowOptions) SetList(value bool) {
	o.setBool(WindowOptionList, value)
}

func (o *WindowOptions) Wrap() bool {
	return o.getBool(WindowOptionWrap)
}

func (o *WindowOptions) SetWrap(value bool) {
	o.setBool(WindowOptionWrap, value)
}

////////////////////////////////////////////////////////////////////////////////

func (o *WindowOptions) getString(name StringOption) string {
	var value string
	o.api.nvim().WindowOption(o.windowID, string(name), &value)
	return value
}

func (o *WindowOptions) setString(name StringOption, value string) {
	o.api.nvim().SetWindowOption(o.windowID, string(name), value)
}

func (o *WindowOptions) getBool(name BoolOption) bool {
	var value bool
	o.api.nvim().WindowOption(o.windowID, string(name), &value)
	return value
}

func (o *WindowOptions) setBool(name BoolOption, value bool) {
	o.api.nvim().SetWindowOption(o.windowID, string(name), value)
}

func (o *WindowOptions) getInt(name IntOption) int {
	var value int
	o.api.nvim().WindowOption(o.windowID, string(name), &value)
	return value
}

func (o *WindowOptions) setInt(name IntOption, value int) {
	o.api.nvim().SetWindowOption(o.windowID, string(name), value)
}
