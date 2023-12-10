package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var displayTables bool

// ----------------------------------------------
// open command
// ----------------------------------------------

var openCmd = &cobra.Command{
	Use:   "open [sqlite db file]",
	Short: "Open a SQLite database file",
	Long:  `This tool is intended to be used specficially with EarthRanger mobile databases`,
	Run: func(cmd *cobra.Command, args []string) {
		var filename = args[0]
		open(filename)
	},
}

// ----------------------------------------------
// funtions
// ----------------------------------------------

func open(file string) {
	db, err := connectToSQLite(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case displayTables:
		tables, err := getTables(*db)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, tableName := range tables {
			fmt.Println(tableName)
		}
	default:
		message := fmt.Sprintf("%s successfully opened", file)
		fmt.Println(message)
	}
}

func connectToSQLite(file string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getTables(db gorm.DB) (tableList []string, err error) {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return nil, err
	}

	return tables, nil
}

// ----------------------------------------------
// initialize
// ----------------------------------------------

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.Flags().BoolVarP(&displayTables, "tables", "t", false, "Display all database tables")
}
