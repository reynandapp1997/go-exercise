package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"crud-golang-mongodb/database"
	"crud-golang-mongodb/handler"
	"crud-golang-mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleUsersList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dbase, err := database.Connect()
	if err != nil {
		handler.ResponseHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res, err := dbase.Collection("student").Find(database.GetContext(), bson.M{})
	if err != nil {
		handler.ResponseHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer res.Close(database.GetContext())

	result := make([]models.Student, 0)
	for res.Next(database.GetContext()) {
		var row models.Student
		err := res.Decode(&row)
		if err != nil {
			handler.ResponseHandler(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		result = append(result, row)
	}
	studentList, err := json.Marshal(result)
	w.Write(studentList)
}

func HandleUserSingle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	dbase, err := database.Connect()
	if err != nil {
		handler.ResponseHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userid, _ := primitive.ObjectIDFromHex(vars["id"])
	res := dbase.Collection("student").FindOne(database.GetContext(), bson.M{"_id": userid})
	var student models.Student
	err = res.Decode(&student)
	if err != nil {
		handler.ResponseHandler(w, "Student not found", http.StatusNotFound)
		return
	}
	result, _ := json.Marshal(student)
	w.Write(result)
}

func HandleInsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dbase, err := database.Connect()
	if err != nil {
		handler.ResponseHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	name := r.FormValue("name")
	grade, _ := strconv.Atoi(r.FormValue("grade"))
	student := models.Student{Name: name, Grade: grade}
	_, err = dbase.Collection("student").InsertOne(database.GetContext(), student)
	if err != nil {
		handler.ResponseHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	handler.ResponseHandler(w, "Success insert student", http.StatusCreated)
}
