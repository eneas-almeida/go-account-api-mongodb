package shared

import (
	"fmt"
	"strings"
)

type Notification struct {
	context string
	errors  []string
}

func CreateNotification(context string) *Notification {
	return &Notification{
		context: context,
	}
}

func (n *Notification) AddError(message string) {
	n.errors = append(n.errors, message)
}

func (n *Notification) HasError() bool {
	return len(n.errors) > 0
}

func (n *Notification) GetErrors() string {
	return fmt.Sprintf("%s %s", n.context, strings.Join(n.errors, ", "))
}
