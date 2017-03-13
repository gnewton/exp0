
package main

import(
      "github.com/icza/gowut/gwu"
      "strconv"
      "log"
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

var pageNumber = 1
var label gwu.Label
//var pages 

func main(){
// Create and build a window
	win := gwu.NewWindow("main", "Test GUI Window")
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HACenter)
	win.SetCellPadding(10)

	// Control panel
	p := gwu.NewHorizontalPanel()
	p.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	p.SetCellPadding(10)
	win.Add(p)

	//p.Add(gwu.NewHtml("<form action=\"main\" method=\"post\" enctype=\"multipart/form-data\"> <input type=\"file\" name=\"fileToUpload\" id=\"fileToUpload\"><input type=\"submit\" value=\"Upload Image\" name=\"submit\"></form>"))

	// Display panel
	p2 := gwu.NewPanel()
	im := gwu.NewImage("", "/guitest/data/" + pages[pageNumber-1])
	p2.Add(im)
	win.Add(p2)


	b := gwu.NewButton("-")
	b.AddEHandlerFunc(func(e gwu.Event) {
	                 if pageNumber > 1{
			 pageNumber = pageNumber -1
			 im.SetUrl("/guitest/data/" + pages[pageNumber-1])		
			 label.SetText(strconv.Itoa(pageNumber))
			 e.MarkDirty(label)
  			 e.MarkDirty(im)	
			 }
			 }, gwu.ETypeClick)
	p.Add(b)


	b = gwu.NewButton("+")
	  b.AddEHandlerFunc(func(e gwu.Event) {
	  		 if pageNumber < len(pages){	   
			 pageNumber = pageNumber +1
 			 im.SetUrl("/guitest/data/" + pages[pageNumber-1])		
 			 label.SetText(strconv.Itoa(pageNumber))
			 e.MarkDirty(label)
 			 e.MarkDirty(im)	
			 }

}, gwu.ETypeClick)
	p.Add(b)

	label = gwu.NewLabel(strconv.Itoa(pageNumber))
	p.Add(label)



// Create and start a GUI server (omitting error check)
	server := gwu.NewServer("guitest", "localhost:8081")
	server.SetText("Test GUI App")
	server.AddWin(win)
	err :=	server.AddStaticDir("data", "data")
	if err != nil{
	    log.Fatal(err)
	}	
	server.Start("") // Also opens windows list in browser
}