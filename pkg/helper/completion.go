package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"native/crus/internal/models"
	"net/http"
	"os"
)

var openaiURL = os.Getenv("OPENAI_URL")
var apiKey = os.Getenv("OPENAI_API_KEY")

type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func GetCompletions(prompt string) (string, error) {

	base := "Membuat panduan kursus dari " + prompt + " terdiri dari judul utama, sub-judul, ringkasan singkat dari kursus tersebut, lama durasi dari kursus tersebut, jumlah sub-judul yang ada, jenis kegiatan dari kursus tersebut seperti pendidikan, olahraga, teknologi dan sejenisya, tipe kegiatan dari kursus tersebut dengan pilihan terbatas pada indoor, outdoor, hybrid. Setiap sub-judul memiliki sebuah materi yang berisi kalimat pembuka yang menjelaskan tentang sub-judul tersebut, panduan langkah demi langkah, kalimat penutup yang dipisahkan oleh baris. Setiap sub-judul memiliki deksripsi yang menjelaskan tentang sub-judul tersebut.Setiap langkah pada panduan dijelaskan secara rinci dan jelas dengan minimal tiap langkah memiliki satu paragraf. Setiap sub-judul memiliki panduan dipisahkan menggunakan baris. Pastikan Setiap kursus memiliki minimal 3 sub-judul. Setiap panduan didalam sub-judul hanya memilki satu panduan didalamnya dan memiliki minimal 3 langkah didalamnya dengan setiap langkah hanya berada pada satu panduan. Setiap langkah pada panduan memiliki nomor serta dipisahkan oleh baris. Buatkan dalam format json dengan title, ringkasan singkat dan sub-judul terpisah. Serta Pastikan nama key durasi sama dengan `duration`, jumlah sub-judul sama dengan `lessons`, jenis kegiatan sama dengan `type_activity` , tipe kegiatan sama dengan `theme_activity` ,judul sama dengan title, ringkasan singkat sama dengan `desc`, sub-judul sama dengan `subtitles` berbentuk list, judul didalam sub-judul sama dengan `topic`, deskripsi didalam sub-judul bernama `shortdesc` ,panduan sama dengan `content` benbentuk object (dictionary), kalimat pembuka pada panduan dengan nama `opening`, langkah pada panduan dengan nama `step` berbentuk list [str, str, ...], kalimat penutup pada panduan dengan nama `closing`. Pastikan output selalu konsisten dan memiliki format json yang selalu sama, nama key json sama dengan nama yang sudah diberikan. hilangkan ```json``` pada output."

	requestBody := models.ChatRequest{
		Model: "gpt-3.5-turbo-1106",
		Messages: []models.ChatMessage{
			{
				Role:    "system",
				Content: base,
			},
		},
	}

	requestJSON, _ := json.Marshal(requestBody)
	requestReader := bytes.NewReader(requestJSON)

	req, err := http.NewRequest("POST", openaiURL, requestReader)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResponse ChatResponse
	if err := json.Unmarshal(body, &chatResponse); err != nil {
		return "", err
	}

	if len(chatResponse.Choices) > 0 {
		return chatResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no completion choices received")
}
