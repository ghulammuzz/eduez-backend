package guide

import (
	"encoding/json"
	"fmt"
	"native/crus/config/db"
	"native/crus/internal/models"
	"native/crus/pkg/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateGuide(c *fiber.Ctx) error {
	var promptReq models.PromptRequest

	// test
	err := c.BodyParser(&promptReq)
	if err != nil {
		return helper.ResponseJson(c, 400, "Error Parsing", err.Error())
	}
	if promptReq.Prompt == "" {
		return helper.ResponseJson(c, 400, "Invalid Request Body", "'prompt' key is missing or empty")
	}
	tx, err := db.DB.Begin()
	if err != nil {
		tx.Rollback()
		return helper.ResponseJson(c, 500, "Error transaction begin", err.Error())
	}
	jsonStr := helper.JsonStr
	// jsonStr, err := helper.GetCompletions(promptReq.Prompt)
	// log.Println("Content from GetCompletions:", jsonStr)
	if err != nil {
		tx.Rollback()
		return helper.ResponseJson(c, 400, "Error Complettions", err.Error())
	}

	var jsonMap map[string]interface{}

	err = json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		tx.Rollback()
		return helper.ResponseJson(c, 500, "Error unmarshal", err.Error())
	}

	expectedKeys := []string{"title", "desc", "type_activity", "theme_activity", "subtitles"}
	for _, key := range expectedKeys {
		if _, ok := jsonMap[key]; !ok {
			tx.Rollback()
			errorMsg := fmt.Sprintf("Key '%s' is missing in the JSON", key)
			return helper.ResponseJson(c, 400, "Missing key in JSON", errorMsg)
		}
	}

	prompt := promptReq.Prompt
	title := jsonMap["title"].(string)
	desc := jsonMap["desc"].(string)
	typeActivity := jsonMap["type_activity"].(string)
	themeActivity := jsonMap["theme_activity"].(string)

	guideID := uuid.New()

	_, err = tx.Exec("INSERT INTO course (id, prompt, title, desciption, type_activity, theme_activity) VALUES (?, ?, ?, ?, ?, ?)",
		guideID.String(), prompt, title, desc, typeActivity, themeActivity)
	if err != nil {
		tx.Rollback()
		return helper.ResponseJson(c, 500, "Error transaction guide", err.Error())
	}
	subQ := jsonMap["subtitles"].([]interface{})
	for _, subtitle := range subQ {
		subtitleID := uuid.New().String()
		sub := subtitle.(map[string]interface{})

		expectedSubtitleKeys := []string{"topic", "shortdesc", "content"}
		for _, key := range expectedSubtitleKeys {
			if _, ok := sub[key]; !ok {
				tx.Rollback()
				errorMsg := fmt.Sprintf("Key '%s' is missing in the subtitle JSON", key)
				return helper.ResponseJson(c, 400, "Missing key in subtitle JSON", errorMsg)
			}
		}

		_, err = tx.Exec("INSERT INTO subtitle (id, topic, shortdesc, course_id) VALUES (?, ?, ?, ?)",
			subtitleID, sub["topic"].(string), sub["shortdesc"].(string), guideID.String())
		if err != nil {
			tx.Rollback()
			return helper.ResponseJson(c, 500, "Error transaction query subtitle", err.Error())
		}
		contentID := uuid.New().String()
		contentMap := sub["content"].(map[string]interface{})

		expectedContentKeys := []string{"opening", "closing", "step"}
		for _, key := range expectedContentKeys {
			if _, ok := contentMap[key]; !ok {
				tx.Rollback()
				errorMsg := fmt.Sprintf("Key '%s' is missing in the content JSON", key)
				return helper.ResponseJson(c, 400, "Missing key in content JSON", errorMsg)
			}
		}

		_, err = tx.Exec("INSERT INTO content (id, opening, closing, subtitle_id) VALUES (?, ?, ?, ?)",
			contentID, contentMap["opening"].(string), contentMap["closing"].(string), subtitleID)
		if err != nil {
			tx.Rollback()
			return helper.ResponseJson(c, 500, "Error transaction query content ", err.Error())
		}
		steps := contentMap["step"].([]interface{})
		for _, step := range steps {
			_, err = tx.Exec("INSERT INTO steps (id, texts, content_id) VALUES (?, ?, ?)",
				uuid.New().String(), step.(string), contentID)
			if err != nil {
				tx.Rollback()
				return helper.ResponseJson(c, 500, "Error transaction step query", err.Error())
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return helper.ResponseJson(c, 500, "Error transaction commit", err.Error())
	}

	return helper.ResponseJson(c, 201, "Berhasil membuat guide", guideID.String())
}
