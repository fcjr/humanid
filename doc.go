// Package humanid provides generates unique(-ish) human readable ids
//
//  Usage
//
// default separator: '-'
// id := humanid.New()
// fmt.Println(id)
//
// custom separator
// hid := humanid.NewProvider("_")
// fmt.Println(hid.New())
//
package humanid //import "github.com/fcjr/humanid"
