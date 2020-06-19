/////////////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2019-2020, Unectio Inc, All Right Reserved.
//
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
/////////////////////////////////////////////////////////////////////////////////

package uauth

import (
	"time"
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

var CapNames = map[string]uint {
	"CapModify":		RoleCapModify,
	"CapAnyProject":	RoleCapAnyProject,
	"CapUserManagement":	RoleCapUserManagement,
	"CapPlatformRef":	RoleCapPlatformRef,
	"CapSetTags":		RoleCapSetTags,
	"CapSharedRepos":	RoleCapSharedRepos,
	"CapBreakRates":	RoleCapBreakRates,
	"CapRouterDomain":	RoleCapRouterDomain,
	"CapAccEverything":	RoleCapAccEverything,
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

func SetRoles(r map[string]*Role) { roles = r }

func GetRole(name string) *Role {
	r, ok := roles[name]
	if !ok {
		r = &Role {
			Name:	x_roleNobody,
			Caps:	util.NewEmptyBitmask(),
		}
	}
	return r
}

func Admin() *Role { return GetRole(x_roleAdmin) }
func Nobody() string { return x_roleNobody }
