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

// Spotify brand properties.
func Spotify() *Brand {
	name := "spotify"
	original := []string{"spotify",}
	whitelist := []string{
		"spotify.com", "spotify.org", "spotify.it", "spotify.de", "spotify.fr",
		"spotify.nl", "spotify.es", "spotify.se", "spotify.no", "spotify.fi",
	}
	suspicious := []string{
	}

	return &Brand{
		Name:       name,
		Original:   original,
		Whitelist:  whitelist,
		Suspicious: suspicious,
	}
}
