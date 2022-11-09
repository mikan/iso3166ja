iso3166ja
=========

Wikipedia 日本語版の [ISO 3166-1](https://ja.wikipedia.org/wiki/ISO_3166-1) のページを解析し、
国コードと国名の対応表を作成するプログラムです。

## 生成済データ

生成済のデータが欲しいかたはこちら:

- CSV 形式: [iso3166ja.csv](iso3166ja.csv) (UTF-8), [iso3166ja-2.csv](iso3166ja-2.csv) (Alpha-2, 国名, 日本語国名のみ)
- JSON 形式: [iso3166ja-a.json](iso3166ja-a.json) (配列), [iso3166ja-m.json](iso3166ja-m.json) (キーが Alpha-2 のマップ)

## 生成方法

### CSV

```
go run ./cmd/iso3166ja -f csv -o iso3166ja
open iso3166ja.csv
```

UTF-8 でエンコードされているのでご注意ください。
Shift_JIS 版は、元データに Shift_JIS へ変換できない文字が含まれているため提供していません。

特定のカラムのみ出力するには、 `-c` を指定してください。

```
go run ./cmd/iso3166ja -f csv -o iso3166ja-2 -c alpha2,name,name_ja
open iso3166ja-2.csv
```

### JSON (配列)

```
go run ./cmd/iso3166ja -f json-array -o iso3166ja-a
open iso3166ja-a.json
```

### JSON (キーが Alpha-2 のマップ)

```
go run ./cmd/iso3166ja -f json-map -o iso3166ja-m
open iso3166ja-m.json
```

## LICENSE

[The 3-Clause BSD License](LICENSE)

## Author

[mikan](https://github.com/mikan)
