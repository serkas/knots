package models__test

import (
	"knots/models"
	"testing"
)

func TestNew(t *testing.T) {
	title := "Title"
	knot := models.Knot{Raw: "# Section \n\n hallo", Title: title}
	if title != knot.Title {
		t.Errorf("Expected %s, got %v", title, knot.Title)
	}
}

func TestValidateTrue(t *testing.T) {
	knot := models.Knot{Raw: "Something"}
	if !knot.Validate() {
		t.Errorf("Got valid knot %v, but validation faild", knot)
	}
}

func TestValidateFalse(t *testing.T) {
	knot := models.Knot{Raw: ""}
	if knot.Validate() {
		t.Errorf("Got invalid knot %v, but validation passed", knot)
	}
}
