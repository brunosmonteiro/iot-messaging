package model

type MessageConfig struct {
	Exchange   string
	Connection ConnectionConfig
	Queues     []QueueDefinition
}

type ConnectionConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type QueueDefinition struct {
	Name       string
	RoutingKey []string
}
