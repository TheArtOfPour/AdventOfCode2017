package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	input := `{{{},{{{{<e""a!!io<"!!!!"!'!!"u!>},<"o!!}i>,<!>!i<!!!!oi!>},<>},{},{{<!>,<{!u!!!>,<!!'<!>!,,!!"!e'o}>}}},{},{<!!!!"a"<!!eo,!>}!!!>,io!>'!o>}},{{{<!!!'ou!!!>!oe!>!<!!i>,{<!>},<uoue>}},{},{{<}}>,<"'',!!}{!!i{>}}},{{{{<'o!>>,<!>!{u}<e!>a<au!!o!>,<a!!""a!!!>">},{{<i{!>,<"!>e!!'!!!!!>},<'o!!!,!e<o<>},{{}}}},{{<e'!!!!!!i"a!>,<,>},{<'!!!><!}!!{!>},<,"u>}}},{{{{<a{o!!!!!>!><<!!!>},<>}},{<!!"!>e>,{<!>},<"}a"''!!!!!!!>o<}!!,>}}},{{{<!,}!"}!!!>,<'!!!!e!>},<!>},<{!>},<,,!!}}!!u>},{{}}},{{{<i!>,<!><>}},{{}}},{{}}},{{{<>},<!>},<eu'!>},<u!''!!!>!!!>!!<>},{<>}},{{},{<!<<!!e!!!>,<>,{}},{{},{{<>},{{<!>!>!!!!a!>,<!!e}e!>{>}}}}}},{{{},{}}}}},{{{<a}a>}},{<!>,<!!!!'!>,<!!!!'o"a!!eea"!>"e!u>,{{{<!!!!{ua>}}}},{{<}!!<,e>},<<!>,<'{!{ei!!!>>}}}},{{{{{<'!a>,<!!i<o!,!!"<i!>,<!>},<o!!e}!!e<!!!>a>},{{<!!}!>,<!i}o!!i!>"!>},<!>},<e'!!<!>!>,<{!>,<>,{}},{{},{{}}},{<!!u'!>,<}!!!!!>!!!>!>},<!>"!!!>!!ui'>}}},{{<u!{<!!u{{!>},<!>},<a!!!!!!,!!!>},<"{>}},{{},{{{<ei!!!{"!>{!a<!><}<!>},<!!e>,{<!!!>'!!a<!!!>},<u}o<aa'!>,<>}},{<",>}},{<!!!>{!!}!e!>,<<{!>},<<!>,<!<!>},<>,{<'!e!>},<!>,<i!!!!!>!>!>!'>}}},{{<<!>!!{>},{{<!oa{>},{<!>i!{>}}},{{{{{{{<!}ou<,!>,<!!a!>},<">,<,}i!!!>!}!>!!!!!>},<e"i!>io,!>},<!!!>}!!i>}}},{{},{<!!!!!!>,{{},{}}}},{<!>,<>}}},{{{<,!>i!!!>},<!>!{ue!!i>},{{{<,!>!>},<}!!{u'!}!>!!!>!>},<>}},{<u!u!!<<{!!<,}e{ua'>}}},{<!>!!!>!!!!!>!io>}},{{{{<e,iu"!>,<!!!>!>e!{'>,{<'e!!'!,u!>},<u>}},{{<o'!!!>,<!>,<!!!>o!>a<u!>},<ao!!a!o},>}}}},{{{{}},{{<!!'!>},<!!!>},<!!aei<,!>>},{<!!!>!!!>"}!oae>},{{<"!!!<!!!>!>,<'!!!>a,<"!>},<o!>>}}},{{{<!>,<'uo{!{>}},<}!!}<!!!!<e}u!!!>'a!>},<>}},{{{<!!i!!!!{<!!!!oo!!a}'!!!>!!'!!>},<!!!>,{oi{!!,!e!!!<>},{{},{}},{{<e<<!>,<!>,<!!!>},<}"{e<!euue"!"o>},{<e!!!!!>oa>}}}},{{{<!>,<'<!!!>eo!>,<!aa},>,<'"!!!!!>!!e"!!!>'}!!>},{{{<"!>},<{a!!'!!!>!!"!'a!!!!!>!!u>},{<'>}}},{{<a!>i!>},<!'e!>!>!!!><e!>!>},<">},{<<>}}},{{<<u!i}!>,<!!!>>},{}},{{<ao'u!>},<'u{',uo>,<!!a'{ue,>},{}}},{{<!>!!e!>,<e>,{{<!>},<{'!>,<!!!>'a,!!!>!'!>,<<!>!!!<!>,<<{">}}},{<!>},<!}oii!o"!>,<>}}},{{<<!>,<}!"!>!>},<!!!>u>}}},{{<ei!!>}},{{{<!>>,{<>}},{{<!,oa}!!!>o{>},{}},{}},{{<">}},{<o!!!>"uaou!>,"'>,{<!>},<<o",!!!>!>u!e!>,<ui!>},<<>,<,!}}}"!!{e!>,<aeaia'!<>}}}}}},{{{}},{{<!>},<!!!>!!!>,o,o!!o!!!><>}},{{<}'!!!>!>!!a}>,{}},{}}}}},{{{{{{<>},{<eui'a!!<!>,<ie!!<}!u!>},<!>},<{!>,<!!>}},{{}},{<!!!>">,<aou!!,!>,<}"{ee>}},{{}}},{{{{<!!!!"uou>}},<<'"!o,e{a!>},<!>,<">},{{<<!>},<!>},<!>},<i!uu<u!!!!ooie!>},<!>{o>},<a>},{{{},<>},{{<o!!!>ao!!{}{u!>},<,i"!!>}}}},{{<">},{{<!!!><!>},<'!!!>"!>},<"!>,<,""!}ua!>"a>}}},{{{},<<!!!>!>,<a!!!>!>},<'oe!!!>,<!>,<!!!!!>!!!!!!i'!uieo>},{{<!!!>!>},<!!o'i}o!!!!ae!!{"!>>,<!>,<!>},<<e,,!!>},{{<!!!!!>u!!!>{!!oi!>>},{}},{{<a"!eii!!!aa,,',!>,<io>},{<!!o!>},<,>}}}}},{{},{{{<i!!!!!>,<!>},<'"}!!!>}u"o!>,<"!}<<!!>},{{<io!>,<!!!!a!!!>e!!!>!,!!!>ae!a!!"<!!}!>,<e>}}},{<<"!>},<!!>,{<u}>}},{{},{<!oo'!>i!>},<!>,,!!'<!!!>e>}}},{{{<!!!>'{>},<!>,<!>},<}!o>},{<}e!!!>!>,'!>,<!!>,{}},{{},<>}},{}},{{{{<'!!!!!>o!!!,!>,<"'>}}},{{{<!>},<>}},{},{{<e!u"'>,{}}}}}},{{{{{{<!<'u}u<!!">},{{<""ie!"{>}}}},{<}!!!!e<'ioa!>},<"!>},<!eo!>},<!>,<!!u>}},{},{{<<<>,{<!{!>},<!!{!!,!>!a!>"{!!!>!{!!!!iei<>,{{<!!a!u'!!i!aeu!!'!>},<'u!!u">},{{<!!!>},<,"oe!<{!,u!!!>!"!!!>>}}}}},{{<iiia>},<<!!!>!!!!,>},{{<!><!!u>},{<!!!>,<'!>"!>,<>}}}},{{{{},{{{<aeu!>,<{e!!!>!!!!!">},<<!!!>},<"!>u!'{',>},<!!!!!>a!"o",!a!'!u}!>},<"a!>},<>}},{<>,<u!>},<{!!!>,o!!aa"!!'!!!!"!!!>ie>}},{{{<!!a!>,<!!u>}},{{{{<!{'!!!!!><{!>ii>}}},<o!!!>!!!>e!!{!>,u!!o!!!!o>},{{<>}}},{{{{<{!!!!i"i<o!!"!!ui}}>,<!<!>"ui"!!<!>o!>},<!!}!>'",>},{<!!e!>>}},{{{{{<!!ouee}o!"!!e!<a!>},<e,}!!'>},<{o<,o"<a!!!>i!!!>},<,!!!'!!o!>,<u!>>},<<"!>,<{o{!}i!!<>}}},{{{{<!!{!!!>'a'"!ei>},{<!!!!!>,<,!!!>'"e!{<!>,<>}}}}}},{{<"!!!>ai'>},{{{<!>,<"ou"u!>}!!'ue!>},<'o}!>},<a!a>}},<i!>},<"i,u!!!>!!!!!>i"{!!!>>}}},{{{{<!a!!!>!o}!!!>},eo!!!>!>},<!>!>},<!>",!!!!u,>,<!>!!ie,o<',<}!>"oe!>,<!!>},{{<!!!>!i{u>},<u!!i!>},<'o!}eu{}!e>},{{<"<u!>!>,<<!!o'!!!!!!!!!!e!,!!!!'!!!>,'!!!>>},<e'!!o>}},{},{{{{{<"!!!>!!!>}o!>},<u}u}!!e!!!>!!e!!e!!!>>}}}},<!>,<o!!eo!!"!!!>>}},{{{{<{'!iu!>,<}!>},<}>},{<'i>}},{<}ea!!{,o!!!>,<'}oi<,!}a">,<!!!>,<!!!!u,"<i>}},{{<!!!>aei{!>,<!>,<'e<>},{}},{{{},<!>,<>}}},{{<>,<o!>,<i!>},<>}}},{{{{<,"oo!>},<!!!>i<'!!{!>a!!i"!"!!>,<!!i!>i!!a!!!!!>{>}},{{<a>}}},{{{<e{!!!!'>,{<!!o"{,!!!!>}},{{<iia{e<<a!!!!!!'>},<!!!>},<!!o!!!>,<!>},<"!!!!!!!>!!!>!!!!uau<>},{<i"a!!!>a!>{oo,!>},<a!!!>!!ii<}>,{<"!ia!><o!!!ao!!oe,,i'>}}},{{{<"o,!>,!a!!e>},{},{{{<>},<!>},<auii!!!><{>},{<!ia<iu'!u!>!!!>},<!>},<o"!!,>}}},{<!!!>},<!!!!!>'!>},<<!!,"!!{}<>,<o!!!>!>!>}!!!>!>!!!>!>,<!>,<i!!uu<ou">},{{<!>u!>!{,a!}!>,<},i!>,<!!!!""a!!!>,<>}}},{},{{},{<}!>e!!>,<!o!>!!!>},<!"!!u!!!>!!'u!>,<!>>}}},{{{{{<!!{"<u}!!!!!>!!o!i"!!!>!>i!!<}o!>},<!!'>}},<<{<>},{{{<!>},<o!>},<i!!!>!>,<'!>},<>}},{},{{<!>{}!>,<!!,,!!,!!">},{}}},{{{{{<"!!,}{a!!>},<,!{eau>},{<!o!oo!i}!>,<!o>}},{{{},{}}}},{<ao>},{<,!!!>a!>,<'>,{}}}},{{<!a'!!!!!,!{!>!>,<}!>},<!!"!!!>}<,>},{{<>},{}},{<""a!!!>>}}},{},{{{{<!>a!>},<<!!ii!!!>!!!,!!>},{<!!oe!{!>},<a'!>,<!!a!,!>>}},{{<!!>},{{<<ouuu,o",a<!!!}!>},<{{!>o!>,<>},<!!!>,<}o>}},{{<!>},<u!!!!o>,{}}}},{{{{{<a,>},{<!a!!<ii"<!"!!!>!!!>!>,<!!!>,aaa!!!>!>,<!!<>}}},{{<!!!>!!'e{!,",!!!>ii!>">},{{},<}u{!!!>!!!>!!"!!!>}!!">}}},{{<!!u<<e!>iu!!!>!>},<!>,<!>,<>},{},{{<i'!>"!!!!!>u!!!!!>!>,i>},<i{u!<,}!!}a!>,!>!>!!{"!!a>}},{{},{}},{{{<!!o!!!!!<u!>,<}!e>},{<a,!>},<!!!!!>!>,<!{!!!>!>!ao"!a>}},{{{{{<>}},{{},<>},{}},{},{<}e>,{}}}},{}}},{{{<}"!!!>e!!!><>}},{{<'!!!!"u'}aou,a>,{<<e!!,'e""!>!>,<{!>},<!!o!>oa>,{{},<u,!!'!!!><">}}},{{<'!!!>e{!>},<!eo!i!><",>},{{<!!!>a,u!>!>,<}u!!}!>},<>},<",!>},<"!a!!"!>!>},<>}},{<o!!"a!>!!}eiii{>,<!!!>!>!>}!>!!!>,<!>!!!!!!i>}},{{{<e}!>,<{!!a!>!>,<!!aoo!>,<ai!>},<>},{<"o}'"!uo<!{'!!!>{'o'"{>,{}}},{<!!"'"!!}!>}>}}},{{{{{<!>},<e!>},<>},{}},{<{'>},{<o!!e!>,<>,{{}}}},{{},{<u,o!!!!i!!!>},<"!>,<{>},{}},{{{<!>">,<'!!,{!>!!!>,>}},{{{<{<!>},<!"!>},<'>,<!>o!!!!!!"<oe{"<!>!>,<!!aa!>},<>},{<e}"<!>},<!>ie!e!!"{auu<oie>}}}},{{{<}>},<uo!!<{"!!}e'o!>,<a!>},<!>,<!!!>!!!>,<>}}},{{{<<!!!!!!!>!>!>}}{{!!o"e>}},{{}},{<!>},<!!!a!!'!!!>!>e>}},{{},{},{{<>}}},{{{<!!!>u!!!>eo!o}u!}!{!>},<}'i!>},<>,<,!!!>},<'!!!!!!!>'uu!!a<!>,<!!!>,<!!!>e>},{<u!!!>"<!!i{aio>}},{{{<!,>},<i>},{<!!!>e!!!>!!!!u}o!!<!!!>">}},{{<>},{<e!>,<!>},<'!>},<<!!!!!>{}!<}"!>,<!>},<!>!>e>},{<{'{!>!!ae!!>}}}}}}},{{{{{{<">},<}oea!>,<}o>},{<!>,<<!oa!!!!}<u,!>!e!!!o>,{}},{{<!!!>!>,<'a{'>}}}},{{{{},{{<>}}},{<!>,<!!!>}eae'!uuo!!}!!<>,{{<"'i>}}},{<o!>},<}ae!!!>e'{!!!><>}},{{{{<}!>},<!!u!!!>e<'a!!!>,<}ee}!>,!>},<"<">,{<e!,!>"!>!!i!!>}},{{{{<>}},{}},{{<uoi!!<!!!>,<{u}>}},{}},{<>,{<!!!!!!{!!!>,<!>o!!!>},<!!"!!{o>}}},{{{<!>u!!!>!eoeo{!!<!{!!!>,<}i!>},<>,<,!}a!>,<oa<'>},{<o{>},{<"e>,{}}},{<!!u!!!{{>,<e!>,<i!!u!!!!!>{!!!>i!!>}},{{{},<!,!!!>,<u!!!>,<'!!!>},<!>},<!>},<!>},<"!>>}},{{<o!!u!!}!!,!u!!"a>},{<<oua<au'!!!!!!,!>i!!,!>},<',>}}}},{{{<'ei!!!>{!>,<i!!""<"!!!>!!!>>},<oauae!!!>},<!!!>ieu!>,<!!u<!!o{!!e!>>},{{{<!ea!!!!!>!<>}}},{<a!!!!a>,{}}},{{{{{<u{>,{{{{<!a!a!>},<!!{!!!>!!"!>,<i<o"!{!!!>!!!>},<a,!!!!,>}}},{}}},{{<!!}!!!!!>!>"{a!!{!>,<!!"u!!{o>},{<>}},{}}},{}}}},{{{{{<!>,<!>},<u!>o!>e!!,>}}},{},{<!>,<'!!aoa!!!>!!!!!!!>!!!>,<',!!a!>i>}},{{<!","!>!!!>},<oeia,!>}{}">,{<!!,,e!!!>,<!>,i"!<oe!><'{!o>}}}},{{{{{<"e!!i!>>,{<!>!>},<iu!>},<<e'!>!!e<o!>,<!>,<<>}}},{<,"o!e<<!!!!!eu!!!>o!io!>,<!!'oe!!'>,{}}},{{},{}},{{},{{{<!>},<!!oa!!i!>,<!!ea!!>},<o!>},<!!!>!!<>},{{}},{{{<!>'!!e!>},<!>,<a}!>},<!!!>!!!>"}}!!i,'>},<{>},{<!>},<<!!!>!>},<!!}>}}},{{{<'i{i'!>!!e!!!>!{ie!!!!!>,<{!>{!!!!a">},{<!}e!{!>},<ui}{o"!>},<e!!}!""e''!>>}},{<<o!!!>{i,!{>,{{{<!>!!,>}},{<!!!>!>!!!!}iu,!>},<{o!>},<ai!o}<"!!!>>}}}}}}},{{{{<!>,<!!{!!!>},<!!<>},{{{{{<!!!>o{>,{}}},{<!!}'e!!!!!ea!>,<{>}}},{{<}!>,<!>,<e!,'">,{{},<>}},{}},{{{<!>!<{,!>iu{'"o>,{<!!>}},{{<!!!>u>},{}},{}},{{<i"!!!>,ee!!''i!><!!'>},<{!>,<!>,,o,o,i!>},<!!!>,<!>}!>},<"!!!>},<{!!!o>},{{{<!e!u!!!'}!!!!!>!!>,{<o!!!>>,{{{},{{<{!!!>'!>,<!>!>!!o!!!>}!!!>>}}},<!>},<e!>},<!>},<!!!>!>u"e"}!a!'a!u{>}}},{{{{<e!!,<!!!!!>,<'{a!>,<}"!>,<i"<!>},<a<>,{}},<}u!>u,a!>!!!>u!!!!!>u!<!!',>},{{<,u<'!>!>},<>}},{{{{},<}!>},<{ee"}"!>!!<>}},{{<!!a,ae!!!>!!e<'>},<!!!>io!'!>"!!i!!}!>,<!!o,i>},{<!!!,u'"!>,<!!!>!>,<!!!>ei!>!!u!>o,!>},<>}}}}},{{<a!!!>{!>,<o!!ea!!},e}!!{>},{{{<e<}>}}}}}}},{{<!!au!uu!!'<"!}!>u<!!,!!a!>"<o'>},<}!>},<>}},{{<!!{ao!!!>},<!>,<!!!!'{a!!!>aa!'!!!>>},{}}},{{{<'}<!!!>!!>,<!>!>},<!>,<}!!!>!!a'!!!>}!>},<!!!>!',"!>}!>},<!>},<>},{{{<!>,<ae{,!a!e!!a!!!>"{}'!!!>!!!>!!!>{>},{<o,!!"}!!!>!>'!>ea!!!>u!!i>,{{<"a<,}!!",'a!!u}a{}<u!>},<'!!>},<!!}i'!><'!,u<!a}!!i!>>}}},{{<ooo!>},<!>!!!!u'au!!u>},{<>,<!!}}!!!!!>!>u{aeai!>,<!!!!!>!!"}!!'e!!'}>}}},{{{<!!u>},{<!>,<}o}>,{<ei!>,<!>},<ue'a}!!!>!!'!>},<!>},<i>}}}}}},{{{{{{<!>},<ie!>ao!!"}o!>},<>}},{}},{<!>,<{!!!!!>}"!>},<}>}},{{<e"o!>,<,!>,<!>!!!>,<<o!>,<!!',e>}}}},{{{{{},{}},{{{<!!!>!!!>,<iuu!o'i{,"o!'ao!!u>}},{{<}<e<!>},<!>,<a<i!>,<!!!>,<a{a>},{<!!!>},<e<o<i!!!>ia>}}},{<e!!}!>,<'<{!>uu!!i<e{i!!!>a>,<!!e'!!!>,<!!!>,<,,<>}},{{{{},{<"!>},<!!!>},<"!>,<!>}ai!!!!a!'u>},{{},{<!!!>o"!>,<!!}o,!ao}!!!>!>!!!!"{!>o}>}}},{<{!!">}},{{{<,oooe>,{<!!!>eo<{!>,<!>},<!u<!!!>'!!>}},{},{}},{{<{,<}!>,<u!<!o>},{{}}}},{{<!>,<!>},<!!!>,<}!!'}<o>},{{<!!!'!!!>!uu''!!{,oeu!>,<!>{>,<,"'>},{{<e!{,o!>},<!!{!oe!>,<!>}"a""!!!>>},{{},{}}}}}},{{<e!"{{!},,o">}}},{{{{},{<'!!"<!!!>,<<!>},<!>,<!!'"">}},<!!!>},<!>}!,,!!ia{>},{{{{{<e>,{<<oe}ii{!e!!!>},<!!<!!!>!!!>>}},{{}}}},<>},{{<e!!!>!!'>,<}!!<<!!!>,<!!!>!!!>,<},i!>!<>},{}},{{<!!!!!>!>,<!!!>'!>,<o!>},<i,!!!>u,}!>}>},{<!>>}}},{{<u,},o!>!!{!>{>},<!o,i!!!e!i!e!>,<aia",!"o'!e>}},{{{<!!!!a,e!>"a!>},<<iu!!!>a!!!>'!>},<}!ii>},{<!,>,{<!>},<!!}<!>!!>}},{{{{<o!>},<>}},<}e!>},<}!!a>}}}}},{{{{{},{<>}},{<>}},{{<!!!!!>"!!!!!>!i!>!>,<ii!>!>,<!>,<}!!!>,<>},<!>},<!!}a!!!>!!i'!>,<>},{{<u!!o!>},<!>},<<,!!!!a!!<i'u"a>,<a<{,!>},<'!u!>,<>},{<<!!a!!!>!>!>},<!>e"i!>>,{<o>}},{<!>,<>}}},{{}},{{<a!!!!o!>!!!u"!>,<i{i}>,{}},{}}}}},{{{{},{<{<!!!>!>,<u',}!>},<>}},{{},<>}},{},{{<i!!!!<oo,o>},{{{{{<!oa}<o>}},{{{{}}}}},{{<e!!!!!!!!!>,<>}}},{{<{ie!!'!o'!!!>,<!>,<,!>!!{!u!>},<>,<!!!>!!!>!>'<!!!>"eu>},{{{<,!eu!io<!!"ui>,{{},<o{eouo<"!>},<!>,<}<!!eu!!,'!!{{{>}}},{{{},{<!>,<!ea<<,!>!>e>}},{<i!!ii{e!>},<!>},<!!ui!>},<'!>,<"e!!!>"">,{<!!!>>}}}},{{}}},{{<>},<<}>}},{{{},<!!!>>},{{},<!!,a}!>}!!oe'>}}}}},{{{{{<!>,<!!!!!>!!">}}},{{{},{}},{{<!'!!a!>},<e!>!}ie!!{!!>}}},{{{<!>o<<!!!!<!!{u!>'!e!<!!!>e!>,<>}},{{{},<!!e<!o!!!>!u!!{u<!><!!>},<e!!i!oi'!>,<a}!>,<u!>,<>},{<aeo,!!,!!!>{!>a<!<{'}{!!!>!>!!,>,{}}}},{{{{{{{<i}!!'!>!>,<<!>,<>},<!>},<uaea!!uo!!!>,<'{!!}!!!>,uo>}},{<!>u!!!!!>>,<e{'i!!<!>},<e!>},<a!!!>!!!>,<!!!!!!!>,<!'ei,>}}},{{{{<!a'!!!!!>,!!!>""<!e<!!!>>}},{<!!!><i,"<!o'<!>,<e!!!!!>o>,{{{{{<'!>},<!!!e!!!>!!<!>o"oe!>},<,e!!'a>,<!>,<,a!!!u}!!!!!!"!!i!!'i!!"u!>},<>},{<<eu>}}},<!!{!>,<!!!>,<i!!!!ee!!!>!!!>>}}}},{{{{{{<}e!>!>o!>!>'i!!"'<,!>},<{e!!i>},<oiu>},{<,!>},<!u<o!>},<}!>,<!!o>,{<<u"u!{,!>!!'!>,<<<>}}},{<e!>!!!!o!!}i,<!!!>!!!!!!!>!}'!>!!}i>,{<<!>,<a'!>},<e!!!>,u!>},<e!!{!>>}},{{<!>!!!>e,!a,a!!!>o}!>},<<!>},<a'!!{!!a>},<!!!!<!!!>{<!>,<u'"!!!>},<<a>}},{{<!>,<{"!aeoa!!>,<!a!>,<>},{<>}}},{{{<!!!>,<!>,<}e"!>,<>},{<{!>,<!>,<e!!e!>,<u",}!!!>>,{}}}},{{<<>}}}}}},{{{{<ie!>,<>},<}u!>},<'ei>},{{},{{<!!i!'o{!{}!>,<!!!>}}!>,<!!!><">,{<oeo<!!!!!>}{{!!!>{>,<">}},{{{<"a!!a!!!}{<>},{{},<}u>}},{{<!>!>'}!!!>!!!>,<!!!><!!{'!!!>!!!>e>},{{{<!>!>,<!>iu>}}}}},{{<a>,{<!!"!>,<"'!>,<!!!>},<u!>},<!!<!'i!"{!!!!!>a!!>}},<<!>,<!>,<}!!!>!>,<!!o!!{!!'i!!!!!>!!!!!>}!u>}},{{<!!'!o{<<!!!>!>,<"!}}}<}ou>},{<!!!!!>},<>}}},{{<,!!o>}}},{{{{<<!!"u!>!!!>},<!>,<!!!>o!>},<}!'!!!>{>},{<>}}}},{{{<!>},<>},{<!!!>},<uae'!!!>}i>}},{{<!u''o<>}},{{{}},{{<!!!!eiua>},<u!!i!>!>},<i!>!!{!!!><'!<!>{!>!!'>}}},{{}}},{{{},{<{,'!>,!!!>{!!!>,<!!!!i{e!!!!!>!!!>'!!!>>,{<<!i!>,<{{!i!!oe'a!>i"!!a<!}}e">}},{<!!!>i!>},<}!!!>,<'!!!>,<a!>,<,!!!!au!>,<o""e>,{<<!>,<!>,<{'>}}},{{<!!ou!!'!>},!!!>ao!!!>>,{<e<<a!{,a{!!!>!>,<!!}!>},<!!!>!>!>!>>}},{<!!{!!!a<!eu!!'!>!!!>!>}>,<,,,{>},{<!!!!!>,!!>}}},{{{{{<euu!!i"{>},{{},<o{!>!}"ea!!!>>}},{{},{<!>,<!!oe}!>},<'{}!!!i!!!>"!>},<!>},<ie!>,<!!>}},{<o>}},{{{<!>},<e!!oe!>,<e!!!>'a!>u!!<,!!!>>},{<!>},<i>}},{{<!>},<}i!!!!!>e!!!!!>o>},{<>}},{}}},{{{{{<!>i}e'!!!>},<o{!"!!!!<!!!<{!o>},<{!>,<<a,{i!,!!!>!>,<>}},{{{<!'!!!>!<{,"o!>,<!!!>!!o>},{<!>i{'!!!!!i'!>,<!!!>}a!!!>},<>}}}}}}}}`
	out := advent9B(input)
	fmt.Printf("Result %d\n", out)
}

