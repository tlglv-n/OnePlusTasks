package main

import (
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"os"
	"strings"
)

func main() {
	// Создание CSV файла
	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Создание нового CSV Writer для записи данных в файл
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Запись заголовков столбцов в CSV файл
	if err := writer.Write([]string{"Rating", "Users", "Subscribers", "Audience", "Authentic", "Engagement"}); err != nil {
		panic(err)
	}

	// Запуск Geziyor, который будет парсить данные в таблицу
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://hypeauditor.com/top-instagram-all-russia/"},
		ParseFunc: parseInstaStars(writer),
	}).Start()
}

// parseInstaStars: функция для извлечения данных с URL в CSV файл
func parseInstaStars(writer *csv.Writer) func(g *geziyor.Geziyor, r *client.Response) {
	return func(g *geziyor.Geziyor, r *client.Response) {
		r.HTMLDoc.Find("div.row__top").Each(func(i int, s *goquery.Selection) {
			if _, ok := s.Find("a.contributor").Attr("href"); ok {
				// Извлечение данных
				rating := strings.TrimSpace(s.Find("div.row-cell.rank").Text())
				users := strings.TrimSpace(s.Find("div.contributor__name-content").Text())
				subscribers := strings.TrimSpace(s.Find("div.row-cell.subscribers").Text())
				audience := strings.TrimSpace(s.Find("div.row-cell.audience").Text())
				authentic := strings.TrimSpace(s.Find("div.row-cell.authentic").Text())
				engagement := strings.TrimSpace(s.Find("div.row-cell.engagement").Text())

				// Запись данных в CSV файл
				if err := writer.Write([]string{rating, users, subscribers, audience, authentic, engagement}); err != nil {
					panic(err)
				}
			}
		})
	}
}
