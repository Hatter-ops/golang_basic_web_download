package main

//import "net/http"
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

	// Get url from command line arguments

	BASE_URL = os.Args[1]

	// Makes dirs where content is downloaded to.
	// root dir is the host name from the url inside it will have css, images and javascript dirs

	make_dirs(BASE_URL)

	// Gets content using http request and returns http.Response type
	resp := get_content(BASE_URL)

	// Reads all of response Body
	raw_body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	// This moves the steam to the beginnig
	resp.Body = io.NopCloser(bytes.NewBuffer(raw_body))

	// =========================================================================
	// Parse conent reads the HTML and finds all of the content with the given element and attribute
	for _, j := range parse_content(resp.Body, "link[type='text/css']", "href") {
		// Sites have multiple types of url so this function is to make them all the same
		css_url := make_url(j)

		//======================================================
		// This block of code is for getting the name for the file that
		// we want to download.
		// It basically parses the name from the url and get the last element
		// by spliting in / and getting the last element.
		file_name_parse, err := url.Parse(css_url)

		if err != nil {
			fmt.Println(err)
		}

		file_url_path := file_name_parse.Path
		file_name_split := strings.Split(file_url_path, "/")
		file_name := file_name_split[len(file_name_split)-1]
		file_path := path.Join("./css", file_name)
		//======================================================
		css_file := get_content(css_url)
		fmt.Println("Downloading ", file_path, "From: ", css_url)
		// This function downloads the given content to a file with the given file name.
		content_to_file(css_file, file_path)

	}
	// =========================================================================

	// This is the same as the last block of code
	resp.Body = io.NopCloser(bytes.NewBuffer(raw_body))
	fmt.Println(resp.Body)
	for _, j := range parse_content(resp.Body, "img", "src") {
		fmt.Println("img")
		image_url := make_url(j)
		fmt.Println("Raw url ", j, "Made url ", image_url)
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

	// This is the same as the last block of code
	resp.Body = io.NopCloser(bytes.NewBuffer(raw_body))
	for _, j := range parse_content(resp.Body, "script", "src") {
		fmt.Println("scripts")
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
	resp.Body = io.NopCloser(bytes.NewBuffer(raw_body))
	make_index(resp.Body)
}

func make_dirs(base_url string) {

	parse_url, err := url.Parse(base_url)

	if err != nil {
		fmt.Println(err)
	}

	// Makes all the paths for the dir
	cur_dir, err := os.Getwd()
	root_dir_name := path.Join(cur_dir, parse_url.Host)
	image_dir_name := path.Join(root_dir_name, "images")
	css_dir_name := path.Join(root_dir_name, "css")
	js_dir_name := path.Join(root_dir_name, "javascript")

	// Makes all the needed dir.
	// Root dir from url, images, css, javascript
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

// Function get the content from the given url.
func get_content(url string) *http.Response {
	// Makes new Get request client with the given url
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}
	// Sets the User-Agent to a Normal chrome one so the request get through possible filters
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")

	// Makes the reqeust
	resp, err := http.DefaultClient.Do(req)

	// if the responce is not 200 return nothing
	if resp.StatusCode != 200 {
		fmt.Println("Error", resp.StatusCode)
		return nil
	}

	if err != nil {
		fmt.Println(err)
	}

	// Returns http response
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

		if parse_passed_url.Host != "" {
			css_url = fmt.Sprintf("https:%s", passed_url)
			return css_url
		}

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

		css_url = fmt.Sprintf("%s://%s%s", parse_base_url.Scheme, parse_base_url.Host, parse_passed_path)
	} else {
		css_url = passed_url
	}

	// fmt.Println(css_url)

	return css_url
}

// Creates and modifies the html document and writes it to index.html file.
// Takes the reader as only arg.
func make_index(index io.ReadCloser) {
	// Init new doc from the given reader
	doc, err := goquery.NewDocumentFromReader(index)

	fmt.Println(doc)

	if err != nil {
		fmt.Println(err)
	}

	// Changes the urls in the html document to the local one.
	change_attr(doc, "img", "src", "images")
	change_attr(doc, "link", "href", "css")
	change_attr(doc, "script", "src", "javascrip")

	// Create index file
	file, err := os.Create("./index.html")

	if err != nil {
		fmt.Println(err)
	}

	// Change document to string
	ht, err := doc.Html()

	if err != nil {
		fmt.Println(err)
	}

	// Write the string document to the index file
	file.WriteString(ht)

}

// Function changes the given attribute in the element to the local file path
// so the content is displayed in the local site.
func change_attr(doc *goquery.Document, find_by_element string, att string, dir string) {
	doc.Find(find_by_element).Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr(att)
		file_name_parse, err := url.Parse(val)
		if err != nil {
			fmt.Println(err)
		}
		file_url_path := file_name_parse.Path
		file_name_split := strings.Split(file_url_path, "/")
		file_name := file_name_split[len(file_name_split)-1]
		file_path := path.Join(dir, file_name)
		s.SetAttr(att, file_path)
	})
}

// Get http responc and the paht where the file will be saved
func content_to_file(resp *http.Response, path string) {

	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
	}

	// Reads the http response content, creates the file and writes the body to the file
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

// Parses content form the given html document.
func parse_content(n io.ReadCloser, find_by_element string, att string) []string {
	content := []string{}
	// Init new doc from the given reader
	doc, err := goquery.NewDocumentFromReader(n)

	if err != nil {
		fmt.Println(err)
	}
	// Finds all element with the given tag and for each found element we get the attribute
	// The path is then appended to the array.
	doc.Find(find_by_element).Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr(att)
		fmt.Println(val)
		content = append(content, val)
	})

	return content
}
