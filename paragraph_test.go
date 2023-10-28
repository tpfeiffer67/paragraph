package paragraph

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWidth(t *testing.T) {
	assert := assert.New(t)
	lns1 := New(4)
	lns1 = append(lns1, "abcde12345")
	lns1 = append(lns1, "")
	lns1 = append(lns1, "¨")
	lns1 = append(lns1, "世界")
	assert.Equal(10, lns1.Width())

	lns2 := New(4)
	assert.Equal(0, lns2.Width())

	lns3 := New(3)
	lns3 = append(lns3, "")
	lns3 = append(lns3, "")
	lns3 = append(lns3, "")
	assert.Equal(0, lns3.Width())

	lns4 := New(3)
	lns4 = append(lns4, " ")
	lns4 = append(lns4, "世界")
	lns4 = append(lns4, "¨")
	assert.Equal(2, lns4.Width())

	lns5 := NewWithGivenLen(3)
	lns5[0] = "hippopotomonstrosesquipédaliophobie"
	lns5[1] = "世界"
	lns5[2] = "¨¨"
	assert.Equal(35, lns5.Width())
}

func TestWriteToFile(t *testing.T) {
	const fileName = "test.txt"
	assert := assert.New(t)
	lns := linesSample2(10)
	lns.WriteToFile(fileName)
	assert.True(compareGoldenFile(fileName))
	os.Remove(fileName)

	// ? is not a valid character for à file name
	err := lns.WriteToFile("?")
	assert.Error(err)
}

func linesSample1() (lns Paragraph) {
	lns = New(3)
	lns = append(lns, "Ceci est une  ligne relativement longue")
	lns = append(lns, "Ligne courte ¨")
	lns = append(lns, "Ceci est la troisième ligne")
	return
}

func linesSample2(n int) (lns Paragraph) {
	lns = New(n)
	// lorem ipsum alsacien https://www.alsacreations.com/page/schnapsum
	if n > 0 {
		lns = append(lns, "Lorem Elsass ipsum gal non hoplageiss")
	}
	if n > 1 {
		lns = append(lns, "vielmols, jetz gehts los picon bière")
	}
	if n > 2 {
		lns = append(lns, "tellus eget Hans quam, Christkindelsmärik auctor,")
	}
	if n > 3 {
		lns = append(lns, "leverwurscht amet gewurztraminer nüdle quam.")
	}
	if n > 4 {
		lns = append(lns, "T'inquiète, ch'ai ramené du schpeck,")
	}
	if n > 5 {
		lns = append(lns, "du chambon, un kuglopf et du schnaps dans mon rucksack.")
	}
	if n > 6 {
		lns = append(lns, "Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et ta wurscht ?")
	}
	if n > 7 {
		lns = append(lns, "Yeuh non che suis au réchime,")
	}
	if n > 8 {
		lns = append(lns, "je ne mange plus que des Grumbeere light et che fais de la chym.")
	}
	return
}

func ExampleMultiStrings_Cut() {
	lns := linesSample1()
	fmt.Println(lns.Cut(30))

	lns = linesSample1()
	fmt.Println(lns.Cut(10))

	lns = linesSample1()
	fmt.Println(lns.Cut(50))

	lns = linesSample1()
	fmt.Println(lns.Cut(1))

	lns = linesSample1()
	fmt.Println(lns.Cut(0))

	lns = linesSample1()
	fmt.Println(lns.Cut(-7))

	lns = NewFromString("Wie geht's les samis ? ")
	fmt.Println(lns.Cut(21))

	//Output:
	//Ceci est une  ligne relativeme
	//Ligne courte ¨
	//Ceci est la troisième ligne
	//
	//Ceci est u
	//Ligne cour
	//Ceci est l
	//
	//Ceci est une  ligne relativement longue
	//Ligne courte ¨
	//Ceci est la troisième ligne
	//
	//C
	//L
	//C
	//
	//Ceci est une  ligne relativement longue
	//Ligne courte ¨
	//Ceci est la troisième ligne
	//
	//Ceci est une  ligne relativement longue
	//Ligne courte ¨
	//Ceci est la troisième ligne
	//
	//Wie geht's les samis
}

