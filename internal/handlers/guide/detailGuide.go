package guide

import (
	"log"
	"native/crus/config/db"
	"native/crus/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

func DetailGuide(c *fiber.Ctx) error {
	id := c.Params("id")

	query := `
        SELECT 
            c.id, c.title, c.desciption, c.type_activity, c.theme_activity,
            s.id AS subtitle_id, s.topic, s.shortdesc,
            ct.id AS content_id, ct.opening, ct.closing,
            st.texts AS steps
        FROM course c
        LEFT JOIN subtitle s ON c.id = s.course_id
        LEFT JOIN content ct ON s.id = ct.subtitle_id
        LEFT JOIN steps st ON ct.id = st.content_id
        WHERE c.id = ?
    `
	rows, err := db.DB.Query(query, id)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching course details"})
	}
	defer rows.Close()

	var courseData struct {
		ID            string `json:"id"`
		Title         string `json:"title"`
		Desc          string `json:"desc"`
		TypeActivity  string `json:"type_activity"`
		ThemeActivity string `json:"theme_activity"`
		Subtitles     []struct {
			ID        string `json:"id"`
			Topic     string `json:"topic"`
			ShortDesc string `json:"shortdesc"`
			Content   struct {
				Opening string   `json:"opening"`
				Closing string   `json:"closing"`
				Steps   []string `json:"steps"`
			} `json:"content"`
		} `json:"subtitles"`
	}

	for rows.Next() {
		var (
			courseID, courseTitle, courseDesc, courseTypeActivity, courseThemeActivity                     string
			subtitleID, subtitleTopic, subtitleShortDesc, contentID, contentOpening, contentClosing, steps string
		)

		err := rows.Scan(
			&courseID, &courseTitle, &courseDesc, &courseTypeActivity, &courseThemeActivity,
			&subtitleID, &subtitleTopic, &subtitleShortDesc,
			&contentID, &contentOpening, &contentClosing, &steps,
		)

		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error scanning database"})
		}

		courseData.ID = courseID
		courseData.Title = courseTitle
		courseData.Desc = courseDesc
		courseData.TypeActivity = courseTypeActivity
		courseData.ThemeActivity = courseThemeActivity

		var subtitleIndex = -1
		for i, subtitle := range courseData.Subtitles {
			if subtitle.ID == subtitleID {
				subtitleIndex = i
				break
			}
		}
		if subtitleIndex == -1 {
			newSubtitle := struct {
				ID        string `json:"id"`
				Topic     string `json:"topic"`
				ShortDesc string `json:"shortdesc"`
				Content   struct {
					Opening string   `json:"opening"`
					Closing string   `json:"closing"`
					Steps   []string `json:"steps"`
				} `json:"content"`
			}{ID: subtitleID, Topic: subtitleTopic, ShortDesc: subtitleShortDesc}
			newSubtitle.Content.Opening = contentOpening
			newSubtitle.Content.Closing = contentClosing
			newSubtitle.Content.Steps = append(newSubtitle.Content.Steps, steps)
			courseData.Subtitles = append(courseData.Subtitles, newSubtitle)
		} else {
			courseData.Subtitles[subtitleIndex].Content.Steps = append(courseData.Subtitles[subtitleIndex].Content.Steps, steps)
		}
	}

	return helper.ResponseJson(c, 200, "Berhail mendapatkan detail", courseData)
}
