# OST Finder

[MO2](https://www.youtube.com/@mo2bluearchive)が作成した YouTube の[プレイリスト](https://www.youtube.com/playlist?list=PLh6Ws4Fpphfqr7VL72Q6HK5Ole9YI54hv)から、指定されたワードを含む動画を検索します。

## 使い方

### 動画の検索

検索ワードを Bot に対してメンションすることで検索できます。

```txt
@ost-finder #10
```

**NOTE**\
`#`の後に英数字を続けることで、OST に付与されている番号で検索することができます。\
**例**：`@ost-finder　#10`

実装上の都合により、`--`に続く文字列は、Bot へのコマンドとして認識されます。\
もし、検索クエリに`--`を含めたい場合は末尾に半角スペースの後`--disable-option`もしくは`--d`をつけてください。

### ヘルプの表示

```
@ost-finder --help
```
