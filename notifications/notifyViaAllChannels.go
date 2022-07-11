package notifications

// NotifyViaAllChannels sends notifications to the admin
// via all notification channels (e.g., email and slack).
func NotifyViaAllChannels(message string){
	SendEmailNotification(message)
	SendSlackNotification(message)
}