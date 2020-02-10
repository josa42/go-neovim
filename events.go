package neovim

const (
	//////////////////////////////////////////////////////////////////////////////
	// Reading

	// starting to edit a file that doesn't exist
	EventBufNewFile = "BufNewFile"

	// starting to edit a new buffer, before reading the file
	EventBufReadPre = "BufReadPre"

	// starting to edit a new buffer, after reading the file
	EventBufRead = "BufRead"

	// starting to edit a new buffer, after reading the file
	EventBufReadPost = "BufReadPost"

	// before starting to edit a new buffer |Cmd-event|
	EventBufReadCmd = "BufReadCmd"

	// before reading a file with a ":read" command
	EventFileReadPre = "FileReadPre"

	// after reading a file with a ":read" command
	EventFileReadPost = "FileReadPost"

	// before reading a file with a ":read" command |Cmd-event|
	EventFileReadCmd = "FileReadCmd"

	// before reading a file from a filter command
	EventFilterReadPre = "FilterReadPre"

	// after reading a file from a filter command
	EventFilterReadPost = "FilterReadPost"

	// before reading from stdin into the buffer
	EventStdinReadPre = "StdinReadPre"

	// After reading from the stdin into the buffer
	EventStdinReadPost = "StdinReadPost"

	//////////////////////////////////////////////////////////////////////////////
	// Writing

	// starting to write the whole buffer to a file
	EventBufWrite = "BufWrite"

	// starting to write the whole buffer to a file
	EventBufWritePre = "BufWritePre"

	// after writing the whole buffer to a file
	EventBufWritePost = "BufWritePost"

	// before writing the whole buffer to a file |Cmd-event|
	EventBufWriteCmd = "BufWriteCmd"

	// starting to write part of a buffer to a file
	EventFileWritePre = "FileWritePre"

	// after writing part of a buffer to a file
	EventFileWritePost = "FileWritePost"

	// before writing part of a buffer to a file |Cmd-event|
	EventFileWriteCmd = "FileWriteCmd"

	// starting to append to a file
	EventFileAppendPre = "FileAppendPre"

	// after appending to a file
	EventFileAppendPost = "FileAppendPost"

	// before appending to a file |Cmd-event|
	EventFileAppendCmd = "FileAppendCmd"

	// starting to write a file for a filter command or diff
	EventFilterWritePre = "FilterWritePre"

	// after writing a file for a filter command or diff
	EventFilterWritePost = "FilterWritePost"

	//////////////////////////////////////////////////////////////////////////////
	// Buffers

	// just after adding a buffer to the buffer list
	EventBufAdd = "BufAdd"

	// just after adding a buffer to the buffer list
	EventBufCreate = "BufCreate"

	// before deleting a buffer from the buffer list
	EventBufDelete = "BufDelete"

	// before completely deleting a buffer
	EventBufWipeout = "BufWipeout"

	// before changing the name of the current buffer
	EventBufFilePre = "BufFilePre"

	// after changing the name of the current buffer
	EventBufFilePost = "BufFilePost"

	// after entering a buffer
	EventBufEnter = "BufEnter"

	// before leaving to another buffer
	EventBufLeave = "BufLeave"

	// after a buffer is displayed in a window
	EventBufWinEnter = "BufWinEnter"

	// before a buffer is removed from a window
	EventBufWinLeave = "BufWinLeave"

	// before unloading a buffer
	EventBufUnload = "BufUnload"

	// just after a buffer has become hidden
	EventBufHidden = "BufHidden"

	// just after creating a new buffer
	EventBufNew = "BufNew"

	// detected an existing swap file
	EventSwapExists = "SwapExists"

	// starting a terminal job
	EventTermOpen = "TermOpen"

	// entering Terminal-mode
	EventTermEnter = "TermEnter"

	// leaving Terminal-mode
	EventTermLeave = "TermLeave"

	// stopping a terminal job
	EventTermClose = "TermClose"

	// after a channel opened
	EventChanOpen = "ChanOpen"

	// after a channel has its state changed
	EventChanInfo = "ChanInfo"

	// Options
	// when the 'filetype' option has been set
	EventFileType = "FileType"

	// when the 'syntax' option has been set
	EventSyntax = "Syntax"

	// after setting any option
	EventOptionSet = "OptionSet"

	// Startup and exit
	// after doing all the startup stuff
	EventVimEnter = "VimEnter"

	// after a UI attaches
	EventUIEnter = "UIEnter"

	// after a UI detaches
	EventUILeave = "UILeave"

	// after the terminal response to t_RV is received
	EventTermResponse = "TermResponse"

	// when using `:quit`, before deciding whether to exit
	EventQuitPre = "QuitPre"

	// when using a command that may make Vim exit
	EventExitPre = "ExitPre"

	// before exiting Nvim, before writing the shada file
	EventVimLeavePre = "VimLeavePre"

	// before exiting Nvim, after writing the shada file
	EventVimLeave = "VimLeave"

	// after Nvim is resumed
	EventVimResume = "VimResume"

	// before Nvim is suspended
	EventVimSuspend = "VimSuspend"

	//////////////////////////////////////////////////////////////////////////////
	// Various

	// after diffs have been updated
	EventDiffUpdated = "DiffUpdated"

	// after the |current-directory| was changed
	EventDirChanged = "DirChanged"

	// Vim notices that a file changed since editing started
	EventFileChangedShell = "FileChangedShell"

	// after handling a file changed since editing started
	EventFileChangedShellPost = "FileChangedShellPost"

	// before making the first change to a read-only file
	EventFileChangedRO = "FileChangedRO"

	// after executing a shell command
	EventShellCmdPost = "ShellCmdPost"

	// after filtering with a shell command
	EventShellFilterPost = "ShellFilterPost"

	// a user command is used but it isn't defined
	EventCmdUndefined = "CmdUndefined"

	// a user function is used but it isn't defined
	EventFuncUndefined = "FuncUndefined"

	// a spell file is used but it can't be found
	EventSpellFileMissing = "SpellFileMissing"

	// before sourcing a Vim script
	EventSourcePre = "SourcePre"

	// after sourcing a Vim script
	EventSourcePost = "SourcePost"

	// before sourcing a Vim script |Cmd-event|
	EventSourceCmd = "SourceCmd"

	// after the Vim window size changed
	EventVimResized = "VimResized"

	// Nvim got focus
	EventFocusGained = "FocusGained"

	// Nvim lost focus
	EventFocusLost = "FocusLost"

	// the user doesn't press a key for a while
	EventCursorHold = "CursorHold"

	// the user doesn't press a key for a while in Insert mode
	EventCursorHoldI = "CursorHoldI"

	// the cursor was moved in Normal mode
	EventCursorMoved = "CursorMoved"

	// the cursor was moved in Insert mode
	EventCursorMovedI = "CursorMovedI"

	// after creating a new window
	EventWinNew = "WinNew"

	// after entering another window
	EventWinEnter = "WinEnter"

	// before leaving a window
	EventWinLeave = "WinLeave"

	// after entering another tab page
	EventTabEnter = "TabEnter"

	// before leaving a tab page
	EventTabLeave = "TabLeave"

	// when creating a new tab page
	EventTabNew = "TabNew"

	// after entering a new tab page
	EventTabNewEntered = "TabNewEntered"

	// after closing a tab page
	EventTabClosed = "TabClosed"

	// after a change was made to the command-line text
	EventCmdlineChanged = "CmdlineChanged"

	// after entering cmdline mode
	EventCmdlineEnter = "CmdlineEnter"

	// before leaving cmdline mode
	EventCmdlineLeave = "CmdlineLeave"

	// after entering the command-line window
	EventCmdwinEnter = "CmdwinEnter"

	// before leaving the command-line window
	EventCmdwinLeave = "CmdwinLeave"

	// starting Insert mode
	EventInsertEnter = "InsertEnter"

	// when typing <Insert> while in Insert or Replace mode
	EventInsertChange = "InsertChange"

	// when leaving Insert mode
	EventInsertLeave = "InsertLeave"

	// when a character was typed in Insert mode, before
	// inserting it
	EventInsertCharPre = "InsertCharPre"

	// when some text is yanked or deleted
	EventTextYankPost = "TextYankPost"

	// after a change was made to the text in Normal mode
	EventTextChanged = "TextChanged"

	// after a change was made to the text in Insert mode
	// when popup menu is not visible
	EventTextChangedI = "TextChangedI"

	// after a change was made to the text in Insert mode
	// when popup menu visible
	EventTextChangedP = "TextChangedP"

	// before loading a color scheme
	EventColorSchemePre = "ColorSchemePre"

	// after loading a color scheme
	EventColorScheme = "ColorScheme"

	// a reply from a server Vim was received
	EventRemoteReply = "RemoteReply"

	// before a quickfix command is run
	EventQuickFixCmdPre = "QuickFixCmdPre"

	// after a quickfix command is run
	EventQuickFixCmdPost = "QuickFixCmdPost"

	// after loading a session file
	EventSessionLoadPost = "SessionLoadPost"

	// just before showing the popup menu
	EventMenuPopup = "MenuPopup"

	// after popup menu changed, not fired on popup menu hide
	EventCompleteChanged = "CompleteChanged"

	// after Insert mode completion is done
	EventCompleteDone = "CompleteDone"

	// to be used in combination with ":doautocmd"
	EventUser = "User"

	// after Nvim receives a signal
	EventSignal = "Signal"
)