func ExampleMultiStrings_Limit() {
	lns := linesSample1()
	fmt.Println(lns.Limit(50))

	lns = linesSample1()
	fmt.Println(lns.Limit(30))

	lns = linesSample1()
	fmt.Println(lns.Limit(10))

	lns = linesSample1()
	fmt.Println(lns.Limit(3))

	lns = linesSample1()
	fmt.Println(lns.Limit(1))

	lns = linesSample1()
	fmt.Println(lns.Limit(0))

	lns = linesSample1()
	fmt.Println(lns.Limit(-7))

	//Output:
	// Ceci est une  ligne relativement longue
	// Ligne courte ¨
	// Ceci est la troisième ligne
	//
	// Ceci est une  ligne
	// relativement longue
	// Ligne courte ¨
	// Ceci est la troisième ligne
	//
	// Ceci est
	// une  ligne
	// relativeme
	// nt longue
	// Ligne
	// courte ¨
	// Ceci est
	// la
	// troisième
	// ligne
	//
	// Cec
	// i
	// est
	// une
	// lig
	// ne
	// rel
	// ati
	// vem
	// ent
	// lon
	// gue
	// Lig
	// ne
	// cou
	// rte
	// ¨
	// Cec
	// i
	// est
	// la
	// tro
	// isi
	// ème
	// lig
	// ne
	//
	// C
	// e
	// c
	// i
	// e
	// s
	// t
	// u
	// n
	// e
	// l
	// i
	// g
	// n
	// e
	// r
	// e
	// l
	// a
	// t
	// i
	// v
	// e
	// m
	// e
	// n
	// t
	// l
	// o
	// n
	// g
	// u
	// e
	// L
	// i
	// g
	// n
	// e
	// c
	// o
	// u
	// r
	// t
	// e
	// ¨
	// C
	// e
	// c
	// i
	// e
	// s
	// t
	// l
	// a
	// t
	// r
	// o
	// i
	// s
	// i
	// è
	// m
	// e
	// l
	// i
	// g
	// n
	// e
	//
	// Ceci est une  ligne relativement longue
	// Ligne courte ¨
	// Ceci est la troisième ligne
	//
	// Ceci est une  ligne relativement longue
	// Ligne courte ¨
	// Ceci est la troisième ligne
}

func ExampleMultiStrings_Fill() {
	lns := linesSample1()
	w := 30
	fmt.Println(lns.Limit(w).PadRight(".", w))

	lns = linesSample1()
	w = 27
	fmt.Println(lns.Limit(w).PadRight("-.¨", w))

	lns = linesSample1()
	w = 27
	fmt.Println(lns.Limit(w).PadRight("", w))

	lns = linesSample1()
	w = 0
	fmt.Println(lns.PadRight(".", w))

	lns = linesSample1()
	w = 1010
	fmt.Println(lns.PadRight(".", w))

	//Output:
	//Ceci est une  ligne...........
	//relativement longue...........
	//Ligne courte ¨................
	//Ceci est la troisième ligne...
	//
	//Ceci est une  ligne-.¨-.¨-.
	//relativement longue-.¨-.¨-.
	//Ligne courte ¨-.¨-.¨-.¨-.¨-
	//Ceci est la troisième ligne
	//
	//Ceci est une  ligne
	//relativement longue
	//Ligne courte ¨
	//Ceci est la troisième ligne
	//
	//Ceci est une  ligne relativement longue
	//Ligne courte ¨
	//Ceci est la troisième ligne
	//
	//Ceci est une  ligne relativement longue
	//Ligne courte ¨
	//Ceci est la troisième ligne
}

func ExampleMultiStrings_Surround() {
	lns := linesSample1()
	fmt.Println(lns.Surround("|", ""))

	lns = linesSample1()
	fmt.Println(lns.Surround("", "|"))

	lns = linesSample1()
	w := 27
	fmt.Println(lns.Limit(w).PadRight(" ", w).Surround("(", ")"))
	//Output:
	// |Ceci est une  ligne relativement longue
	// |Ligne courte ¨
	// |Ceci est la troisième ligne
	//
	// Ceci est une  ligne relativement longue|
	// Ligne courte ¨|
	// Ceci est la troisième ligne|
	//
	// (Ceci est une  ligne        )
	// (relativement longue        )
	// (Ligne courte ¨             )
	// (Ceci est la troisième ligne)
}

