package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/url"
	"strings"
)

func FilterURLs(message *tgbotapi.Message) []string {
	text := message.Text

	links := make([]string, 0)

	inlineKeyboards := message.ReplyMarkup.InlineKeyboard

	for _, inlineKeyboard := range inlineKeyboards {
		for _, inlineKey := range inlineKeyboard {
			if inlineKey.URL == nil {
				continue
			}
			links = append(links, *inlineKey.URL)
		}
	}

	entities := append(message.Entities, message.CaptionEntities...)

	for _, entity := range entities {
		if entity.Type != "text_link" {
			continue
		}
		links = append(links, entity.URL)
	}

	modMessage := strings.Replace(text, "\n", " ", -1)

	words := strings.Split(modMessage, " ")

	words = append(words, links...)

	urls := make([]string, 0)

	for _, word := range words {
		parsedURL, err := url.Parse(word)

		if err != nil || parsedURL.Scheme != "https" || parsedURL.Host != "t.me" || parsedURL.Path != "/proxy" {
			continue
		}

		urls = append(urls, word)
	}
	return urls
}
