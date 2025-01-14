package rod_test

import (
	"errors"
	"io/ioutil"
	"mime"
	"net/http"
	"sync"

	"github.com/7nikhilkamboj/rod"
	"github.com/7nikhilkamboj/rod/lib/proto"
	"github.com/7nikhilkamboj/rod/lib/utils"
	"github.com/ysmood/gson"
)

func (t T) Hijack() {
	s := t.Serve()

	// to simulate a backend server
	s.Route("/", slash("fixtures/fetch.html"))
	s.Mux.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			panic("wrong http method")
		}

		t.Eq("header", r.Header.Get("Test"))

		b, err := ioutil.ReadAll(r.Body)
		t.E(err)
		t.Eq("a", string(b))

		t.HandleHTTP(".html", "test")(w, r)
	})
	s.Route("/b", "", "b")

	router := t.page.HijackRequests()
	defer router.MustStop()

	router.MustAdd(s.URL("/a"), func(ctx *rod.Hijack) {
		r := ctx.Request.SetContext(t.Context())
		r.Req().Header.Set("Test", "header") // override request header
		r.SetBody([]byte("test"))            // override request body
		r.SetBody(123)                       // override request body
		r.SetBody(r.Body())                  // override request body

		type MyState struct {
			Val int
		}

		ctx.CustomState = &MyState{10}

		t.Eq(http.MethodPost, r.Method())
		t.Eq(s.URL("/a"), r.URL().String())

		t.Eq(proto.NetworkResourceTypeXHR, ctx.Request.Type())
		t.Is(ctx.Request.IsNavigation(), false)
		t.Has(ctx.Request.Header("Origin"), s.URL())
		t.Len(ctx.Request.Headers(), 6)
		t.True(ctx.Request.JSONBody().Nil())

		// send request load response from real destination as the default value to hijack
		ctx.MustLoadResponse()

		t.Eq(200, ctx.Response.Payload().ResponseCode)

		// override status code
		ctx.Response.Payload().ResponseCode = http.StatusCreated

		t.Eq("4", ctx.Response.Headers().Get("Content-Length"))
		t.Has(ctx.Response.Headers().Get("Content-Type"), "text/html; charset=utf-8")

		// override response header
		ctx.Response.SetHeader("Set-Cookie", "key=val")

		// override response body
		ctx.Response.SetBody([]byte("test"))
		ctx.Response.SetBody("test")
		ctx.Response.SetBody(map[string]string{
			"text": "test",
		})

		t.Eq("{\"text\":\"test\"}", ctx.Response.Body())
	})

	router.MustAdd(s.URL("/b"), func(ctx *rod.Hijack) {
		panic("should not come to here")
	})
	router.MustRemove(s.URL("/b"))

	router.MustAdd(s.URL("/b"), func(ctx *rod.Hijack) {
		// transparent proxy
		ctx.MustLoadResponse()
	})

	go router.Run()

	t.page.MustNavigate(s.URL())

	t.Eq("201 test key=val", t.page.MustElement("#a").MustText())
	t.Eq("b", t.page.MustElement("#b").MustText())
}

func (t T) HijackContinue() {
	s := t.Serve().Route("/", ".html", `<body>ok</body>`)

	router := t.page.HijackRequests()
	defer router.MustStop()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	router.MustAdd(s.URL("/a"), func(ctx *rod.Hijack) {
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
		wg.Done()
	})

	go router.Run()

	t.page.MustNavigate(s.URL("/a"))

	t.Eq("ok", t.page.MustElement("body").MustText())
	wg.Wait()
}

func (t T) HijackMockWholeResponse() {
	router := t.page.HijackRequests()
	defer router.MustStop()

	router.MustAdd("*", func(ctx *rod.Hijack) {
		ctx.Response.SetHeader("Content-Type", mime.TypeByExtension(".html"))
		ctx.Response.SetBody("<body>ok</body>")
	})

	go router.Run()

	t.page.MustNavigate("http://test.com")

	t.Eq("ok", t.page.MustElement("body").MustText())
}

func (t T) HijackSkip() {
	s := t.Serve()

	router := t.page.HijackRequests()
	defer router.MustStop()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	router.MustAdd(s.URL("/a"), func(ctx *rod.Hijack) {
		ctx.Skip = true
		wg.Done()
	})
	router.MustAdd(s.URL("/a"), func(ctx *rod.Hijack) {
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
		wg.Done()
	})

	go router.Run()

	t.page.MustNavigate(s.URL("/a"))

	wg.Wait()
}

