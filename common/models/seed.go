package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/mvinturis/mailbox-automation/common/db"
)

const STARTING_PROFILE_PATH = "/path/to/profil_default_chrome"

type Seed struct {
	Id bson.ObjectId `bson:"_id"`

	Email        string
	Domain       string
	Password     string
	Phone        string `bson:"phone,omitempty"`
	ProxyIp      string `bson:"proxy_ip,omitempty"`
	ProfilePath  string `bson:"profile_path,omitempty"`
	RecoveryCode string `bson:"recovery_code,omitempty"`
	AuthKey      string `bson:"auth_key,omitempty"`

	Source string
	Tags   []string `bson:"tags,omitempty"`

	LocalEmail string `bson:"local_email,omitempty"`

	Runner string `bson:"runner,omitempty"`

	Status string

	AddedAt time.Time `bson:"added_at"`
}

type seedModel struct{}

var SeedModel = &seedModel{}

func (seedModel) Get(id bson.ObjectId) (s *Seed, err error) {
    db := db.NewSession()
    defer db.Session.Close()

    err = db.C("seeds").Find(bson.M{"_id": id}).One(&s)
    return
}

func (seedModel) GetByEmail(email string) (s *Seed, err error) {
    db := db.NewSession()
    defer db.Session.Close()

    err = db.C("seeds").Find(bson.M{"email": email}).One(&s)
    return
}

func (seedModel) FindOne(q bson.M) (s *Seed, err error) {
    db := db.NewSession()
    defer db.Session.Close()
    
    err = db.C("seeds").Find(q).One(&s)
    return
}

func (seedModel) Find(q bson.M) (seeds []*Seed, err error) {
    db := db.NewSession()
    defer db.Session.Close()
    
    err = db.C("seeds").Find(q).All(&seeds)
    return
}

func (seedModel) Insert(s *Seed) error {
    db := db.NewSession()
    defer db.Session.Close()
    
    s.Id = bson.NewObjectId()
    s.Status = "new"
    s.AddedAt = time.Now()
    
    return db.C("seeds").Insert(s)
}

func (s *Seed) Update(to_update bson.M) error {
    db := db.NewSession()
    defer db.Session.Close()
    
    return db.C("seeds").Update(bson.M{"_id": s.Id}, to_update)
}

func (s *Seed) UpdateStatus(status string) error {
    db := db.NewSession()
    defer db.Session.Close()
    
    return db.C("seeds").Update(bson.M{"_id": s.Id}, bson.M{"$set": bson.M{"status": status}})
}

func (s *Seed) UpdateLoginError(errmsg string) error {
    db := db.NewSession()
    defer db.Session.Close()
    
    return db.C("seeds").Update(bson.M{"_id": s.Id}, bson.M{"$set": bson.M{"login_error": errmsg}})
}

func (s *Seed) Set(to_set bson.M) error {
    db := db.NewSession()
    defer db.Session.Close()
    
    return db.C("seeds").Update(bson.M{"_id": s.Id}, bson.M{"$set": to_set})
}
