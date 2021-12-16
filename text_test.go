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

func TestFulltext(t *testing.T) {
	mw := "Test[[File:Deschd]]Test"
	t.Log(mw)

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetText()
	if strings.TrimSpace(l) != "Test\nFile:Deschd\nTest" {
		t.Error("Error generating full text ", l)
	}
}

func TestAbstractNoFile(t *testing.T) {
	mw := "Test[[File:Deschd]]Test"
	t.Log(mw)

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetAbstract()
	if strings.TrimSpace(l) != "TestTest" {
		t.Error("Error removing file link ", l)
	}
}

func TestAbstractNoHeadline(t *testing.T) {
	mw := "ThisisAbstract\n== Testheadline ==\n NotanAbstract"
	t.Log(mw)

	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}

	l := a.GetAbstract()
	if strings.TrimSpace(l) != "ThisisAbstract" {
		t.Error("Error removing headline ", l)
	}
}
