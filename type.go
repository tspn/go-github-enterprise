package github

import (
	"time"
	"github.com/parnurzeal/gorequest"
)

type Github struct {
	Agent        *gorequest.SuperAgent
	Access_token string
	Base_url     string
	Organize     string
	Repository   string
}

type Error struct {
	Message string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

type PullRequest struct {
	ID int `json:"id"`
	URL string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL string `json:"labels_url"`
	CommentsURL string `json:"comments_url"`
	EventsURL string `json:"events_url"`
	HTMLURL string `json:"html_url"`
	Number int `json:"number"`
	State string `json:"state"`
	Title string `json:"title"`
	Body string `json:"body"`
	User struct {
		   Login string `json:"login"`
		   ID int `json:"id"`
		   AvatarURL string `json:"avatar_url"`
		   GravatarID string `json:"gravatar_id"`
		   URL string `json:"url"`
		   HTMLURL string `json:"html_url"`
		   FollowersURL string `json:"followers_url"`
		   FollowingURL string `json:"following_url"`
		   GistsURL string `json:"gists_url"`
		   StarredURL string `json:"starred_url"`
		   SubscriptionsURL string `json:"subscriptions_url"`
		   OrganizationsURL string `json:"organizations_url"`
		   ReposURL string `json:"repos_url"`
		   EventsURL string `json:"events_url"`
		   ReceivedEventsURL string `json:"received_events_url"`
		   Type string `json:"type"`
		   SiteAdmin bool `json:"site_admin"`
	   } `json:"user"`
	Labels []struct {
		ID int `json:"id"`
		URL string `json:"url"`
		Name string `json:"name"`
		Color string `json:"color"`
		Default bool `json:"default"`
	} `json:"labels"`
	Assignee struct {
		   Login string `json:"login"`
		   ID int `json:"id"`
		   AvatarURL string `json:"avatar_url"`
		   GravatarID string `json:"gravatar_id"`
		   URL string `json:"url"`
		   HTMLURL string `json:"html_url"`
		   FollowersURL string `json:"followers_url"`
		   FollowingURL string `json:"following_url"`
		   GistsURL string `json:"gists_url"`
		   StarredURL string `json:"starred_url"`
		   SubscriptionsURL string `json:"subscriptions_url"`
		   OrganizationsURL string `json:"organizations_url"`
		   ReposURL string `json:"repos_url"`
		   EventsURL string `json:"events_url"`
		   ReceivedEventsURL string `json:"received_events_url"`
		   Type string `json:"type"`
		   SiteAdmin bool `json:"site_admin"`
	   } `json:"assignee"`
	Milestone struct {
		   URL string `json:"url"`
		   HTMLURL string `json:"html_url"`
		   LabelsURL string `json:"labels_url"`
		   ID int `json:"id"`
		   Number int `json:"number"`
		   State string `json:"state"`
		   Title string `json:"title"`
		   Description string `json:"description"`
		   Creator struct {
			       Login string `json:"login"`
			       ID int `json:"id"`
			       AvatarURL string `json:"avatar_url"`
			       GravatarID string `json:"gravatar_id"`
			       URL string `json:"url"`
			       HTMLURL string `json:"html_url"`
			       FollowersURL string `json:"followers_url"`
			       FollowingURL string `json:"following_url"`
			       GistsURL string `json:"gists_url"`
			       StarredURL string `json:"starred_url"`
			       SubscriptionsURL string `json:"subscriptions_url"`
			       OrganizationsURL string `json:"organizations_url"`
			       ReposURL string `json:"repos_url"`
			       EventsURL string `json:"events_url"`
			       ReceivedEventsURL string `json:"received_events_url"`
			       Type string `json:"type"`
			       SiteAdmin bool `json:"site_admin"`
		       } `json:"creator"`
		   OpenIssues int `json:"open_issues"`
		   ClosedIssues int `json:"closed_issues"`
		   CreatedAt time.Time `json:"created_at"`
		   UpdatedAt time.Time `json:"updated_at"`
		   ClosedAt time.Time `json:"closed_at"`
		   DueOn time.Time `json:"due_on"`
	   } `json:"milestone"`
	Locked bool `json:"locked"`
	Comments int `json:"comments"`
	PullRequest struct {
		   URL string `json:"url"`
		   HTMLURL string `json:"html_url"`
		   DiffURL string `json:"diff_url"`
		   PatchURL string `json:"patch_url"`
	   } `json:"pull_request"`
	ClosedAt interface{} `json:"closed_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedBy struct {
		   Login string `json:"login"`
		   ID int `json:"id"`
		   AvatarURL string `json:"avatar_url"`
		   GravatarID string `json:"gravatar_id"`
		   URL string `json:"url"`
		   HTMLURL string `json:"html_url"`
		   FollowersURL string `json:"followers_url"`
		   FollowingURL string `json:"following_url"`
		   GistsURL string `json:"gists_url"`
		   StarredURL string `json:"starred_url"`
		   SubscriptionsURL string `json:"subscriptions_url"`
		   OrganizationsURL string `json:"organizations_url"`
		   ReposURL string `json:"repos_url"`
		   EventsURL string `json:"events_url"`
		   ReceivedEventsURL string `json:"received_events_url"`
		   Type string `json:"type"`
		   SiteAdmin bool `json:"site_admin"`
	   } `json:"closed_by"`
}
