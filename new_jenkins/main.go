package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// JenkinsJob represents a Jenkins job
type JenkinsJob struct {
	Name  string       `json:"name"`
	URL   string       `json:"url"`
	Jobs  []JenkinsJob `json:"jobs,omitempty"`
	Class string       `json:"_class"`
}

// JenkinsAPIResponse represents the response from Jenkins API
type JenkinsAPIResponse struct {
	Jobs []JenkinsJob `json:"jobs"`
}

// GetJenkinsJobs recursively fetches all Jenkins jobs
func GetJenkinsJobs(url, user, token string) ([]JenkinsJob, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+"/api/json", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(user, token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	var apiResponse JenkinsAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	var allJobs []JenkinsJob
	for _, job := range apiResponse.Jobs {
		allJobs = append(allJobs, job)
		if job.Class == "com.cloudbees.hudson.plugins.folder.Folder" {
			subJobs, err := GetJenkinsJobs(job.URL, user, token)
			if err != nil {
				return nil, err
			}
			allJobs = append(allJobs, subJobs...)
		}
	}

	return allJobs, nil
}

func main() {
	jenkinsURL := "https://ci.staging.openxlab.org.cn"
	jenkinsUser := "apx103"
	jenkinsToken := "oooo"

	jobs, err := GetJenkinsJobs(jenkinsURL, jenkinsUser, jenkinsToken)
	if err != nil {
		log.Fatalf("Failed to get Jenkins jobs: %v", err)
	}

	for _, job := range jobs {
		fmt.Printf("Job Name: %s, URL: %s\n", job.Name, job.URL)
	}
}
