package neovim

const (
	GlobalOperatorFunc StringOption = "operatorfunc"
	GlobalSelection    StringOption = "selection"
	GlobalClipboard    StringOption = "clipboard"
)

// WindowOptionWidth

type GlobalOptions struct {
	api *Api
}

func (o GlobalOptions) OperatorFunc() string {
	return o.getString(GlobalOperatorFunc)
}

func (o GlobalOptions) SetOperatorFunc(fn string) {
	o.setString(GlobalOperatorFunc, fn)
}

type GlobalSelectionValue string

const (
	GlobalSelectionOld       GlobalSelectionValue = "old"
	GlobalSelectionInclusive GlobalSelectionValue = "inclusive"
	GlobalSelectionExclusive GlobalSelectionValue = "exclusive"
)

func (o GlobalOptions) Selection() GlobalSelectionValue {
	return GlobalSelectionValue(o.getString(GlobalSelection))
}

func (o GlobalOptions) SetSelection(value GlobalSelectionValue) {
	o.setString(GlobalSelection, string(value))
}

type GlobalClipboardValue string

const (
	GlobalClipboardNone GlobalClipboardValue = ""
	// When included, Vim will use the clipboard register '*'
	// for all yank, delete, change and put operations which
	// would normally go to the unnamed register.  When a
	// register is explicitly specified, it will always be
	// used regardless of whether "unnamed" is in 'clipboard'
	// or not.  The clipboard register can always be
	// explicitly accessed using the "* notation.  Also see
	// |clipboard|.
	GlobalClipboardUnnamed GlobalClipboardValue = "unnamed"

	// A variant of the "unnamed" flag which uses the
	// clipboard register '+' (|quoteplus|) instead of
	// register '*' for all yank, delete, change and put
	// operations which would normally go to the unnamed
	// register.  When "unnamed" is also included to the
	// option, yank and delete operations (but not put)
	// will additionally copy the text into register
	// '*'. See |clipboard|.
	GlobalClipboardUnnamedPlus GlobalClipboardValue = "unnamedplus"
)

func (o GlobalOptions) Clipboard() GlobalClipboardValue {
	return GlobalClipboardValue(o.getString(GlobalClipboard))
}

func (o GlobalOptions) SetClipboard(value GlobalClipboardValue) {
	o.setString(GlobalClipboard, string(value))
}

// set selection=inclusive clipboard-=unnamed clipboard-=unnamedplus

////////////////////////////////////////////////////////////////////////////////

func (o *GlobalOptions) getString(name StringOption) string {
	var value string
	o.api.nvim().Option(string(name), &value)
	return value
}

func (o *GlobalOptions) setString(name StringOption, value string) {
	o.api.nvim().SetOption(string(name), value)
}

func (o *GlobalOptions) getBool(name BoolOption) bool {
	var value bool
	o.api.nvim().Option(string(name), &value)
	return value
}

func (o *GlobalOptions) setBool(name BoolOption, value bool) {
	o.api.nvim().SetOption(string(name), value)
}

func (o *GlobalOptions) getInt(name IntOption) int {
	var value int
	o.api.nvim().Option(string(name), &value)
	return value
}

func (o *GlobalOptions) setInt(name IntOption, value int) {
	o.api.nvim().SetOption(string(name), value)
}
