package config

type Config struct {
	S3 S3Config `yml:"s3" mapstructure:"s3"`
	Notifications NotificationsConfig `yml:"notifications" mapstructure:"notifications"`
}

type S3Config struct {
	NumberOfStoredBackups int `yaml:"number_of_stored_backups" mapstructure:"number_of_stored_backups"`
	DirectoryPrefix string `yaml:"directory_prefix" mapstructure:"directory_prefix"`
}

type NotificationsConfig struct {
	Slack NotificationChannelConfig `yml:"slack" mapstructure:"slack"`
	Email NotificationChannelConfig `yml:"email" mapstructure:"email"`
}

type NotificationChannelConfig struct {
	Enabled bool `yml:"enabled" mapstructure:"enabled"`
	SendOnSuccess bool `yml:"send_on_success" mapstructure:"send_on_success"`
}