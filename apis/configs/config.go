package configs

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JwtExperesIn  int
}

var cfg *conf

func loadConfig(path string) (*conf, error) {

}
