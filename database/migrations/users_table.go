package migrations

import (
	"database/sql"
)

// Up function for users table migration
func UpUsers(db *sql.DB) error {
	_, err := db.Exec(`
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[users]') AND type in (N'U'))
BEGIN
    CREATE TABLE [dbo].[users](
        [id] [int] IDENTITY(1,1) NOT NULL,
        [name] [varchar](255) NOT NULL,
        [email] [varchar](255) NOT NULL,
        [password] [varchar](255) NOT NULL,
        CONSTRAINT [PK_users] PRIMARY KEY CLUSTERED 
        (
            [id] ASC
        )
    )
END
`)
	if err != nil {
		return err
	}

	return nil
}

// Down function for users table migration
func DownUsers(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS users;`)
	if err != nil {
		return err
	}

	return nil
}
