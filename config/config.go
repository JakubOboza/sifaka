package config

import "os"

const (
	DATABASE_FILE_PATH = "data/sifaka.db"
)

func DatabasePath() string {
	envPath := os.Getenv("DATABASE")
	if envPath == "" {
		return DATABASE_FILE_PATH
	}
	return envPath
}

func SlackOAuth() string {
	return os.Getenv("SLACK_OAUTH_TOKEN")
}

func SlackChannelID() string {
	return os.Getenv("SLACK_CHANNEL_ID")
}

func DummyNotificationsEnabled() bool {
	return os.Getenv("DUMMY_NOTIFICATIONS_ENABLED") != ""
}
