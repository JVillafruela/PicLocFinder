package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/urfave/cli"
)

type Config struct {
	bbox string
	// args are the positional (non-flag) command-line arguments.
	args []string
}

func main() {
	conf := Config{}

	app := cli.NewApp()
	app.Name = "plf"
	app.Usage = "Picture Location Finder\n\n   Find geotagged photos according to the location where they were taken"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "bbox",
			Usage:       "bounding box (\"lon1,lat1,lon2,lat2\") ",
			Required:    true,
			Destination: &conf.bbox,
		},
	}

	app.Action = func(c *cli.Context) error {
		conf.args = c.Args()
		err := validateConfig(&conf)
		if err != nil {
			return err
		}
		doWork(&conf)
		return nil
	}

	cli.AppHelpTemplate = fmt.Sprintf(`%s

WEBSITE: https://github.com/JVillafruela/plf

EXAMPLE:
   plf  --bbox="5.68678,45.08596,5.68979,45.08778" E:\OSM\gps\2022\2022-04-16 E:\OSM\gps\2022\2022-04-22
	
	`, cli.AppHelpTemplate)

	app.Setup()
	// do not display command in help (see cli issue #523)
	app.Commands = []cli.Command{}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func validateConfig(config *Config) error {

	if config.bbox == "" {
		return errors.New("missing bbox option")
	}

	_, err := bboxBound(config.bbox)
	if err != nil {
		return errors.New("invalid value for bounding box option. Should be \"lon1,lat1,lon2,lat2\" ")
	}

	if len(config.args) == 0 {
		return errors.New("missing argument")
	}

	for _, v := range config.args {
		path := cleanPath(v)
		if !dirExists(path) {
			return errors.New("Directory does not exist : " + path)
		}
	}

	return nil
}

func doWork(config *Config) {
	//fmt.Printf("config = %+v\n", *config)
	for _, dir := range config.args {
		walkPicFilesInDir(dir, config.bbox)
	}

}

func walkPicFilesInDir(dir string, bbox string) error {
	bound, err := bboxBound(bbox)
	if err != nil {
		return err
	}

	return filepath.WalkDir(dir, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := file.Info()
		if err != nil {
			return err
		}

		if !file.IsDir() && hasPicExtension(info.Name()) {
			fullname := filepath.Clean(path)
			lat, lon, err := PicLocation(fullname)
			if err != nil {
				fmt.Println("File name: ", fullname, "error ", err)
				return nil
			}
			//fmt.Println("DDD file name: ", fullname, lon, lat)
			if MatchLocation(lat, lon, bound) {
				fmt.Println(fullname)
			}
		}
		return nil
	})
}

func hasPicExtension(filename string) bool {
	e := []string{"jpeg", "jpg"} // sorted

	ext := strings.ToLower(strings.Trim(filepath.Ext(filename), "."))
	if ext == "" {
		return false
	}
	i := sort.SearchStrings(e, ext)
	return i < len(e) && e[i] == ext

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cleanPath(path string) string {
	if runtime.GOOS == "windows" {
		path = strings.TrimSuffix(path, `"`)
	}
	return filepath.Clean(path)
}

func dirExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "dirExists "+filename, err.Error())
		return false
	}
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}
