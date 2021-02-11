package chopsticks

import (
	"github.com/go-git/go-git/v5"
)

// gitClone provides a `git clone <url> <dir>` experience
// path must be an absolute path, url must be a valid git repo url
func gitClone(path, url string) (*git.Repository, error) {
	r, err := git.PlainClone(path, false, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		return r, err
	}
	return r, nil
}

func gitPull(path string) error {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		return err
	}

	return nil
}
