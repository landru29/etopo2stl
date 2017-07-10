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
var abovePlan string
var logFunction bool
var offsetS string
var pedestalS string

func init() {
	Cmd.PersistentFlags().StringVarP(&inFile, "inFile", "", "", "xyz file <filename>")
	//Cmd.PersistentFlags().StringVarP(&stlFile, "stlFile", "", "", "out stl file")
	Cmd.PersistentFlags().StringVarP(&xyzFile, "xyzFile", "", "", "out xyz file <filename>")
	Cmd.PersistentFlags().StringVarP(&scaleS, "scale", "", "", "Altitude scale <scale>")
	Cmd.PersistentFlags().StringVarP(&abovePlan, "above", "", "", "Keep above plan <altitude>")
	Cmd.PersistentFlags().BoolVarP(&logFunction, "log", "", false, "apply logarythm")
	Cmd.PersistentFlags().StringVarP(&offsetS, "offset", "", "", "add offset <offset>")
	Cmd.PersistentFlags().StringVarP(&pedestalS, "pedestal", "", "", "add pedestral <thickness>")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "convert",
	Short: "Process conversion",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		scale := float64(1)
		offset := float64(0)

		file, err := os.Open(inFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Read file
		xyz, err := ReadXyz(file)
		if err != nil {
			log.Fatal(err)
		}

		// Above plan
		if len(abovePlan) > 0 {
			plan, err := strconv.ParseFloat(abovePlan, 64)
			if err != nil {
				log.Fatal(err)
			}
			xyz = Above(xyz, plan)
		}

		// Log
		if logFunction {
			xyz = Logarythm(xyz)
		}

		// Offset
		if len(offsetS) > 0 {
			offset, err = strconv.ParseFloat(offsetS, 64)
			if err != nil {
				log.Fatal(err)
			}
			xyz = Offset(xyz, offset)
		}

		// Pedestral
		if len(pedestalS) > 0 {
			thickness, err := strconv.ParseFloat(pedestalS, 64)
			if err != nil {
				log.Fatal(err)
			}
			xyz = Pedestal(xyz, thickness)
		}

		// Scale
		if len(scaleS) > 0 {
			scale, err = strconv.ParseFloat(scaleS, 64)
			if err != nil {
				log.Fatal(err)
			}
			xyz = Scale(xyz, scale)
		}

		// polar to length
		var xyzLength = PolarToLength(xyz)

		// Write to file
		if len(xyzFile) > 0 {
			xyzOut, err := os.Create(xyzFile)
			if err != nil {
				log.Fatal(err)
			}
			defer xyzOut.Close()

			for _, xyzVector := range xyzLength {
				fmt.Fprintf(xyzOut, "%f %f %f\n", xyzVector.U, xyzVector.V, xyzVector.Altitude)
			}

		}
	},
}
