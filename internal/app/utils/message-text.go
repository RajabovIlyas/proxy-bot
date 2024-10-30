package utils

import "fmt"

func GetMessageText(params ProxyParams, channelName string) string {
	return fmt.Sprintf("<b>🇹🇲 Proxy for you 🇹🇲</b> \n\n"+
		"<b>Server :</b> <code>%s</code>\n"+
		"<b>Port :</b> <code>%s</code>\n"+
		"<b>Secret :</b> <code>%s</code>\n\n"+
		"<b>More :</b> %s",
		params.Server,
		params.Port,
		params.Secret,
		channelName,
	)
}
