//go:build windows

package main

import (
	"fmt"
	"os/exec"
)

func notify(title, body string) {
	bin, err := exec.LookPath("powershell")
	if err != nil {
		bin, err = exec.LookPath("powershell.exe")
		if err != nil {
			return
		}
	}

	script := fmt.Sprintf(`[void][Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime]
$xml = [Windows.UI.Notifications.ToastNotificationManager]::GetTemplateContent([Windows.UI.Notifications.ToastTemplateType]::ToastText02)
$text = $xml.GetElementsByTagName("text")
$text.Item(0).AppendChild($xml.CreateTextNode(%q)) | Out-Null
$text.Item(1).AppendChild($xml.CreateTextNode(%q)) | Out-Null
$toast = [Windows.UI.Notifications.ToastNotification]::new($xml)
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier(%q).Show($toast)`, title, body, title)

	cmd := exec.Command(bin, "-NoProfile", "-NonInteractive", "-Command", script)
	_ = cmd.Run()
}
