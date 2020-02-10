package neovim

import (
	"log"
	"os"
)

func SetupLogging() func() {
	f, err := os.OpenFile("/tmp/vim-tree.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)

	return func() { f.Close() }
}
