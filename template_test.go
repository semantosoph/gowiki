/*
Copyright (C) 2021 Sven Windisch <semantosoph@posteo.de>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gowiki

import (
	"strings"
	"testing"
)

func TestTemplateFinder(t *testing.T) {
	mw := "{{Testlate}}"
	t.Log(mw)

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetText()
	if strings.TrimSpace(l) != "" {
		t.Error("Error removing template ", l)
	}
}

func TestMagicMap(t *testing.T) {
	mw := "{{IPA|Deschd}}--{{Testlate}}"
	t.Log(mw)

	identityFunc := func(name string, params map[string]string) string { return params["1"] }
	NoHashFunctionsMap["ipa"] = true
	MagicMap["IPA"] = identityFunc

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetText()
	if strings.TrimSpace(l) != "Deschd--" {
		t.Error("Error in magic map ", l)
	}
}

func TestNilFunc(t *testing.T) {
	mw := "{{DISPLAYTITLE}}"
	t.Log(mw)

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetText()
	if strings.TrimSpace(l) != "" {
		t.Error("NIL entry in magic map not caught", l)
	}
}