func ExampleMultiStrings_Box() {
	lns := linesSample1()
	fmt.Println(lns.Box(BoxSettings{-2, "", LabelAlignLeft, "", LabelAlignLeft}, GetBoxPattern(BoxStyleSingleLine)))
	fmt.Println(lns.Box(BoxSettings{1005, "", LabelAlignLeft, "", LabelAlignLeft}, GetBoxPattern(BoxStyleSingleLine)))

	w := 30
	settings := BoxSettings{w + 2, "-=oOo=-", LabelAlignCenter, "¨", LabelAlignCenter} // +2 because of the Surround
	pattern := GetBoxPattern(BoxStyleDoubleLine)
	fmt.Println(lns.Limit(w).PadRight(".", w).Surround(" ", " ").Box(settings, pattern))

	lns = linesSample1()
	w = 30
	settings = BoxSettings{w, "▅▆▇ TITLE ▇▆▅", LabelAlignCenter, "▁▂▃▃▂▁", LabelAlignCenter}
	pattern = GetBoxPattern(BoxStyleFantasy3)
	fmt.Println(lns.Limit(w).PadRight(".", w).Box(settings, pattern))

	//Output:
	// Ceci est une  ligne relativement longue
	// Ligne courte ¨
	// Ceci est la troisième ligne
	//
	// Ceci est une  ligne relativement longue
	// Ligne courte ¨
	// Ceci est la troisième ligne
	//
	//╔═════════════-=oOo=-════════════╗
	//║ Ceci est une  ligne........... ║
	//║ relativement longue........... ║
	//║ Ligne courte ¨................ ║
	//║ Ceci est la troisième ligne... ║
	//╚════════════════¨═══════════════╝
	//
	// ▁▂▃▃▃▃▃▃▃▃▅▆▇ TITLE ▇▆▅▃▃▃▃▃▃▃▂▁
	// ▌Ceci est une  ligne...........▐
	// ▌relativement longue...........▐
	// ▌Ligne courte ¨................▐
	// ▌Ceci est la troisième ligne...▐
	// ▜▃▂▁▁▁▁▁▁▁▁▁▁▁▂▃▃▂▁▁▁▁▁▁▁▁▁▁▁▂▃▛
}

