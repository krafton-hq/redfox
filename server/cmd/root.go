package cmd

import (
	"io/ioutil"
	"os"
	"sync"

	log_helper "github.com/krafton-hq/golib/log-helper"
	path_utils "github.com/krafton-hq/golib/path-utils"
	"github.com/krafton-hq/red-fox/server/application"
	"github.com/krafton-hq/red-fox/server/application/configs"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

var rootCmd = &cobra.Command{
	Use:   "red-fox",
	Short: "A brief description of your application",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var configPath string
	var debug bool

	rootCmd.Flags().StringVar(&configPath, "config", "./config.yaml", "Application Config Path")
	rootCmd.Flags().BoolVar(&debug, "debug", false, "Print Debug Message Flag")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		log_helper.Initialize(debug, false)

		absPath, err := path_utils.ResolvePathToAbs(configPath)
		if err != nil {
			return err
		}

		buf, err := ioutil.ReadFile(absPath)
		if err != nil {
			return err
		}
		config := &configs.RedFoxConfig{}
		err = yaml.Unmarshal(buf, config)
		if err != nil {
			return err
		}

		app := application.NewApplication(config)
		app.Init()

		wg := &sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
		return nil
	}

}
