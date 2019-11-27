package api

type AppSpec struct {
	Funcs		[]*FunctionImage	`yaml:"functions"`
	Auths		[]*AuthMethodImage	`yaml:"auths,omitempty"`
	Router		*RouterImage		`yaml:"router,omitempty"`
	Refs		[]*SpecRef		`yaml:"refs,omitempty"`
}

type SpecRef struct {
	Name		string			`yaml:"name"`
	Value		string			`yaml:"value"`
}
