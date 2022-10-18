/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/olongfen/gen-id/generator"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gen-id",
	Short: "auto generate id card",
	Long:  "auto generate id card",
}

func main() {

	var (
		count int
		data  []*generator.GeneratorData
	)
	rootCmd.PersistentFlags().IntVar(&count, "count", 1, "gen count (default is 1)")
	execute()
	for i := 0; i < count; i++ {
		data = append(data, generator.NewGeneratorData(nil))
	}
	filePath := fmt.Sprintf("./gen_%d.json", rand.Int())
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	_data, _ := json.Marshal(data)
	file.Write(_data)
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