func ExampleMultiStrings_AutoBox() {
	lns := linesSample1()
	w := 30
	settings := BoxSettings{w, "Oo=-", LabelAlignLeft, "-=xX", LabelAlignRight}
	fmt.Println(settings)
	fmt.Println(lns.AutoBox(settings, GetBoxPattern(BoxStyleSingleLineRounded)))

	lns = lns.Limit(w)
	fmt.Println(lns.AutoBox(BoxSettings{w, "-=oOo=-", LabelAlignCenter, "-=xXx=-", LabelAlignCenter}, GetBoxPattern(BoxStyleDoubleLine)))

	pattern := GetBoxPattern(BoxStyleFantasy4)
	fmt.Println(lns.Surround(" ", " ").AutoBox(BoxSettings{w, "", LabelAlignLeft, "", LabelAlignLeft}, pattern))

	fmt.Println(linesSample1().Cut(8).AutoBox(BoxSettings{w, "Title", LabelAlignLeft, "Status", LabelAlignRight}, GetBoxPattern(-4)))
	fmt.Println(linesSample1().Cut(4).AutoBox(BoxSettings{w, "Title", LabelAlignLeft, "Status", LabelAlignRight}, GetBoxPattern(1000)))

	settings = BoxSettings{w, "", LabelAlignLeft, "", LabelAlignLeft}
	w = 10
	lns = linesSample1().Cut(w)
	for i := 0; i < 4; i++ {
		settings.Width = w + i*2
		lns = lns.AutoBox(settings, GetBoxPattern(BoxStyle(int(BoxStyleBlocksLightShade)+i)))
	}
	fmt.Println(lns)

	pattern = BoxPattern{"", "", "", "", "", "", "", ""}
	fmt.Println(linesSample1().AutoBox(BoxSettings{w, "", LabelAlignLeft, "", LabelAlignLeft}, pattern).Surround("[", "]"))

	//Output:
	// {30 Oo=- LabelAlignLeft -=xX LabelAlignRight}
	// ╭Oo=-───────────────────────────────────╮
	// │Ceci est une  ligne relativement longue│
	// │Ligne courte ¨                         │
	// │Ceci est la troisième ligne            │
	// ╰───────────────────────────────────-=xX╯
	//
	// ╔══════════-=oOo=-══════════╗
	// ║Ceci est une  ligne        ║
	// ║relativement longue        ║
	// ║Ligne courte ¨             ║
	// ║Ceci est la troisième ligne║
	// ╚══════════-=xXx=-══════════╝
	//
	// ▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃
	// █ Ceci est une  ligne         █
	// █ relativement longue         █
	// █ Ligne courte ¨              █
	// █ Ceci est la troisième ligne █
	// █▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃█
	//
	// ┌Title───┐
	// │Ceci est│
	// │Ligne co│
	// │Ceci est│
	// └──Status┘
	//
	// ┌Titl┐
	// │Ceci│
	// │Lign│
	// │Ceci│
	// └Stat┘
	//
	// ██████████████████
	// █▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓█
	// █▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓█
	// █▓▒░░░░░░░░░░░░▒▓█
	// █▓▒░Ceci est u░▒▓█
	// █▓▒░Ligne cour░▒▓█
	// █▓▒░Ceci est l░▒▓█
	// █▓▒░░░░░░░░░░░░▒▓█
	// █▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓█
	// █▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓█
	// ██████████████████
	//
	// [Ceci est une  ligne relativement longue]
	// [Ligne courte ¨]
	// [Ceci est la troisième ligne]
}

func ExampleMultiStrings_Accolades_styleunicode() {
	lns0 := New(0)
	fmt.Println(len(lns0.Accolades(AccoladesStyleUnicode)))

	for i := 1; i < 10; i++ {
		lns := linesSample2(i)
		fmt.Println(lns.Limit(55).PadRight(" ", 55).Accolades(AccoladesStyleUnicode))
	}

	//Output:
	// 0
	// {Lorem Elsass ipsum gal non hoplageiss                  }
	//
	// ⎰Lorem Elsass ipsum gal non hoplageiss                  ⎱
	// ⎱vielmols, jetz gehts los picon bière                   ⎰
	//
	// ⎧Lorem Elsass ipsum gal non hoplageiss                  ⎫
	// ⎫vielmols, jetz gehts los picon bière                   ⎧
	// ⎩tellus eget Hans quam, Christkindelsmärik auctor,      ⎭
	//
	// ⎧Lorem Elsass ipsum gal non hoplageiss                  ⎫
	// ⎭vielmols, jetz gehts los picon bière                   ⎩
	// ⎫tellus eget Hans quam, Christkindelsmärik auctor,      ⎧
	// ⎩leverwurscht amet gewurztraminer nüdle quam.           ⎭
	//
	// ⎧Lorem Elsass ipsum gal non hoplageiss                  ⎫
	// ⎭vielmols, jetz gehts los picon bière                   ⎩
	// ⎫tellus eget Hans quam, Christkindelsmärik auctor,      ⎧
	// ⎪leverwurscht amet gewurztraminer nüdle quam.           ⎪
	// ⎩T'inquiète, ch'ai ramené du schpeck,                   ⎭
	//
	// ⎧Lorem Elsass ipsum gal non hoplageiss                  ⎫
	// ⎪vielmols, jetz gehts los picon bière                   ⎪
	// ⎭tellus eget Hans quam, Christkindelsmärik auctor,      ⎩
	// ⎫leverwurscht amet gewurztraminer nüdle quam.           ⎧
	// ⎪T'inquiète, ch'ai ramené du schpeck,                   ⎪
	// ⎩du chambon, un kuglopf et du schnaps dans mon rucksack.⎭
	//
	// ⎧Lorem Elsass ipsum gal non hoplageiss                  ⎫
	// ⎪vielmols, jetz gehts los picon bière                   ⎪
	// ⎪tellus eget Hans quam, Christkindelsmärik auctor,      ⎪
	// ⎭leverwurscht amet gewurztraminer nüdle quam.           ⎩
	// ⎫T'inquiète, ch'ai ramené du schpeck,                   ⎧
	// ⎪du chambon, un kuglopf et du schnaps dans mon rucksack.⎪
	// ⎪Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et⎪
	// ⎩ta wurscht ?                                           ⎭
	//
	// ⎧Lorem Elsass ipsum gal non hoplageiss                  ⎫
	// ⎪vielmols, jetz gehts los picon bière                   ⎪
	// ⎪tellus eget Hans quam, Christkindelsmärik auctor,      ⎪
	// ⎭leverwurscht amet gewurztraminer nüdle quam.           ⎩
	// ⎫T'inquiète, ch'ai ramené du schpeck,                   ⎧
	// ⎪du chambon, un kuglopf et du schnaps dans mon rucksack.⎪
	// ⎪Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et⎪
	// ⎪ta wurscht ?                                           ⎪
	// ⎩Yeuh non che suis au réchime,                          ⎭
	//
	// ⎧Lorem Elsass ipsum gal non hoplageiss                  ⎫
	// ⎪vielmols, jetz gehts los picon bière                   ⎪
	// ⎪tellus eget Hans quam, Christkindelsmärik auctor,      ⎪
	// ⎪leverwurscht amet gewurztraminer nüdle quam.           ⎪
	// ⎭T'inquiète, ch'ai ramené du schpeck,                   ⎩
	// ⎫du chambon, un kuglopf et du schnaps dans mon rucksack.⎧
	// ⎪Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et⎪
	// ⎪ta wurscht ?                                           ⎪
	// ⎪Yeuh non che suis au réchime,                          ⎪
	// ⎪je ne mange plus que des Grumbeere light et che fais de⎪
	// ⎩la chym.                                               ⎭
}

