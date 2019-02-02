package controller

import (
	model "github.com/wei840222/todo-api/model"

	"encoding/json"
	"net/http"
	
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

var (
	dao = model.Todo{}
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var todo []model.Todo
	todo, err := dao.FindAllTodo()
	if err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, todo)

}

func FindTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := dao.FindTodoById(id)
	if err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, result)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	todo.Id = bson.NewObjectId()
	if err := dao.InsertTodo(todo); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusCreated, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var params model.Todo
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateTodo(params); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := dao.RemoveTodo(id); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	responseWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
