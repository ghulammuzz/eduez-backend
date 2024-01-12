package guide

import (
	"native/crus/config/db"
	"native/crus/internal/models"
	"native/crus/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

func ListMyGuide(c *fiber.Ctx) error {
	var guides []models.ListMyGuide

	// user id
	claimUser := "323"

	query := `SELECT id, title, desciption, type_activity, theme_activity, user_id FROM course WHERE user_id = ?`

	rows, err := db.DB.Query(query, claimUser)

	if err != nil {
		return helper.ResponseJson(c, 500, "Error query", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var guide models.ListMyGuide
		err := rows.Scan(&guide.ID, &guide.Title, &guide.Desc, &guide.TypeActivity, &guide.ThemeActivity)
		if err != nil {
			return helper.ResponseJson(c, 500, "Error Scan", err.Error())
		}
		guides = append(guides, guide)
	}

	return helper.ResponseJson(c, 200, "Sukses My Guide", guides)
}
