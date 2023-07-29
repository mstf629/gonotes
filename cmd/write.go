//the decliration of BasePath DbPath vars in root.go
//the decliration of CheckErr function in root.go 
//the decliration of CheckDb function in root.go 
package cmd

import (
   "log"
	 "fmt"
   "time"
	 "github.com/spf13/cobra"
   "database/sql"
   _ "github.com/mattn/go-sqlite3"
)


var writeCmd = &cobra.Command{
	 Use:   "write",
	 Short: "write new note in datebase",
	 Long:  "",
	 Run: func(cmd *cobra.Command, args []string) {
      CheckDb()

      content, _ := cmd.Flags().GetString("content")
      class, _ := cmd.Flags().GetString("class")
      date, _ := cmd.Flags().GetString("date")
      id, _ := cmd.Flags().GetString("id")

      query := GenWriteQuery(content, class, id, date) 

      DbWriter(query)

      ResetNonColumn()
	 },
}

func init() {
	rootCmd.AddCommand(writeCmd)
	writeCmd.PersistentFlags().StringP("content", "", "", "set class")
	writeCmd.PersistentFlags().StringP("class", "c", "", "set class")
	writeCmd.PersistentFlags().StringP("id", "", "", "set id")
  writeCmd.PersistentFlags().StringP("date", "d", "", "set date in format YYYY:MM:DD")
}

func DbWriter(query string) {
      db, err := sql.Open("sqlite3", DbPath)

      CheckErr(err)
      defer db.Close()

      _, err = db.Exec(query)
      CheckErr(err)
}

func GenWriteQuery(content string, class string ,id string, date string) string {
   var query string
   if content == "" {
      log.Fatalln("you cant write an empty note")
   }

   if date == "" {
      date = time.Now().Format("2006:01:02")
   }

   query = fmt.Sprintf("INSERT INTO notes(content, class, id, date) VALUES('%s', '%s', '%s', '%s');", content, class, id, date)
         
   return query 
}

func ResetNonColumn() {
   var NonRows []string
   db, err := sql.Open("sqlite3",DbPath)
   defer db.Close()
   CheckErr(err)

   rows, err := db.Query("select non from notes;")
   defer rows.Close()
   CheckErr(err)

   for rows.Next() {
      var NonRow string

      err = rows.Scan(&NonRow)
      CheckErr(err)

      NonRows = append(NonRows, NonRow) 
   }


   for i, Non := range NonRows{
      query := fmt.Sprintf("UPDATE notes SET non=%d WHERE non='%s';", i, Non)
      _, err = db.Exec(query)
      CheckErr(err)
   }

}

