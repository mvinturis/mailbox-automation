package models

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/mvinturis/mailbox-automation/common/db"
)

const TASK_ERRORED = "errored"
const TASK_PENDING = "pending"
const TASK_RUNNING = "running"
const TASK_FINISHED = "finished"
const TASK_TIMEOUT = "timeout"

type Task struct {
	Id        bson.ObjectId `bson:"_id"`
	Action    string        // seeds/trigger
	Seed      string        `bson:"seed,omitempty"`
	Behaviour string        `bson:"behaviour,omitempty"`
	Status    string        // new, queued, running, finished, errored
	Error     string
	AddedAt   time.Time `bson:"added_at"`

	Runner string `bson:"runner,omitempty"`

	Params *TaskParams `bson:"params,omitempty"`
}

type TaskParams struct {
	MsgsCount  int    `bson:"msgs_count,omitempty" json:"msgs_count,omitempty"`
	Folder     string `bson:"folder,omitempty" json:"folder,omitempty"`
	SrcFolder  string `bson:"source_folder,omitempty" json:"source_folder,omitempty"`
	DestFolder string `bson:"dest_folder,omitempty" json:"dest_folder,omitempty"`

	FromAddresses        []string `bson:"from_address,omitempty" json:"from_address,omitempty"`
	FromDomains          []string `bson:"from_domain,omitempty" json:"from_domain,omitempty"`
	FriendlyFromContains string   `bson:"friendly_from_contains,omitempty" json:"friendly_from_contains,omitempty"`
	SubjectContains      string   `bson:"subject_contains,omitempty" json:"subject_contains,omitempty"`

	Keyword string `bson:"keyword" json:"keyword"`

	SendToRcpt  string `bson:"send_to_rcpt,omitempty" json:"send_to_rcpt,omitempty"`
	SendSubject string `bson:"send_subject,omitempty" json:"send_subject,omitempty"`
	SendBody    string `bson:"send_body,omitempty" json:"send_body,omitempty"`
}

type taskModel struct{}

var TaskModel = &taskModel{}

func (taskModel) Get(id bson.ObjectId) (s *Task, err error) {
	db := db.NewSession()
	defer db.Session.Close()

	err = db.C("tasks").Find(bson.M{"_id": id}).One(&s)
	return
}

func (taskModel) FindOne(q bson.M) (s *Task, err error) {
	db := db.NewSession()
	defer db.Session.Close()

	err = db.C("tasks").Find(q).One(&s)
	return
}

func (taskModel) Find(q bson.M) (seeds []*Task, err error) {
	db := db.NewSession()
	defer db.Session.Close()

	err = db.C("tasks").Find(q).All(&seeds)
	return
}

func (taskModel) Insert(s *Task) (bson.ObjectId, error) {
	db := db.NewSession()
	defer db.Session.Close()

	s.Id = bson.NewObjectId()
	s.Status = "new"
	s.AddedAt = time.Now()

	err := db.C("tasks").Insert(s)
	return s.Id, err
}

func (s *Task) Update(to_update bson.M) error {
	db := db.NewSession()
	defer db.Session.Close()

	return db.C("tasks").Update(bson.M{"_id": s.Id}, to_update)
}

func (s *Task) UpdateStatus(status string) error {
	db := db.NewSession()
	defer db.Session.Close()

	return db.C("tasks").Update(bson.M{"_id": s.Id}, bson.M{"$set": bson.M{"status": status}})
}

func (tsk *Task) Errored(err error) error {
	db := db.NewSession()
	defer db.Session.Close()

	return db.C("tasks").Update(bson.M{"_id": tsk.Id}, bson.M{"$set": bson.M{"status": TASK_ERRORED, "error": err.Error()}})
}

func (s *Task) Set(to_set bson.M) error {
	db := db.NewSession()
	defer db.Session.Close()

	return db.C("tasks").Update(bson.M{"_id": s.Id}, bson.M{"$set": to_set})
}

func (tsk *Task) PrettyName() string {
	return fmt.Sprintf("%s:%s", tsk.Action, tsk.Behaviour)
}
