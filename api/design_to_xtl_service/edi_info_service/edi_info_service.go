package edi_info_service

import "strings"

func EdiInfo(designName string) (segmentTag, position, subPos string) {
	ediInfoChunks := strings.Split(designName, "-")
	segmWithPosition := ediInfoChunks[0]
	if len(ediInfoChunks) > 1 {
		subPos = ediInfoChunks[1]
	}
	if len(segmWithPosition) > 1 {
		segmentTag = segmWithPosition[:len(segmWithPosition)-2]
		position = segmWithPosition[len(segmWithPosition)-2:]
	}
	return
}
