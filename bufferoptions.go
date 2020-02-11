package neovim

import "github.com/neovim/go-client/nvim"

const (
	BufferOptionModifiable BoolOption   = "modifiable" // bool
	BufferOptionReadOnly   BoolOption   = "readonly"   // bool
	BufferOptionHidden     StringOption = "bufhidden"  // string
	BufferOptionType       StringOption = "buftype"    // string
	BufferOptionFileType   StringOption = "filetype"   // string

	BufferOptionSwapFile  BoolOption   = "swapfile"  // bool
	BufferOptionListed    BoolOption   = "buflisted" // bool
	BufferOptionList      BoolOption   = "list"      // bool? "no"
	BufferOptionSpell     BoolOption   = "spell"     // bool
	BufferOptionListchars StringOption = "listchars" // ""
)

type BufferOptions struct {
	api      *Api
	bufferID nvim.Buffer
}

// modifiable

func (o *BufferOptions) SetModifiable(value bool) {
	o.setBool(BufferOptionModifiable, value)
}

func (o *BufferOptions) Modifiable() bool {
	return o.getBool(BufferOptionModifiable)
}

// readonly

func (o *BufferOptions) SetReadOnly(value bool) {
	o.setBool(BufferOptionReadOnly, value)
}

func (o *BufferOptions) ReadOnly() bool {
	return o.getBool(BufferOptionReadOnly)
}

// bufhidden

type BufferHiddenValue string

const (
	// follow the global 'hidden' option
	BufferHiddenDefault BufferHiddenValue = ""

	// hide the buffer (don't unload it), also when 'hidden' is not set
	BufferHiddenHide BufferHiddenValue = "hide"

	// the buffer, also when 'hidden' is set or using :hide
	BufferHiddenUnload BufferHiddenValue = "unload"

	// the buffer from the buffer list, also when 'hidden' is set or using :hide, like using :bdelete
	BufferHiddenDelete BufferHiddenValue = "delete"

	// out the buffer from the buffer list, also when 'hidden' is set or using :hide, like using :bwipeout
	BufferHiddenWipe BufferHiddenValue = "wipe"
)

func (o *BufferOptions) SetHidden(value BufferHiddenValue) {
	o.setString(BufferOptionHidden, string(value))
}

func (o *BufferOptions) Hidden() string {
	return o.getString(BufferOptionHidden)
}

// buftype

type BufferTypeValue string

const (

	// normal buffer
	BufferTypeDefault BufferTypeValue = ""

	// buffer will always be written with |BufWriteCmd|s
	BufferTypeAcwrite BufferTypeValue = "acwrite"

	// help buffer (do not set this manually)
	BufferTypeHelp BufferTypeValue = "help"

	// buffer is not related to a file, will not be written
	BufferTypeNoFile BufferTypeValue = "nofile"

	// buffer will not be written
	BufferTypeNowrite BufferTypeValue = "nowrite"

	// list of errors |:cwindow| or locations |:lwindow
	BufferTypeQuickFix BufferTypeValue = "quickfix"

	// |terminal-emulator| buffer
	BufferTypeTerminal BufferTypeValue = "terminal"
)

func (o *BufferOptions) SetType(value BufferTypeValue) {
	o.setString(BufferOptionType, string(value))
}

func (o *BufferOptions) Type() string {
	return o.getString(BufferOptionType)
}

// swapfile

func (o *BufferOptions) SetSwapFile(value bool) {
	o.setBool(BufferOptionSwapFile, value)
}

func (o *BufferOptions) SwapFile() bool {
	return o.getBool(BufferOptionSwapFile)
}

// buflisted

func (o *BufferOptions) SetListed(value bool) {
	o.setBool(BufferOptionListed, value)
}

func (o *BufferOptions) SwapListed() bool {
	return o.getBool(BufferOptionListed)
}

// filetype

func (o *BufferOptions) SetFileType(value string) {
	o.setString(BufferOptionFileType, value)
}

func (o *BufferOptions) FileType() string {
	return o.getString(BufferOptionFileType)
}

////////////////////////////////////////////////////////////////////////////////

func (o *BufferOptions) getString(name StringOption) string {
	var value string
	o.api.nvim().BufferOption(o.bufferID, string(name), &value)
	return value
}

func (o *BufferOptions) setString(name StringOption, value string) {
	o.api.nvim().SetBufferOption(o.bufferID, string(name), value)
}

func (o *BufferOptions) getBool(name BoolOption) bool {
	var value bool
	o.api.nvim().BufferOption(o.bufferID, string(name), &value)
	return value
}

func (o *BufferOptions) setBool(name BoolOption, value bool) {
	o.api.nvim().SetBufferOption(o.bufferID, string(name), value)
}
