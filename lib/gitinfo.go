/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-15 21:05:50
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-10-15 21:40:17
 */
package nx

import (
	"regexp"
	"strings"
)

type GitInfoRepo struct {
	currentBranch string
	hash          string
	shortHash     string
	email         string
	name          string
	owner         string
	repo          string
	sshUrl        string
	httpsUrl      string
}

func GitInfo() GitInfoRepo {
	currentBranch, _ := RunShell("git rev-parse --abbrev-ref HEAD")
	hash, _ := RunShell("git rev-parse --verify HEAD")
	shortHash, _ := RunShell("git rev-parse --short HEAD")
	email, _ := RunShell("git config user.email")
	name, _ := RunShell("git config user.name")
	url, _ := RunShell("git config --get remote.origin.url")
	owner, repo := parseGitUrl(url)
	httpsUrl := strings.Replace(url, "git@", "https://", 1)
	httpsUrl = strings.Replace(httpsUrl, ".git", "", 1)

	return GitInfoRepo{
		currentBranch: currentBranch,
		hash:          hash,
		shortHash:     shortHash,
		email:         email,
		name:          name,
		sshUrl:        url,
		owner:         owner,
		repo:          repo,
		httpsUrl:      httpsUrl,
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
