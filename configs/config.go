package configs

var cfg *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        int
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort int
	JWTSecret     string
	JWTExpiresIn  int
}

func LoadConfig(path string) (*conf, error) {

}
