package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main(){
	vCalc := app.New()
	vWindow := vCalc.NewWindow("Cross Product Calculator")
	vector1 := container.NewHBox()
	vector2 := container.NewHBox()
	vector1.Add(widget.NewLabel("Vector 1: "))
	vector2.Add(widget.NewLabel("Vector 2: "))
	v1IHat := widget.NewEntry()
	v1JHat := widget.NewEntry()
	v1KHat := widget.NewEntry()
	v2IHat := widget.NewEntry()
	v2JHat := widget.NewEntry()
	v2KHat := widget.NewEntry()
	vector1.Add(v1IHat)
	vector1.Add(widget.NewLabel("i"))
	vector1.Add(v1JHat)
	vector1.Add(widget.NewLabel("j"))
	vector1.Add(v1KHat)
	vector1.Add(widget.NewLabel("k"))
	vector2.Add(v2IHat)
	vector2.Add(widget.NewLabel("i"))
	vector2.Add(v2JHat)
	vector2.Add(widget.NewLabel("j"))
	vector2.Add(v2KHat)
	vector2.Add(widget.NewLabel("k"))
	content := container.NewVBox()
	content.Add(vector1)
	content.Add(vector2)
	result := widget.NewLabel("Result:")
	content.Add(widget.NewButton("Calculate", func() {
		calculate(v1IHat.Text,v1JHat.Text,v1KHat.Text,v2IHat.Text,v2JHat.Text,v2KHat.Text,result)
	}))
	content.Add(result)
	vWindow.SetContent(content)
	vWindow.ShowAndRun()


}

func calculate(v1IHat string,v1JHat string,v1KHat string,v2IHat string,v2JHat string,v2KHat string, label *widget.Label){
	u1,_ := strconv.ParseFloat(v1IHat,64)
	u2,_ := strconv.ParseFloat(v1JHat,64)
	u3,_ := strconv.ParseFloat(v1KHat,64)
	v1,_ := strconv.ParseFloat(v2IHat,64)
	v2,_ := strconv.ParseFloat(v2JHat,64)
	v3,_ := strconv.ParseFloat(v2KHat,64)
	t1 := u2*v3 - u3*v2
	t2 := u3*v1 - u1*v3
	t3 := u1*v2-u2*v1
	fmtResult := fmt.Sprintf("%f, %f, %f",t1,t2,t3)
	label.SetText("Result: "+fmtResult)


}
