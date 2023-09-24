package services

import (
	"context"
	"database/sql"
	"encoding/json"
)

type Course struct {
	CourseID          string         `json:"course_id"`
	CourseName        string         `json:"course_name"`
	CourseDescription string         `json:"course_description"`
	VideoURL          sql.NullString `json:"video_url"` // Use sql.NullString for nullable columns
	InstructorID      string         `json:"instructor_id"`
}

type UserCourse struct {
	UserID   string `json:"user_id"`
	CourseID string `json:"course_id"`
}

func (c *Course) CreateCourse(course Course) (*Course, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO courses (course_name, course_description, video_url, instructor_id)
			  VALUES ($1, $2, $3, $4) returning *`

	_, err := db.ExecContext(ctx, query,
		course.CourseName,
		course.CourseDescription, course.VideoURL, course.InstructorID)

	if err != nil {
		return nil, err
	}
	return &course, nil

}
func (c *UserCourse) EnrollToCourse(userCourse UserCourse) (*UserCourse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO user_courses (user_id, course_id)
			  VALUES ($1, $2) returning *`

	_, err := db.ExecContext(ctx, query, userCourse.UserID, userCourse.CourseID)

	if err != nil {
		return nil, err
	}
	return &userCourse, nil

}

type NullStringUnmarshaler struct {
	sql.NullString
}

func (n *NullStringUnmarshaler) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	n.String = s
	n.Valid = true

	return nil
}