func ExampleMultiStrings_AutoAccolades() {
	for i := 1; i < 10; i++ {
		lns := linesSample2(i)
		fmt.Println(lns.AutoAccolades(AccoladesStyleUnicode))
	}
	//Output:
	// { Lorem Elsass ipsum gal non hoplageiss }
	//
	// ⎰ Lorem Elsass ipsum gal non hoplageiss ⎱
	// ⎱ vielmols, jetz gehts los picon bière  ⎰
	//
	// ⎧ Lorem Elsass ipsum gal non hoplageiss             ⎫
	// ⎫ vielmols, jetz gehts los picon bière              ⎧
	// ⎩ tellus eget Hans quam, Christkindelsmärik auctor, ⎭
	//
	// ⎧ Lorem Elsass ipsum gal non hoplageiss             ⎫
	// ⎭ vielmols, jetz gehts los picon bière              ⎩
	// ⎫ tellus eget Hans quam, Christkindelsmärik auctor, ⎧
	// ⎩ leverwurscht amet gewurztraminer nüdle quam.      ⎭
	//
	// ⎧ Lorem Elsass ipsum gal non hoplageiss             ⎫
	// ⎭ vielmols, jetz gehts los picon bière              ⎩
	// ⎫ tellus eget Hans quam, Christkindelsmärik auctor, ⎧
	// ⎪ leverwurscht amet gewurztraminer nüdle quam.      ⎪
	// ⎩ T'inquiète, ch'ai ramené du schpeck,              ⎭
	//
	// ⎧ Lorem Elsass ipsum gal non hoplageiss                   ⎫
	// ⎪ vielmols, jetz gehts los picon bière                    ⎪
	// ⎭ tellus eget Hans quam, Christkindelsmärik auctor,       ⎩
	// ⎫ leverwurscht amet gewurztraminer nüdle quam.            ⎧
	// ⎪ T'inquiète, ch'ai ramené du schpeck,                    ⎪
	// ⎩ du chambon, un kuglopf et du schnaps dans mon rucksack. ⎭
	//
	// ⎧ Lorem Elsass ipsum gal non hoplageiss                                ⎫
	// ⎪ vielmols, jetz gehts los picon bière                                 ⎪
	// ⎭ tellus eget Hans quam, Christkindelsmärik auctor,                    ⎩
	// ⎫ leverwurscht amet gewurztraminer nüdle quam.                         ⎧
	// ⎪ T'inquiète, ch'ai ramené du schpeck,                                 ⎪
	// ⎪ du chambon, un kuglopf et du schnaps dans mon rucksack.              ⎪
	// ⎩ Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et ta wurscht ? ⎭
	//
	// ⎧ Lorem Elsass ipsum gal non hoplageiss                                ⎫
	// ⎪ vielmols, jetz gehts los picon bière                                 ⎪
	// ⎪ tellus eget Hans quam, Christkindelsmärik auctor,                    ⎪
	// ⎭ leverwurscht amet gewurztraminer nüdle quam.                         ⎩
	// ⎫ T'inquiète, ch'ai ramené du schpeck,                                 ⎧
	// ⎪ du chambon, un kuglopf et du schnaps dans mon rucksack.              ⎪
	// ⎪ Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et ta wurscht ? ⎪
	// ⎩ Yeuh non che suis au réchime,                                        ⎭
	//
	// ⎧ Lorem Elsass ipsum gal non hoplageiss                                ⎫
	// ⎪ vielmols, jetz gehts los picon bière                                 ⎪
	// ⎪ tellus eget Hans quam, Christkindelsmärik auctor,                    ⎪
	// ⎭ leverwurscht amet gewurztraminer nüdle quam.                         ⎩
	// ⎫ T'inquiète, ch'ai ramené du schpeck,                                 ⎧
	// ⎪ du chambon, un kuglopf et du schnaps dans mon rucksack.              ⎪
	// ⎪ Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et ta wurscht ? ⎪
	// ⎪ Yeuh non che suis au réchime,                                        ⎪
	// ⎩ je ne mange plus que des Grumbeere light et che fais de la chym.     ⎭
}

