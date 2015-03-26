// package ghauth handles authentication to the GitHub API
package ghauth

import (
	"net/http"

	"github.com/google/go-github/github"
)

// Transport holds the basic auth info
type Transport struct {
	Username string
	Password string
}

// RoundTrip returns the default round trip transport
func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.Username, t.Password)
	return http.DefaultTransport.RoundTrip(req)
}

// Client returns a client with the configured transport
func (t *Transport) Client() *http.Client {
	return &http.Client{Transport: t}
}

// GitHubClient provides a pointer to a usable client for the Github API
func GitHubClient() *github.Client {
	t := &Transport{
		Username: "",
		Password: "",
	}
	return github.NewClient(t.Client())
}

// OrgMembers gets a list of users in an organization.
func OrgMembers(org string) ([]string, error) {
	l := &github.ListMembersOptions{
		PublicOnly: false,
		Filter:     "all",
		ListOptions: github.ListOptions{
			PerPage: 1000,
		},
	}
	client := GitHubClient()
	users, _, err := client.Organizations.ListMembers(org, l)
	if err != nil {
		return nil, err
	}
	var u []string
	for _, i := range users {
		u = append(u, *i.Login)
	}
	return u, nil
}
