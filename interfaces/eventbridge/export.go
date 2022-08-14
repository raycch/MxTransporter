package eventbridge

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	eventbridgeConfig "github.com/cam-inc/mxtransporter/config/eventbridge"
	"github.com/cam-inc/mxtransporter/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	MONGO_CHANGESTREAM = "mongo-changestream"
	// reserved for future to handle large payload
	MONGO_CHANGESTREAM_S3 = "mongo-changestream-s3pointer"
)

type (
	eventbridgeClient interface {
		putRecord(ctx context.Context, csEvents []types.PutEventsRequestEntry) error
	}

	EventbridgeImpl struct {
		Eb eventbridgeClient
	}

	EventbridgeClientImpl struct {
		EbClient *eventbridge.Client
	}
)

func (eb *EventbridgeClientImpl) putRecord(ctx context.Context, csEvents []types.PutEventsRequestEntry) error {
	_, err := eb.EbClient.PutEvents(ctx, &eventbridge.PutEventsInput{Entries: csEvents})
	return err
}

func (b *EventbridgeImpl) ExportToEventbridge(ctx context.Context, cs primitive.M) error {
	ebCfg := eventbridgeConfig.EventbridgeConfig()
	clusterTime := cs["clusterTime"].(primitive.Timestamp).T
	time := time.Unix(int64(clusterTime), 0)
	json, err := json.Marshal(cs)
	if err != nil {
		return errors.InternalServerErrorJsonMarshal.Wrap("Failed to marshal back the change stream event json", err)
	}
	jsonStr := string(json)
	events := []types.PutEventsRequestEntry{
		{
			Detail:       &jsonStr,
			DetailType:   &MONGO_CHANGESTREAM,
			EventBusName: &ebCfg.Eventbus,
			Source:       &ebCfg.Source,
			Time:         &time,
		}}

	if err := b.Eb.putRecord(ctx, events); err != nil {
		return errors.InternalServerErrorEventbridgePut.Wrap("Failed to send events to EventBridge.", err)
	}

	return nil
}
