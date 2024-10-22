package repository

import (
	"time"

	"github.com/a6slab/shorter_url/internal/model"
	"github.com/a6slab/shorter_url/internal/utils"
	"github.com/gocql/gocql"
)

type ShorterURLRepository interface {
	ShortenURL(longURL, shortURL string, expiryDate time.Time, userID int) (*model.URLMapping, error)
	GetRedirectURL(shortURL string) (string, error)
}

type shorterURLRepo struct {
	session *gocql.Session
}

func NewShorterURLRepo(session *gocql.Session) ShorterURLRepository {
	return &shorterURLRepo{
		session: session,
	}
}

func (r *shorterURLRepo) ShortenURL(longURL, shortURL string, expiryDate time.Time, userID int) (*model.URLMapping, error) {
	isExist, err := r.IsShortURLExists(shortURL)
	if err != nil {
		return nil, err
	}

	if isExist {
		shortURL = utils.GenerateShortURL(longURL)
	}

	query := `INSERT INTO url_mapping(short_url, long_url, creation_date, expiration_date, user_id, click_count) VALUES(?, ?, ?, ?,?, ?)`
	if err := r.session.Query(query, shortURL, longURL, time.Now(), expiryDate, userID, 0).Exec(); err != nil {
		return &model.URLMapping{}, err
	}
	return &model.URLMapping{
		ShortURL:       shortURL,
		LongURL:        longURL,
		CreationDate:   time.Now(),
		ExpirationDate: expiryDate,
		UserID:         userID,
		ClickCount:     0,
	}, nil
}

func (r *shorterURLRepo) GetRedirectURL(shortURL string) (string, error) {
	var redirectURL string
	query := `SELECT long_url FROM url_mapping WHERE short_url = ?`
	if err := r.session.Query(query, shortURL).Scan(&redirectURL); err != nil {
		return "", err
	}
	return redirectURL, nil
}

func (r *shorterURLRepo) IsShortURLExists(shortURL string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM url_mapping WHERE short_url = ? LIMIT 1`

	if err := r.session.Query(query, shortURL).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}
