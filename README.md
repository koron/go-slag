# slag - struct flag

```go
type Global struct {
  Help bool
  Verbose bool
}

func userAdd(g Global, l struct { Name string }, ...string) error {
    if g.Verbose {
    }
    f, err := os.Open(l.Name)
}

func main() {
    slag.Run(cmd, os.Args)
    slag.Help(cmd)
}
```
