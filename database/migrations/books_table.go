package migrations

import (
	"database/sql"
)

// Up function for books table migration
func UpBooks(db *sql.DB) error {
	_, err := db.Exec(`
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[books]') AND type in (N'U'))
BEGIN
    CREATE TABLE [dbo].[users](
        [id] [int] IDENTITY(1,1) NOT NULL,
        [title] [varchar](255) NOT NULL,
        [author] [varchar](255) NOT NULL,
        [year] [integer] NOT NULL,
        CONSTRAINT [PK_books] PRIMARY KEY CLUSTERED 
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

// Down function for books table migration
func DownBooks(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS books;`)
	if err != nil {
		return err
	}

	return nil
}
