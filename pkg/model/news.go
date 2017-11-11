package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type News struct {
	ID       bson.ObjectId `bson:"_id"`
	Title    string
	Image    string
	Detail   string
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`
}

var newsStorage []*News

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

func CreateNews(news News) error {
	news.ID = bson.NewObjectId()
	news.CreateAt = time.Now()
	news.UpdateAt = news.CreateAt

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").Insert(&news)

	if err != nil {
		return err
	}
	return nil
}

func EditNews(news News) error {
	objectId := bson.ObjectIdHex(id)

	if !objectId.Valid() {
		return fmt.Errorf("Invalid Id")
	}
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").UpdateId(news.ID, news)

	if err != nil {
		return err
	}
	return nil
}
func ListNews() ([]*News, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var news []*News
	err := s.DB(database).C("news").Find(nil).All(news)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func GetNews(id string) (*News, error) {
	objectId := bson.ObjectIdHex(id)

	if !objectId.Valid() {
		return nil, fmt.Errorf("Invalid Id")
	}

	s := mongoSession.Copy()
	defer s.Close()
	var n News
	err := s.DB(database).C("news").Find(id).All(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func DeleteNews(id string) (*News, error) {
	objectId := bson.ObjectIdHex(id)
	if !objectId.Valid() {
		return nil, fmt.Errorf("Invalid Id")
	}

	s := mongoSession.Copy()
	defer s.Close()

	err := s.DB(database).C("news").RemoveId(objectId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
