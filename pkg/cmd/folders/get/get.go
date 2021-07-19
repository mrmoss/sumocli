package get

import (
	"github.com/antihax/optional"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdGet(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var (
		id          string
		isAdminMode bool
	)
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a folder with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			get(id, isAdminMode, client, log)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the identifier of the folder")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	return cmd
}

func get(id string, isAdminMode bool, client *cip.APIClient, log *zerolog.Logger) {
	adminMode := cmdutils.AdminMode(isAdminMode)
	apiResponse, httpResponse, errorResponse := client.GetFolder(id, &types.FolderManagementApiGetFolderOpts{
		IsAdminMode: optional.NewString(adminMode),
	})
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to get folder")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}