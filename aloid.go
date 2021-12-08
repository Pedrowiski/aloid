package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type api_request_t struct {
	Key        string
	File       string
	Api_option string

	Api_user_key string
	Name         string
	Format       string
	Permission   string
	Expire_date  string
	Folder_key   string
}

type flag_parse_t struct {
	Key        *string
	File       *string
	Api_option *string

	Api_user_key *string
	Name         *string
	Format       *string
	Permission   *string
	Expire_date  *string
	Folder_key   *string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func make_request(api_request api_request_t) {
	request_url := "https://pastebin.com/api/api_post.php"

	data := url.Values{
		"api_dev_key":    {api_request.Key},
		"api_option":     {api_request.Api_option},
		"api_paste_code": {api_request.File},

		"api_paste_name":        {api_request.Name},
		"api_paste_format":      {api_request.Format},
		"api_paste_private":     {api_request.Permission},
		"api_paste_expire_date": {api_request.Expire_date},
		"api_folder_key":        {api_request.Folder_key},
	}

	response, err := http.PostForm(request_url, data)
	check(err)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	check(err)

	fmt.Println(string(body))
}

var (
	flag_parse  flag_parse_t
	api_request api_request_t

	help_pointer *string
	help         string
)

func init() {
	flag_parse.Key = flag.String("key", "", "api dev key")
	flag_parse.File = flag.String("file", "", "file name")
	flag_parse.Api_option = flag.String("api_option", "paste", "api option")

	flag_parse.Name = flag.String("name", "", "file name in pastebin")
	flag_parse.Format = flag.String("format", "", "file extension")
	flag_parse.Permission = flag.String("permission", "0", "view permission")
	flag_parse.Expire_date = flag.String("expire", "N", "file duration time")
	flag_parse.Folder_key = flag.String("folder_key", "", "folder key")
	flag_parse.Api_user_key = flag.String("api_user_key", "", "api user key")
	help_pointer = flag.String("help", "", "help option")

	flag.Parse()

	api_request.Key = *flag_parse.Key
	api_request.File = *flag_parse.File
	api_request.Api_option = *flag_parse.Api_option

	api_request.Api_user_key = *flag_parse.Api_user_key
	api_request.Name = *flag_parse.Name
	api_request.Format = *flag_parse.Format
	api_request.Permission = *flag_parse.Permission
	api_request.Expire_date = *flag_parse.Expire_date
	api_request.Folder_key = *flag_parse.Folder_key

	help = *help_pointer
}

func main() {
	switch {
	case api_request.Key == "":
		fmt.Fprintf(os.Stderr, "aloid: %s\n", "missing API key")
		os.Exit(1)
	case api_request.File == "":
		fmt.Fprintf(os.Stderr, "aloid %s\n", "missing file name")
		os.Exit(1)
	}

	file_data, err := os.ReadFile(api_request.File)
	check(err)
	api_request.File = string(file_data)

	make_request(api_request)
}
