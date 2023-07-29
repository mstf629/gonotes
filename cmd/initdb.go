//the decliration of BasePath DbPath vars in root.go
//the decliration of CheckErr function in root.go 
//the decliration of CheckDb function in root.go 
package cmd

import (
	 "fmt"
   "log"
   "os"
	 "github.com/spf13/cobra"
   _ "github.com/mattn/go-sqlite3"
   "database/sql"
)

var initdbCmd = &cobra.Command{
   Use:   "initdb",
	 Short: "create new database",
	 Long: "",
	 Run: func(cmd *cobra.Command, args []string) {
      CreateDb()
      SetupDb()
	 },
}

func init() {
	 rootCmd.AddCommand(initdbCmd)
}

func CreateDb() {
      _, BaseErr:= os.Stat(BasePath)
      _, DbErr := os.Stat(DbPath)

      if os.IsNotExist(BaseErr) && os.IsNotExist(DbErr) {
         os.Mkdir(BasePath, os.ModePerm)
         os.Create(DbPath)
         fmt.Println("db created")
      }else if DbErr != nil{
         os.Create(DbPath)
         fmt.Println("db created")
      }else {
         log.Fatalln("the db is already init")
      }

}

func SetupDb() {
   db, err := sql.Open("sqlite3", DbPath)
   CheckErr(err)
   defer db.Close()
   _, err = db.Exec("CREATE TABLE notes (non INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, content LONGBLOB(4294967295),class CHAR(255) ,id CHAR(255) ,date CHAR(8));")
   CheckErr(err)
}
   
