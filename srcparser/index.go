package srcparser

import (
  "regexp"
)

type ParsedSrc struct {
  Scheme string
  Provider string
  User string
  Repo string
  Tag string
}

func Parse (src string) (*ParsedSrc, error) {

  schemeRe := regexp.MustCompile("(?i)^https?://")
  providerRe := regexp.MustCompile("(?i)^(gh|github|github[.]com)/")
  userRe := regexp.MustCompile("(?i)^([^/]+)/")
  repoRe := regexp.MustCompile("(?i)^([^/]+)/?")
  tagRe := regexp.MustCompile("(?i)#([^#]+)$")

  scheme := ""
  provider := ""
  user := ""
  repo := ""
  tag := ""

  if schemeRe.MatchString(src) {
    scheme = schemeRe.FindString(src)
    src = src[len(scheme):]
  }

  if tagRe.MatchString(src) {
    tTag := tagRe.FindStringSubmatch(src)
    if len(tTag)>0 {
      tag = tTag[1]
      src = src[0:len(src)-len(tTag[0])]
    }
  }

  if providerRe.MatchString(src) {
    tProvider := providerRe.FindStringSubmatch(src)
    if len(tProvider)>0 {
      provider = tProvider[1]
      src = src[len(tProvider[0]):]
      if provider=="github.com" || provider=="gh" {
        provider = "github"
      }
    }
  }

  if userRe.MatchString(src) {
    tUser := userRe.FindStringSubmatch(src)
    if len(tUser)>0 {
      user = tUser[1]
      src = src[len(tUser[0]):]
    }
  }

  if repoRe.MatchString(src) {
    tRepo := repoRe.FindStringSubmatch(src)
    if len(tRepo)>0 {
      repo = tRepo[1]
      src = src[len(tRepo[0]):]
    }
  }

  ret := ParsedSrc{
    Scheme: scheme,
    Provider: provider,
    User: user,
    Repo: repo,
    Tag: tag,
  }

  return &ret, nil
}
