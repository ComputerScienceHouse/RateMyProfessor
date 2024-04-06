package main

type Professor struct {
	Id   int    // we could add a json tag if we want to use a JSON payload but I don't see the point rn? Maybe I'm crazy
	Name string // e.g. 'json:"name"'
}

type Course struct {
	Id          int
	Name        string
	Description string
	College     College
}

type Review struct {
	Id        int
	Rating    int
	Date      string
	Course    Course
	Professor Professor
	User      User
}

type User struct {
	Uid     string // uid from ldap
	Name    string // cn or displayname from ldap
	Reviews []Review
	// There's gonna be more I just don't know what (we probably don't really need more tho)
}

type College struct {
	Id   int
	Name string
}
