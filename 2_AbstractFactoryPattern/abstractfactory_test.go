package abstractfactory

import (
	"fmt"
	"testing"
)

var superFactory *FactoryProducer

func TestMain(m *testing.M) {
	superFactory = NewFactoryProducer()
	m.Run()
}

func TestColorFactory(t *testing.T) {
	colorFactory := superFactory.GetFactory("color")
	redColor := colorFactory.GetColor("red")
	if redColor != nil {
		redColor.Fill()
	} else {
		fmt.Println("redColorFactory test Fail.")
	}
	greenColor := colorFactory.GetColor("green")
	if greenColor != nil {
		greenColor.Fill()
	} else {
		fmt.Println("greenColorFactory test Fail.")
	}
}

func TestShapeFactory(t *testing.T) {
	shapeFactory := superFactory.GetFactory("shape")
	circleShape := shapeFactory.GetShape("circle")
	if circleShape != nil {
		circleShape.Draw()
	} else {
		fmt.Println("circleShapeFactory test Fail.")
	}

	squareShape := shapeFactory.GetShape("square")
	if squareShape != nil {
		squareShape.Draw()
	} else {
		fmt.Println("squareShapeFactory test Fail.")
	}
}
