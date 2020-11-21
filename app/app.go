package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrls()

	log.Fatal(router.Run(":8080"))
}

/*

Request

POST https://api.github.com/user/repos

-H "Authorization: token b2e71fabae3260b83ff32b5629ecbc86d66c6036"

{
    "name": "golang-tut",
    "description": "my golang tut",
    "homepage": "joseberr.io",
    "private": false,
    "has_issues": true,
    "has_projects": true,
    "has_wiki": true
}

Response

201 created

{
    "id": 314260918,
    "node_id": "MDEwOlJlcG9zaXRvcnkzMTQyNjA5MTg=",
    "name": "golang-tut",
    "full_name": "aldarisbm/golang-tut",
    "private": false,
    "owner": {
        "login": "aldarisbm",
        "id": 32185409,
        "node_id": "MDQ6VXNlcjMyMTg1NDA5",
        "avatar_url": "https://avatars1.githubusercontent.com/u/32185409?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/aldarisbm",
        "html_url": "https://github.com/aldarisbm",
        "followers_url": "https://api.github.com/users/aldarisbm/followers",
        "following_url": "https://api.github.com/users/aldarisbm/following{/other_user}",
        "gists_url": "https://api.github.com/users/aldarisbm/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/aldarisbm/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/aldarisbm/subscriptions",
        "organizations_url": "https://api.github.com/users/aldarisbm/orgs",
        "repos_url": "https://api.github.com/users/aldarisbm/repos",
        "events_url": "https://api.github.com/users/aldarisbm/events{/privacy}",
        "received_events_url": "https://api.github.com/users/aldarisbm/received_events",
        "type": "User",
        "site_admin": false
    },
    "html_url": "https://github.com/aldarisbm/golang-tut",
    "description": "my golang tut",
    "fork": false,
    "url": "https://api.github.com/repos/aldarisbm/golang-tut",
    "forks_url": "https://api.github.com/repos/aldarisbm/golang-tut/forks",
    "keys_url": "https://api.github.com/repos/aldarisbm/golang-tut/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/aldarisbm/golang-tut/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/aldarisbm/golang-tut/teams",
    "hooks_url": "https://api.github.com/repos/aldarisbm/golang-tut/hooks",
    "issue_events_url": "https://api.github.com/repos/aldarisbm/golang-tut/issues/events{/number}",
    "events_url": "https://api.github.com/repos/aldarisbm/golang-tut/events",
    "assignees_url": "https://api.github.com/repos/aldarisbm/golang-tut/assignees{/user}",
    "branches_url": "https://api.github.com/repos/aldarisbm/golang-tut/branches{/branch}",
    "tags_url": "https://api.github.com/repos/aldarisbm/golang-tut/tags",
    "blobs_url": "https://api.github.com/repos/aldarisbm/golang-tut/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/aldarisbm/golang-tut/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/aldarisbm/golang-tut/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/aldarisbm/golang-tut/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/aldarisbm/golang-tut/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/aldarisbm/golang-tut/languages",
    "stargazers_url": "https://api.github.com/repos/aldarisbm/golang-tut/stargazers",
    "contributors_url": "https://api.github.com/repos/aldarisbm/golang-tut/contributors",
    "subscribers_url": "https://api.github.com/repos/aldarisbm/golang-tut/subscribers",
    "subscription_url": "https://api.github.com/repos/aldarisbm/golang-tut/subscription",
    "commits_url": "https://api.github.com/repos/aldarisbm/golang-tut/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/aldarisbm/golang-tut/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/aldarisbm/golang-tut/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/aldarisbm/golang-tut/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/aldarisbm/golang-tut/contents/{+path}",
    "compare_url": "https://api.github.com/repos/aldarisbm/golang-tut/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/aldarisbm/golang-tut/merges",
    "archive_url": "https://api.github.com/repos/aldarisbm/golang-tut/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/aldarisbm/golang-tut/downloads",
    "issues_url": "https://api.github.com/repos/aldarisbm/golang-tut/issues{/number}",
    "pulls_url": "https://api.github.com/repos/aldarisbm/golang-tut/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/aldarisbm/golang-tut/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/aldarisbm/golang-tut/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/aldarisbm/golang-tut/labels{/name}",
    "releases_url": "https://api.github.com/repos/aldarisbm/golang-tut/releases{/id}",
    "deployments_url": "https://api.github.com/repos/aldarisbm/golang-tut/deployments",
    "created_at": "2020-11-19T13:45:28Z",
    "updated_at": "2020-11-19T13:45:28Z",
    "pushed_at": "2020-11-19T13:45:30Z",
    "git_url": "git://github.com/aldarisbm/golang-tut.git",
    "ssh_url": "git@github.com:aldarisbm/golang-tut.git",
    "clone_url": "https://github.com/aldarisbm/golang-tut.git",
    "svn_url": "https://github.com/aldarisbm/golang-tut",
    "homepage": "joseberr.io",
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 0,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 0,
    "license": null,
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "main",
    "permissions": {
        "admin": true,
        "push": true,
        "pull": true
    },
    "allow_squash_merge": true,
    "allow_merge_commit": true,
    "allow_rebase_merge": true,
    "delete_branch_on_merge": false,
    "network_count": 0,
    "subscribers_count": 1
}

*/
