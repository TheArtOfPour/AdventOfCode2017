package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDay5(t *testing.T) {
	input := "0 3 0 1 -3"
	out := advent5A(input)
	expected := 5
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "0 3 0 1 -3"
	out = advent5B(input)
	expected = 10
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	input = "2 2 -1 1 -1 1 1 -5 -5 -1 0 -8 -2 -11 -4 -5 -10 -4 -9 -9 1 1 -11 -8 -19 -14 -6 -2 -1 -11 -23 -8 -7 -9 -26 -1 -8 -11 -34 0 -22 -17 -41 -12 -43 -33 -15 0 2 -41 -41 -26 -48 -52 -47 -30 -38 -20 -4 -21 -17 -19 -55 -32 -12 -55 1 -34 -8 -15 -59 -56 -16 -23 -43 -5 -41 -56 -32 -67 -14 0 -28 -32 -7 -54 -19 -9 -24 -63 -2 -60 -5 -78 -11 -84 -50 -36 -72 -14 -30 -4 -62 -6 -1 -69 -17 -33 -32 -45 -71 -87 -71 -60 -19 -80 -11 -106 -45 -27 -23 -51 -77 -67 -103 -17 -98 -109 -91 -125 -68 -39 -34 -96 -49 -64 -38 -105 -31 -100 -89 -108 -69 -36 -94 -38 -124 -123 -79 -92 -42 -14 -87 -68 -17 -36 -21 -54 -98 -79 -142 -25 -60 -112 -99 -64 -15 -78 -37 -64 -15 -129 -32 -102 -74 -112 1 -146 -151 -147 -153 -4 -181 -22 -176 -4 -57 -151 -86 -121 -38 -137 -160 -156 -72 -73 -149 -64 -182 -117 -146 -180 -195 -27 -194 -191 -108 -153 -40 -149 -100 -120 -207 -83 -94 -73 -200 -95 -155 -94 -76 -9 -149 -70 -125 -49 -146 -223 -68 -139 -26 -132 -142 -165 -2 -45 -154 -129 -130 -185 -60 -34 -173 -91 -37 -40 -153 -189 -236 -95 -128 -46 -14 -53 -245 -67 -9 -208 -244 -198 -74 -62 -104 -51 -251 -48 -50 -115 -76 -79 -32 -82 -65 -185 -124 -32 -189 -124 -174 1 -273 -223 -275 -238 -200 -184 -229 -195 -152 -63 -150 -73 -44 -54 -187 -49 -250 -192 -290 -282 -266 -214 -117 -199 -83 -104 -251 -176 -262 -296 -39 -259 -87 -132 -166 -67 -194 1 -294 -8 -3 -264 -217 -228 -233 -241 -294 -210 -72 -307 -259 -33 -101 -103 -235 -100 -110 -253 -292 -134 -269 -52 -265 -15 -29 -272 -126 -210 -151 -308 -40 -40 -112 -268 -185 -346 -237 -287 -34 -302 -41 -25 -191 -29 -170 -95 -315 -278 -160 -220 -99 -126 -224 -33 -350 -76 -138 -340 -284 -268 -128 -238 -197 -93 -110 -120 -190 -140 -64 -217 -296 -103 -363 -199 -254 -233 -190 -282 -136 -174 -309 -61 -206 -18 -105 -111 -163 -287 -188 -145 -294 -251 -398 -265 -273 -50 -250 -376 -5 -357 -6 -8 -198 -20 -82 -158 -122 -196 -97 -183 -48 -428 -36 -88 -424 -35 -380 -109 -209 -323 -394 -102 -276 -153 -229 -320 -391 -7 -328 -127 -430 -102 -372 -447 -222 -401 -184 -183 -49 -239 -413 -101 -187 -289 -12 -418 -248 -279 -318 -134 -443 -272 -456 -143 -3 -209 -276 -414 -189 -302 -238 -241 -106 -332 -375 -400 -476 -9 -95 -412 -52 -127 -442 -278 -25 -446 -411 -39 -55 -80 -234 -361 -223 -384 -283 -47 -164 -18 -38 -87 -393 -93 -380 -493 -73 -150 -241 -378 -211 -516 -349 -520 -38 -397 -406 -16 -461 -276 -448 -316 -376 -156 -369 -216 -431 -309 -400 -135 -523 -40 -508 -87 -25 -151 -355 -141 -3 -495 -153 -438 -343 -161 -66 -455 -70 -248 -278 -548 -300 -337 -290 -551 -200 -68 -540 -476 -395 -245 -318 -424 -112 -556 -541 -94 -148 -542 -100 -120 -199 -569 -471 -298 -16 -453 -469 -50 -500 -84 -435 -579 -287 -522 -77 -83 -347 -437 -171 -231 -139 -350 -357 -221 -214 -224 -148 -125 -385 -255 -38 -320 -254 -517 -532 -80 -286 -58 -97 -390 -309 -548 -319 -323 -238 -297 -12 -312 -517 -434 -466 -103 -621 -448 -503 -72 -601 -287 -61 -577 -87 -143 -33 -482 -275 -529 -340 -279 -130 -512 -63 -109 -528 -22 -549 -317 -375 -377 -385 -23 -191 -138 -509 -40 -565 -559 -14 -547 -28 -159 -153 -585 -508 -582 -431 -580 -637 -561 -513 -243 -420 -298 -485 -132 -613 -157 -521 -596 -61 -420 -498 -577 -563 -354 -662 -264 -273 -111 -597 -466 -389 -345 -306 -102 -57 -596 -1 -45 -12 -619 -47 -43 0 -323 -9 -319 -529 -402 -238 -191 -487 -315 -65 -386 -110 -605 -363 -461 -6 -95 -95 2 -596 -454 -618 -83 -481 -283 -386 -247 -417 -707 -564 -603 -17 -712 -140 -336 -567 -443 -36 -476 -251 -735 -589 -198 -197 -476 -49 -736 -422 -383 -569 -732 -1 -104 -261 -352 -453 -273 -344 -66 -307 -698 -158 -238 -280 -207 -624 -491 -765 -506 -146 -616 -711 -650 -655 -393 -19 -315 -311 -572 -675 -533 -156 -373 -744 -142 -582 -491 -796 -777 -125 -483 -426 -510 -560 -700 -778 -407 -440 -409 -238 -738 -477 -147 -152 -317 -110 -323 -788 -601 -202 -517 -487 -726 -300 -1 -554 -448 -15 -191 -531 -568 -466 -527 -132 -254 -290 -8 -400 -655 -788 -376 -249 -662 -315 -378 -41 -793 -163 -29 -327 -839 -133 -124 -129 -673 -32 -605 -393 -664 -374 -135 -366 -717 -93 -601 -763 -788 -494 -802 -282 -443 -491 -461 -197 -83 -96 -162 -97 -161 -232 -144 -472 -118 -429 -387 -724 -789 -636 -298 -484 -720 -526 -382 -102 -449 -846 -525 -547 -696 -524 -272 -843 -286 -247 -838 -447 -489 -797 -483 -386 -775 -340 -772 -158 -293 -256 -432 -812 -273 -93 -487 -264 -594 -330 -712 -798 -131 -591 -539 -677 -455 -470 -108 -573 -57 -845 -383 -273 -890 -747 -913 -648 -625 -650 -544 -137 -490 -434 -734 -182 -355 -859 -835 -141 -536 -874 -102 -940 -359 -83 -800 -894 -712 -470 -687 -578 -435 -935 -400 -780 -814 -458 -892 -481 -371 -761 -348 -388 -891 -764 -297 -536 -695 -314 -336 -978 -379 -462 -597 -533 -561 -9 -474 -292 -560 -420 -828 -721 -769 -874 -157 -495 -771 -899 -571 -98 -282 -233 -203 -982 -416 -142 -993 -540 -979 -851 -506 -238 -292 -184 -695 -195 -632 -575 -962 -76 -546 -705 -13 -271 -222 -124 -380 2 -1003 -251 -525 -228 -644 -159 -624 -477 -912 -712 -343 -263 -88 -745 -85 -374 -675 -804 -610 -854 -511 -612 -964 -731 -358 -495 -946 -466 -364 -1053 -57 -101 -829 -155 -600"
	out = advent5A(input)
	fmt.Printf("Result 5A %d\n", out)
	out = advent5B(input)
	fmt.Printf("Result 5B %d\n", out)
}

func TestDay6(t *testing.T) {
	input := "0 2 7 0"
	out := advent6A(input)
	expected := 5
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	input = "0 2 7 0"
	out = advent6B(input)
	expected = 4
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	input = "4 1 15 12 0 9 9 5 5 8 7 3 14 5 12 3"
	out = advent6A(input)
	fmt.Printf("Result 6A %d\n", out)
	out = advent6B(input)
	fmt.Printf("Result 6B %d\n", out)
}

