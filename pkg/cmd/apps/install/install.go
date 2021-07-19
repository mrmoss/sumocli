package install

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdAppsInstall(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var (
		destinationFolderId string
		description         string
		logSource           string
		name                string
		uuid                string
	)
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs the app with given UUID in the folder specified.",
		Run: func(cmd *cobra.Command, args []string) {
			installApp(destinationFolderId, description, logSource, name, uuid, client, log)
		},
	}
	cmd.Flags().StringVar(&destinationFolderId, "destinationFolderId", "", "Specify the folder id that the app should be installed into")
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the app")
	cmd.Flags().StringVar(&logSource, "logSource", "", "Specify a log source name (for example _sourceCategory=test)")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the app")
	cmd.Flags().StringVar(&uuid, "uuid", "", "Specify the UUID of the app to install")
	cmd.MarkFlagRequired("destinationFolderId")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("uuid")
	return cmd
}

func installApp(destinationFolderId string, description string, logSource string, name string, uuid string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.InstallApp(types.AppInstallRequest{
		Name:                name,
		Description:         description,
		DestinationFolderId: destinationFolderId,
		DataSourceValues: map[string]string{
			"logsrc": logSource,
		},
	},
		uuid)
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to install app")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}