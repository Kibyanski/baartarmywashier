package core

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	VERSION = "2.4.2"
)

func putAsciiArt(s string) {
	for _, c := range s {
		d := string(c)
		switch string(c) {
		case " ":
			color.Set(color.BgRed)
			d = " "
		case "@":
			color.Set(color.BgBlack)
			d = " "
		case "#":
			color.Set(color.BgHiRed)
			d = " "
		case "W":
			color.Set(color.BgWhite)
			d = " "
		case "_":
			color.Unset()
			d = " "
		case "\n":
			color.Unset()
		}
		fmt.Print(d)
	}
	color.Unset()
}

func printLogo(s string) {
	for _, c := range s {
		d := string(c)
		switch string(c) {
		case "_":
			color.Set(color.FgWhite)
		case "\n":
			color.Unset()
		default:
			color.Set(color.FgHiBlack)
		}
		fmt.Print(d)
	}
	color.Unset()
}

func printUpdateName() {
	nameClr := color.New(color.FgHiRed)
	txt := nameClr.Sprintf("                 - --  Gone Phishing  -- -")
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner1() {
	handleClr := color.New(color.FgHiBlue)
	versionClr := color.New(color.FgGreen)
	textClr := color.New(color.FgHiBlack)
	spc := strings.Repeat(" ", 10-len(VERSION))
	txt := textClr.Sprintf("      by Jerp (") + handleClr.Sprintf("@JerpPanel") + textClr.Sprintf(")") + spc + textClr.Sprintf("version ") + versionClr.Sprintf("%s", VERSION)
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner2() {
	textClr := color.New(color.FgHiBlack)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)
	txt := textClr.Sprintf("                   no ") + red.Sprintf("nginx") + white.Sprintf(" - ") + textClr.Sprintf("pure ") + red.Sprintf("evil")
	fmt.Fprintf(color.Output, "%s", txt)
}

func Banner() {
	fmt.Println()

	putAsciiArt("__                                     __\n")
	putAsciiArt("  @@@@@@            @@@@@@@@@@")
	printLogo(`       _           _____                 _ `)
	fmt.Println()
	putAsciiArt("   @@@@               @@@@@@@@")
	printLogo(`      | |         |  __ \               | |`)
	fmt.Println()
	putAsciiArt("     @           WW    @@@@@@@")
	printLogo(`      | | ___ _ __| |__) |_ _ _ __   ___| |`)
	fmt.Println()
	putAsciiArt("                W##     @@@@@@")
	printLogo(`  _   | |/ _ \  __|  ___/ _  |  _ \ / _ \ |`)
	fmt.Println()
	putAsciiArt("                        @@@@@@")
	printLogo(` | |__| |  __/ |  | |  | (_| | | | |  __/ |`)
	fmt.Println()
	putAsciiArt("     @                 @@@@@")
	printLogo(`  \____/ \___|_|  |_|   \__ _|_| |_|\___|_|`)
	fmt.Println()
	putAsciiArt("    @@@               @@@@@@@@\n")
	putAsciiArt(   @@@@@@@@@@@@@@@@@@@@@@@@@")
	printUpdateName()
	fmt.Println()
	putAsciiArt("  @@@@@@@@@@@@@@@@@@@@@@@@@@@@\n")
	//printOneliner2()
	//fmt.Println()
	putAsciiArt("@@                            ")
	printOneliner1()
	fmt.Println()
	putAsciiArt("__                                     __\n")
	fmt.Println()
}
