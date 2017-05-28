# ch07/ex07

## デフォルト値 `20.0` が `°` を含んでいない理由

`Celsius` は `float64` として定義されているので、`°` を含むことはできない。

```go
// gopl.io/ch7/tempconv/tempconv.go
type Celsius float64
```

## ヘルプメッセージが `°` を含んでいる理由

ヘルプメッセージの出力に使われる関数 `PrintDefaults` は、デフォルト値を文字列に変換して表示する。

```go
// flag/flag.go
func (f *FlagSet) PrintDefaults() {
    ...
    s += fmt.Sprintf(" (default %v)", flag.DefValue)
    ...
}
```

`Celsius` の `String()` の実装が、`float64` 値の末尾に `°C` を挿入した文字列を返すようになっているので、ヘルプメッセージは `°` を含むことになる。

```go
// gopl.io/ch7/tempconv/tempconv.go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```
