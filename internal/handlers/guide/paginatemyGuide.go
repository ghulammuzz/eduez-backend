package guide

import (
	"native/crus/config/db"
	"native/crus/internal/models"
	"native/crus/pkg/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PaginateMyGuide(c *fiber.Ctx) error {
	var guides []models.ListGuide
	var page, pageSize int
	var err error

	user := c.Locals("uid")

	page, err = strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err = strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	query := `SELECT id, title, desciption, type_activity, theme_activity, user_id FROM course  WHERE user_id = ? LIMIT ? OFFSET ?`

	rows, err := db.DB.Query(query, user, pageSize, offset)
	if err != nil {
		return helper.ResponseJson(c, 500, "Error Query", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var guide models.ListGuide
		err := rows.Scan(&guide.ID, &guide.Title, &guide.Desc, &guide.TypeActivity, &guide.ThemeActivity, &user)
		if err != nil {
			return helper.ResponseJson(c, 500, "Error fetching course details", err.Error())
		}
		guides = append(guides, guide)
	}

	totalRows := 0
	countQuery := `SELECT COUNT(*) FROM course`
	err = db.DB.QueryRow(countQuery).Scan(&totalRows)
	if err != nil {
		return helper.ResponseJson(c, 500, "Error counting rows", err.Error())
	}
	baseURL := c.Protocol() + "://" + c.Hostname()
	originalURL := c.OriginalURL()

	finalURL := baseURL + originalURL
	nextLink, prevLink := helper.GetPaginationLinks(finalURL, page, pageSize, totalRows)

	pagination := helper.Pagination{
		Count:    len(guides),
		Next:     nextLink,
		Previous: prevLink,
	}
	return helper.ResponseJson(c, 200, "Success", fiber.Map{"pagination": pagination, "results": guides})
}
