type Operation struct {
	port int16
	path string
	method HttpMethod
	required []string
	optional []string
	mapping map[string]string
	rules map[string]string
}

func (o *Operation) validate(intent Intent) error {
	
}