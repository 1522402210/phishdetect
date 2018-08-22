// PhishDetect
// Copyright (C) 2018  Claudio Guarnieri
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package brand

// Linkedin brand properties.
func Linkedin() *Brand {
	name := "linkedin"
	original := []string{"linkedin"}
	whitelist := []string{"linkedin.com", "linked.in"}
	suspicious := []string{
		"linkedina", "linkedinb", "linkedinc", "linkedind", "linkedine",
		"linkedinf", "linkeding", "linkedinh", "linkedini", "linkedinj",
		"linkedink", "linkedinl", "linkedinm", "linkedinn", "linkedino",
		"linkedinp", "linkedinq", "linkedinr", "linkedins", "linkedint",
		"linkedin", "linkedinv", "linkedinw", "linkedinx", "linkediny",
		"linkedinz", "minkedin", "ninkedin", "hinkedin", "dinkedin", "lhnkedin",
		"lknkedin", "lmnkedin", "lankedin", "lynkedin", "liokedin", "lilkedin",
		"lijkedin", "lifkedin", "linjedin", "liniedin", "linoedin", "lincedin",
		"linkddin", "linkgdin", "linkadin", "linkmdin", "linkudin", "linkeein",
		"linkefin", "linkelin", "linketin", "linkedhn", "linkedkn", "linkedmn",
		"linkedan", "linkedyn", "linkedio", "linkedil", "linkedij", "linkedif",
		"xn--lnkedin-nf1z", "xn--linkdin-uya", "xn--lnkedin-rfb", "llnkedln",
		"xn--lnkedin-ueb", "xn--inkedin-mjb", "xn--linkedn-0nf",
		"xn--lnkedin-2gd", "xn--linkedn-dza", "xn--linkein-2cd",
		"xn--lnkedn-i91ae", "xn--linkdin-us4c", "linkebin", "xn--lnkedin-vnf",
		"llnkedin", "xn--linkedi-9jb", "linkediin", "l1nkedin", "l1nked1n",
		"xn--linedin-5nf", "xn--linkdin-w8a", "xn--linkedn-tza", "lirkedin",
		"xn--lnkedin-j95a", "xn--lnkedn-i6be", "xn--linkdin-1mf",
		"xn--lnkedn-31ce", "xn--linkedn-sf1z", "xn--linkdin-d9a",
		"xn--linkein-ysh", "xn--lnkedn-w8ae", "xn--linkdin-eya",
		"xn--linkdin-eog", "xn--lnkedn-iwae", "lirkedir", "xn--linkdin-edh",
		"xn--lnkedin-7ya", "xn--linkdin-z7a", "xn--linkein-k7a",
		"xn--linedin-sy7e", "linlkedin", "xn--linkedn-zeb", "xn--linkdin-mya",
		"xn--linkedn-7gd", "xn--linkedn-wfb", "xn--lnkedn-3r6vea",
		"xn--likedi-jebe", "xn--linkein-yhi", "linikedin", "xn--lnkedin-3gc",
		"iinkedin", "xn--linkdin-bhg", "xn--inkedin-ihd", "limkedim",
		"xn--lnkedn-p9ae", "limkedin", "linkeclin", "xn--linkedn-8gc",
		"linkedln", "xn--lnkedn-bzee", "xn--lnkedn-3vae", "xn--likedin-4jb",
		"linkedir", "linked1n", "xn--linkdin-t9a", "1inkedin",
		"xn--lnkedin-oza", "linlcedin", "xn--linkedn-o95a", "linkedim",
		"linkedlin", "xn--linkdin-g8a", "l-inkedin", "li-nkedin", "lin-kedin",
		"link-edin", "linke-din", "linked-in", "linkedi-n", "linkedfin",
		"linkeddin", "linkedxin", "linkerdin", "li9nkedin", "linkdedin",
		"lihnkedin", "linkiedin", "lijnkedin", "linkedein", "linked9in",
		"linkedi8n", "linkedcin", "linke3din", "linhkedin", "loinkedin",
		"linbkedin", "linkediun", "liknkedin", "limnkedin", "lkinkedin",
		"linkedkin", "linkedikn", "li8nkedin", "linkjedin", "linkedion",
		"lionkedin", "linkwedin", "linkezdin", "link3edin", "l8inkedin",
		"linkoedin", "linkledin", "linkedijn", "linokedin", "linkexdin",
		"linkedsin", "linkefdin", "linkecdin", "linkredin", "linkedi9n",
		"linke4din", "linkmedin", "linkesdin", "linmkedin", "ljinkedin",
		"linkeedin", "linkzedin", "linkedoin", "linjkedin", "l9inkedin",
		"linksedin", "linkedrin", "linkedjin", "linkewdin", "linkeduin",
		"linked8in", "luinkedin", "libnkedin", "link4edin", "liunkedin",
		"inkedin", "linkein", "lnkedin", "linedin", "linkdin", "linkedn",
		"likedin", "linkedi", "llinkedin", "linkkedin", "linnkedin",
		"liinkedin", "linkedib", "linledin", "linkecin", "link4din", "l9nkedin",
		"linmedin", "linkedon", "link3din", "linkedih", "linkzdin", "linked9n",
		"linkedun", "linkesin", "lunkedin", "linkwdin", "linkrdin", "linkerin",
		"lonkedin", "linkedjn", "pinkedin", "linked8n", "kinkedin", "oinkedin",
		"linksdin", "linkexin", "l8nkedin", "lihkedin", "ljnkedin", "libkedin",
		"l.inkedin", "li.nkedin", "lin.kedin", "link.edin", "linke.din",
		"linked.in", "linkedi.n", "ilnkedin", "lnikedin", "liknedin",
		"linekdin", "linkdein", "linkeidn", "linkedni", "linkodin", "linkidin",
		"lenkedin", "linkeden", "linkedincom",
	}

	return &Brand{
		Name:       name,
		Original:   original,
		Whitelist:  whitelist,
		Suspicious: suspicious,
	}
}
