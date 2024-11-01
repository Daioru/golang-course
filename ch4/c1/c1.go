package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	messageLimit := 100
	if membershipType == "Premium" {
		messageLimit = 1000
	}
	return User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: messageLimit,
		},
	}
}
