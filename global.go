package neovim

type Global struct {
	Vars    Vars
	Options GlobalOptions
	KeyMaps KeyMaps
}

func newGlobal(api *Api) Global {
	return Global{
		Vars:    newGlobalVars(api),
		Options: GlobalOptions{api},
		KeyMaps: newGlobalKeyMaps(api),
	}
}
