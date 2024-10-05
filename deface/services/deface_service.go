package services

import (
	"crypto/md5"
	"deface/config"
	"deface/models"
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/gocolly/colly"
)

func GetWebsiteHTMLHash(url string) (string, error) {
	c := colly.NewCollector()

	var html string
	c.OnHTML("html", func(e *colly.HTMLElement) {
		html = e.Text
	})

	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	hash := fmt.Sprintf("%x", md5.Sum([]byte(html)))
	return hash, nil
}

func TakeScreenshot(url string, filename string) error {
	browser := rod.New().MustConnect()
	page := browser.MustPage(url)
	defer browser.MustClose()

	err := page.MustScreenshot(filename)
	if err != nil {
		return nil
	}

	return nil
}

func CheckWebsiteChanges(url string) error {
	db := config.DB
	var website models.Website

	result := db.First(&website, "url = ?", url)

	newHash, err := GetWebsiteHTMLHash(url)
	if err != nil {
		return err
	}

	if result.RowsAffected > 0 {
		if website.HTMLHash != newHash {
			log.Printf("Website %s has changed!", url)

			screenshotName := fmt.Sprintf("screenshot_%d.png", time.Now().Unix())
			err = TakeScreenshot(url, screenshotName)
			if err != nil {
				return err
			}

			website.HTMLHash = newHash
			website.Screenshot = screenshotName
			website.UpdatedAt = time.Now()

			db.Save(&website)
		} else {
			log.Printf("No changes detected for website %s", url)
		}
	} else {
		// logic insert new website duplicate
		screenshotName := fmt.Sprintf("screenshot_%d.png", time.Now().Unix())
		err = TakeScreenshot(url, screenshotName)
		if err != nil {
			return err
		}

		newWebsite := models.Website{
			URL:        url,
			HTMLHash:   newHash,
			Screenshot: screenshotName,
			UpdatedAt:  time.Now(),
		}
		db.Create(&newWebsite)
		log.Printf("New website %s added to database.", url)
	}

	return nil
}
