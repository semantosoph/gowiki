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

func TestMediaLink(t *testing.T) {
	mw := "[[File:test.png]]"
	t.Log(mw)

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetMedia()
	if l[0].Namespace != "File" || l[0].PageName != "Test.png" {
		t.Error("Error parsing media link ", l)
	}
}

func TestCustomMediaLink(t *testing.T) {
	mw := "[[Datei:test.jpg]]"
	t.Log(mw)

	StandardNamespaces["datei"] = "Datei"
	FileLinkPrefixes = append(FileLinkPrefixes, "[[datei:")

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetMedia()
	if l[0].Namespace != "Datei" || l[0].PageName != "Test.jpg" {
		t.Error("Error parsing media link ", l)
	}
}

func TestExternalLink(t *testing.T) {
	mw := "[[https://test.org|Test]]"
	t.Log(mw)

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetTextLinks()
	if l[0].Text != "Test" || l[0].Link.PageName != "Https://test.org" {
		t.Error("Error parsing ext link ", l)
	}
}

func TestRefRemoval(t *testing.T) {
	mw := "Test<ref name=\"testref\">This is a text reference</ref>Test<ref>{{curly ref}}</ref>Test"
	t.Log(mw)

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetText()
	if strings.TrimSpace(l) != "TestTestTest" {
		t.Error("Error removing ref ", l)
	}
}
