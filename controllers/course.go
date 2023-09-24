package controllers

import (
	"encoding/json"
	"github.com/eif-courses/golab/services"
	"github.com/eif-courses/golab/utils"
	"net/http"
)

var course services.Course

// CreateCourse godoc
// @Summary Create a new course
// @Description Create a new course
// @Tags courses
// @Accept  json
// @Produce  json
// @Param course body services.Course true "Course"
// @Success 200 {object} services.Course
// @Router /courses/course [post]
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var courseData services.Course
	err := json.NewDecoder(r.Body).Decode(&courseData)
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}
	courseCreated, err := course.CreateCourse(courseData)
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}
	err = utils.WriteJSON(w, http.StatusOK, courseCreated)
	utils.ServerErrorHTTP(err, w)
}
