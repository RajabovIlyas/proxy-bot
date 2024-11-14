package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/url"
	"strings"
)

func checkUrlsToProxyUrls(message []string) []string {
	urls := make([]string, 0)

	for _, word := range message {
		parsedURL, err := url.Parse(word)

		if err != nil || parsedURL.Scheme != "https" || parsedURL.Host != "t.me" || parsedURL.Path != "/proxy" {
			continue
		}

		urls = append(urls, word)
	}
	return urls
}

func getUrlsFromEntities(message *tgbotapi.Message) []string {
	links := make([]string, 0)

	entities := append(message.Entities, message.CaptionEntities...)

	for _, entity := range entities {
		if entity.Type != "text_link" {
			continue
		}
		links = append(links, entity.URL)
	}
	return links
}

func getUrlsFromInlineKeyboards(message *tgbotapi.Message) []string {
	links := make([]string, 0)
	if message.ReplyMarkup == nil || message.ReplyMarkup.InlineKeyboard == nil {
		return links
	}

	inlineKeyboards := message.ReplyMarkup.InlineKeyboard

	for _, inlineKeyboard := range inlineKeyboards {
		for _, inlineKey := range inlineKeyboard {
			if inlineKey.URL == nil {
				continue
			}
			links = append(links, *inlineKey.URL)
		}
	}

	return links
}

func getUrlsFromMessage(message *tgbotapi.Message) []string {
	text := message.Text

	modMessage := strings.Replace(text, "\n", " ", -1)

	words := strings.Split(modMessage, " ")

	return words
}

func FilterURLs(message *tgbotapi.Message) []string {

	links := make([]string, 0)

	links = append(links, getUrlsFromInlineKeyboards(message)...)
	links = append(links, getUrlsFromEntities(message)...)
	links = append(links, getUrlsFromMessage(message)...)

	return checkUrlsToProxyUrls(links)
}