func (t T) HijackOnErrorLog() {
	s := t.Serve().Route("/", ".html", `<body>ok</body>`)

	router := t.page.HijackRequests()
	defer router.MustStop()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	var err error

	router.MustAdd(s.URL("/a"), func(ctx *rod.Hijack) {
		ctx.OnError = func(e error) {
			err = e
			wg.Done()
		}
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})

	go router.Run()

	t.mc.stub(1, proto.FetchContinueRequest{}, func(send StubSend) (gson.JSON, error) {
		return gson.New(nil), errors.New("err")
	})

	go func() {
		_ = t.page.Context(t.Context()).Navigate(s.URL("/a"))
	}()
	wg.Wait()

	t.Eq(err.Error(), "err")
}

func (t T) HijackFailRequest() {
	s := t.Serve().Route("/page", ".html", `<html>
	<body></body>
	<script>
		fetch('/a').catch(async (err) => {
			document.title = err.message
		})
	</script></html>`)

	router := t.browser.HijackRequests()
	defer router.MustStop()

	router.MustAdd(s.URL("/a"), func(ctx *rod.Hijack) {
		ctx.Response.Fail(proto.NetworkErrorReasonAborted)
	})

	go router.Run()

	t.page.MustNavigate(s.URL("/page")).MustWaitLoad()

	t.page.MustWait(`document.title === 'Failed to fetch'`)

	{ // test error log
		t.mc.stub(1, proto.FetchFailRequest{}, func(send StubSend) (gson.JSON, error) {
			_, _ = send()
			return gson.JSON{}, errors.New("err")
		})
		_ = t.page.Navigate(s.URL("/a"))
	}
}

func (t T) HijackLoadResponseErr() {
	p := t.newPage().Context(t.Context())
	router := p.HijackRequests()
	defer router.MustStop()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	router.MustAdd("http://test.com/a", func(ctx *rod.Hijack) {
		t.Err(ctx.LoadResponse(&http.Client{
			Transport: &MockRoundTripper{err: errors.New("err")},
		}, true))

		t.Err(ctx.LoadResponse(&http.Client{
			Transport: &MockRoundTripper{res: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(&MockReader{err: errors.New("err")}),
			}},
		}, true))

		wg.Done()

		ctx.Response.Fail(proto.NetworkErrorReasonAborted)
	})

	go router.Run()

	_ = p.Navigate("http://test.com/a")

	wg.Wait()
}

func (t T) HijackResponseErr() {
	s := t.Serve().Route("/", ".html", `ok`)

	p := t.newPage().Context(t.Context())
	router := p.HijackRequests()
	defer router.MustStop()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	router.MustAdd(s.URL("/a"), func(ctx *rod.Hijack) { // to ignore favicon
		ctx.OnError = func(err error) {
			t.Err(err)
			wg.Done()
		}

		ctx.MustLoadResponse()
		t.mc.stub(1, proto.FetchFulfillRequest{}, func(send StubSend) (gson.JSON, error) {
			res, _ := send()
			return res, errors.New("err")
		})
	})

	go router.Run()

	p.MustNavigate(s.URL("/a"))

	wg.Wait()
}

func (t T) HandleAuth() {
	s := t.Serve()

	// mock the server
	s.Mux.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			w.Header().Add("WWW-Authenticate", `Basic realm="web"`)
			w.WriteHeader(401)
			return
		}

		t.Eq("a", u)
		t.Eq("b", p)
		t.HandleHTTP(".html", `<p>ok</p>`)(w, r)
	})
	s.Route("/err", ".html", "err page")

	go t.browser.MustHandleAuth("a", "b")()

	page := t.newPage(s.URL("/a"))
	page.MustElementR("p", "ok")

	wait := t.browser.HandleAuth("a", "b")
	var page2 *rod.Page
	wait2 := utils.All(func() {
		page2, _ = t.browser.Page(proto.TargetCreateTarget{URL: s.URL("/err")})
	})
	t.mc.stubErr(1, proto.FetchContinueRequest{})
	t.Err(wait())
	wait2()
	page2.MustClose()
}