func ExampleMultiStrings_Accolades_styleAscii() {
	lns0 := New(0)
	fmt.Println(len(lns0.Accolades(AccoladesStyleAscii)))

	for i := 1; i < 10; i++ {
		lns := linesSample2(i)
		fmt.Println(lns.Limit(55).PadRight(" ", 55).Accolades(AccoladesStyleAscii))
	}
	//Output:
	// 0
	// <Lorem Elsass ipsum gal non hoplageiss                  >
	//
	// /Lorem Elsass ipsum gal non hoplageiss                  \
	// \vielmols, jetz gehts los picon bière                   /
	//
	//  /Lorem Elsass ipsum gal non hoplageiss                  \
	// < vielmols, jetz gehts los picon bière                    >
	//  \tellus eget Hans quam, Christkindelsmärik auctor,      /
	//
	//  /Lorem Elsass ipsum gal non hoplageiss                  \
	// ▕ vielmols, jetz gehts los picon bière                   ▕
	// < tellus eget Hans quam, Christkindelsmärik auctor,       >
	//  \leverwurscht amet gewurztraminer nüdle quam.           /
	//
	//  /Lorem Elsass ipsum gal non hoplageiss                  \
	// ▕ vielmols, jetz gehts los picon bière                   ▕
	// < tellus eget Hans quam, Christkindelsmärik auctor,       >
	// ▕ leverwurscht amet gewurztraminer nüdle quam.           ▕
	//  \T'inquiète, ch'ai ramené du schpeck,                   /
	//
	//  /Lorem Elsass ipsum gal non hoplageiss                  \
	// ▕ vielmols, jetz gehts los picon bière                   ▕
	// ▕ tellus eget Hans quam, Christkindelsmärik auctor,      ▕
	// < leverwurscht amet gewurztraminer nüdle quam.            >
	// ▕ T'inquiète, ch'ai ramené du schpeck,                   ▕
	//  \du chambon, un kuglopf et du schnaps dans mon rucksack./
	//
	//  /Lorem Elsass ipsum gal non hoplageiss                  \
	// ▕ vielmols, jetz gehts los picon bière                   ▕
	// ▕ tellus eget Hans quam, Christkindelsmärik auctor,      ▕
	// ▕ leverwurscht amet gewurztraminer nüdle quam.           ▕
	// < T'inquiète, ch'ai ramené du schpeck,                    >
	// ▕ du chambon, un kuglopf et du schnaps dans mon rucksack.▕
	// ▕ Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et▕
	//  \ta wurscht ?                                           /
	//
	//  /Lorem Elsass ipsum gal non hoplageiss                  \
	// ▕ vielmols, jetz gehts los picon bière                   ▕
	// ▕ tellus eget Hans quam, Christkindelsmärik auctor,      ▕
	// ▕ leverwurscht amet gewurztraminer nüdle quam.           ▕
	// < T'inquiète, ch'ai ramené du schpeck,                    >
	// ▕ du chambon, un kuglopf et du schnaps dans mon rucksack.▕
	// ▕ Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et▕
	// ▕ ta wurscht ?                                           ▕
	//  \Yeuh non che suis au réchime,                          /
	//
	//  /Lorem Elsass ipsum gal non hoplageiss                  \
	// ▕ vielmols, jetz gehts los picon bière                   ▕
	// ▕ tellus eget Hans quam, Christkindelsmärik auctor,      ▕
	// ▕ leverwurscht amet gewurztraminer nüdle quam.           ▕
	// ▕ T'inquiète, ch'ai ramené du schpeck,                   ▕
	// < du chambon, un kuglopf et du schnaps dans mon rucksack. >
	// ▕ Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et▕
	// ▕ ta wurscht ?                                           ▕
	// ▕ Yeuh non che suis au réchime,                          ▕
	// ▕ je ne mange plus que des Grumbeere light et che fais de▕
	//  \la chym.                                               /

}

