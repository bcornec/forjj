package forjfile

import "github.com/forj-oss/forjj-modules/trace"

type ReposStruct map[string]*RepoStruct

func (r ReposStruct) MarshalYAML() (interface{}, error) {
	to_marshal := make(map[string]*RepoStruct)
	for name, repo := range r {
		if repo == nil {
			gotrace.Error("Unable to save Repository '%s'. Repo data is nil.", name)
			continue
		}
		if ! repo.is_infra {
			to_marshal[name] = repo
		}
	}
	return to_marshal, nil
}

func (r ReposStruct) SetRelapps(relAppName, appName string) (_ error) {
	for _, repo := range r {
		if _, err := repo.SetInternalRelApp(relAppName, appName) ; err != nil {
			return err
		}
	}
	return
}

// AllHasAppWith verify rules on all repos and returned true if all respect the rule.
func (r ReposStruct) AllHasAppWith(rules ...string) (found bool, err error) {
	for _, repo := range r {
		if found, err = repo.HasApps(rules ...); err != nil {
			return
		} else if !found {
			return
		}
	}
	return
}

// HasAppWith return true if at least one repo respect the rule.
func (r ReposStruct) HasAppWith(rules ...string) (found bool, err error) {
	for _, repo := range r {
		if found, err = repo.HasApps(rules ...); err != nil {
			return
		} else if found {
			return
		}
	}
	return
}