func advent1A(test string) (int, error) {
	first := 0
	prev := 0
	total := 0
	for i, rune := range test {
		intRune := int(rune - '0')
		fmt.Printf("%d: %d %d\n", i, intRune, total)
		if i == 0 {
			first = intRune
			prev = intRune
			continue
		}
		if prev == intRune {
			total += intRune
		}
		prev = intRune
	}
	if prev == first {
		total += first
	}

	return total, nil
}

func advent1B(test string) (int, error) {
	fmt.Printf("%s\n", test)
	offset := len(test) / 2
	opposite := 0
	total := 0
	for i, rune := range test[:offset] {
		intRune := int(rune - '0')
		circle := i + offset
		opposite = int(test[circle] - '0')
		if opposite == intRune {
			total += 2 * intRune
		}
		//fmt.Printf("%d: %d %d\n", intRune, opposite, total)
	}

	return total, nil
}

func advent2A(test string) int {
	fmt.Printf("%s\n", test)
	min := 999999999
	max := 0
	sum := 0
	rows := strings.Split(test, ".")
	for _, row := range rows {
		min = 999999999
		max = 0
		numbers := strings.Split(row, " ")
		for _, number := range numbers {
			i, _ := strconv.Atoi(number)
			fmt.Printf("%d\n", i)
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}
		sum += max - min
		//fmt.Printf("%s: %d %d %d\n", row, min, max, sum)
	}

	return sum
}

