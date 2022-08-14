package mongodb

import (
	"os"
	"regexp"

	"github.com/cam-inc/mxtransporter/config/constant"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	MongoDbConnectionUrl     string
	MongoDbDatabase          string
	MongoDbCollection        string
	MongoDbCollectionFilter  *regexp.Regexp
	FullDocument             options.FullDocument
	FullDocumentBeforeChange options.FullDocument
}

var mCfg Mongo

func init() {
	mCfg.MongoDbConnectionUrl = os.Getenv(constant.MONGODB_HOST)
	mCfg.MongoDbDatabase = os.Getenv(constant.MONGODB_DATABASE)
	mCfg.MongoDbCollection = os.Getenv(constant.MONGODB_COLLECTION)
	filter := os.Getenv(constant.MONGODB_COLLECTION_FILTER)
	if filter != "" {
		mCfg.MongoDbCollectionFilter = regexp.MustCompile(filter)
	}
	// see https://github.com/mongodb/mongo-go-driver/blob/95de0fb36ca077bbe9a92b3fbf66ef0d28c6eeae/mongo/integration/unified/client_operation_execution.go#L63
	mCfg.FullDocument = options.FullDocument(os.Getenv(constant.MONGODB_FULLDOCUMENT))
	mCfg.FullDocumentBeforeChange = options.FullDocument(os.Getenv(constant.MONGODB_FULLDOCUMENT_BEFORE_CHANGE))
}

func MongoConfig() Mongo {
	return mCfg
}
