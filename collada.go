// collada.go
// 
// A simple COLLADA file loader used to provide simple meshes from COLLADA
// files.  Over time, I may expand this to import more data from the file,
// but for now, we only import meshes.
//
// Based largely on work done by Stan Steel, see:
// http://www.kryas.com/ as well as information drawn from
// http://www.wazim.com/Collada_Tutorial_1.htm
// and, of course, the COLLADA specification:
// https://collada.org/mediawiki/index.php/COLLADA_-_Digital_Asset_and_FX_Exchange_Schema
//
// Based on Collada format 1.4.1, should work with 1.5.

package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Collada struct {
	//Id                    string `xml:"attr"`
	Version               string `xml:"attr"`
	Library_Geometries    LibraryGeometries
	Library_Visual_Scenes LibraryVisualScenes
}

type LibraryGeometries struct {
	XMLName  xml.Name `xml:"library_geometries"`
	Geometry []Geometry
}

type Geometry struct {
	XMLName xml.Name `xml:"geometry"`
	Id      string   `xml:"attr"`
	Mesh    Mesh
}

type Mesh struct {
	XMLName  xml.Name `xml:"mesh"`
	Source   []Source
	Polylist Polylist
}

type Source struct {
	XMLName     xml.Name   `xml:"source"`
	Id          string     `xml:"attr"`
	Float_array FloatArray `xml:"float_array"`
}

type FloatArray struct {
	XMLName xml.Name `xml:"float_array"`
	Id      string   `xml:"attr"`
	CDATA   string   `xml:"chardata"`
	Count   string   `xml:"attr"`
}

type Polylist struct {
	XMLName xml.Name `xml:"polylist"`
	Id      string   `xml:"attr"`
	Count   string   `xml:"attr"`

	// List of integers, each specifying the number of vertices for one polygon
	VCount string `xml:"vcount"`

	// List of integers that specify the vertex attributes
	P string `xml:"p"`
}

type LibraryVisualScenes struct {
	XMLName     xml.Name `xml:"library_visual_scenes"`
	VisualScene VisualScene
}

type VisualScene struct {
	XMLName xml.Name `xml:"visual_scene"`
}

// Debug functions
func (c *Collada) Debug() {
	fmt.Fprintf(os.Stdout, "*** COLLADA ***\n")
	//fmt.Fprintf(os.Stdout, "* ID: %s\n", c.Id)
	fmt.Fprintf(os.Stdout, "* Version: %s\n", c.Version)
	c.Library_Geometries.Debug()
}

func (l *LibraryGeometries) Debug() {
	fmt.Fprintf(os.Stdout, "*** Library Geometry ***\n")
	fmt.Fprintf(os.Stdout, "* Number of Geometries: %d\n", len(l.Geometry))
	for _, g := range l.Geometry {
		g.Debug()
	}
}

func (g *Geometry) Debug() {
	fmt.Fprintf(os.Stdout, "*** Geometry ***\n")
	fmt.Fprintf(os.Stdout, "* ID: %s\n", g.Id)
	g.Mesh.Debug()
}

func (m *Mesh) Debug() {
	fmt.Fprintf(os.Stdout, "*** Mesh ***\n")
	fmt.Fprintf(os.Stdout, "* Number of Sources: %d\n", len(m.Source))
	for _, s := range m.Source {
		s.Debug()
	}
	m.Polylist.Debug()
}

func (s *Source) Debug() {
	fmt.Fprintf(os.Stdout, "*** Source ***\n")
	fmt.Fprintf(os.Stdout, "* ID: %s\n", s.Id)
	s.Float_array.Debug()
}

func (f *FloatArray) Debug() {
	fmt.Fprintf(os.Stdout, "*** FloatArray ***\n")
	fmt.Fprintf(os.Stdout, "* Id: %s\n", f.Id)
	fmt.Fprintf(os.Stdout, "* CDATA: %s\n", f.CDATA)
	fmt.Fprintf(os.Stdout, "* Count: %s\n", f.Count)
}

func (p *Polylist) Debug() {
	fmt.Fprintf(os.Stdout, "*** Polylist ***\n")
	fmt.Fprintf(os.Stdout, "* ID: %s\n", p.Id)
	fmt.Fprintf(os.Stdout, "* Count: %s\n", p.Count)
	fmt.Fprintf(os.Stdout, "* VCount: %s\n", p.VCount)
	fmt.Fprintf(os.Stdout, "* P: %s\n", p.P)
}