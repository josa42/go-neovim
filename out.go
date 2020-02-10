package neovim

import (
	"fmt"
	"strings"
)

type Out struct {
	api *Api
}

func (o *Out) Print(str string) {
	o.command("echo", escape(str))
}

func (o *Out) Printf(format string, args ...interface{}) {
	o.command("echo", escape(fmt.Sprintf(format, args...)))
}

func (o *Out) Message(str string) {
	o.command("echomsg", escape(str))
}

func (o *Out) Messagef(format string, args ...interface{}) {
	o.command("echomsg", escape(fmt.Sprintf(format, args...)))
}

func (o *Out) Error(str string) {
	o.command("echoerr", escape(str))
}

func (o *Out) Errorf(format string, args ...interface{}) {
	o.command("echoerr", escape(fmt.Sprintf(format, args...)))
}

func (o *Out) command(cmd, str string) {
	if o.api.p.Nvim == nil {
		return
	}
	o.api.p.Nvim.Command(fmt.Sprintf(`%s "%s"`, cmd, escape(str)))
}

func escape(str string) string {
	return strings.ReplaceAll(str, `"`, `\"`)
}

