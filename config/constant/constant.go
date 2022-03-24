package constant

const (
	BIGQUERY_DATASET = "BIGQUERY_DATASET"
	BIGQUERY_TABLE   = "BIGQUERY_TABLE"

	KINESIS_STREAM_NAME   = "KINESIS_STREAM_NAME"
	KINESIS_STREAM_REGION = "KINESIS_STREAM_REGION"

	MONGODB_HOST       = "MONGODB_HOST"
	MONGODB_DATABASE   = "MONGODB_DATABASE"
	MONGODB_COLLECTION = "MONGODB_COLLECTION"

	RESUME_TOKEN_VOLUME_DIR         = "RESUME_TOKEN_VOLUME_DIR"
	RESUME_TOKEN_VOLUME_TYPE        = "RESUME_TOKEN_VOLUME_TYPE"
	RESUME_TOKEN_VOLUME_BUCKET_NAME = "RESUME_TOKEN_VOLUME_BUCKET_NAME"
	RESUME_TOKEN_FILE_NAME          = "RESUME_TOKEN_FILE_NAME"
	RESUME_TOKEN_BUCKET_REGION      = "RESUME_TOKEN_BUCKET_REGION"
	RESUME_TOKEN_SAVE_INTERVAL_SEC  = "RESUME_TOKEN_SAVE_INTERVAL_SEC"

	EXPORT_DESTINATION                    = "EXPORT_DESTINATION"
	PROJECT_NAME_TO_EXPORT_CHANGE_STREAMS = "PROJECT_NAME_TO_EXPORT_CHANGE_STREAMS"

	TIME_ZONE = "TIME_ZONE"

	LOG_LEVEL            = "LOG_LEVEL"
	LOG_FORMAT           = "LOG_FORMAT"
	LOG_OUTPUT_DIRECTORY = "LOG_OUTPUT_DIRECTORY"
	LOG_OUTPUT_FILE      = "LOG_OUTPUT_FILE"
)

/*

	bqCfg.DataSet = os.Getenv("BIGQUERY_DATASET")
	bqCfg.Table = os.Getenv("BIGQUERY_TABLE")


	ksCfg.StreamName = os.Getenv("KINESIS_STREAM_NAME")
	ksCfg.KinesisStreamRegion = os.Getenv("KINESIS_STREAM_REGION")


	mCfg.MongoDbConnectionUrl = os.Getenv("MONGODB_HOST")
	mCfg.MongoDbDatabase = os.Getenv("MONGODB_DATABASE")
	mCfg.MongoDbCollection = os.Getenv("MONGODB_COLLECTION")

	psCfg.MongoDbDatabase = os.Getenv("MONGODB_DATABASE")
	psCfg.MongoDbCollection = os.Getenv("MONGODB_COLLECTION")


	pvDir, pvDirExistence := os.LookupEnv("RESUME_TOKEN_VOLUME_DIR")


	pvType, pvTypeExistence := os.LookupEnv("RESUME_TOKEN_VOLUME_TYPE")

	expDst, expDstExistence := os.LookupEnv("EXPORT_DESTINATION")


	projectID, projectIDExistence := os.LookupEnv("PROJECT_NAME_TO_EXPORT_CHANGE_STREAMS")


	tz, tzExistence := os.LookupEnv("TIME_ZONE")


	l.Level = os.Getenv("LOG_LEVEL")
	l.Format = os.Getenv("LOG_FORMAT")
	l.OutputDirectory = os.Getenv("LOG_OUTPUT_DIRECTORY")
	l.OutputFile = os.Getenv("LOG_OUTPUT_FILE")
*/
