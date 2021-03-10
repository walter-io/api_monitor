package extends

import (
    "fmt"
    "github.com/liushuochen/gotable/color"
    "github.com/liushuochen/gotable/header"
    "reflect"
)

type Set struct {
    base []*header.Header
}

type Options struct {
    ColorController ColorController
}

type Table struct {
    Header *Set
    Value  []map[string]Sequence
    Opts   *Options
}

type ColorController func(field string, val reflect.Value) color.Color

type TableValue string

type Option func(t *Options)


// Sequence sequence for print
type Sequence interface {
    Value() string

    // Actual length except invisible rune
    Len() int

    // Origin value
    OriginValue() string
}


func (tb *Table) AddValue(newValue map[string]Sequence) error {
    for key := range newValue {
        value := reflect.ValueOf(newValue[key])
        clr := tb.Opts.ColorController(key, value)
        newValue[key] = color.ColorfulString(clr, value)
    }

    return tb.addValue(newValue)
}

func (tb *Table) addValue(newValue map[string]Sequence) error {
    for key := range newValue {
        if tb.Header.Exist(key) {
            continue
        } else {
            err := fmt.Errorf("invalid value %s", key)
            return err
        }
    }

    for _, head := range tb.Header.base {
        _, ok := newValue[head.Name]
        if !ok {
            newValue[head.Name] = TableValue(head.Default())
        }
    }

    tb.Value = append(tb.Value, newValue)
    return nil
}

func (set *Set) Exist(element string) bool {
    return set.exist(element) != -1
}

func (set *Set) exist(element string) int {
    for index, data := range set.base {
        if data.Name == element {
            return index
        }
    }

    return -1
}

func (s TableValue) Value() string {
    return string(s)
}

func (s TableValue) Len() int {
    return len(s)
}

func (s TableValue) OriginValue() string {
    return s.Value()
}

func (set *Set) Add(element string) error {
    if set.Exist(element) {
        return fmt.Errorf("value %s has exit", element)
    }

    newHeader := &header.Header{ Name: element }
    set.base = append(set.base, newHeader)
    return nil
}

func DefaultController(field string, val reflect.Value) color.Color {
    return ""
}

func (tb *Table) PrintTable() string {
    if tb.Empty() {
        //fmt.Println("table is empty.")
        return ""
    }

    columnMaxLength := make(map[string]int)
    tag := make(map[string]Sequence)
    taga := make([]map[string]Sequence, 0)
    for _, header := range tb.Header.base {
        columnMaxLength[header.Name] = len(header.Name)
        tag[header.Name] = TableValue("-")
    }

    for _, data := range tb.Value {
        for _, header := range tb.Header.base {
            maxLength := max(len(header.Name), data[header.Name].Len())
            maxLength = max(maxLength, columnMaxLength[header.Name])
            columnMaxLength[header.Name] = maxLength
        }
    }

    // print first line
    taga = append(taga, tag)
    returnStr := printGroup(taga, tb.Header.base, columnMaxLength)

    // print table head
    for index, head := range tb.Header.base {
        itemLen := columnMaxLength[head.Name] + 4
        s, _ := center(TableValue(head.Name), itemLen, " ")
        if index == 0 {
            s = "|" + s + "|"
        } else {
            s = "" + s + "|"
        }
        returnStr = returnStr + s
        //fmt.Print(s)
    }
    returnStr = returnStr + "\n"
    //fmt.Println()

    // print value
    tableValue := taga
    tableValue = append(tableValue, tb.Value...)
    tableValue = append(tableValue, tag)
    returnStr2 := printGroup(tableValue, tb.Header.base, columnMaxLength)
    return returnStr + returnStr2
}

func (tb *Table) Empty() bool {
    if len(tb.Value) == 0 {
        return true
    }
    return false
}

func printGroup(group []map[string]Sequence, header []*header.Header, columnMaxLen map[string]int) string {
    returnStr := ""
    for _, item := range group {
        for index, head := range header {
            itemLen := columnMaxLen[head.Name] + 4
            s := ""
            if item[head.Name].Value() == "-" {
                s, _ = center(item[head.Name], itemLen, "-")
            } else {
                s, _ = center(item[head.Name], itemLen, " ")
            }

            icon := "|"
            if item[head.Name].Value() == "-" {
                icon = "+"
            }

            if index == 0 {
                s = icon + s + icon
            } else {
                s = "" + s + icon
            }
            returnStr = returnStr + s
            //fmt.Print(s)
        }
        returnStr = returnStr + "\n"
        //fmt.Println()
    }
    return returnStr
}

func max(x, y int) int {
    if x >= y {
        return x
    }
    return y
}

func center(str Sequence, length int, fillchar string) (string, error) {
    if len(fillchar) != 1 {
        err := fmt.Errorf("the fill character must be exactly one" +
            " character long")
        return "", err
    }

    if str.Len() >= length {
        return str.Value(), nil
    }

    result := ""
    if isEvenNumber(length - str.Len()) {
        front := ""
        for i := 0; i < ((length - str.Len()) / 2); i++ {
            front = front + fillchar
        }

        result = front + str.Value() + front
    } else {
        front := ""
        for i := 0; i < ((length - str.Len() - 1) / 2); i++ {
            front = front + fillchar
        }

        behind := front + fillchar
        result = front + str.Value() + behind
    }
    return result, nil
}

func isEvenNumber(number int) bool {
    if number%2 == 0 {
        return true
    }
    return false
}