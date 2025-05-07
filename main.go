package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	// エミュレータ接続時は任意のプロジェクトIDを指定できる
	projectID := "example-project"

	emulatorHost := os.Getenv("PUBSUB_EMULATOR_HOST")
	if emulatorHost == "" {
		log.Fatal("PUBSUB_EMULATOR_HOST environment variable not set.")
	}

	// エミュレータ接続用のクライアントオプションを作成する
	opts := []option.ClientOption{
		option.WithEndpoint(emulatorHost),
		option.WithoutAuthentication(), // 認証を無効化
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	}

	// Pub/Subスキーマクライアントを作成する
	client, err := pubsub.NewSchemaClient(ctx, projectID, opts...)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub schema client: %v", err)
	}
	defer client.Close()

	log.Printf("Successfully connected to Pub/Sub emulator (SchemaClient) at %s with project ID %s\n", emulatorHost, projectID)

	if len(os.Args) < 2 {
		log.Fatal("Please specify 'create' or 'update' as a command line argument.")
	}

	command := os.Args[1]
	schemaID := "example"

	switch command {
	case "create":
		// スキーマの作成
		s, err := createSchema(ctx, client, schemaID)
		if err != nil {
			log.Fatalf("Failed to create schema: %v", err)
		}
		log.Printf("Schema created: %v", s)
	case "update":
		// スキーマの更新 (新しいリビジョンの作成)
		s, err := updateSchema(ctx, client, projectID, schemaID)
		if err != nil {
			log.Fatalf("Failed to update schema: %v", err)
		}
		log.Printf("Schema updated: %v", s)
	default:
		log.Fatalf("Invalid command: %s. Please specify 'create' or 'update'.", command)
	}
}

func createSchema(ctx context.Context, client *pubsub.SchemaClient, schemaID string) (*pubsub.SchemaConfig, error) {
	config := pubsub.SchemaConfig{
		Type: pubsub.SchemaAvro,
		Definition: `{
			"type": "record",
			"name": "MyRecord",
			"fields": [
				{
					"name": "id",
					"type": "string"
				},
				{
					"name": "name",
					"type": "string"
				}
			]
		}`,
	}

	s, err := client.CreateSchema(ctx, schemaID, config)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func updateSchema(ctx context.Context, client *pubsub.SchemaClient, projcetID, schemaID string) (*pubsub.SchemaConfig, error) {
	config := pubsub.SchemaConfig{
		Name: fmt.Sprintf("projects/%s/schemas/%s", projcetID, schemaID),
		Type: pubsub.SchemaAvro,
		Definition: `{
			"type": "record",
			"name": "MyRecord",
			"fields": [
				{
					"name": "address",
					"type": "string"
				},
				{
					"name": "email",
					"type": "string"
				}
			]
		}`,
	}

	s, err := client.CommitSchema(ctx, schemaID, config)
	if err != nil {
		return nil, err
	}
	return s, nil
}
