package uauth

import (
	"time"
	"errors"
	"github.com/unectio/db"
	"github.com/unectio/util"
)

const (
	x_roleAdmin string	= "faas.admin"
	x_roleNobody string	= "faas.nobody"

	SelfKeyLifetime time.Duration = 60 * time.Second
)

const (
	/*
	 * Modify objects in a project
	 */
	RoleCapModify = iota
	/*
	 * Work with any object on the platform
	 */
	RoleCapAccEverything
	/*
	 * Work with code only
	 */
	RoleCapAccFnCode
	/*
	 * Access function (and other) logs
	 */
	RoleCapAccLogs
	/*
	 * Claim any project to work on, even if no explicit role
	 * configured for it
	 */
	RoleCapAnyProject
	/*
	 * Create/Modify/Delete users
	 */
	RoleCapUserManagement
	/*
	 * Reference platform objects in aaas, envs, etc.
	 */
	RoleCapPlatformRef
	/*
	 * Set tags on objects with POST methods (see db.Name.Tags)
	 */
	RoleCapSetTags
	/*
	 * Manipulate shared repos
	 */
	RoleCapSharedRepos
	/*
	 * Perform (mostly IO) operations even if the enforced rate is
	 * too high already.
	 */
	RoleCapBreakRates
	/*
	 * Configure domain for router
	 */
	RoleCapRouterDomain
)

var capNames = map[string]uint {
	"CapModify":		RoleCapModify,
	"CapAnyProject":	RoleCapAnyProject,
	"CapUserManagement":	RoleCapUserManagement,
	"CapPlatformRef":	RoleCapPlatformRef,
	"CapSetTags":		RoleCapSetTags,
	"CapSharedRepos":	RoleCapSharedRepos,
	"CapBreakRates":	RoleCapBreakRates,
	"CapRouterDomain":	RoleCapRouterDomain,
	"CapAccEverythig":	RoleCapAccEverything,
	"CapAccFunctionCode":	RoleCapAccFnCode,
	"CapAccLogs":		RoleCapAccLogs,
}

type Role struct {
	Name	string
	Caps	*util.Bitmask
}

func (r *Role)Can(w uint) bool {
	return r.Caps.Check(w)
}

var roles  = make(map[string]*Role)

func GetRole(name string) *Role { return roles[name] }
func Admin() *Role { return GetRole(x_roleAdmin) }
func Nobody() string { return x_roleNobody }

func LoadRoleCaps(c *db.CapsDb) error {
	if _, ok := roles[c.Role]; ok {
		return errors.New("duplicate entry")
	}

	caps, err := parseCaps(c.Caps)
	if err != nil {
		return err
	}

	roles[c.Role] = &Role{
		Name:	c.Role,
		Caps:	caps,
	}

	return nil
}

func parseCaps(caps []string) (*util.Bitmask, error) {
	if len(caps) == 1 && caps[0] == "*" {
		return util.NewFullBitmask(), nil
	}

	ret := util.NewEmptyBitmask()

	for _, c := range caps {
		n, ok := capNames[c]
		if !ok {
			return nil, errors.New("no such cap " + c)
		}

		ret.Set(n)
	}

	return ret, nil
}
