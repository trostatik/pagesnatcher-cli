package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/trostatik/pagesnatcher"
)

var cloneCmd = &cobra.Command{
	Use:   "clone [url]",
	Short: "Clone a webpage and save it locally",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Println("Initiating clone")
		outputDir, _ := cmd.Flags().GetString("output")
		zip, _ := cmd.Flags().GetString("zip")
		profile, _ := cmd.Flags().GetString("profile")
		cleanFlag, _ := cmd.Flags().GetBool("clean")
		cfg := &pagesnatcher.Config{
			OutputDir: outputDir,
			TargetURL: url,
			Clean:     &cleanFlag,
			ZipPath:   zip,
		}
		cfg.Apply(profile)
		srv := pagesnatcher.NewService(cfg)
		if err := srv.DownloadSite(); err != nil {
			panic(err)
		}
		fmt.Println("Site downloaded")
		if err := srv.FixLocalSite(); err != nil {
			panic(err)
		}
		fmt.Println("Local fixes complete")
		if err := srv.FixQueryPaths(); err != nil {
			panic(err)
		}
		if zip != "" {
			srv.CreateZip()
			fmt.Println("Zip complete")
		}
		fmt.Println("Completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	// Define flags and configuration settings.
	cloneCmd.Flags().StringP("zip", "z", "", "Output path for the zipped file")
	cloneCmd.Flags().StringP("output", "o", "", "Output directory for the cloned webpage")
	cloneCmd.Flags().StringP("profile", "p", "none", "Profile to apply for downloading from specific providers. Options are: [none, framer]")
	cloneCmd.Flags().BoolP("clean", "c", false, "Clear output directory before downloading files")

}
