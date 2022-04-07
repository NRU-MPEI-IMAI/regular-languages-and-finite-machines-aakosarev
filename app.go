package main

import (
	"github.com/goccy/go-graphviz"
	"io/ioutil"
	"log"
)

const (
	dotPath_1_1 = "./first_task/dot_files/gr_1_1.dot"
	dotPath_1_2 = "./first_task/dot_files/gr_1_2.dot"
	dotPath_1_3 = "./first_task/dot_files/gr_1_3.dot"
	dotPath_1_4 = "./first_task/dot_files/gr_1_4.dot"
	svgPath_1_1 = "./first_task/svg_files/gr_1_1.svg"
	svgPath_1_2 = "./first_task/svg_files/gr_1_2.svg"
	svgPath_1_3 = "./first_task/svg_files/gr_1_3.svg"
	svgPath_1_4 = "./first_task/svg_files/gr_1_4.svg"

	dotPath_2_1_1 = "./second_task/dot_files/gr_2_1_1.dot"
	dotPath_2_1_2 = "./second_task/dot_files/gr_2_1_2.dot"
	dotPath_2_1   = "./second_task/dot_files/gr_2_1.dot"
	dotPath_2_2_1 = "./second_task/dot_files/gr_2_2_1.dot"
	dotPath_2_2_2 = "./second_task/dot_files/gr_2_2_2.dot"
	dotPath_2_2   = "./second_task/dot_files/gr_2_2.dot"
	dotPath_2_3_1 = "./second_task/dot_files/gr_2_3_1.dot"
	dotPath_2_3_2 = "./second_task/dot_files/gr_2_3_2.dot"
	dotPath_2_3   = "./second_task/dot_files/gr_2_3.dot"
	dotPath_2_4   = "./second_task/dot_files/gr_2_4.dot"
	dotPath_2_5_1 = "./second_task/dot_files/gr_2_5_1.dot"
	dotPath_2_5_2 = "./second_task/dot_files/gr_2_5_2.dot"
	dotPath_2_5   = "./second_task/dot_files/gr_2_5.dot"
	svgPath_2_1_1 = "./second_task/svg_files/gr_2_1_1.svg"
	svgPath_2_1_2 = "./second_task/svg_files/gr_2_1_2.svg"
	svgPath_2_1   = "./second_task/svg_files/gr_2_1.svg"
	svgPath_2_2_1 = "./second_task/svg_files/gr_2_2_1.svg"
	svgPath_2_2_2 = "./second_task/svg_files/gr_2_2_2.svg"
	svgPath_2_2   = "./second_task/svg_files/gr_2_2.svg"
	svgPath_2_3_1 = "./second_task/svg_files/gr_2_3_1.svg"
	svgPath_2_3_2 = "./second_task/svg_files/gr_2_3_2.svg"
	svgPath_2_3   = "./second_task/svg_files/gr_2_3.svg"
	svgPath_2_4   = "./second_task/svg_files/gr_2_4.svg"
	svgPath_2_5_1 = "./second_task/svg_files/gr_2_5_1.svg"
	svgPath_2_5_2 = "./second_task/svg_files/gr_2_5_2.svg"
	svgPath_2_5   = "./second_task/svg_files/gr_2_5.svg"

	dotPath_3_1_1 = "./third_task/dot_files/gr_3_1_1.dot"
	dotPath_3_1   = "./third_task/dot_files/gr_3_1.dot"
	dotPath_3_2_1 = "./third_task/dot_files/gr_3_2_1.dot"
	dotPath_3_2   = "./third_task/dot_files/gr_3_2.dot"
	dotPath_3_3_1 = "./third_task/dot_files/gr_3_3_1.dot"
	dotPath_3_3   = "./third_task/dot_files/gr_3_3.dot"
	dotPath_3_4   = "./third_task/dot_files/gr_3_4.dot"
	svgPath_3_1_1 = "./third_task/svg_files/gr_3_1_1.svg"
	svgPath_3_1   = "./third_task/svg_files/gr_3_1.svg"
	svgPath_3_2_1 = "./third_task/svg_files/gr_3_2_1.svg"
	svgPath_3_2   = "./third_task/svg_files/gr_3_2.svg"
	svgPath_3_3_1 = "./third_task/svg_files/gr_3_3_1.svg"
	svgPath_3_3   = "./third_task/svg_files/gr_3_3.svg"
	svgPath_3_4   = "./third_task/svg_files/gr_3_4.svg"

	dotPath_4_1 = "./fourth_task/dot_files/gr_4_1.dot"
	svgPath_4_1 = "./fourth_task/svg_files/gr_4_1.svg"
)

func ConvertDotToSvg(dotPath, svgPath string) {
	g := graphviz.New()
	b, err := ioutil.ReadFile(dotPath)
	if err != nil {
		log.Fatal(err)
	}
	graph, err := graphviz.ParseBytes(b)
	if err != nil {
		log.Fatal(err)
	}
	if err := g.RenderFilename(graph, graphviz.SVG, svgPath); err != nil {
		log.Fatal(err)
	}
}

func main() {
	ConvertDotToSvg(dotPath_1_1, svgPath_1_1)
	ConvertDotToSvg(dotPath_1_2, svgPath_1_2)
	ConvertDotToSvg(dotPath_1_3, svgPath_1_3)
	ConvertDotToSvg(dotPath_1_4, svgPath_1_4)

	ConvertDotToSvg(dotPath_2_1_1, svgPath_2_1_1)
	ConvertDotToSvg(dotPath_2_1_2, svgPath_2_1_2)
	ConvertDotToSvg(dotPath_2_1, svgPath_2_1)
	ConvertDotToSvg(dotPath_2_2_1, svgPath_2_2_1)
	ConvertDotToSvg(dotPath_2_2_2, svgPath_2_2_2)
	ConvertDotToSvg(dotPath_2_2, svgPath_2_2)
	ConvertDotToSvg(dotPath_2_3_1, svgPath_2_3_1)
	ConvertDotToSvg(dotPath_2_3_2, svgPath_2_3_2)
	ConvertDotToSvg(dotPath_2_3, svgPath_2_3)
	ConvertDotToSvg(dotPath_2_4, svgPath_2_4)
	ConvertDotToSvg(dotPath_2_5_1, svgPath_2_5_1)
	ConvertDotToSvg(dotPath_2_5_2, svgPath_2_5_2)
	ConvertDotToSvg(dotPath_2_5, svgPath_2_5)

	ConvertDotToSvg(dotPath_3_1_1, svgPath_3_1_1)
	ConvertDotToSvg(dotPath_3_1, svgPath_3_1)
	ConvertDotToSvg(dotPath_3_2_1, svgPath_3_2_1)
	ConvertDotToSvg(dotPath_3_2, svgPath_3_2)
	ConvertDotToSvg(dotPath_3_3_1, svgPath_3_3_1)
	ConvertDotToSvg(dotPath_3_3, svgPath_3_3)
	ConvertDotToSvg(dotPath_3_4, svgPath_3_4)

	ConvertDotToSvg(dotPath_4_1, svgPath_4_1)
}
