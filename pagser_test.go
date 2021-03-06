package pagser

import (
	"fmt"
	"testing"
)

const rawPagserHtml = `
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Pagser Example</title>
	<meta name="keywords" content="golang,pagser,goquery,html,page,parser,colly">
</head>

<body>
	<div class="navlink">
		<div class="container">
			<ul class="clearfix">
				<li id=''><a href="/">Index</a></li>
				<li id='2'><a href="/list/web" title="web site">Web page</a></li>
				<li id='3'><a href="/list/pc" title="pc page">Pc Page</a></li>
				<li id='4'><a href="/list/mobile" title="mobile page">Mobile Page</a></li>
			</ul>
		</div>
	</div>
</body>
</html>
`

type PagserData struct {
	Title    string   `pagser:"title"`
	Keywords []string `pagser:"meta[name='keywords']->attrSplit(content)"`
	Navs     []struct {
		ID   int    `pagser:"->attrEmpty(id, -1)"`
		Name string `pagser:"a->text()"`
		Url  string `pagser:"a->attr(href)"`
	} `pagser:".navlink li"`
}

type ConfigData struct {
	Title    string   `query:"title"`
	Keywords []string `query:"meta[name='keywords']@attrSplit(content)"`
	Navs     []struct {
		ID   int    `query:"@attrEmpty(id, -1)"`
		Name string `query:"a@text()"`
		Url  string `query:"a@attr(href)"`
	} `query:".navlink li"`
}

func TestNew(t *testing.T) {
	p := New()

	var data PagserData
	err := p.Parse(&data, rawPagserHtml)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("json: %v\n", prettyJson(data))
}

func TestNewWithConfig(t *testing.T) {
	cfg := Config{
		TagName:    "query",
		FuncSymbol: "@",
		CastError:  true,
		Debug:      true,
	}
	p, err := NewWithConfig(cfg)
	if err != nil {
		t.Fatal(err)
	}

	var data ConfigData
	err = p.Parse(&data, rawPagserHtml)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("json: %v\n", prettyJson(data))
}

func TestNewWithConfigTagNameError(t *testing.T) {
	cfg := Config{
		TagName:    "",
		FuncSymbol: "->",
		CastError:  true,
		Debug:      true,
	}
	_, err := NewWithConfig(cfg)
	if err != nil {
		t.Log(err)
	} else {
		t.Fatal("Result must return error")
	}
}

func TestNewWithConfigFuncSymbolError(t *testing.T) {
	cfg := Config{
		TagName:    "pagser",
		FuncSymbol: "",
		CastError:  true,
		Debug:      true,
	}
	_, err := NewWithConfig(cfg)
	if err != nil {
		t.Log(err)
	} else {
		t.Fatal("Result must return error")
	}
}
