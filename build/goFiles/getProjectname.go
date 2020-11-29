package main

import (
  "regexp"
  "strings"
)

func getProjectname(g Git) (Git) {

//  fmt.Println("enter getProjectname")
//  fmt.Println(g.projectname)
  re := regexp.MustCompile(`.*/`)
  var toCut string = re.FindString(g.Projectname)
  g.Projectname = strings.TrimPrefix(g.Projectname, toCut)
//  fmt.Println(g.projectname)
  g.Projectname = strings.TrimSuffix(g.Projectname, ".git")
//  fmt.Println(g.projectname)

  return g
}
