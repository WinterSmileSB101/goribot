package goribot

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNetIO(t *testing.T) {
	s := NewSpider()
	_ = s.Get(nil, "https://httpbin.org/get?Goribot%20test=hello%20world", func(r *Response) {
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(r.Text), &m)
		if err != nil {
			t.Error("set useragent test", "json load error", err)
		}
		if m["args"].(map[string]interface{})["Goribot test"].(string) != "hello world" {
			fmt.Println(r.Text)
			t.Error("Get test error")
		}
		t.Log("Get test ok")
	})
	_ = s.Post(nil, "https://httpbin.org/post", UrlencodedPostData,
		map[string]string{
			"Goribot test": "hello world",
		},
		func(r *Response) {
			m := make(map[string]interface{})
			err := json.Unmarshal([]byte(r.Text), &m)
			if err != nil {
				t.Error("set useragent test", "json load error", err)
			}
			if m["form"].(map[string]interface{})["Goribot test"].(string) != "hello world" {
				fmt.Println(r.Text)
				t.Error("urlencoded post test error")
			}
			t.Log("UrlencodedPostData test ok")
		})
	_ = s.Post(nil, "https://httpbin.org/post", JsonPostData,
		map[string]string{
			"Goribot test": "hello world",
		},
		func(r *Response) {
			m := make(map[string]interface{})
			err := json.Unmarshal([]byte(r.Text), &m)
			if err != nil {
				t.Error("set useragent test", "json load error", err)
			}
			if m["json"].(map[string]interface{})["Goribot test"].(string) != "hello world" {
				fmt.Println(r.Text)
				t.Error("urlencoded post test error")
			}
			t.Log("JsonPostData test ok")
		})
	s.Run()
}

func TestUaSetting(t *testing.T) {
	s := NewSpider()
	s.UserAgent = "GoRibot test ua"
	_ = s.Get(nil, "https://httpbin.org/user-agent", func(r *Response) {
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(r.Text), &m)
		if err != nil {
			t.Error("set useragent test", "json load error", err)
		}
		if m["user-agent"].(string) != s.UserAgent {
			t.Error(
				"set useragent test error",
				"expected:", "'"+s.UserAgent+"'",
				"got:", "'"+m["user-agent"].(string)+"'")
		}
		t.Log("UA test ok")
	})
	s.Run()
}

func TestHeaderSetting(t *testing.T) {
	s := NewSpider()
	req, err := s.NewGetRequest("https://httpbin.org/headers", func(r *Response) {
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(r.Text), &m)
		if err != nil {
			t.Error("set useragent test", "json load error", err)
		}
		if m["headers"].(map[string]interface{})["Goribot-Test"].(string) != "hello world" ||
			m["headers"].(map[string]interface{})["Cookies"].(string) != "a=1" {
			fmt.Println("TestHeaderSetting", r.Text)
			t.Error("set header test error")
		}
		t.Log("Header test ok")
	})
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("goribot-test", "hello world")
	req.Header.Set("cookies", "a=1")
	s.Crawl(nil, req)
	s.Run()
}
