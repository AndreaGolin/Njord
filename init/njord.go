/*
* @Author: andrea
* @Date:   2018-02-28 08:50:29
* @Last Modified by:   andrea
* @Last Modified time: 2018-02-28 10:46:21
 */
package main

import (
	"github.com/AndreaGolin/njord"
	"os"
)

func main() {
	args := os.Args[1:]
	njord.Init(args)
}
