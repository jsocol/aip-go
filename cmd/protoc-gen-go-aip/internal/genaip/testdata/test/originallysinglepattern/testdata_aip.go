// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc (unknown)
// source: test/originallysinglepattern/testdata.proto

package originallysinglepattern

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type BookMultiPatternResourceName interface {
	fmt.Stringer
	MarshalString() (string, error)
}

func ParseBookMultiPatternResourceName(name string) (BookMultiPatternResourceName, error) {
	switch {
	case resourcename.Match("shelves/{shelf}/books/{book}", name):
		var result ShelvesBookResourceName
		return &result, result.UnmarshalString(name)
	case resourcename.Match("publishers/{publisher}/books/{book}", name):
		var result PublishersBookResourceName
		return &result, result.UnmarshalString(name)
	default:
		return nil, fmt.Errorf("no matching pattern")
	}
}

type BookResourceName struct {
	Shelf string
	Book  string
}

func (n BookResourceName) Validate() error {
	if n.Shelf == "" {
		return fmt.Errorf("shelf: empty")
	}
	if strings.IndexByte(n.Shelf, '/') != -1 {
		return fmt.Errorf("shelf: contains illegal character '/'")
	}
	if n.Book == "" {
		return fmt.Errorf("book: empty")
	}
	if strings.IndexByte(n.Book, '/') != -1 {
		return fmt.Errorf("book: contains illegal character '/'")
	}
	return nil
}

func (n BookResourceName) ContainsWildcard() bool {
	return false || n.Shelf == "-" || n.Book == "-"
}

func (n BookResourceName) String() string {
	return resourcename.Sprint(
		"shelves/{shelf}/books/{book}",
		n.Shelf,
		n.Book,
	)
}

func (n BookResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *BookResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"shelves/{shelf}/books/{book}",
		&n.Shelf,
		&n.Book,
	)
}

type ShelvesBookResourceName struct {
	Shelf string
	Book  string
}

func (n ShelvesBookResourceName) Validate() error {
	if n.Shelf == "" {
		return fmt.Errorf("shelf: empty")
	}
	if strings.IndexByte(n.Shelf, '/') != -1 {
		return fmt.Errorf("shelf: contains illegal character '/'")
	}
	if n.Book == "" {
		return fmt.Errorf("book: empty")
	}
	if strings.IndexByte(n.Book, '/') != -1 {
		return fmt.Errorf("book: contains illegal character '/'")
	}
	return nil
}

func (n ShelvesBookResourceName) ContainsWildcard() bool {
	return false || n.Shelf == "-" || n.Book == "-"
}

func (n ShelvesBookResourceName) String() string {
	return resourcename.Sprint(
		"shelves/{shelf}/books/{book}",
		n.Shelf,
		n.Book,
	)
}

func (n ShelvesBookResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *ShelvesBookResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"shelves/{shelf}/books/{book}",
		&n.Shelf,
		&n.Book,
	)
}

type PublishersBookResourceName struct {
	Publisher string
	Book      string
}

func (n PublishersBookResourceName) Validate() error {
	if n.Publisher == "" {
		return fmt.Errorf("publisher: empty")
	}
	if strings.IndexByte(n.Publisher, '/') != -1 {
		return fmt.Errorf("publisher: contains illegal character '/'")
	}
	if n.Book == "" {
		return fmt.Errorf("book: empty")
	}
	if strings.IndexByte(n.Book, '/') != -1 {
		return fmt.Errorf("book: contains illegal character '/'")
	}
	return nil
}

func (n PublishersBookResourceName) ContainsWildcard() bool {
	return false || n.Publisher == "-" || n.Book == "-"
}

func (n PublishersBookResourceName) String() string {
	return resourcename.Sprint(
		"publishers/{publisher}/books/{book}",
		n.Publisher,
		n.Book,
	)
}

func (n PublishersBookResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *PublishersBookResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"publishers/{publisher}/books/{book}",
		&n.Publisher,
		&n.Book,
	)
}
