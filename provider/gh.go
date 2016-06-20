package provider

import (
  "io"
  "errors"

  "github.com/mh-cbon/go-prebuilt/srcparser"
  "github.com/google/go-github"
)

type Gh struct{}

func (g Gh) get (source *srcparser.ParsedSrc) (io.ReadCloser, error){
  client := github.NewClient(nil)

  if source.Tag=="" {
    latest, _, err := client.Repositories.GetLatestRelease(source.User, source.Repo)
    if err!=nil {
      return nil, err
    }
    if latest==nil {
      return nil, errors.New("This repository does not have any release")
    }
    source.Tag = latest.TagName
  }

  release, _, err := client.Repositories.GetReleaseByTag(source.User, source.Repo, source.Tag)
  if err!=nil {
    return nil, err
  }
  if release==nil {
    return nil, errors.New("This repository does not have any release '"+source.Tag+"'")
  }

  assetRe := regexp.MustCompile("(i?)"+runtime.GOOS+"[.-]"+runtime.GOARCH)
  var asset github.ReleaseAsset
  for _, a := range release.Assets {
    if assetRe.MatchString(a.Name) || assetRe.MatchString(a.Label) {
      asset = a
      break
    }
  }
  if asset==nil {
    return nil, errors.New("This repository does not have any asset for your system/arch '"+runtime.GOOS+"-"+runtime.GOARCH+"'")
  }

  rc, redirect, err := client.Repositories.DownloadReleaseAsset(source.User, source.Repo, asset.Id)
  if err!=nil {
    return nil, err
  }
  if redirect!="" {
    return nil, errors.New("This repository asset returned a redirection to '"+redirect+"'")
  }

  return rc, nil
}
