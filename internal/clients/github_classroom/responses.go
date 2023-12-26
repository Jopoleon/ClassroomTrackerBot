package github_classroom

import "time"

type Assignment struct {
	Id                          int       `json:"id"`
	PublicRepo                  bool      `json:"public_repo"`
	Title                       string    `json:"title"`
	Type                        string    `json:"type"`
	InviteLink                  string    `json:"invite_link"`
	InvitationsEnabled          bool      `json:"invitations_enabled"`
	Slug                        string    `json:"slug"`
	StudentsAreRepoAdmins       bool      `json:"students_are_repo_admins"`
	FeedbackPullRequestsEnabled bool      `json:"feedback_pull_requests_enabled"`
	MaxTeams                    int       `json:"max_teams"`
	MaxMembers                  int       `json:"max_members"`
	Editor                      string    `json:"editor"`
	Accepted                    int       `json:"accepted"`
	Submitted                   int       `json:"submitted"`
	Passing                     int       `json:"passing"`
	Language                    string    `json:"language"`
	Deadline                    time.Time `json:"deadline"`
	Classroom                   struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Archived bool   `json:"archived"`
		Url      string `json:"url"`
	} `json:"classroom"`
}

type AssignmentFull struct {
	Id                          int       `json:"id"`
	PublicRepo                  bool      `json:"public_repo"`
	Title                       string    `json:"title"`
	Type                        string    `json:"type"`
	InviteLink                  string    `json:"invite_link"`
	InvitationsEnabled          bool      `json:"invitations_enabled"`
	Slug                        string    `json:"slug"`
	StudentsAreRepoAdmins       bool      `json:"students_are_repo_admins"`
	FeedbackPullRequestsEnabled bool      `json:"feedback_pull_requests_enabled"`
	MaxTeams                    int       `json:"max_teams"`
	MaxMembers                  int       `json:"max_members"`
	Editor                      string    `json:"editor"`
	Accepted                    int       `json:"accepted"`
	Submitted                   int       `json:"submitted"`
	Passing                     int       `json:"passing"`
	Language                    string    `json:"language"`
	Deadline                    time.Time `json:"deadline"`
	StaterCodeRepository        struct {
		Id            int    `json:"id"`
		FullName      string `json:"full_name"`
		HtmlUrl       string `json:"html_url"`
		NodeId        string `json:"node_id"`
		Private       bool   `json:"private"`
		DefaultBranch string `json:"default_branch"`
	} `json:"stater_code_repository"`
	Classroom struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Archived bool   `json:"archived"`
		Url      string `json:"url"`
	} `json:"classroom"`
}

// Classroom represents a GitHub Classroom
type Classroom struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
	Url      string `json:"url"`
}

type T struct {
	Id          int         `json:"id"`
	Submitted   bool        `json:"submitted"`
	Passing     bool        `json:"passing"`
	CommitCount int         `json:"commit_count"`
	Grade       interface{} `json:"grade"`
	Students    []struct {
		Id        int     `json:"id"`
		Login     string  `json:"login"`
		Name      *string `json:"name"`
		AvatarUrl string  `json:"avatar_url"`
		HtmlUrl   string  `json:"html_url"`
	} `json:"students"`
	Assignment struct {
		Id                          int         `json:"id"`
		PublicRepo                  bool        `json:"public_repo"`
		Title                       string      `json:"title"`
		Type                        string      `json:"type"`
		InviteLink                  string      `json:"invite_link"`
		InvitationsEnabled          bool        `json:"invitations_enabled"`
		Slug                        string      `json:"slug"`
		StudentsAreRepoAdmins       bool        `json:"students_are_repo_admins"`
		FeedbackPullRequestsEnabled bool        `json:"feedback_pull_requests_enabled"`
		MaxTeams                    interface{} `json:"max_teams"`
		MaxMembers                  interface{} `json:"max_members"`
		Editor                      interface{} `json:"editor"`
		Accepted                    int         `json:"accepted"`
		Submissions                 int         `json:"submissions"`
		Passing                     int         `json:"passing"`
		Language                    interface{} `json:"language"`
		Deadline                    interface{} `json:"deadline"`
		Classroom                   struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			Archived bool   `json:"archived"`
			Url      string `json:"url"`
		} `json:"classroom"`
	} `json:"assignment"`
	Repository struct {
		Id            int    `json:"id"`
		FullName      string `json:"full_name"`
		HtmlUrl       string `json:"html_url"`
		NodeId        string `json:"node_id"`
		Private       bool   `json:"private"`
		DefaultBranch string `json:"default_branch"`
	} `json:"repository"`
}