func ExampleMultiStrings_Sort() {
	lns := linesSample2(9)
	fmt.Println(lns.Sort())
	//Output:
	// Allez, s'guelt ! Wotch a kofee avec ton bibalaekaess et ta wurscht ?
	// Lorem Elsass ipsum gal non hoplageiss
	// T'inquiète, ch'ai ramené du schpeck,
	// Yeuh non che suis au réchime,
	// du chambon, un kuglopf et du schnaps dans mon rucksack.
	// je ne mange plus que des Grumbeere light et che fais de la chym.
	// leverwurscht amet gewurztraminer nüdle quam.
	// tellus eget Hans quam, Christkindelsmärik auctor,
	// vielmols, jetz gehts los picon bière
}

func ExampleAlignment() {
	for i := 0; i < 3; i++ {
		fmt.Println(LabelAlign(i))
	}
	//Output:
	// LabelAlignLeft
	// LabelAlignCenter
	// LabelAlignRight
}

func ExampleBoxStyle() {
	for i := 2; i <= BoxStyleMaxIndex; i++ {
		fmt.Println(NewFromString(BoxStyle(i).String()).PadRight(" ", 38).Surround(" ", " ").AutoBox(BoxSettings{40, "", LabelAlignLeft, "", LabelAlignLeft}, GetBoxPattern(BoxStyle(i))))
	}
	//Output:
	// ┌────────────────────────────────────────┐
	// │ BoxStyleSingleLine                     │
	// └────────────────────────────────────────┘
	//
	// ╭────────────────────────────────────────╮
	// │ BoxStyleSingleLineRounded              │
	// ╰────────────────────────────────────────╯
	//
	// ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	// ┃ BoxStyleBold                           ┃
	// ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	//
	// ╒════════════════════════════════════════╕
	// │ BoxStyleSingleVDoubleH                 │
	// ╘════════════════════════════════════════╛
	//
	// ╓────────────────────────────────────────╖
	// ║ BoxStyleSingleHDoubleV                 ║
	// ╙────────────────────────────────────────╜
	//
	// ╔════════════════════════════════════════╗
	// ║ BoxStyleDoubleLine                     ║
	// ╚════════════════════════════════════════╝
	//
	// ▛▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▜
	// ▌ BoxStyleExtraBold                      ▐
	// ▙▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▟
	//
	// ▞▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▚
	// ▌ BoxStyleExtraBoldRounded               ▐
	// ▚▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▞
	//
	// █▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█
	// █ BoxStyleMaxBold                        █
	// █▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█
	//
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	// ░ BoxStyleBlocksLightShade               ░
	// ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
	//
	// ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
	// ▒ BoxStyleBlocksMediumShade              ▒
	// ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
	//
	// ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
	// ▓ BoxStyleBlocksDarkShade                ▓
	// ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
	//
	// ██████████████████████████████████████████
	// █ BoxStyleBlocks                         █
	// ██████████████████████████████████████████
	//
	// ..........................................
	// : BoxStyleDots                           :
	// :........................................:
	//
	// ◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆
	// ◆ BoxStyleDiamonds                       ◆
	// ◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆
	//
	// ╭╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╼╮
	// ╽ BoxStyleFantasy1                       ╿
	// ╰╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╾╯
	//
	// ╱▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔╲
	// │ BoxStyleFantasy2                       │
	// ╲▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁╱
	//
	// ▁▂▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▂▁
	// ▌ BoxStyleFantasy3                       ▐
	// ▜▃▂▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▂▃▛
	//
	// ▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂
	// █ BoxStyleFantasy4                       █
	// █▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂▃▂▁▂█
}

