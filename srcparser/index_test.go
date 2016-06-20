package srcparser

import (
	"testing"
)

func TestIt(t *testing.T) {
  fixtures := map[string]ParsedSrc{
    "http://github.com/some/repo": ParsedSrc{
      Scheme: "http://",
      Provider: "github",
      User: "some",
      Repo: "repo",
      Tag: "",
    },
    "https://github.com/some/repo": ParsedSrc{
      Scheme: "https://",
      Provider: "github",
      User: "some",
      Repo: "repo",
      Tag: "",
    },
    "github.com/some/repo": ParsedSrc{
      Scheme: "",
      Provider: "github",
      User: "some",
      Repo: "repo",
      Tag: "",
    },
    "github/some/repo": ParsedSrc{
      Scheme: "",
      Provider: "github",
      User: "some",
      Repo: "repo",
      Tag: "",
    },
    "gh/some/repo": ParsedSrc{
      Scheme: "",
      Provider: "github",
      User: "some",
      Repo: "repo",
      Tag: "",
    },
    "gh/some/repo#0.0.1": ParsedSrc{
      Scheme: "",
      Provider: "github",
      User: "some",
      Repo: "repo",
      Tag: "0.0.1",
    },
    "some/repo#0.0.1": ParsedSrc{
      Scheme: "",
      Provider: "",
      User: "some",
      Repo: "repo",
      Tag: "0.0.1",
    },
    "repo#0.0.1": ParsedSrc{
      Scheme: "",
      Provider: "",
      User: "",
      Repo: "repo",
      Tag: "0.0.1",
    },
    "#0.0.1": ParsedSrc{
      Scheme: "",
      Provider: "",
      User: "",
      Repo: "",
      Tag: "0.0.1",
    },
    "wtf": ParsedSrc{
      Scheme: "",
      Provider: "",
      User: "",
      Repo: "wtf",
      Tag: "",
    },
  }

  for fixture, expectation := range fixtures {
    res, _ := Parse(fixture)
    if res.Scheme != expectation.Scheme {
      t.Errorf("Expected Scheme=%q, got Scheme=%q in %q\n", expectation.Scheme, res.Scheme, fixture)
    }
    if res.Provider != expectation.Provider {
      t.Errorf("Expected Provider=%q, got Provider=%q in %q\n", expectation.Provider, res.Provider, fixture)
    }
    if res.User != expectation.User {
      t.Errorf("Expected User=%q, got User=%q in %q\n", expectation.User, res.User, fixture)
    }
    if res.Repo != expectation.Repo {
      t.Errorf("Expected Repo=%q, got Repo=%q in %q\n", expectation.Repo, res.Repo, fixture)
    }
    if res.Tag != expectation.Tag {
      t.Errorf("Expected Tag=%q, got Tag=%q in %q\n", expectation.Tag, res.Tag, fixture)
    }
  }
}