func TestDay7(t *testing.T) {
	input := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`
	out := advent7A(input)
	expected := "tknk"
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	// intOut := advent7B(input)
	// intExp := 60
	// if !cmp.Equal(intOut, intExp) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(intOut, intExp))
	// }

	// 	input = `keztg (7)
	// uwbtawx (9)
	// mgyhaax (46)
	// fuvokrr (14) -> pnjbsm, glrua
	// cymmj (257) -> phyzvno, pmfprs, ozgprze, bgjngh
	// goilxo (80)
	// cumfrfc (102) -> yjivxcf, swqkqgz
	// yquljjj (20)
	// ehywag (18)
	// mmtyhkd (21)
	// paglk (98)
	// wtqfs (82)
	// oaynkf (8)
	// cupbfut (78)
	// vpcruoy (70)
	// wmdbo (50)
	// tmbtipi (48)
	// lkopm (9)
	// gluzk (18)
	// prvrg (76)
	// lkdkyk (30) -> oldwss, nadxwf
	// iqsztjd (181) -> hovelvz, pndcqot, naglm, oxxlsk
	// nxdkpuh (217) -> yhcsc, ydmeqtl
	// nxlhjq (306) -> hcwjxe, zixbap
	// vtkgj (89)
	// rzrzage (73)
	// ftegwk (284)
	// lircjh (23)
	// zosskdz (232) -> isrch, bwzvefg, dxodoee
	// dphcbfr (67)
	// pnmvk (180) -> wrabgy, vlfpuo
	// owmjbhg (120) -> szfxhin, czzpk, zwrfiyf
	// oonqts (26)
	// zjaqq (129) -> hopjmyt, cdwkezv
	// hxeoxk (33)
	// csaqixs (1237) -> alzipi, lhxycw, tkeuvp
	// avenz (7)
	// nnhknbl (55) -> owzwbpn, iaonkp
	// bxifcld (86)
	// neyeo (165) -> gxxzwq, fxwez
	// qnjpz (71)
	// qhxmh (61)
	// jmhfgr (139) -> ucuqxgm, hovhxsp
	// tyuhzom (80)
	// pqtboz (207) -> ayvns, codwosk
	// dqyjg (65)
	// nujppls (24)
	// mxbixyi (60)
	// xkzgz (85)
	// oxklzu (2285) -> ehwlw, fptoo, sgobq, eduwet, pqmpnzo
	// fuleuxt (6) -> ljzuyyk, pxydes
	// zktmxll (451) -> txsrez, ewjrko, drtrgwp, kiggy
	// qpxbow (40)
	// rshpnha (36)
	// pqmpnzo (1374) -> hpltoci, oxvwr, vrxeemw
	// wdazzdu (54)
	// kivcyus (53)
	// cvvncju (10)
	// dtkuik (36)
	// opkvs (64)
	// kwjnfg (28)
	// suoiohi (197) -> gluzk, fdhdpw
	// jkaxk (98)
	// zsoro (12)
	// fqvtm (15)
	// nqktjw (14)
	// cbhkkx (116) -> wrprrev, vyoxx
	// rmqdolg (55)
	// mdkes (95)
	// obxansb (343) -> uzyprc, uxaqq
	// pjitmnv (31)
	// vdkrvi (73)
	// nystxqv (35)
	// odgzsnk (73)
	// hehbbo (83)
	// zrthre (30)
	// zwoot (9)
	// mfawmsq (92)
	// sckaqs (1141) -> qpaei, cbhkkx, qezwkkx
	// vxfci (60) -> tdecuga, wssvxr, pchccgz, ypogtw
	// vauwilt (78)
	// qxkas (24) -> mzgyj, xappjar, cgbgm, muarkn
	// eqibqs (20)
	// wuefg (549) -> pwwyeqx, ylidl, qwbfod, mqztoa
	// jmchsu (77)
	// dinng (30)
	// nlpmbrd (37) -> fnzjtvw, qzjyi
	// yjivxcf (19)
	// yhjopn (34)
	// hqxdyv (17568) -> apktiv, ybekxtf, etoxfc
	// wvivxrz (82)
	// wszqat (85) -> eddmyv, edkwqih, mxbixyi
	// xcldsl (78)
	// rnbzlx (35)
	// hibxz (94) -> zgzsu, vzgsgk
	// rgyco (21)
	// lmvvs (22)
	// nezny (13)
	// tpcvq (251)
	// puxgb (15)
	// merrako (8)
	// nzweur (431) -> rfouw, sukktk, rreqg, fmpcnql
	// pwwyeqx (100) -> fmjhlia, yquljjj
	// lccoo (27) -> bmlid, prvrg
	// dudxud (202)
	// cmnzh (49)
	// nitvw (8)
	// dcakuo (21)
	// nnbty (81)
	// kjlxeat (318) -> nmhww, zacsvwk
	// frinpfj (88) -> lwfoqny, tdgel
	// zgqsgm (38) -> phdcpp, qwcrc
	// bvikij (91) -> obxansb, bizbsjd, usggwvu, zrhny, svngfr
	// dnycw (219) -> sibzrx, hdgvs, wnfqg
	// youpfn (38)
	// zixbap (47)
	// mpwldri (55)
	// rfqenw (80)
	// tjtzx (78)
	// zbfut (55)
	// eruxzoi (63)
	// nandmg (344) -> qpqplm, zrthre
	// bizbsjd (159) -> piecd, ghkdvw, caurb
	// jpkwter (19) -> mcjgfx, ujsbt
	// uwpbnv (83)
	// devljb (45)
	// euzztul (30)
	// frmbrb (1660) -> norkse, iitweo, mebwy, sckaqs, xdrge
	// duccgc (15)
	// bmlid (76)
	// clwwv (238) -> jofvyvx, zgjoaiw
	// llmmm (69) -> wvivxrz, pikvdx
	// rstdh (21)
	// afckjn (51)
	// ojqia (22)
	// qtlzten (155) -> fjgsw, uujpt
	// eqxgwfz (40)
	// ljwsi (20)
	// vuxbzm (48)
	// qzuwt (130) -> qnrjqj, bjdtdn
	// lnnwiq (20)
	// pbxpdo (281) -> vabmsx, kwjnfg
	// bfcdy (31)
	// ykpsfj (28)
	// uegcs (210)
	// qezwkkx (74) -> gquil, stfzaxc
	// jkvduo (44)
	// vtylgti (66) -> twgbxu, uukshmq, rlvggmr, aynpr
	// eyyokyd (28)
	// nrbcaqo (45)
	// caocs (35)
	// cfdpxpm (207) -> jhwmc, nezny
	// qwjmobb (28425) -> ohusizx, gqoxv, xatjlb
	// akniuo (129) -> zpczji, tugrmnp
	// kravhjd (17)
	// wtbwbpz (43) -> kybegv, qxhda
	// hovelvz (31)
	// jhwmc (13)
	// wctze (102) -> qwdhdrk, znooxvq, vhrxl, zfhkfwn
	// nmhww (43)
	// fqufwq (58)
	// zyxjg (81)
	// eltbyz (61)
	// ehtsbv (783) -> upaqlj, cckqr, pgprg, ubksf
	// zbrbb (12)
	// vjgvk (28)
	// mqulwk (15)
	// kywmnbd (404)
	// wcuvk (20)
	// ymfls (75)
	// spxwcuv (173) -> iobcvl, xwfbb, wxpauwt
	// eumzi (24)
	// kqoigs (53) -> krfgye, oxklzu, pinipk, ojatf, memkrd
	// alneqju (77)
	// joczsir (313) -> xwkoc, atkmjxg, gurxxfd, axxkh, jmpknjs
	// uwzvy (35)
	// xqxyx (386) -> avenz, keztg
	// erylwj (804) -> wdsbi, ugrhs, fzmaw
	// corfkob (87)
	// sibzrx (67) -> pqccp, audeogd
	// crrfxfn (38)
	// piecd (72)
	// bxefs (22)
	// lnufi (93)
	// qifuph (44)
	// uqccsbh (26)
	// jaqwzi (79) -> zkgoa, juymjz
	// nivpffu (169) -> tnccv, lfqca, sgfco, nnbbrbf, egsgwch
	// dfrkf (49)
	// cdrhm (56)
	// vaylgz (80)
	// ayvns (23)
	// mdddafe (56)
	// fpldxlq (195) -> gjnnmvb, ljwsi
	// eygaz (427) -> ascyv, kmjfxcf, puxgb
	// ymeelep (92)
	// iuzvl (23)
	// tkjeu (41)
	// xdlyd (75) -> cymmj, pbxpdo, vmjbgo, cwttq
	// jsltf (39)
	// ciojx (146) -> ioobamp, ahrfot
	// eqmeu (211) -> anrjxof, nepxnu, mwpbyo, rbzqabo
	// bogvr (202) -> zghrr, bompiu
	// jefztzv (91)
	// fvikm (80) -> zepvwyv, oonqts
	// zdqcu (194) -> ucbuez, nqktjw
	// pdvolf (75)
	// mkmci (40) -> ggaxx, xvzlrw
	// sqmfis (35)
	// chrqi (74)
	// tvgytpm (49)
	// bjoyw (29)
	// nkfvkp (62)
	// xbtswv (7446) -> iehfo, xcrkb, qksclw
	// qomxhp (721) -> agufw, djgzb, jxbksoh, twnfzz, ucxgom
	// ibmiu (9)
	// atmzoso (6) -> drosj, wcrrrlf
	// fuqvw (56)
	// jfaca (49)
	// yulga (213)
	// dxodoee (7)
	// gethyvd (39)
	// hjxcpi (30)
	// jlcgqt (55)
	// lzouo (144) -> ubjgijo, rnbzlx
	// djyxrkb (78)
	// bscpyic (61) -> nktmu, sqpdsk
	// sxmdnhl (31)
	// qbdafi (25)
	// vwbxg (35)
	// rlvggmr (63)
	// kiggy (27) -> pgsokae, ottiad, eruxzoi, zhttn
	// ocpngbz (73) -> hqzay, ewzryd, ipjbc, xjnhqlg
	// movmxq (216) -> kfxhl, ulpkj
	// mzgyj (85)
	// orwbdwn (52)
	// ixyqeq (6735) -> xgwjcx, nkkgyl, sykwd
	// htsjndf (211)
	// ndhsa (82)
	// jmpknjs (80) -> gdrcfwr, wgivp
	// hieel (65)
	// htdwe (25)
	// qrfuvjh (9)
	// ubjgijo (35)
	// vobnpuq (32)
	// elgyjo (141) -> ineoncq, pfdmmg
	// zsnge (71)
	// zbwxa (28)
	// ogczchc (31)
	// njdpm (53)
	// cpsce (84)
	// ftdylco (19)
	// zrwfi (22)
	// hyfuy (252) -> pmscqw, ecaph
	// nayudfl (320) -> ssvsso, zrwfi
	// kykfb (72) -> euzztul, vxrtejs
	// ggpjxwv (9)
	// aostqf (29)
	// zujltb (13) -> hatvlca, ppmrgga, cjoya, bogvr, gtbpbl, ocwkc, qzuwt
	// ajbtn (18)
	// fzmaw (20) -> qjixqo, fkhxkeg, uqccsbh
	// akmbqb (108) -> kqkzsm, grgsn
	// slrdn (55)
	// nrjwctg (96)
	// norkse (661) -> jmhfgr, anwxvv, ptwhbm, znubct, djrrc, hgmjvpp
	// bwzvefg (7)
	// peuadz (8)
	// kvmqsdx (308) -> zvtoom, twvdhg
	// bvnjiou (32)
	// lnwuqu (159) -> yeqnq, fqvtm
	// pksfx (54)
	// xatua (97)
	// tbrznk (37)
	// ucuqxgm (31)
	// hwjhf (78)
	// pinipk (7229) -> pqewl, zujltb, kfcowx
	// yyhzd (12) -> kvqspmf, dtkuik, wsvir
	// imyvlyt (38)
	// oldwss (96)
	// paegovu (86)
	// knjlz (83)
	// oevbo (23)
	// yeucm (98)
	// usvkq (56)
	// hatvlca (69) -> gpogy, eyrfvtl, subwna
	// qkhtsa (208) -> pyxgmtu, pqgpuv, pudnxf, dilqo, juqbdco, flufot, fzdyvo
	// gzvjxk (1397) -> wctze, zoijv, fuleuxt
	// bkwcwf (326) -> jazkpl, kfcgv, qpjctjw
	// vqghhbs (157) -> lsrvhoi, livmxo
	// fzdyvo (186) -> gwfrqr, tcgffi
	// rqkfkxw (47)
	// bntdh (76)
	// rfmiqz (158) -> omlwg, uhwvnbg
	// zorvsm (50)
	// edzgraw (83)
	// iwsknb (345) -> qhqnfsp, vrwkr
	// vyoxx (63)
	// livmxo (27)
	// foabep (92)
	// dbuccip (28)
	// oojme (73) -> ugkxkqe, keucu, zjqeu
	// cmmqbz (29868) -> frmbrb, ixyqeq, njatvu
	// syjvwzt (6039) -> bigkiu, wjipa, pmbnia
	// uxaqq (16)
	// vboha (79)
	// irsfgz (94)
	// lsrvhoi (27)
	// zdtvktq (99) -> ezrix, lyhfj
	// hgbkwjv (32)
	// inxivh (261) -> mzzxjcu, rkaxx
	// nfrtom (44)
	// xhzylb (97)
	// nuxqskl (39)
	// qhfxqrr (65)
	// foaayon (78)
	// rhtdtxv (234) -> rplxsw, fjqomax
	// kppxrk (73)
	// qjixqo (26)
	// sgngx (75)
	// ycsbsyn (87) -> gmeueu, wdhmsi, zrqqtx
	// kqkzsm (61)
	// irjvpam (39) -> ajozeez, xxlbk, nfwlplx
	// ldfofo (84)
	// mygcpku (84)
	// caurb (72)
	// gpysit (22)
	// gxfij (171) -> sxnsqj, ksxixz
	// mxbyg (39)
	// jlshk (29)
	// havdbe (132) -> dqyjg, wvdapsm
	// ocwkc (86) -> dtzws, fleszz
	// rqrtz (83)
	// eddmyv (60)
	// ecaph (56)
	// ebgsk (60) -> rshpnha, puexzf
	// lymsa (44) -> qcyypa, vbdcxx
	// mozvs (54) -> dfrkf, cmnzh
	// kpghxz (17)
	// wrpqf (83)
	// dzjsx (57)
	// vrzbmt (14)
	// dryngd (29) -> shawokt, elkflt, tjtzx, auiiuv
	// ittmm (28)
	// zqurr (284) -> fqufwq, htsuyvw
	// ktejrze (69)
	// pudnxf (192) -> aostqf, szwpt
	// tqjlm (74)
	// codwosk (23)
	// pmscqw (56)
	// cmcto (441)
	// edoycls (93) -> vvnzr, imnvt, vuxbzm, rhzco
	// zrhny (177) -> dcwfs, dppwsec
	// vffew (46)
	// hcwjxe (47)
	// dinlw (342) -> zzjuf, lnwuqu, gxfij, kmqurp
	// lgtpaqk (75)
	// fluwt (65)
	// syuoewb (288) -> uwbtawx, avjkgl, nkycb
	// anrjxof (6)
	// vrwkr (5)
	// pngubp (76)
	// vvnhe (89)
	// rnlaw (45) -> wszqat, ptizsk, mofyda, ttolm, velktz, nzsnkla, hhdzz
	// nadxwf (96)
	// sxpan (31)
	// gdrcfwr (71)
	// vabmsx (28)
	// tsiyp (52)
	// fptoo (1398) -> gintpbf, cklkizf, kjgtfqc
	// obxrn (756) -> yyhzd, afatio, uqjge
	// znkchkk (50)
	// hvapnf (121) -> siruccf, tgpxvyr
	// htsuyvw (58)
	// yifny (122) -> dcakuo, xhtzti
	// uvvqcxz (79)
	// suiyl (773) -> hjfwn, thknml, gkijw
	// shawokt (78)
	// ezdiq (37)
	// khrqmbo (9)
	// gkjbikb (79) -> nrjwctg, iamrpx
	// ysskib (76)
	// zufoz (93) -> smrvw, kkkjsil
	// uusiaqf (84)
	// cjxyt (69) -> deczzuu, ymeelep
	// gcefq (33)
	// tgebda (247) -> ngthpc, bnvlsm, afckjn
	// ljmve (55)
	// rqdcuf (144) -> xzxvwzf, bjoyw
	// sjxaxv (76)
	// jiybx (88)
	// wxpauwt (20)
	// clsbdm (10)
	// fmpcnql (56) -> rjvfwcu, xjjdapk, nwtlu
	// hchmn (258) -> bremy, vpcruoy
	// labnsw (157) -> wnwsbdq, jqjkgv
	// hvpcvdf (95)
	// xatjlb (40) -> aiunbee, keoaqjb, jxkofob, plbtdq
	// dflvw (86)
	// vmcgj (233)
	// jxwyzy (13)
	// mebwy (1107) -> uegez, dpiyhv, vdzuw, icpwoof
	// btqebbp (156)
	// mmvszxj (71)
	// gpysm (85)
	// dppwsec (99)
	// iodveqh (89)
	// xcycu (116) -> dnkiyf, njdpm
	// ijictm (38)
	// ineay (19)
	// hvdpzbv (73)
	// cckqr (148) -> rqhfhc, jcyciyq, jbqvpsb, uapwikr
	// opowpvm (265) -> aewie, phrume
	// ptwhbm (69) -> omzlzx, suxpc
	// bpdixmv (83)
	// mzlmr (43)
	// ezrix (79)
	// jovly (61)
	// audeogd (99)
	// hjfwn (168) -> wtqfs, uamje
	// wdhbym (184) -> pksfx, bgahvfu, krlep, pkmmfc
	// nmmtdv (83)
	// uftwml (50)
	// ycyyvdy (17)
	// dqdpop (53) -> ggllv, zatwq
	// egqmgr (37)
	// txsrez (99) -> iiznokx, dobmve
	// lqmdutk (25)
	// mfecx (33)
	// kldww (23)
	// dxbxha (135) -> uvswv, ydmrd
	// xmvdh (41)
	// ctinwus (96)
	// arkedwb (237) -> qsvckew, merrako
	// iaonkp (89)
	// uthjdye (124) -> awoedy, wuisqnk, qsfpaj
	// yzbccbx (36) -> fxemb, osryil, nkfvkp
	// xvlymgg (33)
	// wbbaxr (10)
	// wcrrrlf (98)
	// bclse (38) -> slgjon, tzltv
	// krzjli (33)
	// dnkiyf (53)
	// cklkizf (219)
	// pqccp (99)
	// bgahvfu (54)
	// gotqku (33)
	// vregap (88)
	// qjcqm (138) -> vdkrvi, qoijc
	// phocd (23)
	// lerycoe (84)
	// aidbql (35)
	// qwvmczr (187) -> youpfn, kkukqoj
	// achhc (30)
	// qxnkp (68)
	// znkzp (94) -> otpnx, drtxytc, ntonira, zyxjg
	// ewzryd (92)
	// xxssj (168) -> nhnrpz, rzwfp
	// lnaoiv (70) -> awqat, dnaoe, qombesj, eygaz
	// qzjyi (88)
	// aevhbim (49)
	// jeuzbj (44)
	// msskvkg (60)
	// douvhy (133) -> mqulwk, nnkmf
	// qrfgzm (552) -> lrtds, yrpat, tvjwhhy, atmzoso, dudxud, rqdcuf, xxssj
	// alankuj (75)
	// juymjz (77)
	// gcqaj (69)
	// ozwpmzc (43) -> bknogxz, ldfofo
	// gintpbf (59) -> byckty, tagyci
	// yrhid (222) -> phocd, wtzkvm, kldww, xugyewm
	// keucu (15) -> kwozg, xfmxo, rwepl
	// zicelk (67)
	// xwwxswj (40) -> frnhsjr, mmvszxj, qsmvtif
	// lcfyznt (151) -> klmvmt, bgmmb
	// ehwlw (55) -> tgebda, ehypcwy, xqxyx, vxfci, nxlhjq
	// uegez (150) -> rnjfcg, tusgzei
	// pjaxkr (92)
	// mdprv (179) -> paanydj, vtaejs
	// avjkgl (9)
	// sshfxeu (11) -> goilxo, ltichp
	// alzipi (53) -> cpsce, ezspcab
	// wxoqoa (100) -> mfvul, pwatre
	// trmwq (10)
	// ltmwbs (33)
	// kmdzlmk (235) -> mzauq, ycsbsyn, srqclhj, shhfy, xcycu, zsnatp
	// dfjvh (61)
	// zknesz (39)
	// ezglz (289) -> kyjgf, oaynkf
	// pqixzdw (128) -> nvvca, cwxgqrk, duccgc, sbnsoxi
	// zpczji (60)
	// rathd (135) -> rsbune, wmdbo, uftwml
	// pchccgz (85)
	// rzwfp (17)
	// umgzhmp (150) -> eqibqs, rwgpi, wcuvk
	// ouqtjz (127) -> ejkei, sdgdv
	// sstvew (52)
	// qxhda (96)
	// cvstod (96)
	// rbyiay (94)
	// vwzmkq (63) -> vhjcue, nnbty
	// dtzws (83)
	// nyepy (363) -> jsltf, nuxqskl
	// iiwkm (84) -> fluwt, clcixal
	// ucezr (10)
	// zvtoom (55)
	// fleszz (83)
	// hdtcizc (9)
	// wzatr (57)
	// teqswj (40) -> advhon, qhfxqrr, hieel
	// usuaiv (45)
	// eygdy (86) -> pqtboz, kmodn, dqklqo, tukyh
	// janrdqf (81)
	// jjukf (723) -> vqghhbs, ozwpmzc, htsjndf, wsupp
	// udmez (91)
	// lsvqox (1205) -> tpcvq, fvqvgw, jusfet
	// zjymq (33)
	// wuisqnk (34)
	// pxydes (62)
	// qegmu (84) -> hzcanxq, xfbosce, glyxk, fvlmvtk, akmbqb, lymsa, fbrwdf
	// mckrfxj (38)
	// keylghg (60)
	// dqwdx (215) -> hseaxj, ubxke
	// fzjgh (179) -> bafdzbu, uanvh
	// qsvckew (8)
	// pvhrim (52)
	// mzauq (98) -> sxpan, pjitmnv, ogczchc, aruczxj
	// nxwjvx (53)
	// agufw (229) -> wbbaxr, knrgg
	// urpnw (47) -> khtjh, dflvw, otxphme
	// klmvmt (17)
	// skgwztg (233)
	// ghaxvq (9)
	// sgobq (60) -> lsspa, srgmt, wofung, rathd, joaed, fzjgh, pbkfd
	// qaveutv (44)
	// dnaoe (80) -> vohepl, mymke, paglk, pahwxj
	// edgdvfa (23)
	// mrfmn (94)
	// hovhxsp (31)
	// hqlyg (47) -> aidbql, tfhij, krvyy
	// zuybvj (32)
	// kquxfy (1294) -> srcyajr, neako, vtuqq
	// htsuvhg (253)
	// tagyci (80)
	// ldnnoag (35)
	// qkougo (76)
	// fvmhrf (38)
	// spvcd (84)
	// tcbpqu (70)
	// yzxqp (76)
	// ihmqs (66) -> oydsj, qheany
	// neako (48) -> panao, fcufg
	// dkgmsse (18)
	// zupym (33)
	// pwumvgy (80)
	// ognax (83)
	// oxvwr (203) -> wixlvcp, wqgcqb
	// ysltqk (65) -> mrfmn, hrjgbc
	// dtqzu (61) -> tgnqn, corfkob
	// vmhwy (177) -> yaudo, eqxgwfz
	// pwatre (16)
	// uijtrw (247) -> cbdczg, idjhk
	// stfzaxc (84)
	// refya (15)
	// xfbosce (188) -> rstdh, nqzwt
	// kbrxrks (56)
	// fplihc (88)
	// fkoqh (251) -> ohvifiy, npckah
	// lwfoqny (61)
	// paanydj (39)
	// xzxvwzf (29)
	// bpebim (180) -> vvnohc, jlyaty, kixaco
	// aiunbee (1214) -> fpldxlq, eqmeu, teqswj, ngtkzm, yndyrey, wtbwbpz, eoqtf
	// wrdgs (220)
	// oxxlsk (31)
	// atlrd (269) -> dgszhd, dkgmsse
	// svtwviu (69) -> iiwkm, ieuwo, carbhvi, lzouo, cklpcr, fkkzg, ouubjrl
	// mfvul (16)
	// ndpwgdy (36)
	// rhzco (48)
	// nzetqt (8)
	// xvzlrw (87)
	// utsob (14)
	// rsbune (50)
	// tcgffi (32)
	// cwttq (9) -> mughfl, ponlfyf, tuyyte, ndhsa
	// wejvpzr (29)
	// ozgprze (20)
	// eemnlc (83)
	// omlwg (28)
	// hpmuqvl (37) -> gawck, yeucm
	// ohusizx (9151) -> fsokbvd, oojme, bonjgrt
	// mqllnxu (206)
	// cnqrxk (38)
	// qpjctjw (32)
	// wttpvzh (23)
	// tgnqn (87)
	// eyrfvtl (61)
	// mzzxjcu (27)
	// velktz (145) -> dxzlkz, keylghg
	// hvdhkw (210)
	// zafde (41)
	// qmwvc (558) -> hibxz, xqbvg, bgjzy, mvswhtp
	// knrgg (10)
	// xrbuyn (38)
	// wsupp (35) -> ecgpjx, jiybx
	// krlep (54)
	// gvajgxp (40)
	// yaudo (40)
	// prdkf (96) -> chrqi, tqjlm
	// suxpc (66)
	// pnmfw (51)
	// jldaz (64)
	// rucse (85)
	// ivcedgz (46)
	// zytvsav (21)
	// tytka (53)
	// pfdmmg (11)
	// hmumqsz (47)
	// cncicpd (93)
	// rrixk (27)
	// uzyprc (16)
	// ljctbd (13)
	// llgoq (76) -> rgwpu, movmxq, hvifpbk, hiiqp, kszkv, bjtza, hfttss
	// gqoxv (1646) -> bvikij, dcqpq, kquxfy, qrfgzm, qomxhp
	// apktiv (8441) -> obxrn, kqaksir, itjrw, supnaxw, wuefg, twtcx
	// nicjj (56)
	// nzmaui (18) -> wmxobe, eipmjyb, jovly, tjblqzk
	// ezspcab (84)
	// bclrfac (31)
	// jrcng (318)
	// pbkfd (157) -> pcdvdp, jldaz
	// vmqlqrp (46)
	// ynvqpm (92)
	// rnjfcg (20)
	// qhqnfsp (5)
	// ajozeez (92)
	// shciqu (127) -> ehywag, ajbtn
	// uujpt (29)
	// znooxvq (7)
	// mwblvo (257) -> ixmfiwz, jlcgqt, rmqdolg
	// rrazh (233) -> clsbdm, uavnq
	// ipjbc (92)
	// jovkydi (89)
	// rkfsa (102) -> pestqep, vnvkvb
	// owzwbpn (89)
	// cragbdx (19)
	// vtpldh (19)
	// rplxsw (82)
	// cbvamfw (47) -> wdnebh, vpqqbz
	// vhrxl (7)
	// cqvyvl (122) -> eemnlc, hehbbo
	// nmxmtaa (188) -> bvahtih, asklr
	// kfcgv (32)
	// pztxq (339) -> pqbar, bclrfac
	// gwfrqr (32)
	// dcqpq (1411) -> dqdpop, gtrgqb, lcfyznt
	// ksxixz (9)
	// itjrw (9) -> jjjpzm, mhofo, znljf, wrdgs, dcknzvh
	// omwrb (859) -> pqixzdw, hoqtxuf, weenw
	// rwgpi (20)
	// xoqxg (73)
	// rdyda (5)
	// dolng (83)
	// osvmh (77)
	// pqewl (185) -> hchmn, rhtdtxv, xzgejmu, fumvuu
	// dcknzvh (138) -> pdviq, hnpndnp
	// vjmwqzj (35)
	// zhttn (63)
	// lyuys (284) -> zicelk, joernlg
	// jjjpzm (164) -> vjgvk, yaiinhu
	// siruccf (37)
	// tzkuvl (87)
	// jpbiodh (10)
	// jodrf (56) -> lktgac, ysskib, qkougo
	// egvza (251) -> gotqku, viqlepb
	// jdpcmb (6)
	// ejlgch (133) -> mmoea, vvbcb
	// ecgpjx (88)
	// pdviq (41)
	// igqvq (288)
	// pndcqot (31)
	// kmjfxcf (15)
	// rqhfhc (12)
	// cgbgm (85)
	// jjmfi (33)
	// vpizq (37)
	// zynpfpv (73)
	// vnyllno (37)
	// sfyad (69)
	// sooqm (69)
	// dxdltnx (80) -> xtexo, ohpvt, acjtzxw, pnmfw
	// gcmpmnt (166) -> stqwvs, xmgkswu
	// jwgrqmj (44)
	// gcksx (73)
	// uswphji (1472) -> nzweur, pocxcw, xdlyd, joczsir, kymbpe, rhwgdsv, omwrb
	// lhifvp (49)
	// sgfco (305)
	// wqgcqb (12)
	// rgjxp (116) -> nujppls, eumzi
	// omzlzx (66)
	// jeplz (96) -> gpysit, gzmagb
	// bgjzy (78) -> ibmiu, ghaxvq, lkopm, khrqmbo
	// kvqsn (69) -> bpdixmv, knjlz, nmmtdv, edzgraw
	// ugkuxzz (314)
	// pryjaeo (21)
	// jzpwsg (78) -> paegovu, bxifcld, oijrg, wyftg
	// flufot (182) -> ftoskhn, afzlar
	// cmqaaw (83)
	// fvqvgw (26) -> ymfls, lgtpaqk, cumus
	// gurxxfd (46) -> qaveutv, bwtusip, jeuzbj, qifuph
	// jqtlcm (380) -> njeahp, sfyoyp
	// bywpbd (215) -> qrfuvjh, nrczybn
	// pcpnk (28)
	// pnjbsm (69)
	// lhxycw (161) -> achhc, ttwnws
	// qczhlyf (14)
	// sldytqh (49)
	// fmjhlia (20)
	// soqvass (53)
	// qptfj (21)
	// izuebwg (37)
	// iaafo (1335) -> iggtk, gkjbikb, gcvpqk
	// bkzdepq (56)
	// nqzwt (21)
	// frnhsjr (71)
	// nfwlplx (92)
	// egsgwch (279) -> jxwyzy, sbnot
	// ntonira (81)
	// ahrfot (9)
	// zjbiqnt (54)
	// tpozmfd (27)
	// jopepd (27)
	// kkukqoj (38)
	// panao (88)
	// szsntp (88) -> uvvqcxz, vboha, xgzexud, aaqcpt
	// lfqca (65) -> kieka, rfqenw, srbuc
	// ttwnws (30)
	// iobcvl (20)
	// gmeueu (45)
	// jazkpl (32)
	// carbhvi (118) -> vxrzo, ujlhns
	// jlyaty (11)
	// jwbxyd (75)
	// mlmzsqc (78)
	// sqpdsk (86)
	// gkijw (268) -> hgbkwjv, stmzb
	// sebxl (52)
	// anwxvv (59) -> qnjpz, kbibi
	// cylfwm (81)
	// sbnot (13)
	// cdcnp (5)
	// sohzsx (80)
	// gfmgl (47)
	// cbdczg (47)
	// uamje (82)
	// phrume (20)
	// jhpijl (185) -> mojjl, yhjopn
	// czzpk (17)
	// mwpdil (37)
	// lizssx (47)
	// otxphme (86)
	// vvnohc (11)
	// xmusk (10)
	// pqgpuv (210) -> lnnwiq, qqosg
	// iqumjrz (19)
	// khtjh (86)
	// sqfgpm (41)
	// qpqplm (30)
	// pqlav (8)
	// nzhprt (93) -> sohzsx, vaylgz
	// iehfo (624) -> fcsfrg, cbvamfw, dtqzu
	// vhfwgvx (83)
	// tbkhho (97)
	// uqymxu (58) -> egqmgr, izuebwg
	// xzgejmu (188) -> cwiwn, vtadj, tcbpqu
	// gjnnmvb (20)
	// ppmlbky (544) -> kvqsn, kcgui, chetpy, pztxq
	// jqqojdl (33)
	// qksclw (69) -> inxivh, dxbxha, irjvpam, syuoewb
	// axxkh (222)
	// tnccv (259) -> wttpvzh, edgdvfa
	// cumus (75)
	// jxkofob (1875) -> mrizfl, zgqsgm, bclse, ciojx, rgjxp, yifny
	// uydtye (8)
	// tuyyte (82)
	// twgbxu (63)
	// ulpkj (40)
	// lunefek (35)
	// pestqep (60)
	// kvxnjmq (25)
	// siitl (61)
	// chutuh (13)
	// gtjbuae (71)
	// fbrwdf (17) -> jzjrtvd, gtjbuae, zsnge
	// tukyh (71) -> ggzym, rfvehkx
	// xqpffoe (68) -> elthna, mwpdil, vpizq
	// iepmwr (176) -> irsfgz, rbyiay
	// lntikva (47)
	// kmodn (253)
	// ufsmed (81)
	// psrqv (54)
	// xluwv (30)
	// qombesj (380) -> rdghpyb, ivcedgz
	// shhfy (76) -> zynpfpv, kppxrk
	// fvlmvtk (130) -> ilpuzq, jzndr
	// gdqtqg (155) -> lhifvp, jfaca
	// osgjrb (69)
	// uqjge (88) -> nzetqt, nitvw, wxemcge, uydtye
	// ppafb (30)
	// opduu (38)
	// fpank (10)
	// sngnjdu (15)
	// ndidblx (52)
	// asklr (19)
	// wmxobe (61)
	// vvnzr (48)
	// usggwvu (375)
	// sukktk (202) -> oevbo, sitjrw
	// mqztoa (26) -> dzjsx, wzatr
	// kfxhl (40)
	// jsjexer (258) -> ljctbd, chutuh
	// bremy (70)
	// xugyewm (23)
	// djrrc (39) -> egcat, wdazzdu, zjbiqnt
	// juqbdco (144) -> vggsia, tytka
	// rfouw (28) -> yxzzaaq, zbfut, mpwldri, ljmve
	// svngfr (83) -> odgzsnk, rzrzage, lbzddgt, hvdpzbv
	// viqlepb (33)
	// pgsokae (63)
	// cobwyky (76)
	// zgjoaiw (25)
	// ohvifiy (6)
	// simmjy (65) -> mlksi, mfawmsq
	// zrqqtx (45)
	// yrpat (120) -> zlkzyr, zafde
	// nevkec (53)
	// gvxhl (125) -> liymk, psrqv
	// ohpvt (51)
	// kmqurp (155) -> ycyyvdy, slmxb
	// acrsmhw (99) -> cncicpd, tiwakam
	// yrdsj (149) -> mmtyhkd, omshqjl, pryjaeo, zytvsav
	// pahwxj (98)
	// nhnrpz (17)
	// mfzvywf (14)
	// fhphiyb (45183) -> ixktgj, zfzum, yzyzmzw
	// eduwet (1794) -> tzkuvl, jcmte, rndomrv
	// akrgqe (74)
	// bzjctqm (115) -> dphcbfr, sauzcee
	// xqbvg (104) -> cdcnp, wixfh
	// ljzuyyk (62)
	// egusv (173) -> qpxbow, gvajgxp
	// grmhpg (75)
	// gtbpbl (252)
	// fkdukrv (331) -> zbrbb, zsoro
	// aondpve (53)
	// vgsnxlm (23)
	// hcptw (44)
	// zkgoa (77)
	// dasyq (66)
	// skumcr (55)
	// afatio (30) -> hjxcpi, qmddtb, ppafb
	// ioobamp (9)
	// wrprrev (63)
	// nrczybn (9)
	// hoxdnl (32)
	// ndcfy (19)
	// tdecuga (85)
	// ugqqv (70) -> akrgqe, aeawrmy, ozjbwv
	// ravjm (83)
	// tsizsce (52)
	// uyopxvt (56)
	// djgzb (79) -> gpysm, rucse
	// ozbfh (64)
	// rwepl (73)
	// vkfrqct (56)
	// mansh (22)
	// bnsivw (336) -> vobnpuq, hoxdnl
	// omshqjl (21)
	// lktgac (76)
	// hgrua (66)
	// vrxeemw (49) -> vvnhe, nyejrrv
	// ydfajlk (88) -> qgjqyh, alneqju, jmchsu
	// ewjrko (167) -> kbrxrks, cygdj
	// oeftx (49) -> bpebim, sdbmop, nlpmbrd, yulga, qtlzten, hidrx
	// ypogtw (85)
	// eoqtf (97) -> ktejrze, osgjrb
	// cbbcj (46)
	// dvlyc (16) -> arkedwb, nzhprt, zjaqq, cjxyt, gdqtqg, jhpijl, egusv
	// tvjwhhy (120) -> tkjeu, vaeuo
	// cahxjyq (52)
	// muarkn (85)
	// zzjuf (95) -> rqkfkxw, lizssx
	// cygdj (56)
	// vlagh (5)
	// vohepl (98)
	// pgprg (152) -> jqdmk, bxefs
	// yrwyc (5) -> ocpngbz, ybfew, nyepy, cmcto
	// asfqik (180) -> orcyokv, tuyzi
	// jcyciyq (12)
	// twnfzz (149) -> vbateme, olydb
	// jofvyvx (25)
	// tazfdb (43)
	// jcmte (87)
	// ttolm (173) -> vmqlqrp, vffew
	// fndeqk (27)
	// xydxd (56)
	// elthna (37)
	// uspimc (85)
	// zatwq (66)
	// mwpbyo (6)
	// joaed (205) -> ycxlhqo, mgzox, nmemj, ljmkz
	// vtadj (70)
	// ilpuzq (50)
	// ubksf (106) -> nrbcaqo, shnrfq
	// mmnejdx (10)
	// srqclhj (90) -> jqqojdl, krzjli, ltmwbs, mtrnyd
	// rpegyn (23)
	// nzsnkla (265)
	// zfhkfwn (7)
	// hrjgbc (94)
	// stqwvs (28)
	// etkdg (38)
	// fumvuu (86) -> xcldsl, memrqz, ryfse, vauwilt
	// uavnq (10)
	// yaiinhu (28)
	// mlksi (92)
	// oyguh (32) -> ipxkky, lqmdutk, hrtgwt, kvxnjmq
	// ceeem (25)
	// pqbar (31)
	// aeawrmy (74)
	// fcufg (88)
	// hjnhdkp (27)
	// xmxzvhr (14) -> cvstod, jxcnmmo
	// mketjaw (45)
	// srbuc (80)
	// qqosg (20)
	// ynrctr (140)
	// kbibi (71)
	// lalfu (50)
	// iamrpx (96)
	// xgfsqq (309) -> qwvmczr, ladcna, fkoqh
	// dobmve (90)
	// nvvca (15)
	// vtaejs (39)
	// xgzexud (79)
	// tmvkuob (15)
	// foxjx (60)
	// ladcna (97) -> hbdwzm, cmqaaw
	// klnew (68)
	// vtdwe (37)
	// eayhoi (58) -> tvgytpm, pvlirjn
	// nccoxir (52)
	// wibfie (90) -> uqymxu, oyguh, wxoqoa, alomcle, fvikm, kykfb, ebgsk
	// yhcsc (34)
	// ptzbutd (38)
	// xmgkswu (28)
	// ajhio (301) -> jjmfi, mfecx, zjymq
	// lasluxv (87) -> etkdg, mckrfxj
	// msyvug (76)
	// weenw (188)
	// drtrgwp (185) -> hmumqsz, lntikva
	// gyojd (214)
	// zwrfiyf (17)
	// gquil (84)
	// erghvss (45)
	// mqmgirb (251) -> buffpzu, qpwiok
	// xaali (31)
	// clcixal (65)
	// pyxgmtu (169) -> hjnhdkp, jopepd, fndeqk
	// grgsn (61)
	// finkdao (78)
	// njeahp (10)
	// hidrx (175) -> ineay, vtpldh
	// qcyypa (93)
	// gnwlorj (48)
	// xjjdapk (64)
	// vxrtejs (30)
	// uukshmq (63)
	// hzcanxq (214) -> peuadz, pqlav
	// ptizsk (10) -> cbirs, syurke, ihpmp
	// cjoya (196) -> pcpnk, dbuccip
	// qpaei (86) -> djyxrkb, cupbfut
	// tjblqzk (61)
	// xjnhqlg (92)
	// qscpjx (39)
	// xptjtlm (108) -> utsob, lmxnloe, vrzbmt
	// rgwpu (161) -> xjuoid, oszlto, mketjaw
	// erdxj (299) -> mdprv, zdtvktq, labnsw, vmhwy
	// stmzb (32)
	// vxjjxhz (49)
	// xxlbk (92)
	// fxwez (7)
	// hbdwzm (83)
	// bjdtdn (61)
	// wnfqg (99) -> ravjm, dolng
	// sauzcee (67)
	// nkfrt (98)
	// rfvehkx (91)
	// fnzjtvw (88)
	// wdhmsi (45)
	// slmxb (17)
	// dkoxq (128) -> hvpcvdf, mdkes
	// qheany (80)
	// lpziczm (229) -> xydxd, fuqvw
	// gzmagb (22)
	// bgjngh (20)
	// krfgye (51) -> gzvjxk, dvlyc, niwzp, olqbic, iaixlte, tycqst, kfqrgj
	// dcwfs (99)
	// ixktgj (955) -> eabhh, iqvhov, lemnz
	// mbtsi (300) -> gbixdw, lalfu
	// dqklqo (57) -> jkaxk, nkfrt
	// qwsan (44)
	// iggtk (235) -> gyuyb, cjjbqxl
	// ascyv (15)
	// gtrgqb (17) -> coaznz, spvcd
	// awoedy (34)
	// mtrnyd (33)
	// wofung (133) -> ptzbutd, jlpxjeo, cnqrxk, crrfxfn
	// dpiyhv (86) -> tsizsce, gfdyf
	// surri (214)
	// tusgzei (20)
	// kzcyjtb (78) -> sqmfis, ppfmc, ldnnoag
	// memrqz (78)
	// javiioo (87) -> gnwlorj, tmbtipi
	// dzqrtc (206) -> zbwxa, ykpsfj
	// jtpgzp (31)
	// yndyrey (189) -> iuzvl, poinih
	// hkazlkt (235) -> adxeaf, rkfsa, yzbccbx, lkdkyk, gcmpmnt, zdqcu
	// bjqkaya (78) -> szsntp, nandmg, kjlxeat, kywmnbd
	// guywt (55)
	// qwcrc (63)
	// asdlfku (80)
	// vhjcue (81)
	// hiiqp (148) -> lhvyfsb, dltungt, tbrznk, vnyllno
	// deczzuu (92)
	// mymke (98)
	// subwna (61)
	// kmcad (1022) -> bzjctqm, akniuo, simmjy
	// uufxdt (328) -> hyfuy, iepmwr, asfqik, qxkas, nayudfl
	// ejkei (89)
	// afzlar (34)
	// buffpzu (33)
	// kyjgf (8)
	// bigkiu (748) -> uegcs, umgzhmp, frinpfj, nurclfr, hvdhkw
	// ryfse (78)
	// olztzl (25)
	// tgkbxa (33)
	// nktmu (86)
	// zlkzyr (41)
	// pmgwwzy (191) -> rgyco, qptfj
	// jfmts (56)
	// zacsvwk (43)
	// znljf (175) -> refya, tmvkuob, sngnjdu
	// kwozg (73)
	// eklqi (90)
	// cdwkezv (62)
	// ugmlsij (121) -> kpghxz, kravhjd
	// awqat (372) -> fdmbkt, znkchkk
	// ydgzjs (110) -> ymyzead, yzxqp
	// tugrmnp (60)
	// fmarl (5688) -> xbtswv, jvqwi, slgiv, uswphji, syjvwzt
	// memkrd (7478) -> nivpffu, bjqkaya, qegmu
	// jvfmwp (27)
	// qhqyt (66)
	// lbzddgt (73)
	// qsmvtif (71)
	// ytmpzku (49) -> jwgrqmj, nfrtom, lfiqwye, jkvduo
	// tdgel (61)
	// shnrfq (45)
	// hhrqz (129) -> rrixk, jvfmwp
	// ojatf (9518) -> qmwvc, wibfie, dnycw
	// yxzzaaq (55)
	// wjipa (1180) -> mqllnxu, xmxzvhr, qbynkmw
	// cwxgqrk (15)
	// aotajs (133) -> ucezr, mmnejdx, jpbiodh
	// lsspa (60) -> paldf, grmhpg, exjsjxd
	// ineoncq (11)
	// eipmjyb (61)
	// fjgsw (29)
	// wnwsbdq (50)
	// hdgvs (37) -> msyvug, bntdh, sjxaxv
	// npckah (6)
	// swqkqgz (19)
	// eabhh (379) -> jzpwsg, bkwcwf, mwblvo
	// ggaxx (87)
	// rjvfwcu (64)
	// bwtusip (44)
	// slgjon (63)
	// vlpqaxo (66)
	// mmoea (11)
	// xnltw (93)
	// tgpxvyr (37)
	// phdcpp (63)
	// dltungt (37)
	// ljvrcj (83)
	// upaqlj (44) -> cobwyky, pngubp
	// hwriv (97)
	// hopjmyt (62)
	// sizmuwa (180) -> sooqm, sfyad
	// jzndr (50)
	// ppmrgga (222) -> cvvncju, xmusk, trmwq
	// uvswv (90)
	// xstrla (46)
	// pcdvdp (64)
	// kszkv (164) -> hgrua, qhqyt
	// niwzp (855) -> nnhknbl, ichznto, llmmm, hpmuqvl
	// kybegv (96)
	// rxcbsdp (30) -> ynvqpm, pjaxkr
	// ihpmp (85)
	// hseaxj (51)
	// fuzhuf (38)
	// plbtdq (633) -> gkuma, jrcng, dkoxq, vtylgti, qifoay, sizmuwa, pjxttn
	// esooc (151) -> vtdwe, ezdiq
	// jvqwi (33) -> csaqixs, rnlaw, hvkhvwl, rkpcoen, nptoou, ypvztl
	// luqyr (81)
	// ggzym (91)
	// fcsfrg (113) -> siitl, dkntzn
	// nnbbrbf (135) -> uspimc, xkzgz
	// cwiwn (70)
	// zragt (58)
	// glyxk (22) -> pvhrim, ndidblx, sebxl, sstvew
	// rhwgdsv (481) -> ugkuxzz, yyjqegu, yrhid
	// bmltf (23)
	// gfdyf (52)
	// cbirs (85)
	// xhtzti (21)
	// advhon (65)
	// xcrkb (864) -> ugmlsij, jpkwter, ejlgch
	// kfqrgj (347) -> igqvq, cqvyvl, jzgzd, ashxu, clwwv
	// ouubjrl (36) -> iodveqh, wanom
	// nptoou (269) -> yrdsj, bywpbd, lmcqig, skgwztg, suoiohi, gvxhl, rgdvmy
	// jqdmk (22)
	// ljmkz (20)
	// dxzlkz (60)
	// svhqwim (9)
	// nurclfr (30) -> usuaiv, zyjxpjz, devljb, erghvss
	// vdzuw (146) -> ostrl, lmvvs
	// fsokbvd (262) -> owmjbhg, sshfxeu, zufoz
	// cdvkf (6)
	// lrtds (202)
	// ybfew (377) -> zuybvj, bvnjiou
	// lfiqwye (44)
	// htbjl (57) -> czdpe, gcqaj
	// cjjbqxl (18)
	// vucfp (53)
	// bnvlsm (51)
	// ybekxtf (10703) -> erylwj, xgfsqq, eygdy, dinlw
	// ashxu (96) -> ctinwus, ypdpmwi
	// qpwiok (33)
	// nrdlfop (84)
	// sbhiqhs (27) -> qxnkp, klnew
	// ujlhns (48)
	// ostrl (22)
	// vmjbgo (155) -> udmez, jefztzv
	// rkaxx (27)
	// pocxcw (823) -> blhgy, xptjtlm, kerjqr, ekfewi
	// gkuma (134) -> kvyjyt, foabep
	// ugrhs (32) -> mansh, hpkvfr, ojqia
	// gbixdw (50)
	// phkthld (75)
	// mofyda (109) -> finkdao, mlmzsqc
	// lxucxpm (73) -> iytknc, qhxmh, eltbyz, dfjvh
	// qgjqyh (77)
	// ucxgom (249)
	// mmufeg (13) -> gcksx, oybdjt, xoqxg, masykz
	// yexxg (319)
	// bvahtih (19)
	// lesmi (204) -> qwsan, hcptw
	// pmbnia (94) -> jsjexer, qjcqm, ftegwk, dxdltnx, negkd, jodrf
	// knmac (36)
	// hnpndnp (41)
	// qoijc (73)
	// lemnz (60) -> egvza, mqmgirb, dqwdx, whglbyy, lxucxpm
	// ihtfh (116) -> fplihc, vregap
	// lmxnloe (14)
	// whglbyy (201) -> vnkqnhh, zragt
	// vbateme (50)
	// inxav (77)
	// sykwd (860) -> jeplz, sttlom, ynrctr, cumfrfc
	// aaqcpt (79)
	// mwpatsx (64) -> nrdlfop, lerycoe, mygcpku, uusiaqf
	// uanvh (53)
	// pvlirjn (49)
	// kerjqr (150)
	// wrabgy (32)
	// xjzpum (148) -> vjmwqzj, uwzvy, vwbxg
	// tycqst (1112) -> ytmpzku, vwzmkq, esooc
	// gxxzwq (7)
	// knjbw (27)
	// hpkvfr (22)
	// xappjar (85)
	// oybdjt (73)
	// qmddtb (30)
	// krvyy (35)
	// mughfl (82)
	// hoqtxuf (74) -> xrbuyn, imyvlyt, fvmhrf
	// xtexo (51)
	// bknogxz (84)
	// rreqg (144) -> nccoxir, cahxjyq
	// icpwoof (172) -> wxhsqe, svhqwim
	// kszyl (113) -> sqfgpm, xmvdh
	// zyjxpjz (45)
	// ipxkky (25)
	// jxcnmmo (96)
	// ghkdvw (72)
	// cwphzk (25)
	// mgzox (20)
	// gcvpqk (143) -> ozbfh, opkvs
	// bompiu (25)
	// ckisc (103) -> alankuj, phkthld
	// ngtkzm (67) -> vkfrqct, usvkq, jfmts
	// qifoay (318)
	// rkpcoen (1788) -> uyopxvt, bkzdepq
	// vtuqq (30) -> tbkhho, xhzylb
	// ekfewi (96) -> knjbw, tpozmfd
	// paldf (75)
	// zghrr (25)
	// ggllv (66)
	// vlfpuo (32)
	// sdgdv (89)
	// osryil (62)
	// puexzf (36)
	// smrvw (39)
	// ozjbwv (74)
	// naglm (31)
	// aewie (20)
	// etoxfc (59) -> uxdeg, hcceg, ppmlbky, uufxdt, qbcscs, llgoq, iaafo
	// wxemcge (8)
	// xwkoc (63) -> ijcfw, aondpve, nevkec
	// byckty (80)
	// kvyjyt (92)
	// pkmmfc (54)
	// sxnsqj (9)
	// wtzkvm (23)
	// kfcowx (352) -> edoycls, pqdfbrl, nxdkpuh, acrsmhw, rrkqend
	// ngthpc (51)
	// nkycb (9)
	// otpnx (81)
	// jusfet (39) -> kivcyus, nxwjvx, vucfp, soqvass
	// mvswhtp (28) -> tazfdb, mzlmr
	// exjsjxd (75)
	// ugkxkqe (159) -> qbdafi, htdwe, olztzl
	// pikvdx (82)
	// vnvkvb (60)
	// sitjrw (23)
	// yyjqegu (94) -> guywt, skumcr, whoqvq, slrdn
	// qnrjqj (61)
	// szfxhin (17)
	// iiznokx (90)
	// vbpyoqo (97)
	// qwdhdrk (7)
	// iqvhov (120) -> atlrd, mmufeg, ouqtjz, xgitr, ezglz
	// wdnebh (94)
	// ylidl (82) -> wejvpzr, jlshk
	// qsfpaj (34)
	// qbcscs (1680) -> yazsie, eayhoi, btqebbp
	// auiiuv (78)
	// ixmfiwz (55)
	// hcceg (1083) -> dkoxwi, iwsknb, fkdukrv
	// kqaksir (152) -> ydfajlk, tyvhfhz, yexxg
	// wixlvcp (12)
	// tfhij (35)
	// vnkqnhh (58)
	// dkntzn (61)
	// oggnstj (279) -> ydgzjs, havdbe, dzqrtc, nzmaui
	// yzyzmzw (583) -> yrwyc, suiyl, kmcad
	// hfttss (5) -> xatua, hwriv, vbpyoqo
	// kieka (80)
	// czdpe (69)
	// vxrzo (48)
	// mpmue (9)
	// ujsbt (68)
	// oijrg (86)
	// vaeuo (41)
	// qbynkmw (114) -> mgyhaax, cbbcj
	// srgmt (165) -> msskvkg, foxjx
	// ycxlhqo (20)
	// pmfprs (20)
	// zbsxee (35)
	// ypvztl (1363) -> xqpffoe, neyeo, lccoo
	// hwrwnh (90)
	// szwpt (29)
	// mhofo (66) -> inxav, osvmh
	// masykz (73)
	// gawck (98)
	// drtxytc (81)
	// wyftg (86)
	// ssvsso (22)
	// koeqohh (275) -> zupym, hxeoxk
	// kcgui (335) -> kguwwze, gcefq
	// olqbic (911) -> ihtfh, ugqqv, lesmi
	// sttlom (62) -> zknesz, gethyvd
	// acjtzxw (51)
	// supnaxw (194) -> iqsztjd, urpnw, opowpvm
	// jzjrtvd (71)
	// rbzqabo (6)
	// keoaqjb (59) -> jqtlcm, wdhbym, mwpatsx, mbtsi, bnsivw, zqurr, ajhio
	// kkkjsil (39)
	// jzgzd (188) -> zorvsm, ainstbs
	// iitweo (96) -> ckisc, ysltqk, zosskdz, xjzpum, rrazh, htsuvhg, xwwxswj
	// xyzbni (58) -> xnltw, lnufi
	// rgdvmy (53) -> hwrwnh, eklqi
	// hhdzz (227) -> cragbdx, iqumjrz
	// mcjgfx (68)
	// jlpxjeo (38)
	// zfzum (16) -> lnaoiv, qkhtsa, lsvqox
	// njatvu (26) -> ehtsbv, hkazlkt, tvgwgq, zktmxll, jjukf, kmdzlmk, svtwviu
	// wdsbi (98)
	// ieuwo (196) -> hdtcizc, ggpjxwv
	// tuyzi (92)
	// dkoxwi (35) -> pwumvgy, asdlfku, aedcmjb, tyuhzom
	// ichznto (221) -> cdvkf, jdpcmb
	// tyvhfhz (179) -> nystxqv, caocs, zbsxee, lunefek
	// uxdeg (443) -> eewnnkx, dryngd, lpziczm, koeqohh, uijtrw
	// egcat (54)
	// kvqspmf (36)
	// ydmrd (90)
	// syurke (85)
	// tkeuvp (159) -> jtpgzp, bfcdy
	// ypdpmwi (96)
	// nmemj (20)
	// eewnnkx (341)
	// bafdzbu (53)
	// wgivp (71)
	// tvgwgq (835) -> prdkf, pnmvk, xyzbni
	// slgiv (7452) -> oeftx, erdxj, oggnstj
	// gpogy (61)
	// xdrge (889) -> shciqu, aotajs, douvhy, lasluxv, sbhiqhs, elgyjo
	// vbdcxx (93)
	// oydsj (80)
	// pjxttn (168) -> jwbxyd, isymgjd
	// edkwqih (60)
	// wanom (89)
	// hqzay (92)
	// xwfbb (20)
	// phyzvno (20)
	// glrua (69)
	// ijcfw (53)
	// alomcle (112) -> fpank, rfwvc
	// lmcqig (71) -> ufsmed, cylfwm
	// ltichp (80)
	// tzltv (63)
	// thknml (89) -> luqyr, janrdqf, wkxkv
	// zsnatp (118) -> orwbdwn, tsiyp
	// zoijv (130)
	// aynpr (63)
	// isymgjd (75)
	// pqdfbrl (117) -> nicjj, cdrhm, mdddafe
	// bonjgrt (319) -> fuvokrr, hqlyg, mozvs
	// vggsia (53)
	// ucbuez (14)
	// wixfh (5)
	// dgszhd (18)
	// joernlg (67)
	// elkflt (78)
	// lhvyfsb (37)
	// liymk (54)
	// xgitr (139) -> vhfwgvx, rqrtz
	// wxhsqe (9)
	// wssvxr (85)
	// bgmmb (17)
	// ehypcwy (334) -> xvlymgg, tgkbxa
	// qwbfod (68) -> ndpwgdy, knmac
	// nepxnu (6)
	// ftoskhn (34)
	// ponlfyf (82)
	// atkmjxg (162) -> xluwv, dinng
	// rrkqend (223) -> xaali, sxmdnhl
	// sbnsoxi (15)
	// fxemb (62)
	// sfyoyp (10)
	// lyhfj (79)
	// nnkmf (15)
	// xfmxo (73)
	// dilqo (84) -> ljvrcj, ognax
	// jxbksoh (157) -> xstrla, ykudi
	// fkhxkeg (26)
	// wsvir (36)
	// chetpy (363) -> ndcfy, ftdylco
	// iytknc (61)
	// hvifpbk (202) -> mhkxu, gfmgl
	// gyuyb (18)
	// whoqvq (55)
	// twtcx (524) -> htbjl, kszyl, hvapnf
	// wkxkv (81)
	// ainstbs (50)
	// ppfmc (35)
	// hvkhvwl (1351) -> hhrqz, kzcyjtb, javiioo
	// yeqnq (15)
	// aruczxj (31)
	// poinih (23)
	// kymbpe (745) -> ihmqs, uthjdye, nmxmtaa
	// kixaco (11)
	// oszlto (45)
	// ydmeqtl (34)
	// mojjl (34)
	// adxeaf (212) -> rdyda, vlagh
	// xjuoid (45)
	// rndomrv (87)
	// mhkxu (47)
	// znubct (69) -> dasyq, vlpqaxo
	// fdmbkt (50)
	// coaznz (84)
	// caxbgqp (23)
	// uhwvnbg (28)
	// orcyokv (92)
	// zgzsu (10)
	// xgwjcx (22) -> cfdpxpm, pmgwwzy, bscpyic, vmcgj, jaqwzi, spxwcuv
	// vpqqbz (94)
	// wvdapsm (65)
	// uapwikr (12)
	// olydb (50)
	// kjgtfqc (69) -> pdvolf, sgngx
	// cklpcr (58) -> hwjhf, foaayon
	// ottiad (63)
	// bjtza (250) -> caxbgqp, rpegyn
	// fdhdpw (18)
	// jqjkgv (50)
	// kguwwze (33)
	// blhgy (100) -> cwphzk, ceeem
	// aedcmjb (80)
	// drosj (98)
	// jbqvpsb (12)
	// isrch (7)
	// vzgsgk (10)
	// fjqomax (82)
	// twvdhg (55)
	// vvbcb (11)
	// yazsie (78) -> qscpjx, mxbyg
	// nkkgyl (166) -> kvmqsdx, znkzp, lyuys
	// mrizfl (108) -> ittmm, eyyokyd
	// fkkzg (145) -> vgsnxlm, bmltf, lircjh
	// zepvwyv (26)
	// idjhk (47)
	// hrtgwt (25)
	// iaixlte (717) -> surri, mkmci, rfmiqz, gyojd, rxcbsdp
	// ubxke (51)
	// imnvt (48)
	// hgmjvpp (35) -> wrpqf, uwpbnv
	// tiwakam (93)
	// rfwvc (10)
	// zjqeu (120) -> fuzhuf, opduu, ijictm
	// ymyzead (76)
	// veboyvy (61) -> qwjmobb, fmarl, hqxdyv, fhphiyb, cmmqbz, kqoigs
	// nwtlu (64)
	// rdghpyb (46)
	// srcyajr (77) -> vxjjxhz, aevhbim, sldytqh
	// hpltoci (199) -> qczhlyf, mfzvywf
	// sdbmop (35) -> vtkgj, jovkydi
	// nyejrrv (89)
	// negkd (266) -> zwoot, mpmue
	// ykudi (46)`
	// 	out = advent7A(input)
	// 	fmt.Printf("Result 7A %s\n", out)
	// intOut := advent7B(input)
	// fmt.Printf("Result 7B %d\n", intOut)
}

