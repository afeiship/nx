/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-15 21:05:50
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-11-15 09:19:32
 */
package nx

import (
	"regexp"
	"strings"
)

type GitInfoRepo struct {
	CurrentBranch string `json:"current_branch"`
	Hash          string `json:"hash"`
	ShortHash     string `json:"short_hash"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Owner         string `json:"owner"`
	Repo          string `json:"repo"`
	SshUrl        string `json:"ssh_url"`
	HttpsUrl      string `json:"https_url"`
	LatestVersion string `json:"latest_version"`
}

func GitInfo() GitInfoRepo {
	currentBranch, _ := RunShell(&ShellOptions{Command: "git rev-parse --abbrev-ref HEAD"})
	hash, _ := RunShell(&ShellOptions{Command: "git rev-parse --verify HEAD"})
	shortHash, _ := RunShell(&ShellOptions{Command: "git rev-parse --short HEAD"})
	email, _ := RunShell(&ShellOptions{Command: "git config user.email"})
	name, _ := RunShell(&ShellOptions{Command: "git config user.name"})
	url, _ := RunShell(&ShellOptions{Command: "git config --get remote.origin.url"})
	owner, repo := parseGitUrl(url)
	httpsUrl := strings.Replace(url, "git@", "https://", 1)
	httpsUrl = strings.Replace(httpsUrl, ".git", "", 1)
	latestTag, _ := RunShell(&ShellOptions{Command: "git describe --tags --abbrev=0"})
	latestVersion := strings.TrimSpace(latestTag)

	return GitInfoRepo{
		CurrentBranch: strings.TrimSpace(currentBranch),
		Hash:          strings.TrimSpace(hash),
		ShortHash:     strings.TrimSpace(shortHash),
		Email:         strings.TrimSpace(email),
		Name:          strings.TrimSpace(name),
		Owner:         owner,
		Repo:          repo,
		SshUrl:        strings.TrimSpace(url),
		HttpsUrl:      strings.TrimSpace(httpsUrl),
		LatestVersion: latestVersion,
	}
}

func parseGitUrl(url string) (owner string, repo string) {
	if strings.Contains(url, "github.com") {
		r, _ := regexp.Compile(`github.com[:/](.*?)/(.*?)\.git`)
		owner = r.FindStringSubmatch(url)[1]
		repo = r.FindStringSubmatch(url)[2]
	}
	return owner, repo
}
