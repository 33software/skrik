package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"

// 	"audio-stream-golang/routes"
// )

// func setupApp() *fiber.App {
// 	app := fiber.New()
// 	// Настраиваем маршруты
// 	routes.SetupUserRoutes(app)
// 	return app
// }

// func TestAllEndpoints(t *testing.T) {
// 	app := setupApp()
// 	// Тест для эндпоинта POST /api/users
// 	t.Run("TestCreateUser", func(t *testing.T) {
// 		user := map[string]interface{}{
// 			"username": "john_doe",
// 			"email":    "john@example.com",
// 			"password": "123456",
// 		}
// 		userJSON, _ := json.Marshal(user)
// 		req := httptest.NewRequest("POST", "/api/users", bytes.NewReader(userJSON))
// 		req.Header.Set("Content-Type", "application/json")
// 		resp, err := app.Test(req)
// 		if err != nil {
// 			t.Fatalf("Failed to send request: %v", err)
// 		}
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 		var response map[string]interface{}
// 		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 			t.Fatalf("Failed to decode response: %v", err)
// 		}
// 		assert.Equal(t, "john_doe", response["username"])
// 		assert.Equal(t, "john@example.com", response["email"])
// 	})

// 	// Тест для эндпоинта GET /api/users
// 	t.Run("TestGetUser", func(t *testing.T) {
// 		req := httptest.NewRequest("GET", "/api/users?userid=123", nil)
// 		resp, err := app.Test(req)
// 		if err != nil {
// 			t.Fatalf("Failed to send request: %v", err)
// 		}
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 		body, _ := io.ReadAll(resp.Body)
// 		assert.Equal(t, "Hello 123", string(body))
// 	})

// 	t.Run("TestGetUserNoID", func(t *testing.T) {
// 		req := httptest.NewRequest("GET", "/api/users", nil)
// 		resp, err := app.Test(req)
// 		if err != nil {
// 			t.Fatalf("Failed to send request: %v", err)
// 		}
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 		body, _ := io.ReadAll(resp.Body)
// 		assert.Equal(t, "Hello World", string(body))
// 	})

// 	// Тест для эндпоинта PUT /api/users
// 	t.Run("TestUpdateUser", func(t *testing.T) {
// 		user := map[string]interface{}{
// 			"userid":   123,
// 			"username": "john_doe_updated",
// 			"email":    "john_updated@example.com",
// 		}
// 		userJSON, _ := json.Marshal(user)
// 		req := httptest.NewRequest("PUT", "/api/users?userid=123", bytes.NewReader(userJSON))
// 		req.Header.Set("Content-Type", "application/json")
// 		resp, err := app.Test(req)
// 		if err != nil {
// 			t.Fatalf("Failed to send request: %v", err)
// 		}
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 		var response map[string]interface{}
// 		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 			t.Fatalf("Failed to decode response: %v", err)
// 		}
// 		assert.Equal(t, float64(123), response["userid"])
// 		assert.Equal(t, "john_doe_updated", response["username"])
// 		assert.Equal(t, "john_updated@example.com", response["email"])
// 	})

// 	// Тест для эндпоинта DELETE /api/users
// 	t.Run("TestDeleteUser", func(t *testing.T) {
// 		req := httptest.NewRequest("DELETE", "/api/users?userid=123", nil)
// 		resp, err := app.Test(req)
// 		if err != nil {
// 			t.Fatalf("Failed to send request: %v", err)
// 		}
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)
// 		body, _ := io.ReadAll(resp.Body)
// 		assert.Equal(t, "Delete 123", string(body))
// 	})

// 	t.Run("TestDeleteUserNoID", func(t *testing.T) {
// 		req := httptest.NewRequest("DELETE", "/api/users", nil)
// 		resp, err := app.Test(req)
// 		if err != nil {
// 			t.Fatalf("Failed to send request: %v", err)
// 		}
// 		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
// 	})
// }
