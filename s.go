package main

import (
	"fmt"
	"os"

	// Load necessary providers.
	_ "github.com/zquestz/s/providers/500px"
	_ "github.com/zquestz/s/providers/8tracks"
	_ "github.com/zquestz/s/providers/aliexpress"
	_ "github.com/zquestz/s/providers/allocine"
	_ "github.com/zquestz/s/providers/amazon"
	_ "github.com/zquestz/s/providers/archpkg"
	_ "github.com/zquestz/s/providers/archwiki"
	_ "github.com/zquestz/s/providers/arstechnica"
	_ "github.com/zquestz/s/providers/arxiv"
	_ "github.com/zquestz/s/providers/atmospherejs"
	_ "github.com/zquestz/s/providers/aur"
	_ "github.com/zquestz/s/providers/baidu"
	_ "github.com/zquestz/s/providers/bandcamp"
	_ "github.com/zquestz/s/providers/bgr"
	_ "github.com/zquestz/s/providers/bigbasket"
	_ "github.com/zquestz/s/providers/bing"
	_ "github.com/zquestz/s/providers/buzzfeed"
	_ "github.com/zquestz/s/providers/cnn"
	_ "github.com/zquestz/s/providers/codepen"
	_ "github.com/zquestz/s/providers/coursera"
	_ "github.com/zquestz/s/providers/cplusplus"
	_ "github.com/zquestz/s/providers/cppreference"
	_ "github.com/zquestz/s/providers/crates"
	_ "github.com/zquestz/s/providers/crunchyroll"
	_ "github.com/zquestz/s/providers/debianpkg"
	_ "github.com/zquestz/s/providers/dict"
	_ "github.com/zquestz/s/providers/digg"
	_ "github.com/zquestz/s/providers/diigo"
	_ "github.com/zquestz/s/providers/dockerhub"
	_ "github.com/zquestz/s/providers/dribbble"
	_ "github.com/zquestz/s/providers/duckduckgo"
	_ "github.com/zquestz/s/providers/dumpert"
	_ "github.com/zquestz/s/providers/engadget"
	_ "github.com/zquestz/s/providers/explainshell"
	_ "github.com/zquestz/s/providers/facebook"
	_ "github.com/zquestz/s/providers/flickr"
	_ "github.com/zquestz/s/providers/flipkart"
	_ "github.com/zquestz/s/providers/foursquare"
	_ "github.com/zquestz/s/providers/giphy"
	_ "github.com/zquestz/s/providers/gist"
	_ "github.com/zquestz/s/providers/github"
	_ "github.com/zquestz/s/providers/gmail"
	_ "github.com/zquestz/s/providers/go"
	_ "github.com/zquestz/s/providers/godoc"
	_ "github.com/zquestz/s/providers/goodreads"
	_ "github.com/zquestz/s/providers/google"
	_ "github.com/zquestz/s/providers/googledocs"
	_ "github.com/zquestz/s/providers/googleplus"
	_ "github.com/zquestz/s/providers/hackernews"
	_ "github.com/zquestz/s/providers/ietf"
	_ "github.com/zquestz/s/providers/ifttt"
	_ "github.com/zquestz/s/providers/imdb"
	_ "github.com/zquestz/s/providers/imgur"
	_ "github.com/zquestz/s/providers/inbox"
	_ "github.com/zquestz/s/providers/instagram"
	_ "github.com/zquestz/s/providers/kickasstorrents"
	_ "github.com/zquestz/s/providers/libgen"
	_ "github.com/zquestz/s/providers/linkedin"
	_ "github.com/zquestz/s/providers/lmgtfy"
	_ "github.com/zquestz/s/providers/macports"
	_ "github.com/zquestz/s/providers/magnetdl"
	_ "github.com/zquestz/s/providers/mdn"
	_ "github.com/zquestz/s/providers/medium"
	_ "github.com/zquestz/s/providers/metacpan"
	_ "github.com/zquestz/s/providers/msdn"
	_ "github.com/zquestz/s/providers/naver"
	_ "github.com/zquestz/s/providers/netflix"
	_ "github.com/zquestz/s/providers/nhaccuatui"
	_ "github.com/zquestz/s/providers/npm"
	_ "github.com/zquestz/s/providers/npmsearch"
	_ "github.com/zquestz/s/providers/npr"
	_ "github.com/zquestz/s/providers/nvd"
	_ "github.com/zquestz/s/providers/overstock"
	_ "github.com/zquestz/s/providers/packagist"
	_ "github.com/zquestz/s/providers/phandroid"
	_ "github.com/zquestz/s/providers/php"
	_ "github.com/zquestz/s/providers/pinterest"
	_ "github.com/zquestz/s/providers/postgresql"
	_ "github.com/zquestz/s/providers/python"
	_ "github.com/zquestz/s/providers/quora"
	_ "github.com/zquestz/s/providers/qwant"
	_ "github.com/zquestz/s/providers/reddit"
	_ "github.com/zquestz/s/providers/regex"
	_ "github.com/zquestz/s/providers/rottentomatoes"
	_ "github.com/zquestz/s/providers/rubygems"
	_ "github.com/zquestz/s/providers/shodan"
	_ "github.com/zquestz/s/providers/soundcloud"
	_ "github.com/zquestz/s/providers/spotify"
	_ "github.com/zquestz/s/providers/stackoverflow"
	_ "github.com/zquestz/s/providers/steam"
	_ "github.com/zquestz/s/providers/taobao"
	_ "github.com/zquestz/s/providers/thepiratebay"
	_ "github.com/zquestz/s/providers/theregister"
	_ "github.com/zquestz/s/providers/torrentz"
	_ "github.com/zquestz/s/providers/twitchtv"
	_ "github.com/zquestz/s/providers/twitter"
	_ "github.com/zquestz/s/providers/unity3d"
	_ "github.com/zquestz/s/providers/upcloud"
	_ "github.com/zquestz/s/providers/vimeo"
	_ "github.com/zquestz/s/providers/wikipedia"
	_ "github.com/zquestz/s/providers/wolframalpha"
	_ "github.com/zquestz/s/providers/yahoo"
	_ "github.com/zquestz/s/providers/yandex"
	_ "github.com/zquestz/s/providers/youtube"
	_ "github.com/zquestz/s/providers/zhihu"

	"github.com/zquestz/s/cmd"
)

func main() {
	setupSignalHandlers()

	if err := cmd.SearchCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
