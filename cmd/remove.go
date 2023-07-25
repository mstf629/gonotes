//the decliration of BasePath DbPath vars in root.go
//the decliration of CheckErr function in root.go 
//the decliration of CheckDb function in root.go 
package cmd

import (
   "log"
	 "fmt"
	 "github.com/spf13/cobra"
   _ "github.com/mattn/go-sqlite3"
   "database/sql"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	 Use:   "remove",
	 Short: "remove note from database",
	 Long: "",
	 Run: func(cmd *cobra.Command, args []string) {
      non, _ := cmd.Flags().GetString("non")

      query := GenRemoveQuery(non)
      RemoveNote(query)
       
	 },
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.PersistentFlags().StringP("non", "n", "", "set number of note to remove it")

}

func GenRemoveQuery(non string) string {
   var query string
   if non == "" {
      log.Fatalln("please set number of note that you want remove it")
   }
   query = fmt.Sprintf("DELETE FROM notes WHERE non=%s ;", non)

   return query
}

func RemoveNote(query string) {
   db, err := sql.Open("sqlite3", DbPath)
   CheckErr(err)
   defer db.Close()
   
   _, err = db.Exec(query)
   CheckErr(err)
   fmt.Println("note removed")
}
