package main

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (info authenticationInfo) getBasicAuth() string {
	return "Authentification: Basic " + info.username + ":" + info.password
}
