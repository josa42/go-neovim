package neovim

type Mode string

const (
	ModeAll Mode = ""

	// n	Normal
	ModeNormal Mode = "n"

	// no	Operator-pending
	// nov	Operator-pending (forced characterwise |o_v|)
	// noV	Operator-pending (forced linewise |o_V|)
	// noCTRL-V Operator-pending (forced blockwise |o_CTRL-V|)
	// niI	Normal using |i_CTRL-O| in |Insert-mode|
	// niR	Normal using |i_CTRL-O| in |Replace-mode|
	// niV	Normal using |i_CTRL-O| in |Virtual-Replace-mode|

	// v	Visual by character
	ModeVisual Mode = "v"

	// V	Visual by line
	ModeVisualLine Mode = "V"

	// CTRL-V   Visual blockwise
	// s	Select by character
	// S	Select by line
	// CTRL-S   Select blockwise

	// i	Insert
	ModeInsert Mode = "i"

	// ic	Insert mode completion |compl-generic|
	// ix	Insert mode |i_CTRL-X| completion

	// R	Replace |R|
	ModeReplace Mode = "R"

	// Rc	Replace mode completion |compl-generic|
	// Rv	Virtual Replace |gR|
	// Rx	Replace mode |i_CTRL-X| completion
	// c	Command-line editing
	// cv	Vim Ex mode |gQ|
	// ce	Normal Ex mode |Q|
	// r	Hit-enter prompt
	// rm	The -- more -- prompt
	// r?	|:confirm| query of some sort
	// !	Shell or external command is executing

	// t	Terminal mode: keys go to the job
	ModeTerminal Mode = "t"
)
