package model

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

// News type
type News struct {
	ID        string
	Title     string
	Image     string
	Detail    string
	CreateAt  time.Time
	UpdatedAt time.Time
}

var (
	//newsStorage []*News
	newsStorage []News
	mutexNews   sync.RWMutex
)

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

// CreateNews ..
func CreateNews(news News) {
	news.ID = generateID()
	news.CreateAt = time.Now()
	news.UpdatedAt = news.CreateAt
	mutexNews.Lock()
	defer mutexNews.Unlock()
	newsStorage = append(newsStorage, news)
}

// ListNews ...
func ListNews() []*News {
	mutexNews.RLock()
	defer mutexNews.RUnlock()
	r := make([]*News, len(newsStorage))
	for i := range newsStorage {
		n := newsStorage[i]
		r[i] = &n
	}
	return r
}

// GetNews ...
func GetNews(id string) *News {
	mutexNews.Lock()
	defer mutexNews.Unlock()
	for _, news := range newsStorage {
		if news.ID == id {
			n := news
			return &n
		}
	}
	return nil
}

// DeleteNews a a
func DeleteNews(id string) {
	mutexNews.Lock()
	defer mutexNews.Unlock()
	for i, news := range newsStorage {
		if news.ID == id {
			newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
		}
	}
}
