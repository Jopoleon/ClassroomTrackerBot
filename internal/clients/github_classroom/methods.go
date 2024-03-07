package github_classroom

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"

	"github.com/k0kubun/pp/v3"
)

func (c *ClassroomClient) makeRequest(endpoint string, method string) ([]byte, error) {
	req, err := http.NewRequest(method, c.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Handle non-200 responses
		return nil, fmt.Errorf("API request error: %s", resp.Status)
	}
	return io.ReadAll(resp.Body)
}

//	curl -L \
//	 -H "Accept: application/vnd.github+json" \
//	 -H "Authorization: Bearer <YOUR-TOKEN>" \
//	 -H "X-GitHub-Api-Version: 2022-11-28" \
//	 https://api.github.com/classrooms
//
// ListClassrooms lists all classrooms accessible by the token
func (c *ClassroomClient) ListClassrooms() ([]Classroom, error) {

	body, err := c.makeRequest("/classrooms", http.MethodGet)
	if err != nil {
		return nil, err
	}

	var classrooms []Classroom
	if err := json.Unmarshal(body, &classrooms); err != nil {
		return nil, err
	}

	return classrooms, nil
}

// GetClassroomAssignments List assignments for a classroom
// Lists GitHub Classroom assignments for a classroom. Assignments will
// only be returned if the current user is an administrator of the GitHub Classroom.
//
//	curl -L \
//	 -H "Accept: application/vnd.github+json" \
//	 -H "Authorization: Bearer <YOUR-TOKEN>" \
//	 -H "X-GitHub-Api-Version: 2022-11-28" \
//	 https://api.github.com/classrooms/CLASSROOM_ID/assignments
func (c *ClassroomClient) GetClassroomAssignments(classroom Classroom) ([]Assignment, error) {
	endpoint := fmt.Sprintf("/classrooms/%d/assignments", classroom.ID)
	body, err := c.makeRequest(endpoint, "GET")
	if err != nil {
		return nil, err
	}
	var assignments []Assignment
	if err := json.Unmarshal(body, &assignments); err != nil {
		return nil, err
	}

	return assignments, nil
}

//	curl -L \
//	 -H "Accept: application/vnd.github+json" \
//	 -H "Authorization: Bearer <YOUR-TOKEN>" \
//	 -H "X-GitHub-Api-Version: 2022-11-28" \
//	 https://api.github.com/assignments/ASSIGNMENT_ID
func (c *ClassroomClient) GetAssignmentInfo(assignmentID int) (*AssignmentFull, error) {
	endpoint := fmt.Sprintf("/assignments/%d", assignmentID)
	body, err := c.makeRequest(endpoint, "GET")
	if err != nil {
		return nil, err
	}
	var assignments AssignmentFull
	if err := json.Unmarshal(body, &assignments); err != nil {
		return nil, err
	}

	return &assignments, nil
}

// GetAcceptedAssignments List accepted assignments for an assignment
// Lists any assignment repositories that have been created by students accepting
// a GitHub Classroom assignment. Accepted assignments will only be returned
// if the current user is an administrator of the GitHub Classroom for the assignment.
//
//	curl -L \
//	 -H "Accept: application/vnd.github+json" \
//	 -H "Authorization: Bearer <YOUR-TOKEN>" \
//	 -H "X-GitHub-Api-Version: 2022-11-28" \
//	 https://api.github.com/assignments/ASSIGNMENT_ID/accepted_assignments
//
// https://docs.github.com/en/rest/classroom/classroom?apiVersion=2022-11-28#list-accepted-assignments-for-an-assignment
func (c *ClassroomClient) GetAcceptedAssignments(assignmentID int) (*AssignmentFull, error) {
	endpoint := fmt.Sprintf("/assignments/%d/accepted_assignments", assignmentID)
	body, err := c.makeRequest(endpoint, "GET")
	if err != nil {
		return nil, err
	}
	err = os.WriteFile("GetAcceptedAssignments.json", body, fs.ModePerm)
	if err != nil {
		return nil, err
	}
	var assignments AssignmentFull
	if err := json.Unmarshal(body, &assignments); err != nil {
		return nil, err
	}

	pp.Println(string(body))

	return &assignments, nil
}
