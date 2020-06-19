package datastore

/* DB is an interface for the app database. Any struct used as the app database
* must implement these functions */
type DB interface {
	Init(config DatabaseConfig) error // Initialize pgx connection
	//AllMaintenanceRequests()
}
type DatabaseConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

// MaintenanceRequest models the maintenance_request table
type MaintenanceRequest struct {
	ID            int
	Request       string
	UserSubmitted int
	//DateSubmitted time.Time    May need to change if things are read in time.Time
	DateSubmitted string
}
