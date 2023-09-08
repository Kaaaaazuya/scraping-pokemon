# scraping-pokemon

このリポジトリは、ウェブサイトからポケモンの日本語と英語の名前をスクレイピングして、JSON 形式で出力する Go 言語のプロジェクトです。

## 使用方法

まず、このリポジトリをクローンまたはダウンロードします。

```bash
git clone https://github.com/Kaaaaazuya/scraping-pokemon.git
cd scraping-pokemon
```

## 必要なパッケージをインストールします。

```bash
go get -u github.com/gocolly/colly/v2
```

## スクレイピングを実行し、結果を pokemons.json に出力します。

```bash
go run .
```
