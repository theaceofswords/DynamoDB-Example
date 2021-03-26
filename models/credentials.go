package models

type Cred struct {
	AWS_ACCESS_KEY_ID     string `yaml:"aws_id"`
	AWS_SECRET_ACCESS_KEY string `yaml:"aws_secret_key"`
	Region                string `yaml: "region"`

	
}
