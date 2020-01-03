package mysql

import (
	"database/sql"
	"errors"
	"mrmambo.dev/snippetbox/pkg/models"
	"time"
)

type Articles struct {
	DB *sql.DB
}

func (a *Articles) Get(id int) (*models.Article, error) {
	row := a.DB.QueryRow(`SELECT id, title, slug, created, expires FROM articles where id = ?`, id)

	article := &models.Article{}
	err := row.Scan(&article.ID, &article.Title, &article.Slug, &article.Created, &article.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrorNoRecord
		} else {
			return nil, err
		}
	}

	return article, nil
}

func (a *Articles) All() ([]*models.Article, error) {

	r, err := a.DB.Query(`SELECT id, title, slug FROM articles`)
	if err != nil {
		return nil, err
	}

	var articles []*models.Article

	for r.Next() {
		article := &models.Article{}
		err = r.Scan(&article.ID, &article.Title, &article.Slug)
		articles = append(articles, article)
	}

	return articles, nil
}

func (a *Articles) Create(title, slug string) error {
	_, err := a.DB.Exec(`INSERT INTO articles (title, slug, created, expires) VALUES  (?, ?, ?, ?)`, title, slug, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}
