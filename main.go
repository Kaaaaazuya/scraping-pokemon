package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Pokemon struct {
	No     int    `json:"-"`  // 図鑑番号
	JPName string `json:"jp"` // 日本語名
	ENName string `json:"en"` // 英語名
}

func main() {
	// Colly collector を作成
	c := colly.NewCollector()

	// Pokemon のスライスを保存するための変数
	var pokemons []Pokemon

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			cells := el.ChildTexts("td")
			if len(cells) >= 8 {
				no, err := strconv.Atoi(strings.TrimSpace(cells[0]))
				if err != nil {
					fmt.Print("error")
					return
				}
				jpName := strings.TrimSpace(cells[1])
				enName := getFirstEnglishName(strings.TrimSpace(cells[2]))
				if jpName != "" && enName != "" {
					pokemons = append(pokemons, Pokemon{No: no, JPName: jpName, ENName: enName})
				}
			}
		})
	})

	// エラー処理
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://wiki.xn--rckteqa2e.com/wiki/%E3%83%9D%E3%82%B1%E3%83%A2%E3%83%B3%E3%81%AE%E5%A4%96%E5%9B%BD%E8%AA%9E%E5%90%8D%E4%B8%80%E8%A6%A7")
	if err != nil {
		log.Fatal(err)
	}

	// No に基づいてソート
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].No < pokemons[j].No
	})

	// json 形式に変換
	jsonOutput, err := json.MarshalIndent(pokemons, "", "  ")
	if err != nil {
		log.Fatal("Error encoding to JSON:", err)
	}

	// ファイルに json を出力
	file, err := os.Create("pokemons.json")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	file.Write(jsonOutput)
	fmt.Println("JSON data written to pokemons.json")
}

func getFirstEnglishName(raw string) string {
	parts := strings.Fields(raw) // スペースで分割
	if len(parts) > 0 {
		name := parts[0]
		return strings.Trim(name, "()") // 括弧を削除
	}
	return ""
}
