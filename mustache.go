package paragraph

import (
	"errors"
	"fmt"

	"github.com/alexkappa/mustache"
)

func (linesIn Paragraph) MustacheNoErr(m map[string]interface{}) (linesOut Paragraph) {
	linesOut, _ = linesIn.Mustache(m)
	return
}

func (linesIn Paragraph) Mustache(m map[string]interface{}) (linesOut Paragraph, err error) {
	linesOut = NewWithGivenLen(len(linesIn))
	moustache := mustache.New()

	for i, line := range linesIn {

		errt := moustache.ParseString(line)
		if errt != nil {
			if err == nil {
				err = errors.New("lines.mustache")
			}
			err = fmt.Errorf("%s;%w", err, errt)
		}

		l, errt := moustache.RenderString(m)
		if errt != nil {
			if err == nil {
				err = errors.New("lines.mustache")
			}
			err = fmt.Errorf("%s;%w", err, errt)
			linesOut[i] = linesIn[i]
		}
		linesOut[i] = l
	}
	return
}
