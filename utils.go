package main

//all other util function go here

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// ShowNotification is a utility function that shows either a system notification or an in-app dialog.
// If useSystemNotification is true, it sends a system notification; otherwise, it shows an in-app dialog.
func ShowNotification(app fyne.App, window fyne.Window, title, message string, useSystemNotification bool) {
	if useSystemNotification {
		// Send a system notification
		app.SendNotification(&fyne.Notification{
			Title:   title,
			Content: message,
		})
	} else {
		// Show an in-app dialog notification
		dialog.ShowInformation(title, message, window)
	}
}
