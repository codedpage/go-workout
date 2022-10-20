package main
 
import (
    "encoding/csv"
    "os"
)
 
func main() {
    file, err := os.OpenFile("csv_apend.csv", os.O_CREATE|os.O_WRONLY, 0777)
    defer file.Close()
 
    if err != nil {
        os.Exit(1)
    }
 
 
    z := append([]string{"aa","bb"},[]string{"xx", "yy", "zz"}...)


    csvWriter := csv.NewWriter(file)
    strWrite := [][]string{z}
    csvWriter.WriteAll(strWrite)
    csvWriter.Flush()
}