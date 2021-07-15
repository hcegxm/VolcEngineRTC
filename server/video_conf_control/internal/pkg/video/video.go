package video

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/volcengine/VolcEngineRTC/server/video_conf_control/internal/config"
	"github.com/volcengine/VolcEngineRTC/server/video_conf_control/pkg/httpclient"
)

// only choose wanted url, not all param
type response struct {
	Result result `json:"result"`
}

type result struct {
	Data data `json:"Data"`
}

type data struct {
	Status       int        `json:"Status"`
	VideoID      string     `json:"VideoID"`
	PlayInfoList []playInfo `json:"PlayInfoList"`
}

type playInfo struct {
	MainPlayUrl string `json:"MainPlayUrl"`
}

func GetVideoURL(ctx context.Context, vids []string) map[string]string {
	tokenURL := t.get()
	logs.CtxInfo(ctx, "token is: %v", tokenURL)
	res := make(map[string]string)

	for _, vid := range vids {
		code, body, err := httpclient.DoRequest(config.Config.VideoObtainingURL+"?"+tokenURL+"&video_id="+vid, http.MethodGet, nil)
		logs.CtxInfo(ctx, "code %v body: %v, err: %v", code, string(body), err)
		if err != nil || code != http.StatusOK {
			continue
		}

		url := parseURL(body)
		logs.CtxInfo(ctx, "url: %v", url)
		if url != "" {
			res[vid] = url
		}
	}
	return res
}

func parseURL(data []byte) string {
	var res response
	if err := json.Unmarshal(data, &res); err != nil {
		logs.Warnf("unmarshal video info error: %v", err)
		return ""
	}

	return res.Result.Data.PlayInfoList[0].MainPlayUrl
}