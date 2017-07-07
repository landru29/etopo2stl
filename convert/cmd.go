package convert

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var inFile string
var stlFile string
var xyzFile string
var scaleS string

func init() {
	Cmd.PersistentFlags().StringVarP(&inFile, "inFile", "", "", "xyz file")
	Cmd.PersistentFlags().StringVarP(&stlFile, "stlFile", "", "", "out stl file")
	Cmd.PersistentFlags().StringVarP(&xyzFile, "xyzFile", "", "", "out xyz file")
	Cmd.PersistentFlags().StringVarP(&scaleS, "scale", "", "", "Altitude scale")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "convert",
	Short: "Process conversion",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		scale := float64(1)

		if len(scaleS) > 0 {
			scale, err = strconv.ParseFloat(scaleS, 64)
			if err != nil {
				log.Fatal(err)
			}
		}

		file, err := os.Open(inFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		xyz, err := ReadXyz(file)
		if err != nil {
			log.Fatal(err)
		}

		var xyzLength = Map(xyz)

		if len(xyzFile) > 0 {
			xyzOut, err := os.Create(xyzFile)
			if err != nil {
				log.Fatal(err)
			}
			defer xyzOut.Close()

			for _, xyzVector := range xyzLength {
				fmt.Fprintf(xyzOut, "%f %f %f\n", xyzVector.U, xyzVector.V, xyzVector.Altitude*scale)
			}

		}
	},
}
