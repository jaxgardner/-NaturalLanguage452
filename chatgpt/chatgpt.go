package chatgpt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func GetChatResponse(userInput string) string {
	apiKey := os.Getenv("OPENAI_API_KEY")
	baseUrl := "https://api.openai.com/v1/chat/completions"

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-4o-mini",
		"store": true,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": userInput,
			},
		},
	})

	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	// Read and parse the response
	body, _ := io.ReadAll(resp.Body)

	// Parse the JSON response
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	choices, _ := result["choices"].([]interface{})

	content, _ := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return trimCodeBlock(content)
}

func trimCodeBlock(content string) string {
	if len(content) > 7 && content[:6] == "```sql" && content[len(content)-3:] == "```" {
		return content[6 : len(content)-3] // Remove "```sql" and trailing "```"
	}

	if len(content) > 6 && content[:3] == "```" && content[len(content)-3:] == "```" {
		return content[3 : len(content)-3]
	}

	return content
}
