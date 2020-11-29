package main

import (
  "testing"
)

func TestGetProjectname(t *testing.T) {

  g1 := Git{"hash", "https://github.com/user/projectname.git"}
  g1 = getProjectname(g1)

  if g1.Projectname != "projectname" {
    t.Errorf("getProjectname failed, github repo")
  }

  g2 := Git{"hash", "https://git.company.com/user/projectname.git"}
  g2 = getProjectname(g2)

  if g2.Projectname != "projectname" {
    t.Errorf("getProjectname failed, private git server")
  }

  g3 := Git{"hash", "git@github.com:user/projectname.git"}
  g3 = getProjectname(g3)

  if g3.Projectname != "projectname" {
    t.Errorf("getProjectname failed, ssh github")
  }

  g4 := Git{"hash", "git@git.company.com:user/projectname.git"}
  g4 = getProjectname(g4)

  if g4.Projectname != "projectname" {
    t.Errorf("getProjectname failed, ssh private git server")
  }

  g5 := Git{"hash", ""}
  g5 = getProjectname(g5)

  if g5.Projectname != "" {
    t.Errorf("getProjectname failed, no Url given")
  }
}
