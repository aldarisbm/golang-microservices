package github

/*

{
    "name": "golang-tut",
    "description": "my golang tut",
    "homepage": "joseberr.io",
    "private": false,
    "has_issues": true,
    "has_projects": true,
    "has_wiki": true
}

*/

type CreateRepoRequest struct {
	Name        string
	Description string
	Homepage    string
	Private     bool
	HasIssues   bool
	HasProjects bool
	HasWiki     bool
}

//this response has an owner json field that is represented by the RepoOwner struct
type CreateRepoResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

//nested JSON
type RepoOwner struct {
	ID      int64  `json:"id"`
	URL     string `json:"url"`
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

/* this could also be represented as:

type CreateRepoResponse struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	FullName string    `json:"full_name"`
	Owner    struct{
		ID      int64  `json:"id"`
		URL     string `json:"url"`
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"owner"`
}

but we lose a lot of readability and the ability to create an owner without a RepoResponse
*/

type RepoPermissions struct {
	Admin   bool `json:"admin"`
	CanPush bool `json:"push"`
	CanPull bool `json:"pull"`
}
