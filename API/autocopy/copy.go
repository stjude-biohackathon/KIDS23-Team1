package autocopy

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"os/exec"
)

const HotUrl = "https://hackathonhot.blob.core.windows.net"
const HotName = "hackathonhot"
const ColdName = "hackathoncold"
const ColdUrl = "https://hackathoncold.blob.core.windows.net"
const ColdAccess = "mlnImajad6aaYtsDIpzuy+dTgzyE4XZzFlbZXqyu3k3AOGDapUrAFkJuFCEylqAfnPtjfeYJqKt/+AStCnbtxA=="
const hotAccess = "+Um064T8oiGZQHg3ZWhKdQDCyOMpK+LfjQ+mDNN6RxVnW/UQOwO00Nu+d0C2gLOWVX+ijo7BWMZL+AStoGyjfQ=="
const ColdSAS = "?sp=racwdli&st=2023-05-05T15:01:29Z&se=2023-05-07T23:01:29Z&spr=https&sv=2022-11-02&sr=c&sig=KsH4leNtLE6HAgC%2BdImZ%2BTqkOp50c3wUrU%2B9SIssxrM%3D"
const HotSAS = "?sp=racwdlmeop&st=2023-05-05T18:28:12Z&se=2023-05-08T02:28:12Z&spr=https&sv=2022-11-02&sr=c&sig=RPmnsNdVo2VyTTldttRBGyOA7CLItCy5JDjfiixZxzs%3D"

func CopyBlobBetweenAccounts(srcAccountName, srcSAS, tier, srcContainerName, srcBlobName, dstAccountName, dstSAS, dstContainerName, dstBlobName string) error {
	cmd := exec.Command("az", "storage", "blob", "copy", "start",
		"--account-name", dstAccountName,
		"--rehydrate-priority", "standard",
		"--sas-token", dstSAS,
		"--source-container", srcContainerName,
		"--source-blob", srcBlobName,
		"--source-account-name", srcAccountName,
		"--source-sas", srcSAS,
		"--destination-blob", dstBlobName,
		"--destination-container", dstContainerName,
		"--tier", tier)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to copy blob: %v, output: %s", err, string(output))
	}
	return nil
}

func DoCopies() {
	ctx := context.Background()
	credential, _ := azidentity.NewDefaultAzureCredential(nil)
	coldContainer, _ := container.NewClient(ColdUrl+"/clingen", credential, nil)
	blobsToCopy := make([]string, 0)
	pager := coldContainer.NewListBlobsFlatPager(nil)
	for pager.More() {
		resp, _ := pager.NextPage(ctx)
		for _, bl := range resp.Segment.BlobItems {
			blobsToCopy = append(blobsToCopy, *bl.Name)
		}
	}
	fmt.Println("made it here")
	totalBlobs := 0
	for _, name := range blobsToCopy {
		err := CopyBlobBetweenAccounts(
			ColdName,
			ColdSAS,
			"Hot",
			"clingen",
			name,
			HotName,
			HotSAS,
			"clingen",
			name)
		if err != nil {
			fmt.Println(err)
		}
		totalBlobs++
	}
	fmt.Println("total blobs copied: ", totalBlobs)
}
