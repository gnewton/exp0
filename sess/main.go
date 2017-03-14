package main

import (
	"github.com/icza/gowut/gwu"
	"log"
	"strconv"
	"time"
)

var pages = []string{"journal.pcbi.1005413-0.png",
	"journal.pcbi.1005413-1.png",
	"journal.pcbi.1005413-2.png",
	"journal.pcbi.1005413-3.png",
	"journal.pcbi.1005413-4.png",
	"journal.pcbi.1005413-5.png",
	"journal.pcbi.1005413-6.png",
	"journal.pcbi.1005413-7.png",
	"journal.pcbi.1005413-8.png",
	"journal.pcbi.1005413-9.png",
	"journal.pcbi.1005413-10.png",
	"journal.pcbi.1005413-11.png",
	"journal.pcbi.1005413-12.png",
	"journal.pcbi.1005413-13.png",
	"journal.pcbi.1005413-14.png",
	"journal.pcbi.1005413-15.png",
	"journal.pcbi.1005413-16.png",
	"journal.pcbi.1005413-17.png",
	"journal.pcbi.1005413-18.png",
	"journal.pcbi.1005413-19.png",
	"journal.pcbi.1005413-20.png",
	"journal.pcbi.1005413-21.png",
	"journal.pcbi.1005413-22.png",
	"journal.pcbi.1005413-23.png",
	"journal.pcbi.1005413-24.png"}

type sessHandler struct{}

var pageNumber = 1
var label gwu.Label

//var pages

func main() {
	// Create and start a GUI server (omitting error check)
	server := gwu.NewServer("guitest", "localhost:8081")
	server.SetText("Test GUI App")
	//server.AddWin(win)
	err := server.AddStaticDir("data", "data")
	if err != nil {
		log.Fatal(err)
	}

	server.AddSessCreatorName("main", "Login Window")
	server.AddSHandler(sessHandler{})

	server.Start("") // Also opens windows list in browser
}

func (h sessHandler) Removed(s gwu.Session) {}

func (h sessHandler) Created(s gwu.Session) {
	// Create and build a window
	win := gwu.NewWindow("main", "Test GUI Window")
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HACenter)
	win.SetCellPadding(10)

	s.AddWin(win)

	// Control panel
	p := gwu.NewHorizontalPanel()
	p.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	p.SetCellPadding(10)
	win.Add(p)

	//p.Add(gwu.NewHtml("<form action=\"main\" method=\"post\" enctype=\"multipart/form-data\"> <input type=\"file\" name=\"fileToUpload\" id=\"fileToUpload\"><input type=\"submit\" value=\"Upload Image\" name=\"submit\"></form>"))

	numPages := len(pages)

	im := gwu.NewImage("", "/guitest/data/"+pages[pageNumber-1])
	// Timer to check page number
	timer := gwu.NewTimer(time.Millisecond * 2000000)
	timer.SetRepeat(true)
	timer.AddEHandlerFunc(func(e gwu.Event) {
		if pageNumber < numPages {
			im.SetUrl("/guitest/data/" + pages[pageNumber-1])
			label.SetText(strconv.Itoa(pageNumber))
			e.MarkDirty(label)
			e.MarkDirty(im)
			log.Println("Timer update")

		}
	}, gwu.ETypeStateChange)

	win.Add(timer)

	// Display panel
	p2 := gwu.NewPanel()

	p2.Add(im)
	win.Add(p2)

	bminus := gwu.NewButton("-")
	var bplus gwu.Button
	bminus.AddEHandlerFunc(func(e gwu.Event) {
		if pageNumber > 1 {
			pageNumber = pageNumber - 1
			im.SetUrl("/guitest/data/" + pages[pageNumber-1])
			label.SetText(strconv.Itoa(pageNumber))
			e.MarkDirty(label)
			e.MarkDirty(im)

			if pageNumber == 1 {
				bminus.SetEnabled(false)
				e.MarkDirty(bminus)
			} else {
				if pageNumber < numPages {
					bplus.SetEnabled(true)
					e.MarkDirty(bplus)
				}
			}
		}

	}, gwu.ETypeClick)
	p.Add(bminus)

	bplus = gwu.NewButton("+")
	bplus.AddEHandlerFunc(func(e gwu.Event) {
		if pageNumber < len(pages) {
			pageNumber = pageNumber + 1
			im.SetUrl("/guitest/data/" + pages[pageNumber-1])
			label.SetText(strconv.Itoa(pageNumber))
			e.MarkDirty(label)
			e.MarkDirty(im)
			if pageNumber == numPages {
				bplus.SetEnabled(false)
				e.MarkDirty(bplus)
			} else {
				if pageNumber > 1 {
					bminus.SetEnabled(true)
					e.MarkDirty(bminus)
				}
			}
		}

	}, gwu.ETypeClick)
	p.Add(bplus)

	label = gwu.NewLabel(strconv.Itoa(pageNumber))
	p.Add(label)

}
