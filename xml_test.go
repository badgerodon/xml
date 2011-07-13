package xml

import (
	"fmt"
	"testing"
)

func TestXML(t *testing.T) {
	e := E("rss", A("version", "2.0"),
		E("channel",
			E("title", T("RSS Title with < in it ")),
			E("description", T("This is an example of an RSS feed")),
			E("link", T("http://www.someexamplerssdomain.com/main.html")),
			E("lastBuildDate", T("Mon, 06 Sep 2010 00:01:00 +0000")),
			E("pubDate", T("Mon, 06 Sep 2009 16:45:00 +0000")),
		),
	)
	
	fmt.Println(e)
}
