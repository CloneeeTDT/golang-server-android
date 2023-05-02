package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"golang-server-android/config"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DictionaryResponse []struct {
	Word      string `json:"word"`
	Phonetics []struct {
		Audio     string `json:"audio"`
		SourceURL string `json:"sourceUrl,omitempty"`
		License   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"license,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Synonyms   []any  `json:"synonyms"`
			Antonyms   []any  `json:"antonyms"`
			Example    string `json:"example"`
		} `json:"definitions"`
		Synonyms []string `json:"synonyms"`
		Antonyms []any    `json:"antonyms"`
	} `json:"meanings"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
	SourceUrls []string `json:"sourceUrls"`
}

type GoogleAPI struct {
	Sentences []struct {
		Trans              string `json:"trans"`
		Orig               string `json:"orig"`
		Backend            int    `json:"backend"`
		ModelSpecification []struct {
		} `json:"model_specification"`
		TranslationEngineDebugInfo []struct {
			ModelTracking struct {
				CheckpointMd5 string `json:"checkpoint_md5"`
				LaunchDoc     string `json:"launch_doc"`
			} `json:"model_tracking"`
		} `json:"translation_engine_debug_info"`
	} `json:"sentences"`
	Src        string  `json:"src"`
	Confidence float64 `json:"confidence"`
	Spell      struct {
	} `json:"spell"`
	LdResult struct {
		Srclangs            []string  `json:"srclangs"`
		SrclangsConfidences []float64 `json:"srclangs_confidences"`
		ExtendedSrclangs    []string  `json:"extended_srclangs"`
	} `json:"ld_result"`
}

type Speech2TextGoogleResponse struct {
	Results []struct {
		Alternatives []struct {
			Transcript string  `json:"transcript"`
			Confidence float64 `json:"confidence"`
		} `json:"alternatives"`
		ResultEndTime string `json:"resultEndTime"`
		LanguageCode  string `json:"languageCode"`
	} `json:"results"`
	TotalBilledTime string `json:"totalBilledTime"`
	RequestId       string `json:"requestId"`
}

type Speech2TextGoogleBody struct {
	Audio struct {
		Content string `json:"content"`
	} `json:"audio"`
	Config struct {
		EnableAutomaticPunctuation bool   `json:"enableAutomaticPunctuation"`
		Encoding                   string `json:"encoding"`
		LanguageCode               string `json:"languageCode"`
		Model                      string `json:"model"`
		SampleRateHertz            int    `json:"sampleRateHertz"`
	} `json:"config"`
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func getWordFromDAPI(word string) *DictionaryResponse {
	client := httpClient()
	uri := "https://api.dictionaryapi.dev/api/v2/entries/en/" + url.QueryEscape(word)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// cast response body to struct
	var dictionaryResponse DictionaryResponse
	err = json.NewDecoder(res.Body).Decode(&dictionaryResponse)
	if err != nil {
		return nil
	}
	return &dictionaryResponse
}

func getWordFromGoogleAPI(word string) *GoogleAPI {
	client := httpClient()
	uri := "https://translate.google.com/translate_a/single?client=at&dt=t&dt=rm&dj=1&sl=en&tl=vi&q=" + url.QueryEscape(word)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// cast response body to struct
	var googleAPIResponse GoogleAPI
	err = json.NewDecoder(res.Body).Decode(&googleAPIResponse)
	if err != nil {
		return nil
	}
	return &googleAPIResponse
}

func GetExamples(word string) []string {
	var result []string
	dictionaryResponse := getWordFromDAPI(word)
	for _, response := range *dictionaryResponse {
		for _, meaning := range response.Meanings {
			for _, definition := range meaning.Definitions {
				if definition.Example != "" {
					result = append(result, definition.Example)
					break
				}
			}
		}
	}
	return result
}

func GetTranslate(from string, to string, text string) *GoogleAPI {
	if strings.Contains(from, "-") || strings.Contains(to, "-") {
		from = strings.Split(from, "-")[0]
		to = strings.Split(to, "-")[0]
	}
	client := httpClient()
	uri := fmt.Sprintf("https://translate.google.com/translate_a/single?client=at&dt=t&dt=rm&dj=1&sl=%s&tl=%s&q=%s", from, to, url.QueryEscape(text))
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// cast response body to struct
	var googleResponse GoogleAPI
	err = json.NewDecoder(res.Body).Decode(&googleResponse)
	if err != nil {
		return nil
	}
	return &googleResponse
}

func SearchWord(word string) *GoogleAPI {
	return getWordFromGoogleAPI(word)
}

func GetAudio(tl string, text string) string {
	uri := fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&client=tw-ob&tl=%s&q=%s", tl, url.QueryEscape(text))
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ""
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	b64encoded := base64.StdEncoding.EncodeToString(buf)

	result := fmt.Sprintf("data:audio/mpeg;base64, %s", b64encoded)
	return result
}

func GetTextFromAudio(audio string, langCode string) (*Speech2TextGoogleResponse, error) {
	vConfig := config.GetConfig()
	apiKey := vConfig.GetString("google.key")
	body := Speech2TextGoogleBody{}
	body.Audio.Content = audio
	body.Config.EnableAutomaticPunctuation = true
	body.Config.Encoding = "OGG_OPUS"
	body.Config.LanguageCode = langCode
	body.Config.Model = "default"
	jsonBody, err := json.Marshal(body)
	client := httpClient()
	uri := fmt.Sprintf("https://speech.googleapis.com/v1p1beta1/speech:recognize?key=%s", apiKey)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// cast response body to struct
	var googleAPIResponse Speech2TextGoogleResponse
	if res.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(bodyBytes))
	}
	err = json.NewDecoder(res.Body).Decode(&googleAPIResponse)
	if err != nil {
		return nil, err
	}
	return &googleAPIResponse, nil
}
