package core

// Config : struct of the config file
type Config struct {
	RoutesFile    string
	NoSQL	      int    `json:"nosql"`
	NoSQLConnType string `json:"nosql_conn_type"`
	NoSQLAddress  string `json:"nosql_address"`
	NoSQLPort     string `json:"nosql_port"`
	NoSQLPassword string `json:"nosql_password"`
	UseSQL	      int    `json:"sql"`
	AddressSQL    string `json:"sql_address"`
	UsernameSQL   string `json:"sql_username"`
	PasswordSQL   string `json:"sql_password"`
	DatabaseSQL   string `json:"sql_database"`
	AuthTable     string `json:"auth_table"`
	AuthUserField string `json:"auth_user_field"`
	AuthPassField string `json:"auth_password_field"`
	AuthLevel     string `json:"auth_level_field"`
	Secret	      string `json:"secret"`
	ServerAddress string `json:"server_address"`
	ServerPort    string `json:"server_port"`
	HTTPS	      int    `json:"https"`
	HTTPSOnly     int    `json:"https_only"`
	CertsDir      string `json:"certs_dir"`
	LogFile       string `json:"logfile"`
}

// Route : struct for the routes config file
type Route struct {
	Name	    string `json:"name"`
	Method	    string `json:"method"`
	Pattern     string `json:"pattern"`
	HandlerFunc string `json:"handler"`
	Level	    int    `json:"auth_req"`
}
