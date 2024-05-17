package integrations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GithubClient struct {
	baseUrl    string
	authToken  string
	httpClient *http.Client
}

const APPROVE_ACTION string = "APPROVE"

func NewGithubClient(authToken string) *GithubClient {
	return &GithubClient{
		baseUrl:    "https://api.github.com",
		authToken:  authToken,
		httpClient: &http.Client{},
	}
}

func (client *GithubClient) buildHeaders() map[string]string {
	return map[string]string{
		"Authorization":        fmt.Sprintf("Bearer %s", client.authToken),
		"X-GitHub-Api-Version": "2022-11-28",
		"Accept":               "application/vnd.github+json",
	}
}

func (client *GithubClient) buildRequest(rtype string, path string, params map[string]string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", client.baseUrl, path)
	headers := client.buildHeaders()
	bodyJson, _ := json.Marshal(params)

	request, err := http.NewRequest(rtype, url, bytes.NewReader(bodyJson))

	if err != nil {
		return nil, err
	}

	for hkey, hvalue := range headers {
		request.Header.Add(hkey, hvalue)
	}

	return request, nil
}

func (client *GithubClient) send(request *http.Request) (*http.Response, error) {
	return client.httpClient.Do(request)
}

func (client *GithubClient) sendPost(path string, params map[string]string) error {
	request, err := client.buildRequest("POST", path, params)
	if err != nil {
		return err
	}

	response, err := client.send(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		bodyData, _ := io.ReadAll(response.Body)
		return errors.New(string(bodyData))
	}

	return nil
}

func (client *GithubClient) ApprovePullRequest(repository string, pullRequestId string, bodyComment string) error {
	path := fmt.Sprintf("/repos/%s/pulls/%s/reviews", repository, pullRequestId)
	params := map[string]string{
		"body":  bodyComment,
		"event": APPROVE_ACTION,
	}
	return client.sendPost(path, params)
}
