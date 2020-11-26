package controllers

import "net/http"

//CreateUser creates the user on api.
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating the user..."))
}

//ListAllUser return all users.
func ListAllUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating the user..."))
}

//ListUser return the specified user.
func ListUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating the user..."))
}

//UpdateUser update the specified user.
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating the user..."))
}

//DeleteUser delete the specified user.
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating the user..."))
}
