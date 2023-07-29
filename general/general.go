package general

type CheckCategory int

const (
	LOGGING CheckCategory = iota + 1
	ENCRYPTION
	GENERAL_SECURITY
	NETWORKING
	IAM
	BACKUP_AND_RECOVERY
	CONVENTION
	SECRETS
	KUBERNETES
	APPLICATION_SECURITY
	SUPPLY_CHAIN
	API_SECURITY
)
