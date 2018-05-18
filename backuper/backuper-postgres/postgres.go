package postgres

//Backuper struct
type Backuper struct {
	Host     string
	Port     string
	User     string
	Password string
}

//Backup method
func (pb *Backuper) Backup() error {

	return nil
}
