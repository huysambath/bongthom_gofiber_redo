package translate

import (
	// Community Packages
	"fmt"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
	"log"
	"path/filepath"

	// Internal Packages
	"admin-api/pkg/logs"
	error_responses "admin-api/pkg/responses"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

var bundle *goi18n.Bundle

func Init() *error_responses.ErrorResponse {

	// Default language
	bundle = goi18n.NewBundle(language.English)

	// Tell the bundle how to read .yaml files
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// List of all translation files to load
	localeFiles := []string{
		"pkg/i18n/localize/en.yaml",
		"pkg/i18n/localize/km.yaml",
		"pkg/i18n/localize/zh.yaml",
	}

	for _, file := range localeFiles {
		_, err := bundle.LoadMessageFile(filepath.Join(file))
		if err != nil {
			log.Printf("Error loading local file %s: %v", file, err)
			logs.NewCustomLog("translate_error", err.Error(), "error")
			return &error_responses.ErrorResponse{
				MessageID: "ErrorLoadMessage",
				Err:       err,
			}
		} 
	}
	return nil
}

func TranslateWithError(c fiber.Ctx, key string, templateData ...map[string]any) (string, *error_responses.ErrorResponse) {
	// If bundle not load yet, it call Init() automatic
	if bundle == nil {
		if errResp := Init(); errResp != nil {
			logs.NewCustomLog("I18nNotInit", errResp.ErrorString(), "error")
			return "", &error_responses.ErrorResponse{
				MessageID: key,
				Err:       fmt.Errorf("translation service is unavailable"),
			}
		}
	}

	// Default "en"
	lang := c.Get("Accept-Language", "en")

	localizer := goi18n.NewLocalizer(bundle, lang)

	data := map[string]any{}
	if len(templateData) > 0 && templateData[0] != nil {
		data = templateData[0]
	}

	msg, err := localizer.Localize(&goi18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: data,
	})

	//If translate fail return error response
	if err != nil {
		log.Printf("Error localizing message ID %s: %v", key, err)
		logs.NewCustomLog("TranslationNotFound", err.Error(), "error")
		return "", &error_responses.ErrorResponse{
			MessageID: key,
			Err:       fmt.Errorf("Translation not found"),
		}
	}
	return msg, nil
}


// This function return only string
func Translate(c fiber.Ctx, key string) string {
	// If translate fail it return key
	msg, err := TranslateWithError(c, key)
	if err != nil {
		return key
	}
	return msg
}
