# neovim

**Experimental** neovim SDK for go.

---

## Examples

```go
func main() {
  neovim.Register(&plugin{})
}

func plugin struct {}

func (p *plugin) Register(api novim.RegisterApi) {
  api.Function("Hello", func() {
    api.Out.Messagef("Hallo %s!", "Welt")
  })
}
```

## License

[MIT © Josa Gesell](LICENSE)

