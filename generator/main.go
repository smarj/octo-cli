package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const routesTimestampFormat = "20060102T150405Z0700"

type (
	genCli struct {
		Run            genCliRun            `cmd:"" help:"generate production code"`
		UpdateRoutes   genCliUpdateRoutes   `cmd:"" help:"update routes.json with the latest"`
		UpdateTestdata genCliUpdateTestdata `cmd:"" help:"updates routes.json and services in generator/testdata"`
	}

	genCliUpdateRoutes struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		RoutesURL  string `default:"https://octokit.github.io/routes/index.json"`
	}

	genCliRun struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		OutputPath string `type:"existingdir" default:"./services"`
		Verify     bool   `help:"Verify a new run won't change anything"`
	}

	genCliUpdateTestdata struct{}
)

func (k *genCliRun) Run() error {
	if k.Verify {
		diffs, err := verify(k.RoutesPath, k.OutputPath)
		if err != nil {
			return errors.New("error verifying")
		}
		if len(diffs) > 0 {
			return fmt.Errorf("some files did not match: %v", diffs)
		}
	} else {
		Generate(k.RoutesPath, k.OutputPath)
	}
	return nil
}

func updateRoutes(routesURL, routesPath string) error {
	resp, err := http.Get(routesURL)
	if err != nil {
		return errors.Wrapf(err, "failed getting %q", routesURL)
	}

	err = writeRoutesJSON(routesPath, resp)
	if err != nil {
		return errors.Wrap(err, "failed writing routesJSON")
	}

	err = updateLastModified(routesPath, resp)
	return errors.Wrap(err, "failed updateing last modified")
}

func writeRoutesJSON(routesPath string, resp *http.Response) error {
	outFile, err := os.Create(routesPath)
	if err != nil {
		return errors.Wrapf(err, "failed creating file %q", routesPath)
	}
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed writing to file %q", routesPath)
	}
	err = resp.Body.Close()
	if err != nil {
		return errors.Wrap(err, "failed closing response body")
	}
	return errors.Wrapf(err, "failed closing file %q", outFile.Close())
}

func updateLastModified(routesPath string, resp *http.Response) error {
	extension := filepath.Ext(routesPath)
	if extension == "" {
		return errors.Errorf("routesPath must have a file extension (preferable .json)")
	}
	lmPath := strings.TrimSuffix(routesPath, extension) + "-last-modified.txt"
	lmHeader := resp.Header.Get("last-modified")
	lmTime, err := time.Parse(time.RFC1123, lmHeader)
	if err != nil {
		return errors.Wrapf(err, "couldn't parse last-modified time %q", lmHeader)
	}
	lmFile, err := os.Create(lmPath)
	if err != nil {
		return errors.Wrapf(err, "failed creating file %q", lmPath)
	}
	_, err = lmFile.WriteString(lmTime.Format(routesTimestampFormat))
	if err != nil {
		return errors.Wrapf(err, "failed writing to file %q", lmPath)
	}
	return errors.Wrapf(lmFile.Close(), "failed closing file %q", lmPath)
}

func (k *genCliUpdateRoutes) Run() error {
	return updateRoutes(k.RoutesURL, k.RoutesPath)
}

func (k *genCliUpdateTestdata) Run() error {
	url := "https://octokit.github.io/routes/index.json"
	routesPath := "generator/testdata/routes.json"
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "")
	}
	outFile, err := os.Create(routesPath)
	if err != nil {
		return errors.Wrap(err, "")
	}
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}

	Generate(routesPath, "generator/testdata/services")

	return errors.Wrap(err, "")
}

func main() {
	k := kong.Parse(&genCli{})
	err := k.Run()
	k.FatalIfErrorf(err)
}
