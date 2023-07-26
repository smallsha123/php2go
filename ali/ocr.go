package ali

import (
	ocr "github.com/alibabacloud-go/ocr-api-20210707/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func (s *BaseSvc) Ocr(url string) (*ocr.RecognizeBusinessLicenseResponse, error) {
	client, _ := ocr.NewClient(s.config)
	recognizeBusinessLicenseRequest := &ocr.RecognizeBusinessLicenseRequest{Url: &url}
	runtime := &util.RuntimeOptions{}
	resp, tryErr := func() (resp *ocr.RecognizeBusinessLicenseResponse, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		return client.RecognizeBusinessLicenseWithOptions(recognizeBusinessLicenseRequest, runtime)
	}()

	return resp, tryErr
}
