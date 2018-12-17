package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"

	"github.com/pa-m/sklearn/cluster"
	"github.com/pa-m/sklearn/datasets"
	"github.com/pa-m/sklearn/preprocessing"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot/vg/draw"

	"github.com/pkg/errors"
)

func findClusterMembers(labels []int, X *mat.Dense) error {
	/*
		for
			  X := mat.NewDense(NSamples, 2, []float64{1, 2, 33, 3, 644, 7, 5555, 5})
			  X.At(i, j) // At returns the element at row i, column j.
			  X.At(0, 0) == 1
			  X.At(0, 1) == 2
			  X.At(0, 2) == Does NOT exist
			  X.At(1, 0) == 33
			  X.At(3, 1) == 5
			  X.At(3, 2) == Does not exist
			  X.At(4, 0) == Does not exist
	*/

	// cluster_members :=  map[string][]int

	rows, columns := X.Caps()
	fmt.Println("rows, columns", rows, columns)
	for k := range labels {
		for i := 0; i < columns; i++ {
			zAt := X.At(k, i)
			fmt.Println("zAt", zAt)
		}

	}

	// cluster_members =  map[string]int    //{"noise": []}
	// for val in labels:
	//     if val != -1:
	//         cluster_members.update({str(val): []})

	// for k, v in enumerate(labels):
	//     if v == -1:
	//         # noise
	//         cluster_members["noise"].append(X[k])
	//     for cluster_members_key in cluster_members.keys():
	//         if str(v) == cluster_members_key:
	//             cluster_members[cluster_members_key].append(X[k])
	return nil
}

var visualDebug = flag.Bool("visual", false, "output images for benchmarks and test data")

func main() {
	*visualDebug = true
	// Save the plot to a PNG file.
	pngfile := "GolangExampleDBSCAN.png"

	// adapted from http://scikit-learn.org/stable/_downloads/plot_dbscan.ipynb
	// Generate sample data
	centers := mat.NewDense(3, 2, []float64{1, 1, -1, -1, 1, -1})
	NSamples := 750
	X, _ := datasets.MakeBlobs(&datasets.MakeBlobsConfig{NSamples: NSamples, Centers: centers, ClusterStd: .3}) //RandomState: rand.New(rand.NewSource(0)),

	X, _ = preprocessing.NewStandardScaler().FitTransform(X, nil)

	db := cluster.NewDBSCAN(&cluster.DBSCANConfig{Eps: .3, MinSamples: 10, Algorithm: "kd_tree"})
	db.Fit(X, nil)
	coreSampleMask := make([]bool, len(db.Labels))
	for sample := range db.CoreSampleIndices {
		coreSampleMask[sample] = true
	}
	labels := db.Labels
	labelsmap := make(map[int]int)
	for _, l := range labels {
		labelsmap[l] = l
	}
	nclusters := len(labelsmap)
	if _, ok := labelsmap[-1]; ok {
		nclusters--
	}
	fmt.Printf("Estimated number of clusters: %d\n", nclusters)

	findClusterMembers(labels, X)

	if *visualDebug {

		// plot result
		p, err := plot.New()
		if err != nil {
			err = errors.Wrap(err, "error instantiating plot")
			log.Fatalf("%+v", err)

		}
		p.Title.Text = fmt.Sprintf("Estimated number of clusters: %d", nclusters)
		for cl := range labelsmap {
			var data plotter.XYs
			for sample := 0; sample < NSamples; sample++ {
				if labels[sample] == cl {
					data = append(data, struct{ X, Y float64 }{X.At(sample, 0), X.At(sample, 1)})
				}
			}
			s, err := plotter.NewScatter(data)
			if err != nil {
				err = errors.Wrap(err, "error instantiating plotter.NewScatter")
				log.Fatalf("%+v", err)
			}
			var color0 color.RGBA
			switch cl {
			case -1:
				color0 = color.RGBA{0, 0, 0, 255}
			case 0:
				color0 = color.RGBA{176, 0, 0, 255}
			case 1:
				color0 = color.RGBA{0, 176, 0, 255}
			case 2:
				color0 = color.RGBA{0, 0, 176, 255}
			}
			s.GlyphStyle.Color = color0
			s.GlyphStyle.Shape = draw.CircleGlyph{}
			p.Add(s)
			// p.Legend.Add(fmt.Sprintf("scatter %d", cl), s)

		}

		if err := p.Save(6*vg.Inch, 4*vg.Inch, pngfile); err != nil {
			err = errors.Wrap(err, "error saving png")
			log.Fatalf("%+v", err)
		}

	}
}
