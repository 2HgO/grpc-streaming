package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type vidRange [2]int

func (v *vidRange) UnmarshalXMLAttr(attr xml.Attr) error {
	vals := strings.Split(attr.Value, "-")
	if len(vals) != 2 {
		return errors.New("Invalid attribute")
	}
	i, iErr := strconv.Atoi(vals[0])
	j, jErr := strconv.Atoi(vals[1])
	if iErr != nil || jErr != nil {
		return errors.New("Invalid attribute")
	}
	*v = [2]int{i, j}
	return nil
}

func (v *vidRange) Begin() int {
	return v[0]
}
func (v *vidRange) End() int {
	return v[1] + 1
}
func (v vidRange) String() string {
	return fmt.Sprintf("%d - %d", v[0], v[1])
}

type initSegment struct {
	Range vidRange `xml:"range,attr"`
}
type dataSegment struct {
	Range vidRange `xml:"mediaRange,attr"`
}
type videoManifest struct {
	XMLName xml.Name      `xml:"MPD"`
	Init    initSegment   `xml:"Period>AdaptationSet>Representation>SegmentList>Initialization"`
	Data    []dataSegment `xml:"Period>AdaptationSet>Representation>SegmentList>SegmentURL"`
}
