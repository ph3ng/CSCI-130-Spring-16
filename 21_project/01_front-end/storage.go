package main

import (
	"io"

	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"
)

func putFile(ctx context.Context, name string, rdr io.Reader) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}

	writer.ContentType = "image/jpeg"
	io.Copy(writer, rdr)
	// check for errors on io.Copy in production code!
	return writer.Close()
}
