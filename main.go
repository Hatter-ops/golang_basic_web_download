package main

//import "net/http"
import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var BASE_URL string

func main() {

	BASE_URL = ""

	make_dirs(BASE_URL)

	resp := get_content(BASE_URL)

	// // bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// // resp.Body.Close() //  must close
	// // body := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	var buf bytes.Buffer
	tee_body := io.TeeReader(resp.Body, &buf)

	make_index(tee_body)

	for _, j := range parse_link_func(&buf) {
		css_url := make_url(j)
		fmt.Println(css_url)
		file_name_parse, err := url.Parse(css_url)

		if err != nil {
			fmt.Println(err)
		}

		file_url_path := file_name_parse.Path
		file_name_split := strings.Split(file_url_path, "/")
		file_name := file_name_split[len(file_name_split)-1]
		file_path := path.Join("./css", file_name)
		css_file := get_content(css_url)
		fmt.Println("Downloading ", file_path, "From: ", css_url)
		content_to_file(css_file, file_path)

	}

	for _, j := range parse_img_func(&buf) {
		image_url := make_url(j)
		file_name_parse, err := url.Parse(image_url)

		if err != nil {
			fmt.Println(err)
		}

		file_url_path := file_name_parse.Path
		file_name_split := strings.Split(file_url_path, "/")
		file_name := file_name_split[len(file_name_split)-1]
		file_path := path.Join("./images", file_name)
		css_file := get_content(image_url)
		fmt.Println("Downloading ", file_path, "From: ", image_url)
		content_to_file(css_file, file_path)
	}

	for _, j := range parse_script_func(&buf) {
		fmt.Println(j)
		script_url := make_url(j)
		fmt.Println(script_url)
		file_name_parse, err := url.Parse(script_url)

		if err != nil {
			fmt.Println(err)
		}

		file_url_path := file_name_parse.Path
		file_name_split := strings.Split(file_url_path, "/")
		file_name := file_name_split[len(file_name_split)-1]
		file_path := path.Join("./javascript", file_name)
		css_file := get_content(script_url)
		fmt.Println("Downloading ", file_path, "From: ", script_url)
		content_to_file(css_file, file_path)
	}

}

func make_dirs(base_url string) {

	parse_url, err := url.Parse(base_url)

	if err != nil {
		fmt.Println(err)
	}

	cur_dir, err := os.Getwd()

	root_dir_name := path.Join(cur_dir, parse_url.Host)
	image_dir_name := path.Join(root_dir_name, "images")
	css_dir_name := path.Join(root_dir_name, "css")
	js_dir_name := path.Join(root_dir_name, "javascript")
	fmt.Println(root_dir_name, image_dir_name, css_dir_name)
	err = os.Mkdir(root_dir_name, 0777)

	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir(image_dir_name, 0777)

	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir(css_dir_name, 0777)

	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir(js_dir_name, 0777)

	if err != nil {
		fmt.Println(err)
	}

	os.Chdir(root_dir_name)

}

func get_content(url string) *http.Response {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != 200 {
		fmt.Println("Error", resp.StatusCode)
		return nil
	}

	if err != nil {
		fmt.Println(err)
	}

	return resp

}

func make_url(passed_url string) string {

	var css_url string

	parse_base_url, err := url.Parse(BASE_URL)

	parse_passed_url, err := url.Parse(passed_url)

	if err != nil {
		fmt.Println(err)
	}

	if !parse_passed_url.IsAbs() {
		base_url_path := parse_base_url.Path

		parse_passed_path := parse_passed_url.Path

		//fmt.Println(parse_passed_path == "")

		if parse_passed_path != "" {
			if string(parse_passed_path[0]) != "/" {
				if filepath.Ext(base_url_path) != "" {
					base_url_path = filepath.Dir(base_url_path)
				}
				parse_passed_path = base_url_path + "/" + parse_passed_path

			}
		}

		if filepath.Ext(base_url_path) != "" {
			base_url_path = filepath.Dir(base_url_path)
		}

		//fmt.Println(base_url_host, base_url_path, parse_passed_host, parse_passed_path, passed_url)

		css_url = fmt.Sprintf("%s://%s%s", parse_base_url.Scheme, parse_base_url.Host, parse_passed_path)
	} else {
		css_url = passed_url
	}

	// fmt.Println(css_url)

	return css_url
}

func make_index(index io.Reader) {

	bytes, err := io.ReadAll(index)

	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create("./index.html")

	if err != nil {
		fmt.Println(err)
	}

	file.Write(bytes)

}

func content_to_file(resp *http.Response, path string) {

	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(body)

	if err != nil {
		fmt.Println(err)
	}

}

func parse_link_func(n io.Reader) []string {
	css := []string{}
	doc, err := goquery.NewDocumentFromReader(n)

	if err != nil {
		fmt.Println(err)
	}

	doc.Find("link[type='text/css']").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("href")
		fmt.Println(val)
		css = append(css, val)
	})

	return css
}

func parse_img_func(n io.Reader) []string {
	img := []string{}
	doc, err := goquery.NewDocumentFromReader(n)

	if err != nil {
		fmt.Println(err)
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s)
		val, _ := s.Attr("src")
		fmt.Println(val)
		img = append(img, val)
	})

	return img
}

func parse_script_func(n io.Reader) []string {
	script := []string{}
	doc, err := goquery.NewDocumentFromReader(n)

	if err != nil {
		fmt.Println(err)
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s)
		val, _ := s.Attr("src")
		fmt.Println(val)
		script = append(script, val)
	})

	return script
}
