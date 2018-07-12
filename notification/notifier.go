package notification

import "context"

// Notifier receives and notify a line.
type Notifier interface {
	// Notify notifies a line event.
	Notify(ctx context.Context, line string) error
}

type notifier []Notifier

// NewNotifier creates a notifier instance from list of sub notifiers.
func NewNotifier(clients []Notifier) notifier {
	return notifier(clients)
}

// Notify notifies a line event, returns error on first error occur.
func (n notifier) Notify(ctx context.Context, line string) error {
	for _, sub := range n {
		if err := sub.Notify(ctx, line); err != nil {
			return err
		}
	}

	return nil
}
