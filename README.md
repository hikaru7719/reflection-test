# reflection-test
reflect packageのテスト

## func Printf(format string, a ...interface{}) (n int, err error)
https://play.golang.org/p/TbPs9hGNCuM
このコードが動くのではないのか？
interfaceの可変長引数やしいけるやろ。
しかしながら
（そもそもgovetで怒られます。）
%!d(string=abc)という実行結果に。
しかしながらメソッドのシグネチャはintreface型の可変長引数
なぜ、型情報が復元されているのか...
謎は深まる。

## Reflect
1. Reflectionは、interface値からreflectionオブジェクトに変換できる。

reclectはinterface型から型情報を復元できる。
reflectパッケージにはTypeとValueがある。
reflect.TypeOfとreflect.ValueOfを呼び出すことでinterface値からreflect.Typeとreflect.Valueを取り出します。

reflect.TypeOf関数もValueOf関数もメソッドの引数はinterface型
TypeとValueの両方がKindメソッドをもっており、Uint、Float64、Slice、等を示す定数を返します


reflect.Intとかがある。
