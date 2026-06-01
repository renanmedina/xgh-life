package translation

type TranslationCollection map[string]string

var TRANSLATIONS = map[string]TranslationCollection{
	"pt-BR": TranslationCollection{
		"footer_text": "Desenvolvido e mantido por ",
	},
	"en-US": TranslationCollection{
		"footer_text": "Developed and maintained by ",
	},
}

const DEFAULT_LANGUAGE = "pt-BR"

func Get(language string, key string) string {
	translations, ok := TRANSLATIONS[language]
	if !ok {
		return TRANSLATIONS[DEFAULT_LANGUAGE][key]
	}

	return translations[key]
}
