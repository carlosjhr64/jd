package jd

import "testing"
import "time"

func TestYMD(test *testing.T) {

  number := YMD2J(2016, 1, 13)
  if number != 2457401 {
    test.Error("Expected 2457401 for 2016-01-13, got ", number)
  }

  year, month, day := J2YMD(number)
  if year != 2016 || month != 1 || day != 13 {
    test.Error("Expected 2016 01 13, got ", year, month, day)
  }

  number, err := ToNumber("2016-01-13")
  if err != nil {
    test.Error("Did not expect an error, got", err)
  }
  if number != 2457401 {
    test.Error("Expected 2457401, got", number)
  }

  date := ToDate(number)
  if date != "2016-01-13" {
    test.Error("Expected 2016-01-13, got", date)
  }

  t, _ := time.Parse("2006-01-02", "2016-01-13")
  number = Number(t)
  if number != 2457401 {
    test.Error("Expected 2457401, got", number)
  }
}
