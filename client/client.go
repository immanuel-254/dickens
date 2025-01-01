package main

import (
	"C"
	"bytes"
	"dickens/server/views"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// const url = "http://localhost:8080"

var client = &http.Client{
	Timeout: 10 * time.Second,
}

//export SignUp
func SignUp(url string, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.UseCreate.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the POST request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url, route), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export LogIn
func LogIn(url string, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.GetPubKey.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the POST request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url, route), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export LogOut
func LogOut(url string, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.DeletePubKey.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the POST request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url, route), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read and display the response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			message["error"] = err.Error()
			return message
		}

		message["status"] = resp.Status
		message["body"] = string(body)
		return message
	}

	message["status"] = resp.Status

	return message
}

//export ChangeEmail
func ChangeEmail(url string, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.ChangeEmail.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the PUT request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", url, route), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export ResetPassword
func ResetPassword(url string, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.ResetPassword.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the PUT request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", url, route), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export ReadProfile
func ReadProfile(url string, id int64) map[string]string {
	message := make(map[string]string)

	route := views.ProfileRead.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the GET request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s%v", url, route, id), nil)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export ListProfile
func ListProfile(url string) map[string]string {
	message := make(map[string]string)

	route := views.ProfileList.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the GET request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", url, route), nil)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export UpdateUser
func UpdateUser(url string, id int64, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.UserUpdate.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the PUT request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s%v", url, route, id), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export UpdateProfile
func UpdateProfile(url string, id int64, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.ProfileUpdate.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the PUT request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s%v", url, route, id), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export CreateCategory
func CreateCategory(url string, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.CategoryCreate.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the POST request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url, route), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export ReadCategory
func ReadCategory(url string, id int64) map[string]string {
	message := make(map[string]string)

	route := views.CategoryRead.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the GET request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s%v", url, route, id), nil)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export ListCategory
func ListCategory(url string) map[string]string {
	message := make(map[string]string)

	route := views.CategoryList.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the GET request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", url, route), nil)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export UpdateCategory
func UpdateCategory(url string, id int64, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.CategoryUpdate.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the PUT request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s%v", url, route, id), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export CreateBlog
func CreateBlog(url string, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.BlogCreate.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the POST request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url, route), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export ReadBlog
func ReadBlog(url string, id int64) map[string]string {
	message := make(map[string]string)

	route := views.BlogRead.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the GET request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s%v", url, route, id), nil)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export ListBlog
func ListBlog(url string) map[string]string {
	message := make(map[string]string)

	route := views.BlogList.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the GET request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", url, route), nil)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export UpdateBlog
func UpdateBlog(url string, id int64, data map[string]string) map[string]string {
	message := make(map[string]string)
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	route := views.BlogUpdate.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the PUT request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s%v", url, route, id), bytes.NewBuffer(jsonData))
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	// Read and display the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message["error"] = err.Error()
		return message
	}

	message["status"] = resp.Status
	message["body"] = string(body)
	return message
}

//export DeleteBlog
func DeleteBlog(url string, id int64) map[string]string {
	message := make(map[string]string)

	route := views.BlogDelete.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the PUT request
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%v", url, route, id), nil)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	message["status"] = resp.Status
	return message
}

//export DeleteCategory
func DeleteCategory(url string, id int64) map[string]string {
	message := make(map[string]string)

	route := views.CategoryDelete.Route

	if len(route) > 0 {
		route = route[1:] // Remove the first character
	}

	// Create the PUT request
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%v", url, route, id), nil)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		message["error"] = err.Error()
		return message
	}
	defer resp.Body.Close()

	message["status"] = resp.Status
	return message
}

func main() {}
