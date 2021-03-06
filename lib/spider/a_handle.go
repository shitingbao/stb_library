// recheck 待定修改

package spider

import (
	"net/http"

	"golang.org/x/net/html"
)

type aNode struct {
	Src  string
	Node *html.Node
	Resp *http.Response
}

func NewANode(resp *http.Response, n *html.Node) htmlNode {
	return &aNode{
		Node: n,
		Resp: resp,
	}
}

func (a *aNode) Handle(ch chan *imgNode) error {
	if a.Node.Data != aNodeSign {
		return nil
	}
	for _, at := range a.Node.Attr {
		if at.Key != hrefNodeSign {
			continue
		}
		link, err := a.Resp.Request.URL.Parse(at.Val)
		if err != nil {
			return err
		}
		l := link.String()
		// if core.Rds.HGet(nodeSign, l).Val() != "" { // recheck
		// 	return nil
		// }
		// core.Rds.HSet(nodeSign, l, nodeSign)
		SpiderLoad(l, ch)
		return nil
	}
	return nil
}