func ExampleAccoladesStyle() {
	for i := 0; i < 3; i++ {
		fmt.Println(AccoladesStyle(i))
	}
	//Output:
	// AccoladesStyleNone
	// AccoladesStyleAscii
	// AccoladesStyleUnicode
}

func TestMultiStrings_AccoladesStyleFromString(t *testing.T) {
	assert := assert.New(t)
	var v AccoladesStyle

	v, err := AccoladesStyleFromString("AccoladesStyleAscii")
	assert.NoError(err)
	assert.Equal(AccoladesStyleAscii, v)

	v, err = AccoladesStyleFromString("Ascii")
	assert.NoError(err)
	assert.Equal(AccoladesStyleAscii, v)

	v, err = AccoladesStyleFromString("Unicode")
	assert.NoError(err)
	assert.Equal(AccoladesStyleUnicode, v)

	v, err = AccoladesStyleFromString("AccoladesStyle")
	assert.Equal(AccoladesStyle(0), v)
	assert.Error(err)

	v, err = AccoladesStyleFromString("foo")
	assert.Equal(AccoladesStyle(0), v)
	assert.Error(err)

	v, err = AccoladesStyleFromString("")
	assert.Equal(AccoladesStyleNone, v)
	assert.Error(err)

	s := fmt.Sprint(AccoladesStyleUnicode)
	assert.Equal("AccoladesStyleUnicode", s)
}

func ExampleMultiStrings_Append() {
	lns := linesSample2(2)
	lns2 := linesSample2(2)
	fmt.Println(lns.Append(lns2))
	fmt.Println(lns.Append(lns2).AutoBox(BoxSettings{1, "", LabelAlignLeft, "", LabelAlignRight}, GetBoxPattern(BoxStyleSingleLine)))
	fmt.Println(lns.Append(NewFromString("T'inquiète, ch'ai ramené du schpeck\ndu chambon et un kuglopf.")))
	//Output:
	// Lorem Elsass ipsum gal non hoplageiss
	// vielmols, jetz gehts los picon bière
	// Lorem Elsass ipsum gal non hoplageiss
	// vielmols, jetz gehts los picon bière
	//
	// ┌─────────────────────────────────────┐
	// │Lorem Elsass ipsum gal non hoplageiss│
	// │vielmols, jetz gehts los picon bière │
	// │Lorem Elsass ipsum gal non hoplageiss│
	// │vielmols, jetz gehts los picon bière │
	// └─────────────────────────────────────┘
	//
	// Lorem Elsass ipsum gal non hoplageiss
	// vielmols, jetz gehts los picon bière
	// T'inquiète, ch'ai ramené du schpeck
	// du chambon et un kuglopf.
}

func compareGoldenFile(fileName string) bool {
	got, err := os.ReadFile(fileName)
	if err != nil {
		return false
	}
	wanted, err := os.ReadFile(filepath.Join("testdata", fileName+".golden"))
	if err != nil {
		return false
	}
	return bytes.Equal(got, wanted)
}
