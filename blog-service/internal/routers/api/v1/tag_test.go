package v1

import (
	"blog-service/internal/service"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestTag_Create(t *testing.T) {
	tests := []service.CreateTagRequest{
		{Name: "go", CreatedBy: "weirdo"},
		{Name: "python", CreatedBy: "weirdo"},
		{Name: "java", CreatedBy: "weirdo"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			bytesData, _ := json.Marshal(tt)
			resp, _ := http.Post("http://localhost:9999/api/v1/tags",
				"application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
			defer resp.Body.Close()
		})
	}
}

func TestTag_List(t *testing.T) {
	url := "http://localhost:9999/api/v1/tags"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("page", "1")
	req.Header.Add("page_size", "1")
	http.DefaultClient.Do(req)
}
