package presto_client

// PrestoClient describes a interface for querying Presto.
type PrestoClient interface {
	Healthcheck() error

	WorkersCount() (int, error)
}
