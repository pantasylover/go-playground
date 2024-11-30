package config

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Config struct {
	People []Person `json:"people"`
}

func (cfg *Config) SetDefaults() {
	cfg.People = []Person{
		{
			FirstName: "Christian",
			LastName:  "Deacon",
			Age:       26,
		},
		{
			FirstName: "John",
			LastName:  "Doe",
			Age:       45,
		},
		{
			FirstName: "Eric",
			LastName:  "Smith",
			Age:       21,
		},
	}
}
