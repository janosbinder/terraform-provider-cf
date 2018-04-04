package cloudfoundry

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-cf/cloudfoundry/repo"
)

var repoManager *repo.RepoManager = repo.NewRepoManager()

// getRepositoryFromConfig -
func getRepositoryFromConfig(d *schema.ResourceData) (repository repo.Repository, err error) {

	var (
		version     string
		versionType repo.VersionType
	)

	if v, ok := d.Get("git").([]interface{}); ok && len(v) > 0 {
		gitArgs := v[0].(map[string]interface{})

		var (
			arg, repoURL               string
			user, password, privateKey *string
		)

		repoURL = gitArgs["url"].(string)

		if arg = gitArgs["branch"].(string); len(arg) > 0 {
			version = arg
			versionType = repo.GitVersionTypeBranch
		}
		if arg = gitArgs["tag"].(string); len(arg) > 0 {
			version = arg
			versionType = repo.GitVersionTypeTag
		}
		if arg = gitArgs["user"].(string); len(arg) > 0 {
			s := arg
			user = &s
		}
		if arg = gitArgs["password"].(string); len(arg) > 0 {
			s := arg
			password = &s
		}
		if arg = gitArgs["key"].(string); len(arg) > 0 {
			s := arg
			privateKey = &s
		}

		if repository, err = repoManager.GetGitRepository(repoURL, user, password, privateKey); err != nil {
			return
		}

	} else if v, ok := d.Get("github_release").([]interface{}); ok && len(v) > 0 {
		githubArgs := v[0].(map[string]interface{})

		var (
			arg, ghOwner, ghRepo, archiveName string
			token                             *string
		)

		ghOwner = githubArgs["owner"].(string)
		ghRepo = githubArgs["repo"].(string)
		archiveName = githubArgs["filename"].(string)
		version = githubArgs["version"].(string)
		versionType = repo.DefaultVersionType

		if arg = githubArgs["token"].(string); len(arg) > 0 {
			token = &arg
		}

		if repository, err = repoManager.GetGithubRelease(ghOwner, ghRepo, archiveName, token); err != nil {
			return
		}
	}
	err = repository.SetVersion(version, versionType)
	return
}
