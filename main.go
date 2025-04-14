package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

/*
	storageAccountUrl       = "https://sad8ayv7a7test1.blob.core.windows.net/"
	containerName = "ct-hosting-cts-d8ay-v7a7-test1"
	blobName      = "test-down.txt"
	localFile    = "./test-down.txt"
*/
func main() {
	storageAccountUrl   := flag.String("sa", "", "The Storage Account URL")
	containerName :=  flag.String("container", "", "The container Name")
	blobName      :=  flag.String("remote", "", "The Blob name to download")
	localFile    :=  flag.String("local", "", "The local file to create")
	flag.Parse()

	// authenticate with Azure Active Directory
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}


	// create a client for the specified storage account
	client, err := azblob.NewClient(*storageAccountUrl, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create or open a local file where we can download the blob
	file, err := os.Create(*localFile)
	if err != nil {
		log.Fatal(err)
	}

	// Download the blob to the local file
	_, err = client.DownloadFile(context.TODO(), *containerName, *blobName, file, nil)
	if err != nil {
		log.Fatal(err)
	}
}
