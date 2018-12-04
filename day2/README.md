## 可变参数
- 只能是参数列表中的最后一个参数
- slice的语法糖而已
```go
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintln(a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```

## 指针
- *,&用法同c
- 零值为nil
- go不支持指针运算