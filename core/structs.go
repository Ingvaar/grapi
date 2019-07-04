package core

// Files : struct with path to config files
type Files struct {
	Routes string
	Config string
}

// Config : struct of the config file
type Config struct {
	Files		 Files
	Cache		 int	`json:"cache"`
	CacheConnType	 string `json:"cache_conn_type"`
	CacheAddress	 string `json:"cache_address"`
	CachePort	 string `json:"cache_port"`
	CachePassword	 string `json:"cache_password"`
	Database	 int	`json:"db"`
	DatabaseAddress  string `json:"db_address"`
	DatabaseUsername string `json:"db_username"`
	DatabasePassword string `json:"db_password"`
	DatabaseName	 string `json:"db_name"`
	AuthTable	 string `json:"auth_table"`
	AuthUserField	 string `json:"auth_user_field"`
	AuthPassField	 string `json:"auth_password_field"`
	AuthLevel	 string `json:"auth_level_field"`
	Secret		 string `json:"secret"`
	ServerAddress	 string `json:"server_address"`
	ServerPort	 string `json:"server_port"`
	HTTPS		 int	`json:"https"`
	HTTPSOnly	 int	`json:"https_only"`
	CertsDir	 string `json:"certs_dir"`
	LogFile		 string `json:"logfile"`
}

// Route : struct for the routes config file
type Route struct {
	Name	    string `json:"name"`
	Method	    string `json:"method"`
	Pattern     string `json:"pattern"`
	HandlerFunc string `json:"handler"`
	Level	    int    `json:"auth_req"`
}
