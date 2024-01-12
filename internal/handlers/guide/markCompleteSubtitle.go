package guide

import (
	"database/sql"
	"fmt"
	"native/crus/config/db"
	"native/crus/pkg/helper"

	"github.com/gofiber/fiber/v2"
)


func MarkCompleteSubtitle(c *fiber.Ctx) error {

	type CompleteSubtitle struct {
		SubtitleID string `json:"subtitle_id"`
		CourseID   string `json:"course_id"`
	}

	var completeSubtitle CompleteSubtitle
	if err := c.BodyParser(&completeSubtitle); err != nil {
		return helper.ResponseJson(c, 400, "Error parsing", nil)
	}

	query := "UPDATE subtitle SET is_done = 1 WHERE id = ?"

	result, err := db.DB.Exec(query, completeSubtitle.SubtitleID)
	if err != nil {
		return helper.ResponseJson(c, 500, "Error Query", err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return helper.ResponseJson(c, 404, "Subtitle ID not found", nil)
	}

	var totalSubtitles int
	var completedSubtitles sql.NullInt64

	countQuery := `SELECT COUNT(id) AS total_subtitles, SUM(CASE WHEN is_done = 1 THEN 1 ELSE 0 END) AS completed_subtitles FROM subtitle WHERE course_id = ?`

	err = db.DB.QueryRow(countQuery, completeSubtitle.CourseID).Scan(&totalSubtitles, &completedSubtitles)
	fmt.Println("query total subtitle : ", totalSubtitles)
	fmt.Println("query complete :", completedSubtitles.Int64)
	if err != nil {
		if err == sql.ErrNoRows {
			totalSubtitles = 0
			if !completedSubtitles.Valid {
				completedSubtitles.Int64 = 0
			}
			// completedSubtitles.Int64 = completedSubtitles.Int64
		} else {
			return helper.ResponseJson(c, 500, "Error checking subtitles", err.Error())
		}
	}
	fmt.Println("query total subtitle : ", totalSubtitles)
	fmt.Println("query complete :", completedSubtitles.Int64)

	if totalSubtitles > 0 && int(completedSubtitles.Int64) == totalSubtitles {
		courseUpdateQuery := "UPDATE course SET is_done = 1 WHERE id = ?"
		_, err := db.DB.Exec(courseUpdateQuery, completeSubtitle.CourseID)
		if err != nil {
			return helper.ResponseJson(c, 500, "Error updating course status", err.Error())
		}
	}

	return helper.ResponseJson(c, 200, "Sukses mengupdate subtitle", nil)
}
