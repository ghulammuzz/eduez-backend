package guide

import (
	"native/crus/config/db"
	"native/crus/pkg/helper"
	"native/crus/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ListGuide(c *fiber.Ctx) error {
	var guides []models.ListGuide

	query := `SELECT id, title, desciption, type_activity, theme_activity FROM course`

	rows, err := db.DB.Query(query)
	if err != nil {
		return helper.ResponseJson(c, 500, "Error Query", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var guide models.ListGuide
		err := rows.Scan(&guide.ID, &guide.Title, &guide.Desc, &guide.TypeActivity, &guide.ThemeActivity)
		if err != nil {
			return helper.ResponseJson(c, 500, "Error fetching course details", err.Error())
		}
		guides = append(guides, guide)
	}

	return helper.ResponseJson(c, 200, "Success", guides)
}
