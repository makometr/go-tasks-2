package main

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

func main() {
	// uri := "https://sun9-7.userapi.com/impg/P5qwKdi3VEJvVXv6KvTrTs4ij9s9hHeyAMWr6w/d0lgFdt9UOY.jpg?size=806x1080&quality=96&sign=8d98653ac22de3ebf2b2f7864776df88&type=album"
	// uri := "http://example.com/download.mp4"
	uri := "https://www.iana.org/domains/reserved"

	c := newDownloadCounter()
	c.download(uri, 2)

}

type counter struct {
	names map[string]struct{}
	m     *sync.Mutex
}

func newDownloadCounter() *counter {
	return &counter{names: make(map[string]struct{}), m: &sync.Mutex{}}
}

func (c *counter) isNew(uri string) bool {
	c.m.Lock()
	defer c.m.Unlock()
	if _, found := c.names[uri]; found {
		return false
	}
	c.names[uri] = struct{}{}
	return true
}

func (c *counter) download(uri string, maxLevel int) {
	fmt.Println("dowload::", uri, maxLevel)
	if maxLevel < 0 {
		return
	}
	if _, e := url.ParseRequestURI(uri); e != nil {
		fmt.Printf("invalid url %v\n", uri)
		return
	}

	if !c.isNew(uri) {
		return
	}

	res, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
		return

	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	filename := parseFilename(uri, res.Header)
	out, err := os.Create(filename)
	if err != nil {
		return
	}
	defer out.Close()

	fmt.Println("start")
	size, _ := out.Write(data)
	fmt.Printf("finished, size: %s. \n", byteUnitString(int64(size)))

	links := getSubLinks(data)
	var wg sync.WaitGroup
	for _, sublink := range links {
		wg.Add(1)
		go func(link string, level int) {
			defer wg.Done()
			c.download(link, level)
		}(sublink, maxLevel-1)
	}
	wg.Wait()

}

func getSubLinks(data []byte) []string {
	re := regexp.MustCompile(`(http|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	result := re.FindAll(data, -1)

	subUris := make([]string, len(result))
	for i := 0; i < len(result); i++ {
		subUris = append(subUris, string(result[i]))
		// fmt.Println(string(result[i]))
	}
	return subUris
}

func getFilename(url string, mediaType string) string {
	n := path.Base(url)

	if mediaType == "" && (n == "" || n == ".") {
		return "goget-download"
	}

	name1 := cutAfter(n, "#")
	name := cutAfter(name1, "?")

	if path.Ext(name) == "" && mediaType != "" {
		return name + "." + mediaType
	}

	return name
}

func parseFilename(uri string, header http.Header) string {
	contentType := header.Get("Content-Type")
	mimeType, _, _ := mime.ParseMediaType(contentType)
	mediaType := cutBefore(mimeType, "/")

	filename := getFilename(uri, mediaType)

	return filename
}

func cutAfter(s, sep string) string {
	if strings.Contains(s, sep) {
		return strings.Split(s, sep)[0]
	}

	return s
}

func cutBefore(s, sep string) string {
	if strings.Contains(s, sep) {
		return strings.Split(s, sep)[1]
	}

	return s
}

var byteUnits = []string{"B", "KB", "MB", "GB", "TB", "PB"}

func byteUnitString(n int64) string {
	var unit string
	size := float64(n)
	for i := 1; i < len(byteUnits); i++ {
		if size < 1000 {
			unit = byteUnits[i-1]
			break
		}

		size = size / 1000
	}

	return fmt.Sprintf("%.3g %s", size, unit)
}
