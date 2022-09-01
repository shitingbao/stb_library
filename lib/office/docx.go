package office

import (
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey("9be68b870d264caf7842deb88abbe70719fb27dcdc389ba826909ac696c44981")
	if err != nil {
		panic(err)
	}
}

// CreateDocx
// 写入文件名，页眉，内容段落，生成一个 docx 文档
// filename 为完整的路径文件名称，eg ： D://test/fl/aa.docx
func CreateDocx(filename, title string, contents []string) {
	doc := document.New()
	defer doc.Close()

	for _, content := range contents {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(content)
	}

	hdr := doc.AddHeader()
	para := hdr.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	run := para.AddRun()
	// run.AddTab()
	run.AddText(title)

	para = doc.AddParagraph()
	section := para.Properties().AddSection(wml.ST_SectionMarkNextPage)
	section.SetHeader(hdr, wml.ST_HdrFtrDefault)

	doc.SaveToFile(filename)
}