func advent2B(test string) int {
	fmt.Printf("%s\n", test)
	sum := 0
	rows := strings.Split(test, ".")
	for _, row := range rows {
		numbers := strings.Split(row, " ")
		for _, number1 := range numbers {
			found := false
			for _, number2 := range numbers {
				n1, _ := strconv.Atoi(number1)
				n2, _ := strconv.Atoi(number2)
				if n1 == n2 {
					continue
				}
				if n1 > n2 {
					if n1%n2 == 0 {
						sum += n1 / n2
						found = true
						break
					}
				}
				if n2%n1 == 0 {
					sum += n2 / n1
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		//fmt.Printf("%s: %d %d %d\n", row, min, max, sum)
	}

	return sum
}

func advent3A(test int) int {
	fmt.Printf("%d\n", test)
	half := 572 / 2
	var twoD [572][572]int
	x := half
	y := half
	number := 1
	twoD[x][y] = number
	done := false
	ru := true
	for i := 1; i <= test; i++ {
		if ru {
			// R
			for j := 0; j < i; j++ {
				if number >= test {
					done = true
					break
				}
				x++
				twoD[x][y] = number
				number++
				fmt.Printf("R | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			// U
			for j := 0; j < i; j++ {
				if number >= test {
					done = true
					break
				}
				y--
				twoD[x][y] = number
				number++
				fmt.Printf("U | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			ru = false
		} else {
			// L
			for j := 0; j < i; j++ {
				if number >= test {
					done = true
					break
				}
				x--
				twoD[x][y] = number
				number++
				fmt.Printf("L | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			// D
			for j := 0; j < i; j++ {
				if number >= test {
					done = true
					break
				}
				y++
				twoD[x][y] = number
				number++
				fmt.Printf("D | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			ru = true
		}

	}
	// R1 U1 L2 D2 R3 U3 L4 D4
	//fmt.Printf("%s: %d %d %d\n", row, min, max, sum)

	x = x - half
	if x < 0 {
		x *= -1
	}
	y = y - half
	if y < 0 {
		y *= -1
	}

	return x + y
}

func advent3B(test int) int {
	fmt.Printf("%d\n", test)
	half := 574 / 2
	var twoD [574][574]int
	x := half
	y := half
	number := 1
	twoD[x][y] = number
	done := false
	ru := true
	for i := 1; i <= test; i++ {
		if ru {
			// R
			for j := 0; j < i; j++ {
				x++
				number = twoD[x+1][y] + twoD[x+1][y+1] + twoD[x][y+1] + twoD[x-1][y+1] + twoD[x-1][y] + twoD[x-1][y-1] + twoD[x][y-1] + twoD[x+1][y-1]
				if number > test {
					done = true
					break
				}
				twoD[x][y] = number
				fmt.Printf("R | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			// U
			for j := 0; j < i; j++ {
				y--
				number = twoD[x+1][y] + twoD[x+1][y+1] + twoD[x][y+1] + twoD[x-1][y+1] + twoD[x-1][y] + twoD[x-1][y-1] + twoD[x][y-1] + twoD[x+1][y-1]
				if number > test {
					done = true
					break
				}
				twoD[x][y] = number
				fmt.Printf("U | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			ru = false
		} else {
			// L
			for j := 0; j < i; j++ {
				x--
				number = twoD[x+1][y] + twoD[x+1][y+1] + twoD[x][y+1] + twoD[x-1][y+1] + twoD[x-1][y] + twoD[x-1][y-1] + twoD[x][y-1] + twoD[x+1][y-1]
				if number > test {
					done = true
					break
				}
				twoD[x][y] = number
				fmt.Printf("L | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			// D
			for j := 0; j < i; j++ {
				y++
				number = twoD[x+1][y] + twoD[x+1][y+1] + twoD[x][y+1] + twoD[x-1][y+1] + twoD[x-1][y] + twoD[x-1][y-1] + twoD[x][y-1] + twoD[x+1][y-1]
				if number > test {
					done = true
					break
				}
				twoD[x][y] = number
				fmt.Printf("D | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			ru = true
		}

	}
	// R1 U1 L2 D2 R3 U3 L4 D4
	//fmt.Printf("%s: %d %d %d\n", row, min, max, sum)

	return number
}

func advent4A(test string) int {
	count := 0
	rows := strings.Split(test, ".")
	for _, row := range rows {
		valid := true
		words := strings.Split(row, " ")
		for i, word1 := range words {
			for _, word2 := range words[i+1:] {
				if word1 == word2 {
					fmt.Printf("Duplicate %s\n", word1)
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			fmt.Printf("Valid\n")
			count++
		}
	}
	return count
}

func advent4B(test string) int {
	count := 0
	rows := strings.Split(test, ".")
	for _, row := range rows {
		valid := true
		words := strings.Split(row, " ")
		for i, word1 := range words {
			for _, word2 := range words[i+1:] {
				if len(word1) == len(word2) {
					word1 = SortString(word1)
					word2 = SortString(word2)
					if word1 == word2 {
						fmt.Printf("Duplicate %s %s\n", word1, word2)
						valid = false
						break
					}
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			fmt.Printf("Valid\n")
			count++
		}
	}
	return count
}

func advent5A(test string) int {
	steps := 0
	index := 0
	instructionStrings := strings.Split(test, " ")
	var instructions []int
	for _, s := range instructionStrings {
		d, _ := strconv.Atoi(s)
		instructions = append(instructions, d)
	}

	for {
		if index > len(instructions)-1 || index < 0 {
			break
		}
		instruction := instructions[index]
		//fmt.Printf("Current instruction %d index %d\n", instruction, index)
		instructions[index]++
		index += instruction
		steps++
	}
	return steps
}

func advent5B(test string) int {
	steps := 0
	index := 0
	instructionStrings := strings.Split(test, " ")
	var instructions []int
	for _, s := range instructionStrings {
		d, _ := strconv.Atoi(s)
		instructions = append(instructions, d)
	}

	for {
		if index > len(instructions)-1 || index < 0 {
			break
		}
		instruction := instructions[index]
		//fmt.Printf("Current instruction %d index %d\n", instruction, index)
		if instruction >= 3 {
			instructions[index]--
		} else {
			instructions[index]++
		}
		index += instruction
		steps++
	}
	return steps
}

func advent6A(test string) int {
	steps := 0
	var combinations [][]byte
	bankStrings := strings.Split(test, " ")
	var banks []int
	for _, s := range bankStrings {
		d, _ := strconv.Atoi(s)
		banks = append(banks, d)
	}

	//fmt.Printf("%v\n", banks)
	for {
		h := sha1.New()
		bankStrings = []string{}
		for _, d := range banks {
			s := strconv.Itoa(d)
			bankStrings = append(bankStrings, s)
		}
		h.Write([]byte(strings.Join(bankStrings, " ")))
		hash := h.Sum(nil)
		//fmt.Printf("%x\n", hash)
		for _, combination := range combinations {
			hashString := string(hash)
			combinationString := string(combination)
			if hashString == combinationString {
				return steps
			}
		}
		combinations = append(combinations, hash)

		maxIndex := 0
		for index, bank := range banks[1:] {
			if bank > banks[maxIndex] {
				maxIndex = index + 1
				//fmt.Printf("%d > %d %d\n", bank, banks[maxIndex], index+1)
			}
		}
		//fmt.Printf("maxIndex: %d\n", maxIndex)
		blocks := banks[maxIndex]
		banks[maxIndex] = 0
		currentIndex := maxIndex + 1
		for i := 0; i < blocks; i++ {
			if currentIndex >= len(banks) {
				currentIndex = 0
			}
			banks[currentIndex]++
			currentIndex++
		}
		//fmt.Printf("%v\n", banks)
		steps++
	}
}

func advent6B(test string) int {
	steps := 0
	var combinations [][]byte
	bankStrings := strings.Split(test, " ")
	var banks []int
	for _, s := range bankStrings {
		d, _ := strconv.Atoi(s)
		banks = append(banks, d)
	}

	//fmt.Printf("%v\n", banks)
	for {
		h := sha1.New()
		bankStrings = []string{}
		for _, d := range banks {
			s := strconv.Itoa(d)
			bankStrings = append(bankStrings, s)
		}
		h.Write([]byte(strings.Join(bankStrings, " ")))
		hash := h.Sum(nil)
		//fmt.Printf("%x\n", hash)
		for i, combination := range combinations {
			//fmt.Printf("%d\n", len(combinations)-i)
			hashString := string(hash)
			combinationString := string(combination)
			if hashString == combinationString {
				return len(combinations) - i
			}
		}
		combinations = append(combinations, hash)

		maxIndex := 0
		for index, bank := range banks[1:] {
			if bank > banks[maxIndex] {
				maxIndex = index + 1
				//fmt.Printf("%d > %d %d\n", bank, banks[maxIndex], index+1)
			}
		}
		//fmt.Printf("maxIndex: %d\n", maxIndex)
		blocks := banks[maxIndex]
		banks[maxIndex] = 0
		currentIndex := maxIndex + 1
		for i := 0; i < blocks; i++ {
			if currentIndex >= len(banks) {
				currentIndex = 0
			}
			banks[currentIndex]++
			currentIndex++
		}
		//fmt.Printf("%v\n", banks)
		steps++
	}
}

//Node bi-directional tree node
type Node struct {
	parent   *Node
	name     string
	weight   int
	children []Node
}

func advent7A(test string) string {
	var nodes []Node
	scanner := bufio.NewScanner(strings.NewReader(test))
	var temp Node
	for scanner.Scan() {
		s := scanner.Text()
		nameAndWeight := s
		if strings.Contains(s, "->") {
			stringParts := strings.Split(s, " -> ")
			nameAndWeight = stringParts[0]
			childrenNames := strings.Split(stringParts[1], ", ")
			for _, childName := range childrenNames {
				var childNode Node
				childNode.name = childName
				temp.children = append(temp.children, childNode)
			}
		}
		nameParts := strings.Split(nameAndWeight, " (")
		weight, _ := strconv.Atoi(strings.TrimRight(nameParts[1], ")"))
		temp.name = nameParts[0]
		temp.weight = weight
		nodes = append(nodes, temp)
		//fmt.Printf("%v\n", temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	root := findRoot(nodes)
	//fmt.Printf("%v\n", root)

	return root.name
}

func advent7B(test string) int {
	var nodes []Node
	nodeMap := make(map[string]Node)
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		var temp Node
		s := scanner.Text()
		nameAndWeight := s
		if strings.Contains(s, "->") {
			stringParts := strings.Split(s, " -> ")
			nameAndWeight = stringParts[0]
			childrenNames := strings.Split(stringParts[1], ", ")
			for _, childName := range childrenNames {
				var childNode Node
				childNode.name = childName
				temp.children = append(temp.children, childNode)
			}
		}
		nameParts := strings.Split(nameAndWeight, " (")
		weight, _ := strconv.Atoi(strings.TrimRight(nameParts[1], ")"))
		temp.name = nameParts[0]
		temp.weight = weight
		nodes = append(nodes, temp)
		nodeMap[temp.name] = temp
		//fmt.Printf("%v\n", temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	root := findRoot(nodes)
	buildTree(&root, nodeMap)
	//fmt.Printf("%v\n", nodeMap)
	result := updateWeights(&root, 0)

	return result
}

func updateWeights(root *Node, result int) int {
	if result != 0 {
		return result
	}
	if len(root.children) == 0 {
		return root.weight
	}
	var weights []int
	sum := 0
	for _, child := range root.children {
		weight := updateWeights(&child, 0)
		weights = append(weights, weight)
		sum += weight
	}
	for i, weight := range weights {
		count := 0
		for _, weightSearch := range weights {
			if weight == weightSearch {
				count++
			}
		}
		if count == 1 {
			offset := 0
			if i == 0 {
				offset = weight - weights[i+1]
			} else {
				offset = weight - weights[i-1]
			}
			//fmt.Printf("\ninvalid weight! : %d %d %d %d %d\n", root.children[i].weight, offset, weight, root.weight, sum)
			panic(root.children[i].weight - offset)
		}
	}
	root.weight += sum
	sum = root.weight
	//fmt.Printf("\nweights : %v\n", weights)
	return sum
}

func buildTree(root *Node, nodeMap map[string]Node) {
	deroot := *root
	newRoot := nodeMap[deroot.name]
	root = &newRoot
	//fmt.Printf("\nroot : %v\n", root)
	if len(root.children) == 0 {
		//fmt.Printf("leaf : %v\n", root)
	} else {
		for i := range root.children {
			//fmt.Printf("child : %v\n", root.children[i])
			root.children[i] = nodeMap[root.children[i].name]
			root.children[i].parent = root
			buildTree(&root.children[i], nodeMap)
		}
	}
}

func findRoot(nodes []Node) Node {
	for _, node := range nodes {
		isParent := true
		for _, parent := range nodes {
			for _, child := range parent.children {
				if child.name == node.name {
					isParent = false
				}
			}
		}
		if isParent {
			return node
		}
	}
	return nodes[0]
}
