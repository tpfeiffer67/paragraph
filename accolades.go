package paragraph

func (linesIn Paragraph) Accolades(style AccoladesStyle) (linesOut Paragraph) {
	if style == AccoladesStyleNone {
		return linesIn
	}

	var left []string
	var right []string
	l := len(linesIn)
	linesOut = NewWithGivenLen(l)
	switch style {
	case AccoladesStyleAscii:
		left, right = getAccolades2(l)
	default:
		left, right = getAccolades(l)
	}
	for i := 0; i < l; i++ {
		linesOut[i] = left[i] + linesIn[i] + right[i]
	}
	return
}

func (linesIn Paragraph) AutoAccolades(style AccoladesStyle) (linesOut Paragraph) {
	if style == AccoladesStyleNone {
		return linesIn
	}

	w := linesIn.Width()
	return linesIn.PadRight(" ", w).Surround(" ", " ").Accolades(style)
}

func getAccolades(l int) (left []string, right []string) {
	left = make([]string, l, l)
	right = make([]string, l, l)

	switch l {
	case 0: // nothing to do
	case 1:
		left[0] = "{"
		right[0] = "}"
	case 2:
		left[0] = "⎰"
		left[1] = "⎱"
		right[0] = "⎱"
		right[1] = "⎰"
	case 3:
		left[0] = "⎧"
		left[1] = "⎫"
		left[2] = "⎩"
		right[0] = "⎫"
		right[1] = "⎧"
		right[2] = "⎭"
	case 4: // could have been processed in default section
		left[0] = "⎧"
		left[1] = "⎭"
		left[2] = "⎫"
		left[3] = "⎩"
		right[0] = "⎫"
		right[1] = "⎩"
		right[2] = "⎧"
		right[3] = "⎭"
	default:
		p := (l - 2) / 2
		left[0] = "⎧"
		left[p] = "⎭"
		left[p+1] = "⎫"
		left[l-1] = "⎩"
		right[0] = "⎫"
		right[p] = "⎩"
		right[p+1] = "⎧"
		right[l-1] = "⎭"
		for i := 1; i < l-1; i++ {
			if left[i] == "" {
				left[i] = "⎪"
				right[i] = "⎪"
			}
		}
	}
	return left, right
}

func getAccolades2(l int) (left []string, right []string) {
	left = make([]string, l, l)
	right = make([]string, l, l)

	switch l {
	case 0: // nothing to do
	case 1:
		left[0] = "<"
		right[0] = ">"
	case 2:
		left[0] = `/`
		left[1] = `\`
		right[0] = `\`
		right[1] = `/`
	case 3:
		left[0] = ` /`
		left[1] = `< `
		left[2] = ` \`
		right[0] = `\`
		right[1] = ` >`
		right[2] = `/`
	default:
		p := 1 + (l-2)/2
		left[0] = ` /`
		left[p] = `< `
		left[l-1] = ` \`
		right[0] = `\`
		right[p] = ` >`
		right[l-1] = `/`
		for i := 1; i < l-1; i++ {
			if left[i] == `` {
				left[i] = `▕ `
				right[i] = `▕`
			}
		}
	}
	return left, right
}
