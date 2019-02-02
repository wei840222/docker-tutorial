package model

import "github.com/globalsign/mgo/bson"

type Todo struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
}

const (
	db         = "Todo"
	collection = "TodoModel"
)

func (m *Todo) InsertTodo(todo Todo) error {
	return Insert(db, collection, todo)
}

func (m *Todo) FindAllTodo() ([]Todo, error) {
	var result []Todo
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

func (m *Todo) FindTodoById(id string) (Todo, error) {
	var result Todo
	err := FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
	return result, err
}

func (m *Todo) UpdateTodo(todo Todo) error {
	return Update(db, collection, bson.M{"_id": todo.Id}, todo)
}

func (m *Todo) RemoveTodo(id string) error {
	return Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
