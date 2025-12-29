package trans

// https://github.com/g17ui/translate/blob/master/src/main.rs

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type Lang struct {
	Code string
	Name string
}

var Langs = []Lang{
	{"af", "Afrikaans"},
	{"sq", "Albanian"},
	{"am", "Amharic"},
	{"ar", "Arabic"},
	{"hy", "Armenian"},
	{"az", "Azerbaijani"},
	{"eu", "Basque"},
	{"be", "Belarusian"},
	{"bn", "Bengali"},
	{"bs", "Bosnian"},
	{"bg", "Bulgarian"},
	{"ca", "Catalan"},
	{"ceb", "Cebuano"},
	{"zh-CN", "Chinese (Simplified)"},
	{"zh-TW", "Chinese (Traditional)"},
	{"co", "Corsican"},
	{"hr", "Croatian"},
	{"cs", "Czech"},
	{"da", "Danish"},
	{"nl", "Dutch"},
	{"en", "English"},
	{"eo", "Esperanto"},
	{"et", "Estonian"},
	{"fi", "Finnish"},
	{"fr", "French"},
	{"fy", "Frisian"},
	{"gl", "Galician"},
	{"ka", "Georgian"},
	{"de", "German"},
	{"el", "Greek"},
	{"gu", "Gujarati"},
	{"ht", "Haitian Creole"},
	{"ha", "Hausa"},
	{"haw", "Hawaiian"},
	{"he", "Hebrew"},
	{"hi", "Hindi"},
	{"hmn", "Hmong"},
	{"hu", "Hungarian"},
	{"is", "Icelandic"},
	{"ig", "Igbo"},
	{"id", "Indonesian"},
	{"ga", "Irish"},
	{"it", "Italian"},
	{"ja", "Japanese"},
	{"jv", "Javanese"},
	{"kn", "Kannada"},
	{"kk", "Kazakh"},
	{"km", "Khmer"},
	{"ko", "Korean"},
	{"ku", "Kurdish"},
	{"ky", "Kyrgyz"},
	{"lo", "Lao"},
	{"la", "Latin"},
	{"lv", "Latvian"},
	{"lt", "Lithuanian"},
	{"lb", "Luxembourgish"},
	{"mk", "Macedonian"},
	{"mg", "Malagasy"},
	{"ms", "Malay"},
	{"ml", "Malayalam"},
	{"mt", "Maltese"},
	{"mi", "Maori"},
	{"mr", "Marathi"},
	{"mn", "Mongolian"},
	{"my", "Myanmar (Burmese)"},
	{"ne", "Nepali"},
	{"no", "Norwegian"},
	{"ny", "Nyanja (Chichewa)"},
	{"or", "Odia (Oriya)"},
	{"ps", "Pashto"},
	{"fa", "Persian"},
	{"pl", "Polish"},
	{"pt", "Portuguese"},
	{"pa", "Punjabi"},
	{"ro", "Romanian"},
	{"ru", "Russian"},
	{"sm", "Samoan"},
	{"gd", "Scots Gaelic"},
	{"sr", "Serbian"},
	{"st", "Sesotho"},
	{"sn", "Shona"},
	{"sd", "Sindhi"},
	{"si", "Sinhala (Sinhalese)"},
	{"sk", "Slovak"},
	{"sl", "Slovenian"},
	{"so", "Somali"},
	{"es", "Spanish"},
	{"su", "Sundanese"},
	{"sw", "Swahili"},
	{"sv", "Swedish"},
	{"tl", "Tagalog (Filipino)"},
	{"tg", "Tajik"},
	{"ta", "Tamil"},
	{"tt", "Tatar"},
	{"te", "Telugu"},
	{"th", "Thai"},
	{"tr", "Turkish"},
	{"tk", "Turkmen"},
	{"uk", "Ukrainian"},
	{"ur", "Urdu"},
	{"ug", "Uyghur"},
	{"uz", "Uzbek"},
	{"vi", "Vietnamese"},
	{"cy", "Welsh"},
	{"xh", "Xhosa"},
	{"yi", "Yiddish"},
	{"yo", "Yoruba"},
	{"zu", "Zulu"},
}

// data имеет такой формат [[["Hello World","Привет, мир",null,null,10]],null,"ru",null,null,null,1,[],[["ru"],null,[1],["ru"]]]

type Translate struct {
	Target     string
	UserAgents []string
}

func New(lang string) *Translate {
	agents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:94.0) Gecko/20100101 Firefox/94.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36",
	}
	return &Translate{
		Target:     lang,
		UserAgents: agents,
	}
}

func (t *Translate) getUserAgent() string {
	if len(t.UserAgents) == 0 {
		return ""
	}
	rand := rand.New(
		rand.NewSource(
			time.Now().UnixNano(),
		),
	)
	return t.UserAgents[rand.Intn(len(t.UserAgents))]
}

func (t *Translate) generateURL(text string, source string) (string, error) {
	// "https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=%s&dt=t&q=%s",
	u, err := url.Parse("https://translate.googleapis.com/translate_a/single")
	if err != nil {
		return "", fmt.Errorf("failed to encode url: %w", err)
	}
	q := u.Query()
	q.Set("client", "gtx")
	q.Set("sl", source)
	q.Set("tl", t.Target)
	q.Set("dt", "t")
	q.Set("q", text)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (t *Translate) Translate(text string, source ...string) (string, []string, error) {
	if len(source) == 0 {
		source = append(source, "auto")
	}
	url, err := t.generateURL(text, source[0])
	if err != nil {
		return "", nil, fmt.Errorf("url error: %w", err)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create req: %w", err)
	}

	req.Header.Set("User-Agent", t.getUserAgent())

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("failed to send http: %w", err)
	}
	defer resp.Body.Close()

	var data any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", nil, fmt.Errorf("failed to parse json: %w", err)
	}

	// [ [ [ "перевод", "оригинал", ... ], ... ], null, "detected_lang", ... ]
	root, ok := data.([]any)
	if !ok || len(root) < 3 {
		return "", nil, fmt.Errorf("invalid json")
	}

	// detected_lang = response[2]
	lang, ok := root[2].(string)
	if !ok {
		return "", nil, fmt.Errorf("failed to detect lang")
	}

	// translations = response[0][i][0]
	first, ok := root[0].([]any)
	if !ok {
		return "", nil, fmt.Errorf("inavlid translate field")
	}

	var translations []string
	for _, item := range first {
		inner, ok := item.([]any)
		if !ok || len(inner) == 0 {
			continue
		}
		if s, ok := inner[0].(string); ok {
			translations = append(translations, s)
		}
	}

	return lang, translations, nil
}
