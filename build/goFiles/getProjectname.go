/* This function parses the projectname out of the given repository url.
 * Also works with ssh.
 */

package main

import (
  "regexp"
  "strings"
)

func getProjectname(g Git) (Git) {

  re := regexp.MustCompile(`.*/`)
  var toCut string = re.FindString(g.Projectname)
  g.Projectname = strings.TrimPrefix(g.Projectname, toCut)
  g.Projectname = strings.TrimSuffix(g.Projectname, ".git")

  return g
}
