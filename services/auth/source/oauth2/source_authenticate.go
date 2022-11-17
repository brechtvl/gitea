// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package oauth2

import (
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/services/auth/source/db"
)

// Authenticate falls back to the db authenticator
func (source *Source) Authenticate(user *user_model.User, login, password string) (*user_model.User, error) {
	return db.Authenticate(user, login, password)
}

func (source *Source) IsOAuthForced() bool {
	return source.ForceOAuth
}

// NB: Oauth2 does not implement LocalTwoFASkipper for password authentication
// as its password authentication drops to db authentication
