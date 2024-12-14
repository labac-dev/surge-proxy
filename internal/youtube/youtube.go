package youtube

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/labac-dev/surge-proxy/internal/common"
	"github.com/labac-dev/surge-proxy/internal/youtube/gen/response/browse"
	"github.com/labac-dev/surge-proxy/internal/youtube/gen/response/next"
	"github.com/labac-dev/surge-proxy/internal/youtube/gen/response/search"
	"github.com/labac-dev/surge-proxy/internal/youtube/gen/response/watch"
)

func ModifyResponse(res *http.Response) error {
	if res.Request.URL.Path == "/youtubei/v1/browse" {
		return wrappedResponse(res, cleanBrowse)
	}

	if res.Request.URL.Path == "/youtubei/v1/search" {
		return wrappedResponse(res, cleanSearch)
	}

	if res.Request.URL.Path == "/youtubei/v1/get_watch" {
		return wrappedResponse(res, cleanWatch)
	}

	if res.Request.URL.Path == "/youtubei/v1/next" {
		return wrappedResponse(res, cleanNext)
	}

	return nil
}

func wrappedResponse(res *http.Response, cleaner func([]byte) ([]byte, error)) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	body, err = cleaner(body)
	if err != nil {
		return err
	}

	return common.Response(res, body)
}

// Remove the Shorts from the Browse response
func cleanBrowse(body []byte) ([]byte, error) {
	data := &browse.Browse{}

	err := proto.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	for _, tab := range data.GetContent().GetSingleColumnResultsRenderer().GetTabs() {
		sectionList := tab.GetTabRenderer().GetContent().GetSectionListRenderer().GetSectionListSupportedRenderers()
		cleanSectionListSupportedRenderers(sectionList)
	}

	sectionList := data.GetOnResponseReceivedAction().GetSectionListRenderer().GetSectionListSupportedRenderers()
	cleanSectionListSupportedRenderers(sectionList)

	return proto.Marshal(data)
}

// Remove the Shorts from the Search response
func cleanSearch(body []byte) ([]byte, error) {
	data := &search.Search{}

	err := proto.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	var (
		content          = data.GetContent().GetSectionListRenderer().GetSectionListSupportedRenderers()
		responseReceived = data.GetOnResponseReceivedCommand().GetAppendContinuationItemsAction().GetSectionListSupportedRenderers()
	)

	cleanSectionListSupportedRenderers(content)
	cleanSectionListSupportedRenderers(responseReceived)

	return proto.Marshal(data)
}

// Remove the Shorts from the Watch response
func cleanWatch(body []byte) ([]byte, error) {
	data := &watch.Watch{}

	err := proto.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	for _, content := range data.GetContents() {
		var (
			next             = content.GetNext()
			content          = next.GetContent().GetNextResult().GetContent().GetSectionListRenderer().GetSectionListSupportedRenderers()
			responseReceived = next.GetOnResponseReceivedAction().GetSectionListRenderer().GetSectionListSupportedRenderers()
		)

		cleanSectionListSupportedRenderers(content)
		cleanSectionListSupportedRenderers(responseReceived)
	}

	return proto.Marshal(data)
}

// Remove the Short from the Next response
func cleanNext(body []byte) ([]byte, error) {
	data := &next.Next{}

	err := proto.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	var (
		content          = data.GetContent().GetNextResult().GetContent().GetSectionListRenderer().GetSectionListSupportedRenderers()
		responseReceived = data.GetOnResponseReceivedAction().GetSectionListRenderer().GetSectionListSupportedRenderers()
		relatedContent   = data.GetOnResponseReceivedAction().GetRelatedChipContinuation().GetSectionListSupportedRenderers()
	)

	cleanSectionListSupportedRenderers(content)
	cleanSectionListSupportedRenderers(responseReceived)
	cleanSectionListSupportedRenderers(relatedContent)

	return proto.Marshal(data)
}

// Support functions
func cleanSectionListSupportedRenderers(sectionList []*browse.SectionListSupportedRenderer) {
	for i, section := range sectionList {
		if section.GetShelfRenderer() != nil {
			sectionList[i] = &browse.SectionListSupportedRenderer{}
		}

		cleanRichItemContent(section.GetItemSectionRenderer().GetRichItemContent())
	}
}

func cleanRichItemContent(richItemList []*browse.RichItemContent) {
	for i, item := range richItemList {
		var (
			videoRenderer = item.GetVideoWithContextRenderer().GetVideoRendererContent()
			layoutEml     = videoRenderer.GetRenderInfo().GetLayoutRender().GetEml()
			videoContext  = videoRenderer.GetVideoInfo().GetVideoContext()
			videoEml      = videoContext.GetVideoRender().GetEml()
			payload       = videoContext.GetVideoContent().GetPayload()
		)

		if strings.Contains(layoutEml, "shorts") || strings.Contains(layoutEml, "shelf") {
			richItemList[i] = &browse.RichItemContent{}
		}

		if strings.Contains(videoEml, "shorts") || strings.Contains(videoEml, "shelf") {
			richItemList[i] = &browse.RichItemContent{}
		}

		if bytes.Contains(bytes.ToLower(payload), []byte("shorts")) {
			richItemList[i] = &browse.RichItemContent{}
		}
	}
}
