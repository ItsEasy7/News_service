package news

import (
	"Gogogo/internal/utils"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func GetAllNews() ([]News, error) {
	rows, err := utils.DB.News.Query(`
	SELECT ar.id as ID, ar.title as Title, ar.file_path as Content, im.file_path as Image
	FROM articles ar
	LEFT JOIN images im ON ar.id = im.article_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println(rows)
	var newsList []News
	for rows.Next() {
		var news News
		if err := rows.Scan(&news.ID, &news.Title, &news.Content, &news.Image); err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}
	fmt.Println(newsList)
	return newsList, nil
}

func AddNews(news News) error {
	_, err := utils.DB.News.Exec("INSERT INTO news (title, content, image) VALUES ($1, $2, $3)", news.Title, news.Content, news.Image)
	return err
}

func UpdateNewsByID(news News) error {
	result, err := utils.DB.News.Exec("UPDATE news SET title=$1, content=$2, image=$3 WHERE id=$4", news.Title, news.Content, news.Image, news.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("новость не найдена")
	}
	return nil
}

func DeleteNewsByID(id int) error {
	result, err := utils.DB.News.Exec("DELETE FROM news WHERE id=$1", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("новость не найдена")
	}
	return nil
}
