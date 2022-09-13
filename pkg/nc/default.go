package nc

var (
	shared = NewNotificationCenter()
)

func Default() *NotificationsCenter {
	return shared
}
