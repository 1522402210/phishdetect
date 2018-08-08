package brand

// NYTimes brand properties.
func NYTimes() *Brand {
	name := "nytimes"
	original := []string{"nytimes", "newyorktimes", "nytco"}
	whitelist := []string{
		"nytimes.com", "newyorktimes.com", "nytco.com", "nyt.net",
	}
	suspicious := []string{
		"nytimes", "nytimesa", "nytimesb", "nytimesc", "nytimesd", "nytimese",
		"nytimesf", "nytimesg", "nytimesh", "nytimesi", "nytimesj", "nytimesk",
		"nytimesl", "nytimesm", "nytimesn", "nytimeso", "nytimesp", "nytimesq",
		"nytimesr", "nytimess", "nytimest", "nytimesu", "nytimesv", "nytimesw",
		"nytimesx", "nytimesy", "nytimesz", "oytimes", "lytimes", "jytimes",
		"fytimes", "nqtimes", "nitimes", "n9times", "nyuimes", "nyvimes",
		"nypimes", "nydimes", "ny4imes", "nythmes", "nytkmes", "nytmmes",
		"nytames", "nytymes", "nytiles", "nytioes", "nytiies", "nytiees",
		"nyti-es", "nytimds", "nytimgs", "nytimas", "nytimms", "nytimus",
		"nytimer", "nytimeq", "nytimew", "nytimec", "nytime3", "xn--nytime-8ib",
		"xn--ntimes-wrf", "xn--nyties-s3c", "nyt1mes", "xn--nyimes-qrf",
		"xn--nyties-6s7b", "xn--nytime-8k2a", "xn--ntimes-bza",
		"xn--nyimes-qkb", "xn--ntimes-3xe", "nytlmes", "nytines",
		"xn--nytmes-5va", "xn--nytmes-k6b", "xn--nytims-fhg", "xn--nytims-tva",
		"xn--nytims-04a", "xn--nytims-7of", "xn--nytime-86c", "xn--nytmes-k91a",
		"xn--ytimes-heb", "xn--nytime-nj9d", "xn--nytims-f5a", "xn--nytims-fva",
		"xn--nytmes-51c", "xn--nytims-t3a", "xn--nytime-nvf", "xn--nytmes-5r6v",
		"xn--nytmes-dze", "xn--nytims-fvf", "rytimes", "xn--ntimes-ieg",
		"xn--nytmes-kwa", "nytirnes", "nytirres", "xn--nytmes-y8a",
		"xn--nyties-lqf", "xn--nytims-th8b", "xn--nytims-m4a",
		"xn--nyties-605b", "xn--nytims-mva", "xn--ntimes-p9c", "xn--nytmes-r9a",
		"nytinnes", "xn--nytims-mye", "xn--nyimes-j1e", "xn--nytims-73a",
		"n-ytimes", "ny-times", "nyt-imes", "nyti-mes", "nytim-es", "nytime-s",
		"nyytimes", "nytuimes", "nytimpes", "nytoimes", "nyftimes", "nxytimes",
		"nystimes", "nyttimes", "nyhtimes", "nytiomes", "nytimjes", "nytimews",
		"ny6times", "nytiumes", "nyztimes", "ny5times", "ngytimes", "nytyimes",
		"nyrtimes", "nytimkes", "nytimnes", "nytkimes", "nyti8mes", "nyt9imes",
		"nytimles", "nytime3s", "nsytimes", "nytimses", "nytim3es", "nytim4es",
		"nyt8imes", "ntytimes", "nytijmes", "nytimwes", "nygtimes", "nyt5imes",
		"nyxtimes", "nytinmes", "nytjimes", "nytfimes", "n7ytimes", "nytrimes",
		"nytikmes", "nytimdes", "nytimeds", "nyatimes", "nyutimes", "nytimezs",
		"nyti9mes", "nytimzes", "nhytimes", "nytimers", "nytimres", "n6ytimes",
		"nytgimes", "nytipmes", "ny7times", "nytilmes", "nuytimes", "nytime4s",
		"naytimes", "nytzimes", "nyt6imes", "nytims", "nyimes", "nytmes",
		"nytime", "ytimes", "nyties", "ntimes", "nytiimes", "nytimees",
		"nnytimes", "nytimmes", "ngtimes", "nytimea", "nutimes", "nyt8mes",
		"nyzimes", "nytipes", "nytimzs", "nytimrs", "nytimss", "nytim3s",
		"bytimes", "natimes", "nttimes", "nytumes", "n7times", "nyrimes",
		"n6times", "ny6imes", "nytimez", "hytimes", "nytimex", "nytimee",
		"nytimed", "nyt9mes", "nytimey", "nstimes", "nytomes", "nytjmes",
		"nyyimes", "ny5imes", "nytikes", "nhtimes", "nytimws", "nyfimes",
		"nytim4s", "nytijes", "nygimes", "n.ytimes", "ny.times", "nyt.imes",
		"nyti.mes", "nytim.es", "nytime.s", "yntimes", "ntyimes", "nyitmes",
		"nytmies", "nytiems", "nytimse", "nytimis", "nytemes", "nytimos",
		"nytime",
	}

	suspicious = append(suspicious, []string{
		"newyorktimesa", "newyorktimesb", "newyorktimesc", "newyorktimesd",
		"newyorktimese", "newyorktimesf", "newyorktimesg", "newyorktimesh",
		"newyorktimesi", "newyorktimesj", "newyorktimesk", "newyorktimesl",
		"newyorktimesm", "newyorktimesn", "newyorktimeso", "newyorktimesp",
		"newyorktimesq", "newyorktimesr", "newyorktimess", "newyorktimest",
		"newyorktimesu", "newyorktimesv", "newyorktimesw", "newyorktimesx",
		"newyorktimesy", "newyorktimesz", "oewyorktimes", "lewyorktimes",
		"jewyorktimes", "fewyorktimes", "ndwyorktimes", "ngwyorktimes",
		"nawyorktimes", "nmwyorktimes", "nuwyorktimes", "nevyorktimes",
		"neuyorktimes", "nesyorktimes", "negyorktimes", "ne7yorktimes",
		"newxorktimes", "newqorktimes", "newiorktimes", "new9orktimes",
		"newynrktimes", "newymrktimes", "newykrktimes", "newygrktimes",
		"newyosktimes", "newyopktimes", "newyovktimes", "newyozktimes",
		"newyobktimes", "newyo2ktimes", "newyorjtimes", "newyoritimes",
		"newyorotimes", "newyorctimes", "newyorkuimes", "newyorkvimes",
		"newyorkpimes", "newyorkdimes", "newyork4imes", "newyorkthmes",
		"newyorktkmes", "newyorktmmes", "newyorktames", "newyorktymes",
		"newyorktiles", "newyorktioes", "newyorktiies", "newyorktiees",
		"newyorkti-es", "newyorktimds", "newyorktimgs", "newyorktimas",
		"newyorktimms", "newyorktimus", "newyorktimer", "newyorktimeq",
		"newyorktimew", "newyorktimec", "newyorktime3", "xn--newyrktimes-unj",
		"xn--newyortimes-ndi", "xn--nwyorktims-rmbi", "xn--nwyorktims-b2ii",
		"nevvyorktimes", "xn--newyorktims-nqb", "xn--neyorktimes-8tl",
		"xn--neworktimes-jpj", "xn--nwyorktims-enbi", "rewyorktimes",
		"xn--nwyorktims-ilbi", "xn--neworktimes-59h", "xn--nwyorktimes-gsb",
		"newyorktines", "xn--newyrktimes-m99e", "xn--newyorktime-xvj",
		"xn--nwyorktimes-rrb", "xn--nwyorktims-0nbi", "xn--newyoktimes-x35e",
		"xn--nwyorktims-x7ai", "xn--newyorktmes-51e", "xn--newyorkties-bnj",
		"xn--newyorktime-tbf", "xn--newyorktims-xkj", "xn--newyorktmes-vcb",
		"xn--newyorktmes-pkd", "xn--newyorktmes-u91c", "xn--nwyorktimes-ppb",
		"newyorlctimes", "xn--neworktimes-7hb", "xn--newyoktimes-3jj",
		"xn--newyorktims-kvj", "xn--neyorktimes-kim", "xn--newyorkimes-bic",
		"xn--neworktimes-1ff", "xn--newyrktimes-5he", "xn--newyrktimes-sbl",
		"xn--newyorktime-d55i", "xn--ewyorktimes-h6b", "xn--newyorktims-1bi",
		"xn--newyoktimes-l8e", "xn--newyorkimes-bpj", "xn--newyrktimes-x89e",
		"xn--newyorkties-hv5e", "xn--nwyorktimes-2qb", "xn--newyorktmes-tyb",
		"xn--newyorkties-x4e", "xn--newyoktimes-x8e", "xn--nwyorktimes-okj",
		"xn--newyorktims-kbb", "xn--newyorktims-kwk", "xn--nwyorktims-cnhi",
		"xn--newyoktimes-y9e", "xn--nwyorktimes-z19e", "newyorktlmes",
		"xn--nwyorktimes-bvj", "xn--newyorktmes-6zb", "xn--neworktimes-krk",
		"newyorktirnes", "xn--newyorktmes-kdb", "newyorktirres",
		"xn--newyortimes-lt4i", "xn--newyorktims-0rb", "xn--nwyorktimes-zbb",
		"newy0rktimes", "xn--nwyorktimes-bwk", "xn--newyorktims-wbb",
		"xn--newyorkties-rw8e", "xn--nwyorktimes-sbi", "mewyorktimes",
		"xn--newyorktmes-ddi", "newyorlktimes", "xn--nwyorktims-b7ai",
		"xn--newyrktimes-h3c", "xn--newyorktime-ofc", "xn--newyorktims-8bb",
		"xn--newyorkimes-3gi", "xn--nwyorktims-xq3ei", "xn--newyrktimes-bfi",
		"xn--newyorktmes-0z92b", "xn--newyrktimes-ufb", "xn--newyrktimes-yum",
		"newyoriktimes", "xn--newyorktims-psb", "xn--nwyorktims-b0ji",
		"xn--newyrktimes-unj", "xn--neyorktimes-szj", "xn--newyorktims-ypb",
		"newyorktinnes", "xn--newyorktims-crb", "xn--nwyorktimes-nbb",
		"xn--nwyorktimes-eqb", "xn--nwyorktimes-bbb", "xn--newyrktimes-teb",
		"xn--nwyorktims-4lbi", "xn--newyorktime-3s2c", "xn--newyrktimes-bfi",
		"xn--nwyorktims-m7ai", "xn--nwyorktims-ksii", "xn--newyorktims-819e",
		"newyorkt1mes", "n-ewyorktimes", "ne-wyorktimes", "new-yorktimes",
		"newy-orktimes", "newyo-rktimes", "newyor-ktimes", "newyork-times",
		"newyorkt-imes", "newyorkti-mes", "newyorktim-es", "newyorktime-s",
		"newyorktime3s", "newyorktipmes", "newyorkjtimes", "n3ewyorktimes",
		"newtyorktimes", "newyorktijmes", "newyo0rktimes", "newyorkrtimes",
		"newyorktgimes", "newyorktzimes", "newyoerktimes", "newyor4ktimes",
		"newyorkftimes", "newy7orktimes", "newyorktjimes", "newykorktimes",
		"nzewyorktimes", "neqwyorktimes", "newyorktimews", "newyotrktimes",
		"newyoirktimes", "ne3wyorktimes", "neawyorktimes", "newyofrktimes",
		"newxyorktimes", "newyo5rktimes", "newyorkti8mes", "newyaorktimes",
		"newhyorktimes", "nwewyorktimes", "newyorktimdes", "newyorktimers",
		"newyo9rktimes", "newwyorktimes", "nexwyorktimes", "newqyorktimes",
		"newyorktimezs", "newyodrktimes", "nrewyorktimes", "neweyorktimes",
		"newyporktimes", "neewyorktimes", "ne4wyorktimes", "newy6orktimes",
		"newyormktimes", "newyokrktimes", "newyorktimpes", "newyorkt6imes",
		"nedwyorktimes", "newy9orktimes", "newsyorktimes", "newyorktimzes",
		"newyorktimkes", "newyorktime4s", "new6yorktimes", "newyorkytimes",
		"newyorktim4es", "newyork6times", "newyorktimles", "n4ewyorktimes",
		"newyorjktimes", "newyorkotimes", "newyorfktimes", "newyorktilmes",
		"newyorktimwes", "newyiorktimes", "neswyorktimes", "newyorkmtimes",
		"newyorktim3es", "newyorkt5imes", "newyorktoimes", "newyorktiumes",
		"newylorktimes", "newyorktimeds", "newyorkt8imes", "newyorktikmes",
		"newyorktrimes", "newysorktimes", "newgyorktimes", "nezwyorktimes",
		"newytorktimes", "newyork5times", "newyxorktimes", "newyorkti9mes",
		"newyorktfimes", "newyorktiomes", "newyhorktimes", "newyorktyimes",
		"new3yorktimes", "newyuorktimes", "newyorktinmes", "newyo4rktimes",
		"newygorktimes", "nerwyorktimes", "newyordktimes", "newyorkgtimes",
		"newyorktimres", "new7yorktimes", "ndewyorktimes", "newyorkitimes",
		"newyorkztimes", "newyorkt9imes", "newyorkltimes", "newyortktimes",
		"nsewyorktimes", "newyorektimes", "newuyorktimes", "newyorktimses",
		"newyoprktimes", "newyor5ktimes", "ne2wyorktimes", "newyorktimnes",
		"newy0orktimes", "newayorktimes", "newyolrktimes", "newyorktkimes",
		"new2yorktimes", "newyorktimjes", "newyorktuimes", "newyoroktimes",
		"newyorktime", "newyorkimes", "nwyorktimes", "neyorktimes",
		"newyorkties", "newyrktimes", "newyoktimes", "newyorktims",
		"neworktimes", "ewyorktimes", "newyorktmes", "newyortimes",
		"nnewyorktimes", "newyyorktimes", "newyorktimees", "newyorrktimes",
		"newyorkktimes", "newyoorktimes", "newyorktiimes", "newyorkttimes",
		"newyorktimmes", "newsorktimes", "newyorktimrs", "newyork6imes",
		"newyorktim3s", "newyorktimex", "neeyorktimes", "newyorktikes",
		"neqyorktimes", "newhorktimes", "newyorkt8mes", "newyorltimes",
		"nrwyorktimes", "newyorkrimes", "newyprktimes", "newyofktimes",
		"neayorktimes", "newyorktomes", "newyorktjmes", "newyorkzimes",
		"newyo5ktimes", "newy9rktimes", "n3wyorktimes", "nwwyorktimes",
		"newylrktimes", "ne2yorktimes", "newyorktimey", "newyorktumes",
		"newyorktimws", "newyorktim4s", "newyodktimes", "newyorkt9mes",
		"newyorkyimes", "newyorkgimes", "newyorktimss", "newyorktipes",
		"newyormtimes", "newuorktimes", "newyo4ktimes", "newyotktimes",
		"new7orktimes", "newyorktimed", "n4wyorktimes", "newaorktimes",
		"newtorktimes", "newyorktimea", "newyorkfimes", "new6orktimes",
		"newgorktimes", "bewyorktimes", "newyirktimes", "ne3yorktimes",
		"newyork5imes", "newyorktimzs", "newyorktimee", "nzwyorktimes",
		"nswyorktimes", "nexyorktimes", "newyoektimes", "newyorktimez",
		"newyorktijes", "hewyorktimes", "n.ewyorktimes", "ne.wyorktimes",
		"new.yorktimes", "newy.orktimes", "newyo.rktimes", "newyor.ktimes",
		"newyork.times", "newyorkt.imes", "newyorkti.mes", "newyorktim.es",
		"newyorktime.s", "enwyorktimes", "nweyorktimes", "neyworktimes",
		"newoyrktimes", "newyroktimes", "newyokrtimes", "newyortkimes",
		"newyorkitmes", "newyorktmies", "newyorktiems", "newyorktimse",
		"newyorktimos", "niwyorktimes", "newyurktimes", "newyarktimes",
		"newyorktimis", "newyerktimes", "newyorktemes", "nowyorktimes",
		"newyorktime",
	}...)

	return &Brand{
		Name:       name,
		Original:   original,
		Whitelist:  whitelist,
		Suspicious: suspicious,
	}
}