func TestDay8(t *testing.T) {
	input := `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`
	out := advent8A(input)
	expected := 1
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	out = advent8B(input)
	expected = 10
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	input = `d dec 683 if qn == 0
d dec -220 if h == 0
rak dec -875 if rak > -9
isy dec 250 if wf == 0
cie dec 20 if rak > 870
isy inc 93 if wf >= -5
o dec 739 if bok < 8
rak inc -605 if mxg <= 9
rak inc 668 if rfw > -8
rfw dec 214 if h > -7
j dec 649 if wf != 4
bok dec -712 if cie >= -22
s dec 151 if rxb == 0
bok dec -656 if d <= -463
pf dec -435 if brr != -10
pf dec 115 if rxb < 6
uxr dec -574 if brr == -3
h inc -34 if s == -151
rxb inc -919 if rak == 938
s inc 627 if o <= -748
rxb dec -456 if rfw != -214
rak dec -687 if x <= 8
d dec 292 if bok >= 1363
mxg inc 665 if o == -739
brr dec 531 if bok == 1368
isy dec -890 if x != 7
o inc -649 if bok > 1361
erb dec 656 if j != -648
vso inc -882 if wf <= -2
rxb inc 978 if brr >= -539
pf dec -176 if wf <= 8
rxb dec -647 if hsn == 0
vso inc -2 if isy <= 725
brr dec -661 if x != 4
uxr dec 913 if x <= 9
hsn inc -784 if rxb <= 706
pf dec -795 if hsn <= -782
pf dec 421 if rfw >= -217
o inc -88 if s < -147
rak inc 991 if hsn == -784
brr dec -27 if bok != 1372
uxr dec 705 if x >= -2
tss inc 130 if h != -35
vso inc -800 if wf != 8
tss inc 301 if pf != 866
ie inc -311 if cie == -20
mxg inc -20 if vso >= -808
cie dec -142 if bok == 1368
rfw dec -141 if rxb != 706
isy dec -826 if j > -656
cie inc 252 if pf >= 870
hsn dec 972 if bs < 2
hsn inc -576 if rfw >= -219
bok dec 439 if h >= -40
cie dec 614 if h == -34
cie dec 890 if cie == -240
wf inc -251 if pf <= 875
isy inc -507 if o == -1476
bok dec 619 if bok < 928
rxb inc 361 if rfw == -224
rfw dec -654 if s < -141
mxg dec -488 if isy >= 1047
pf dec 721 if mxg >= 1135
h inc -4 if rxb == 706
hsn dec -966 if rfw >= 431
isy dec -422 if h < -44
rfw inc -379 if qn == 0
pf inc -78 if cie >= -1136
o inc 902 if isy > 1042
erb dec -529 if brr <= 160
x dec 280 if o <= -572
isy dec -266 if d < -751
j dec -260 if isy != 1313
rak dec 330 if qn != 9
tss inc -113 if ie <= -311
isy dec -186 if erb > -135
o dec 839 if pf <= 800
x inc -906 if uxr > -1622
hsn dec 353 if vso > -808
cie dec -734 if bs > -9
rak dec -690 if rfw == 61
tss inc 355 if rfw == 54
isy dec 507 if erb != -127
uxr dec -901 if hsn >= -1726
tss inc 787 if cie == -396
j inc -851 if d < -753
d dec -929 if pf < 790
h inc -328 if mxg <= 1140
j dec -381 if erb != -122
o dec 810 if rxb <= 696
brr dec 207 if x == -1186
rak dec -697 if erb < -135
ie dec 96 if vso <= -798
vso inc 739 if rak < 2968
o inc 890 if mxg == 1133
wf dec 453 if pf <= 785
rxb inc -950 if bok > 923
brr inc 160 if isy != 1504
isy inc 767 if rak < 2967
bok dec 297 if bs != 2
qn dec -686 if erb == -127
uxr dec -835 if h >= -370
rxb dec -315 if pf < 799
h inc -97 if j != -864
bs inc -380 if erb < -117
rfw dec -444 if tss <= 1108
brr inc 951 if wf <= -251
erb dec -665 if h != -465
qn inc -350 if isy != 1504
h inc 603 if bs < -373
vso inc -545 if uxr != 111
o inc 110 if brr > 894
s dec -924 if rfw <= 509
hsn dec -114 if rfw != 507
tss dec -441 if erb != 529
d inc -791 if bok > 630
rxb dec 576 if pf >= 800
x dec 797 if j >= -859
cie inc 828 if rfw <= 507
s inc 945 if mxg != 1127
mxg inc 122 if qn <= 685
rak inc 778 if pf <= 786
tss inc -795 if pf < 794
vso dec -412 if d >= -1553
vso inc 371 if bs != -385
wf inc 443 if uxr != 109
wf dec 366 if rak == 2976
tss inc 608 if vso >= -571
x inc 316 if qn != 694
o inc -570 if mxg > 1136
qn inc 346 if bok < 636
erb dec 947 if erb != 529
uxr dec -673 if h != 147
pf inc -252 if d == -1546
ie dec -157 if o < -406
vso inc -46 if rfw > 499
qn dec -915 if vso <= -617
bok dec 288 if rak < 2983
o inc -842 if uxr == 783
rfw dec 867 if mxg > 1136
s dec -639 if uxr == 791
uxr dec 825 if erb < -406
rxb inc 414 if j >= -866
d dec 695 if vso > -612
mxg dec 119 if uxr == -34
tss inc -168 if bs > -379
tss inc 379 if mxg <= 1017
mxg dec 629 if ie != -250
rak inc 803 if d != -2245
pf dec -393 if rxb > 488
ie dec 574 if tss <= 1740
bok dec 498 if brr <= 907
rxb inc -999 if x < -1663
h inc -716 if s != 2359
vso inc -635 if cie > 430
tss inc 712 if mxg > 1009
x dec 303 if mxg == 1016
x dec -426 if rfw <= 510
rxb dec -99 if wf <= -174
ie dec 955 if x == -1241
ie inc -6 if qn < 1036
o inc 231 if pf != 541
d inc -967 if bok < -147
bs dec -86 if qn <= 1037
rak dec 255 if x != -1240
rfw dec 143 if d >= -3204
qn dec 175 if hsn > -1614
j inc -550 if isy == 1504
cie dec -553 if d != -3209
wf dec 353 if hsn == -1605
vso dec -563 if d <= -3205
o inc 593 if wf == -535
vso dec -519 if vso != -677
rfw inc 122 if s >= 2357
wf inc -889 if vso <= -166
ie dec -73 if rak >= 3531
pf dec 79 if d <= -3201
hsn dec 164 if cie > 982
rak dec 237 if rxb != -419
qn dec -585 if tss > 2442
rfw dec -186 if rfw >= 623
uxr dec 520 if rfw < 822
x dec 447 if x > -1251
vso inc -796 if pf < 464
tss dec 34 if rxb == -415
x inc 603 if cie != 994
uxr dec -234 if h <= -577
rak inc 381 if h < -571
j inc 692 if pf >= 455
pf inc -279 if pf != 461
bs inc -401 if h == -576
o inc -629 if mxg != 1018
hsn dec -619 if x < -1086
bok dec -91 if qn >= 1433
bs dec 660 if isy < 1495
x dec -908 if vso != -960
rak inc -16 if x >= -185
mxg dec -367 if mxg == 1010
bs inc 487 if o > -802
mxg inc -608 if o == -811
brr inc -730 if d < -3198
mxg dec -872 if qn != 1438
bs inc 467 if cie >= 983
isy inc 504 if cie < 981
rxb inc 356 if isy >= 1501
pf inc 427 if mxg < 1279
brr inc 967 if isy < 1510
rxb dec 561 if j > -720
hsn dec -175 if d != -3218
wf dec 700 if hsn >= -1602
h inc 637 if brr >= 1137
uxr inc 964 if wf >= -1232
rxb inc -739 if uxr > 401
erb dec -494 if qn == 1450
rxb inc -226 if qn != 1436
brr inc 131 if bs <= -220
rak inc -899 if mxg == 1274
tss dec -737 if bs != -238
j dec 700 if wf == -1220
j dec -823 if rak > 3654
j inc 466 if d > -3211
ie dec -521 if o == -811
d inc -898 if qn > 1440
isy dec 197 if rfw <= 815
erb dec -911 if vso > -952
j inc -835 if cie >= 983
erb dec -315 if uxr != 416
brr inc -673 if erb <= -99
erb inc -72 if vso <= -949
vso inc 165 if mxg <= 1285
qn dec -870 if vso >= -799
pf dec -574 if j == -1086
pf inc -858 if erb > -173
erb inc -243 if cie >= 985
rxb inc -927 if x == -177
rxb inc -260 if vso < -792
tss inc 62 if rak > 3647
tss dec -675 if vso <= -792
qn inc 807 if rfw <= 810
bok inc -360 if d > -4101
s dec 56 if brr != 1260
wf inc 855 if j == -1096
j inc 698 if s <= 2304
isy dec -459 if hsn <= -1591
s dec -577 if bok > -71
wf dec -469 if j != -397
pf dec 127 if tss != 3897
o inc 442 if rak >= 3651
d dec 941 if j <= -383
brr dec 809 if rak <= 3652
o inc 161 if rfw != 820
h dec 284 if uxr <= 418
j inc -805 if pf < 475
uxr inc 731 if uxr > 409
isy dec -589 if isy <= 1772
s inc 358 if h <= -221
wf inc 680 if rxb == -2519
brr dec 919 if ie != -1270
tss dec -77 if h > -226
qn dec 953 if hsn == -1594
s inc -553 if rfw == 813
isy dec 389 if qn < 1367
qn dec 676 if vso != -782
mxg dec 328 if uxr > 1135
wf dec 788 if hsn < -1599
uxr inc -361 if brr == -459
isy dec -775 if rxb > -2504
uxr inc 680 if vso < -788
tss dec 368 if tss <= 3967
pf inc -494 if qn <= 681
tss inc 84 if bok == -63
rfw dec -146 if rfw >= 804
erb dec -175 if pf >= 484
pf dec -655 if hsn <= -1590
j dec 273 if pf != 1133
rak inc -305 if wf > -764
o inc -234 if rxb != -2503
rfw dec -499 if uxr < 1470
bs inc -102 if qn < 687
j dec 698 if o > -435
cie inc -460 if cie > 977
wf inc 134 if rfw <= 1463
uxr dec -390 if h == -223
uxr dec -834 if bok >= -64
ie inc 538 if s == 2688
cie inc 504 if o > -445
bok inc 102 if ie < -1255
pf dec -65 if erb != -408
vso dec -317 if vso >= -791
x dec -139 if brr >= -459
s inc 772 if bok < 34
bs dec 736 if j <= -664
j dec -352 if vso != -794
s dec -656 if cie == 1029
rfw inc -919 if x <= -44
o inc 101 if vso < -793
bok dec 37 if bok == 39
hsn dec 413 if erb > -414
d inc 686 if rak != 3337
d dec -164 if j < -300
brr dec 968 if bok >= 1
rxb dec 177 if cie != 1029
j inc 200 if o >= -448
brr inc -665 if vso <= -792
vso inc 581 if rak != 3345
vso inc 491 if wf != -621
rxb inc 611 if rxb >= -2513
bok inc -77 if x != -39
cie dec 354 if qn <= 691
ie dec -328 if uxr != 2693
qn inc 629 if cie <= 679
mxg dec -941 if x > -46
tss inc -153 if x >= -46
o inc -459 if h == -222
ie inc -744 if s >= 3333
tss inc 62 if brr != -2086
tss inc -27 if bs == -330
vso dec -435 if rak != 3341
isy dec -751 if x != -38
hsn inc 700 if mxg != 1898
hsn inc -383 if wf == -624
rfw dec -697 if rfw >= 1458
brr inc -446 if ie > -1682
d dec 43 if rfw < 2152
pf dec -798 if rxb == -1909
vso inc -517 if x != -38
rxb dec -595 if rxb < -1892
s dec -820 if mxg <= 1884
o dec -502 if uxr >= 2686
wf inc -541 if rak < 3340
d inc 285 if tss != 3572
vso dec -901 if brr != -2543
uxr dec 402 if rxb != -1306
erb dec -927 if erb > -416
tss dec -62 if rxb != -1312
o inc -194 if qn >= 1319
hsn inc 734 if isy <= 1972
tss dec 553 if qn > 1308
d inc -33 if vso == 1616
bok dec 775 if tss <= 3080
hsn dec 119 if cie >= 669
x inc -50 if j <= -106
erb dec -422 if qn >= 1309
bs dec -573 if rxb <= -1303
uxr dec 431 if uxr <= 2692
rxb inc 174 if uxr <= 2255
tss inc -613 if wf >= -627
brr inc 166 if d != -3936
rak inc -947 if brr <= -2365
bs dec -504 if qn == 1312
cie inc 259 if wf > -630
hsn inc -290 if hsn == -1075
isy dec -778 if mxg <= 1899
rxb inc -146 if x >= -90
bs inc -812 if cie > 929
cie dec -71 if d <= -3943
vso dec 235 if d >= -3935
qn inc 192 if rak >= 2391
rfw inc 403 if erb == 940
rak inc -881 if rxb >= -1285
x dec 30 if rak > 1510
s inc 541 if rxb < -1268
hsn inc -622 if tss != 2453
rak inc -778 if erb <= 946
hsn dec 300 if wf <= -628
hsn inc -72 if ie != -1688
erb inc -145 if erb > 936
x inc -99 if o < -440
vso inc 866 if bok > -842
vso dec -308 if uxr <= 2256
rfw dec 487 if uxr != 2260
rfw dec -21 if mxg > 1881
h inc 866 if rxb <= -1273
erb dec 234 if d >= -3951
erb dec -658 if qn == 1504
o dec 771 if rfw == 2092
h inc -843 if brr < -2377
ie inc -528 if bok >= -851
qn dec 187 if rfw == 2086
isy inc -308 if d <= -3946
rfw inc -163 if vso < 1930
isy inc 815 if h == 643
mxg dec 606 if h < 642
mxg dec 165 if vso > 1923
rfw dec -530 if uxr <= 2253
ie inc 129 if mxg == 1726
uxr dec -505 if h != 638
bok inc -535 if bok >= -851
o dec 941 if j != -104
rak dec -133 if rfw >= 2460
bs inc -461 if mxg >= 1727
j dec -938 if brr != -2368
erb inc 737 if x >= -207
tss dec -280 if hsn < -2051
s dec 651 if bs == -65
rak inc -979 if vso < 1934
pf dec -487 if hsn != -2060
qn inc 384 if j != 830
j inc -389 if o < -2147
hsn dec 14 if s > 3224
rxb dec -710 if uxr <= 2766
hsn inc 278 if s != 3229
tss inc 376 if brr < -2370
d inc 306 if j != 435
wf inc -561 if tss > 3122
h dec -154 if bs >= -65
x inc 396 if hsn != -2073
mxg dec -806 if cie < 1013
hsn dec 941 if hsn <= -2078
erb inc 18 if pf > 1677
x inc 63 if rak > -244
rfw inc 883 if tss == 3124
s inc -615 if erb < 1241
s inc -696 if vso < 1933
qn dec 120 if pf <= 1678
rxb dec -680 if rxb <= -563
wf inc -880 if ie > -2087
tss dec 422 if d <= -3641
d dec 270 if cie > 1004
h dec 993 if tss > 3114
isy dec 490 if o != -2144
bok dec -649 if h > -200
bs dec -505 if o != -2164
hsn inc -863 if rfw == 2459
bok inc -307 if bs <= 442
x dec 961 if rxb > 104
pf inc -606 if brr >= -2376
mxg dec 409 if qn >= 1880
rxb inc -962 if ie == -2079
bs dec 847 if hsn >= -2944
hsn inc 543 if erb == 1237
d dec -752 if ie != -2086
h inc 90 if rak != -228
j dec 620 if cie < 1003
o inc 827 if o != -2156
erb inc -601 if bs >= -408
mxg inc -616 if o >= -1335
h dec -583 if rfw < 2468
rfw inc -727 if erb <= 643
s dec 178 if rak < -237
rfw inc -781 if x >= -1122
hsn inc 503 if brr < -2365
h dec 738 if hsn != -1888
vso inc 10 if rfw != 961
tss dec 455 if j != 430
pf dec -654 if bok >= -1047
mxg dec -953 if hsn == -1890
rfw inc -859 if cie <= 1008
rxb inc -957 if x >= -1112
erb inc 775 if isy >= 3061
rfw inc -574 if rxb == -850
erb inc 304 if s <= 1746
bok dec 703 if bs <= -413
d dec 788 if bs <= -404
uxr inc -714 if bok > -1035
bs dec -244 if h <= -264
h dec -530 if vso == 1934
s inc -792 if bs != -407
j inc 384 if rfw >= -484
rfw dec -694 if cie <= 1013
hsn inc -53 if uxr != 2762
rxb dec 83 if j != 834
rxb inc 849 if o >= -1329
o dec -244 if j < 816
x inc 812 if rxb <= -78
pf dec -950 if brr == -2372
x inc -169 if hsn == -1943
mxg inc -673 if vso != 1929
tss dec -932 if pf == 2682
o dec 976 if pf <= 2685
mxg inc -90 if mxg >= 1783
pf inc -466 if pf <= 2690
rxb dec -592 if uxr < 2760
s dec 904 if pf <= 2216
o inc -958 if qn != 1888
d inc 317 if ie != -2077
d inc 732 if wf > -1511
tss dec -200 if tss <= 3584
mxg inc -894 if mxg >= 1692
uxr dec -883 if rak > -248
d dec 576 if x >= -481
mxg inc -418 if d <= -3478
rxb dec -435 if tss == 3594
rxb dec -200 if uxr != 3646
j dec 70 if d != -3473
cie inc 524 if d > -3469
tss dec 988 if rak <= -233
bok inc 567 if pf < 2226
tss inc -486 if isy > 3064
erb inc -952 if brr < -2380
uxr inc 723 if rfw <= 214
x dec -510 if ie != -2078
d inc 395 if rfw <= 212
erb inc 592 if isy > 3075
ie inc -61 if d >= -3080
rak inc 400 if qn >= 1898
erb dec 449 if cie > 997
ie inc -918 if pf > 2215
s dec -711 if rfw <= 220
rxb dec 519 if wf != -1504
ie dec -59 if rxb < 1145
erb dec 425 if tss == 2120
cie inc -306 if isy <= 3069
x dec 635 if o >= -2296
wf dec 815 if tss == 2120
d dec -366 if pf >= 2210
cie inc -787 if o < -2302
erb dec -305 if brr != -2374
rfw dec -915 if pf > 2216
bs inc -516 if hsn >= -1947
rxb inc 845 if tss != 2120
erb dec -201 if qn == 1888
h inc 297 if tss < 2129
h dec -860 if s > 1541
isy dec -59 if hsn > -1949
mxg inc -266 if wf > -2316
vso dec -578 if d != -2703
qn dec 420 if o <= -2299
ie dec 307 if wf == -2319
pf dec -948 if bs >= -929
j dec 927 if h == 1416
vso dec 879 if d > -2717
j dec -543 if uxr != 4374
d dec -20 if o != -2308
brr inc 606 if erb > 1340
hsn dec 547 if rak != -244
cie dec -865 if uxr < 4366
bs inc -408 if o <= -2294
wf inc 938 if rfw < 207
pf inc 633 if j != 1289
rak dec -871 if d <= -2687
hsn inc 311 if tss < 2125
wf dec -231 if qn <= 1477
mxg dec 617 if rxb < 1137
hsn inc 20 if ie > -3315
uxr inc -810 if hsn < -2154
d dec -738 if mxg > 798
o inc 513 if d < -1960
s dec 953 if mxg <= 812
s inc 664 if cie < 775
h dec 18 if bok >= -480
erb dec 43 if rak > 631
vso inc 662 if rak > 632
mxg inc -938 if wf >= -2083
brr inc -359 if rxb > 1136
s dec -402 if rxb < 1144
h dec 59 if x < 48
mxg inc 842 if bs == -1331
ie inc -748 if vso >= 2298
mxg inc -176 if erb > 1301
s inc -345 if bs == -1335
cie inc 52 if brr < -2125
j inc -243 if wf > -2095
d dec 992 if qn != 1468
h inc 302 if tss == 2120
rfw dec 431 if hsn >= -2168
hsn inc -930 if wf > -2098
x dec -477 if cie != 767
erb dec 922 if pf <= 3797
j dec 838 if cie <= 782
mxg inc -96 if bs < -1324
cie dec 707 if hsn > -3099
mxg inc 258 if h < 1655
hsn dec -441 if qn >= 1461
j dec 518 if tss >= 2115
tss dec -760 if ie == -3306
ie dec 21 if d != -1948
j dec -990 if s > 997
bs inc -889 if bok >= -482
s dec 732 if hsn == -2646
pf dec -401 if hsn <= -2639
rak inc 211 if isy != 3128
brr inc -868 if uxr == 3554
vso dec 811 if uxr <= 3563
pf dec 232 if rxb >= 1137
bok inc -193 if vso < 1490
d dec 817 if qn == 1468
rak inc 282 if bs < -2216
o inc 74 if pf < 3972
isy dec -602 if rxb <= 1148
pf inc -6 if cie <= 69
uxr inc 835 if rfw >= -228
tss inc 790 if bok <= -661
rak dec -719 if bok <= -662
cie inc 978 if j > -298
d dec 475 if brr > -2994
rxb dec -348 if pf >= 3976
o dec 62 if j < -301
brr dec 486 if bok > -675
isy dec -256 if hsn == -2648
rxb dec -536 if tss < 3671
ie dec 91 if vso < 1490
h inc 53 if h >= 1644
qn inc 621 if d > -3247
hsn dec -433 if x <= 523
x dec 703 if tss >= 3673
x inc -557 if isy >= 3978
bs inc -180 if h < 1709
pf inc -25 if hsn <= -2207
bs dec 761 if o > -2294
uxr dec 989 if brr > -3482
erb dec 385 if h > 1703
hsn dec 339 if bok < -664
isy inc -176 if s > 989
hsn inc 405 if isy == 3810
pf dec -849 if uxr < 3402
bok inc 274 if rak > 1631
j inc 433 if tss >= 3668
ie dec 385 if o <= -2290
rak inc 529 if tss >= 3670
isy dec -467 if o < -2285
erb inc 765 if erb >= 7
vso dec -935 if wf >= -2084
tss inc 21 if brr != -3487
s dec 631 if erb > 5
hsn inc 709 if h < 1709
isy inc 754 if tss != 3691
rak dec -730 if rfw <= -216
erb inc 33 if bs <= -3156
hsn inc 972 if rxb <= 1681
d inc 902 if mxg <= 1634
cie inc -544 if d == -2343
wf dec -706 if bs <= -3160
j dec 564 if isy < 4282
ie inc -832 if o > -2298
rfw inc 737 if pf >= 4795
o inc 558 if o <= -2285
erb dec 749 if bs > -3163
pf dec -93 if h > 1700
o dec 85 if isy > 4271
brr inc -713 if x <= -45
x dec 659 if d < -2344
tss dec -316 if mxg < 1638
pf inc -542 if bok < -388
tss dec 572 if h <= 1708
s inc -19 if cie != -464
j dec -262 if hsn <= -463
rfw dec -775 if rxb > 1677
s inc 564 if ie > -4639
o inc -484 if brr == -3479
uxr dec -651 if hsn <= -463
wf dec 866 if qn == 2094
ie inc 635 if pf < 4339
cie inc 248 if pf == 4341
rxb dec 233 if o > -2299
bok inc 914 if cie != -232
hsn dec -748 if x == -47
cie dec 322 if uxr <= 4046
mxg inc 79 if brr <= -3477
tss dec 939 if x == -42
pf dec 773 if s < 1551
bok inc -145 if mxg >= 1712
uxr inc -237 if tss > 2494
d dec 187 if mxg >= 1705
ie inc -280 if o > -2305
rak inc -757 if isy > 4273
s dec -922 if isy != 4278
pf dec 868 if qn <= 2089
pf dec 838 if j >= -174
ie inc 293 if rfw > 554
bok dec -846 if o >= -2303
o dec -869 if qn == 2089
rak inc 960 if cie <= -224
vso dec -192 if o != -1443
o dec -477 if rxb != 1675
bs dec 256 if uxr > 3813
brr dec 224 if bok > 1360
qn dec -294 if erb <= -726
h dec 790 if ie >= -4628
erb inc -964 if rak != 3104
rfw dec 417 if s >= 2462
wf inc -396 if pf == 1862
wf dec 765 if s > 2457
bs inc -213 if mxg < 1711
rfw dec 18 if o > -963
isy dec 99 if uxr != 3808
uxr dec -548 if mxg > 1707
hsn dec 404 if brr == -3703
cie dec -420 if x <= -49
erb dec -623 if rxb < 1682
h dec 519 if rxb > 1682
o dec -58 if rak < 3106
vso inc 459 if rxb == 1679
hsn dec 843 if rfw < 126
mxg inc 820 if pf < 1864
hsn inc 60 if ie >= -4629
brr dec 25 if rxb == 1679
uxr dec 148 if wf > -2545
s dec 416 if o == -897
hsn inc -427 if ie != -4625
s inc -925 if erb >= -1066
tss dec -413 if d != -2530
j dec 346 if tss <= 2500
uxr inc -160 if wf != -2542
wf inc -718 if bs == -3630
j inc -518 if brr != -3728
hsn dec -489 if bok <= 1372
bok dec -28 if tss == 2496
bok dec -106 if x != -34
d dec 905 if erb == -1060
bok dec -303 if s <= 1529
vso inc 650 if uxr > 4053
cie inc -209 if ie < -4627
erb dec 272 if d <= -3435
h dec 100 if wf != -3253
ie inc -783 if d != -3425
wf dec -712 if vso == 2785
pf dec 133 if ie > -5403
brr dec 218 if bs < -3625
qn dec -420 if rfw < 127
bok dec -633 if s < 1546
cie inc -932 if ie > -5412
wf dec 515 if pf > 1857
mxg inc 454 if ie <= -5411
x inc -77 if hsn == -1593
o inc 746 if brr <= -3950
x dec 872 if erb >= -1332
vso inc 268 if h != 815
h dec -436 if bs != -3624
erb dec -543 if bok < 2141
ie dec -375 if tss == 2496
x dec -176 if rxb > 1672
pf dec 926 if rfw <= 124
bs inc -134 if pf <= 938
isy dec 664 if rxb != 1684
s inc 584 if isy >= 3523
vso dec 913 if rak >= 3098
bs inc 745 if wf == -3064
rfw inc 309 if qn == 2515
uxr dec -496 if isy <= 3515
hsn dec -633 if isy <= 3514
brr dec 502 if j <= -521
brr inc -521 if j != -519
tss dec 302 if o != -891
hsn inc 340 if ie != -5030
cie dec 959 if j > -522
wf inc 417 if erb <= -781
isy dec -376 if d <= -3444
h inc -131 if cie != -2112
bs inc -510 if vso < 3055
x inc -780 if rxb > 1676
ie dec 702 if tss >= 2187
rfw inc 590 if h <= 1124
wf dec -967 if hsn > -969
vso dec -104 if s > 1536
hsn inc 197 if brr <= -4468
qn dec -309 if ie == -5730
x inc 694 if brr > -4472
isy dec -704 if vso > 3164
j inc -127 if x >= -904
pf inc -579 if uxr != 4550
bok inc -666 if rak > 3092
vso dec -926 if hsn <= -969
brr dec 71 if brr > -4467
s dec 10 if o > -904
j inc -338 if x == -901
cie dec -630 if x >= -910
cie dec -716 if ie < -5737
wf inc 613 if wf <= -1678
s dec -409 if bs <= -3539
rfw inc -505 if d < -3432
bok inc 912 if bs == -3529
isy inc -992 if mxg > 2525
bok dec 100 if tss <= 2196
s inc 712 if s > 1520
cie dec -381 if brr != -4465
brr dec 793 if cie > -1099
bok dec -387 if bok < 2283
s dec 812 if mxg > 2520
hsn dec 422 if isy >= 2519
brr dec -597 if wf <= -1063
tss inc -539 if h == 1119
d dec 159 if hsn == -1382
h dec 474 if wf <= -1067
h dec -34 if vso < 3161
rxb inc 714 if pf != 936
wf dec -564 if pf > 926
d inc 534 if erb >= -781
h inc 808 if j == -982
brr inc 976 if qn == 2509
o dec 623 if rfw > 201
hsn dec -794 if h < 1492
uxr inc 649 if mxg <= 2533
cie inc 156 if rak == 3096
j dec -392 if hsn == -588
tss inc -161 if rfw != 213
rfw dec 864 if qn > 2512
d inc -499 if cie == -950
rxb inc -849 if rfw > 203
x inc -478 if pf == 939
j inc -343 if uxr > 5194
qn inc 504 if d < -4087
pf inc -430 if s > 1423
o inc -709 if x <= -911
mxg inc 594 if rxb != 824
h dec -885 if h == 1487
brr inc -663 if bok <= 2657
tss inc 504 if bs != -3532
brr inc -100 if s < 1432
bs inc 133 if rfw == 206
hsn inc 535 if rfw >= 200
hsn inc -82 if h != 2372
mxg dec 254 if rak != 3090
wf inc -401 if o <= -1513
d dec 577 if qn < 3019
qn inc 94 if bok > 2660
pf dec -990 if ie <= -5732
bok dec 181 if cie < -945
qn dec 431 if bok <= 2489
mxg inc 128 if qn >= 2673
d dec -233 if tss > 1996
rak dec 364 if ie != -5732
tss inc 174 if o == -1521
tss dec -463 if isy != 2530
rak inc -238 if bs >= -3403
brr dec 447 if j == -933
ie dec 853 if cie > -952
bs dec 350 if bs < -3393
brr dec -996 if d <= -4434
cie dec -749 if tss < 2634
mxg inc -484 if bok == 2484
vso dec -135 if x == -899
d inc -426 if rxb < 832
h dec -40 if uxr != 5199
rak inc -480 if brr == -2445
h inc 818 if s == 1428
hsn dec -708 if hsn < -51
rfw inc -155 if rxb != 838
pf inc -3 if o < -1511
qn dec -144 if ie <= -6579
hsn dec 166 if s > 1418
rfw inc 216 if bok == 2484
h inc 65 if isy <= 2523
uxr inc -405 if ie <= -6592
h inc -85 if wf <= -896
rfw dec 104 if s == 1428
ie dec -144 if vso != 3157
ie dec -877 if rfw < 156
uxr dec 836 if isy >= 2514
mxg inc -753 if uxr >= 4358
j inc -695 if x >= -901
rak dec -399 if wf >= -910
d dec 289 if vso != 3157
hsn inc 482 if j == -1628
rfw dec 10 if bs >= -3741
vso inc 803 if d != -4864
vso inc -335 if qn == 2820
o dec 527 if rak != 2776
tss dec 799 if rfw > 170
j dec -763 if hsn > 967
tss dec 259 if pf <= 1498
wf inc -260 if rxb <= 820
x inc -560 if hsn < 980
j inc 126 if wf <= -898
uxr inc -665 if j != -736
vso inc -59 if h >= 3170
j dec 88 if isy < 2528
brr inc 622 if uxr < 3692
wf dec -81 if isy >= 2521
rfw inc 497 if tss > 2370
d inc 411 if erb <= -789
cie inc -375 if mxg == 1761
bok inc -998 if hsn <= 972
pf inc 224 if mxg != 1757
x inc 551 if rxb <= 835
qn inc -388 if brr != -2442
rfw inc -79 if rak < 2779
hsn dec 756 if j > -831
wf dec -67 if ie > -6589
hsn dec -175 if tss < 2381
tss dec -861 if rak == 2777
o dec -461 if rak >= 2772
rxb dec 916 if erb <= -795
pf dec 379 if erb == -789
isy inc -155 if bok == 1486
erb dec 790 if o == -1587
brr inc 445 if tss < 3242
rxb inc 884 if hsn > 394
tss dec 591 if x == -913
j dec 298 if cie == -1325
ie dec -68 if mxg <= 1764
mxg dec 387 if vso < 3570
tss dec -86 if bs != -3748
rxb inc 238 if hsn >= 390
rak inc -28 if bok == 1486
rak inc -640 if mxg == 1374
qn dec -83 if tss >= 3314
tss dec -238 if bok != 1476
tss dec -101 if o != -1587
hsn dec -537 if tss == 3561
wf inc -859 if brr == -2000
tss inc -35 if uxr >= 3698
cie inc -970 if rak <= 2112
rfw inc 971 if h < 3172
wf dec 887 if rfw < 1555
bok inc 20 if o <= -1592
mxg inc -326 if cie > -2298
mxg inc -684 if j > -1116
brr inc 639 if qn >= 2510
rxb dec -332 if rxb < 1069
uxr dec 189 if brr > -1369
rak dec -567 if rfw == 1552
cie dec 407 if uxr > 3505
h dec -365 if o < -1585
cie dec -990 if s == 1428
rfw dec 368 if ie == -6517
vso dec -76 if wf != -2502
erb inc 334 if bs >= -3747
o inc 583 if qn == 2515
pf inc 7 if rxb > 1399
o inc 166 if qn >= 2512
wf inc 492 if x <= -905
uxr inc -566 if d == -4452
bok inc -919 if isy >= 2364
hsn inc 517 if rak != 2677
j dec -916 if d != -4454
bok dec 116 if rak > 2674
vso dec -477 if x < -907
qn dec -964 if erb == -1245
j dec 987 if qn > 3478
rak dec 317 if d < -4442
isy dec 1 if cie > -1717
j dec 95 if mxg == 1046
h dec -927 if s <= 1435
o dec 812 if ie >= -6518
brr inc -499 if uxr >= 2936
rak dec 716 if cie == -1712
x inc 482 if pf == 1350
cie inc -272 if mxg != 1049
isy inc -839 if brr == -1870
o inc -259 if erb <= -1239
j dec 859 if x != -910
isy inc 493 if tss != 3521
rxb dec -673 if mxg >= 1043
vso dec 35 if uxr <= 2947
rxb dec -128 if bok < 459
bok inc 645 if isy >= 2853
bok inc 384 if s < 1430
h inc -709 if brr != -1865
o inc 638 if cie != -1980
rak dec 9 if pf != 1345
erb dec -772 if rfw != 1176
ie inc -817 if rfw == 1184
vso dec 21 if ie == -7334
s inc -829 if uxr > 2936
isy inc -205 if cie <= -1984
mxg dec -404 if o == -1271
x inc 633 if mxg < 1462
vso inc 795 if bok <= 1481
cie inc 613 if h == 3753
s inc 156 if qn != 3478
wf inc 908 if erb > -467
x inc -340 if j < -1191
bok dec -790 if uxr > 2940
d inc 588 if vso >= 4774
ie dec 711 if rxb < 2209
j dec -687 if isy >= 2657
brr inc -979 if bok == 2270
uxr dec -130 if isy < 2649
h inc -403 if pf == 1341
pf dec -156 if ie == -8045
o dec 828 if bok <= 2276
bs inc -379 if d != -3864
vso inc 570 if isy != 2654
brr dec -208 if brr > -2837
d dec 364 if wf != -2005
pf inc -459 if d <= -4219
wf dec 449 if tss != 3536
s dec -477 if wf >= -2464
j dec 366 if vso == 4782
s dec 240 if rxb != 2201
hsn dec -892 if h > 3752
ie dec -551 if o <= -2098
wf inc -574 if tss < 3531
s dec -972 if isy >= 2664
isy dec -632 if pf <= 1042
uxr inc -733 if vso > 4779
bok inc -541 if o == -2099
pf inc -723 if cie >= -1378
brr inc 993 if isy != 3285
isy inc -859 if brr == -1851
qn inc -925 if rfw != 1184
brr dec 118 if j != -1572
rak inc -711 if hsn <= 2343
o dec 896 if rfw <= 1184
qn inc -332 if h <= 3759
rak dec -684 if mxg == 1452
qn dec 945 if uxr <= 2216
d inc 971 if rfw < 1189
bok dec -80 if o != -2995
hsn dec 288 if uxr >= 2219
x inc -986 if bok > 1725
o inc -640 if tss > 3527
rak dec -210 if isy == 3287
qn dec -937 if j > -1570
d dec -796 if ie >= -7497
rfw dec 869 if d != -2456
j dec 941 if erb == -473
vso inc -366 if j > -2509
h inc -526 if d != -2461`
	out = advent8A(input)
	fmt.Printf("Result 8A %d\n", out)
	out = advent8B(input)
	fmt.Printf("Result 8B %d\n", out)
}

func TestDay9(t *testing.T) {
	input := "{}"
	out := advent9A(input)
	expected := 1
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{{}}}"
	out = advent9A(input)
	expected = 6
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{},{}}"
	out = advent9A(input)
	expected = 5
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{{},{},{{}}}}"
	out = advent9A(input)
	expected = 16
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{<a>,<a>,<a>,<a>}"
	out = advent9A(input)
	expected = 1
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{<ab>},{<ab>},{<ab>},{<ab>}}"
	out = advent9A(input)
	expected = 9
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{<!!>},{<!!>},{<!!>},{<!!>}}"
	out = advent9A(input)
	expected = 9
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = "{{<a!>},{<a!>},{<a!>},{<ab>}}"
	out = advent9A(input)
	expected = 3
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}

	// input = ""
	// out = advent9A(input)
	// fmt.Printf("Result 9A %d\n", out)
	// out = advent9B(input)
	// fmt.Printf("Result 9B %d\n", out)
}

func TestAdvent(t *testing.T) {
}
